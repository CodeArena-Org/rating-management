package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/CodeArena-Org/rating-management/db"
	gm "github.com/CodeArena-Org/rating-management/grpc_mngm"
	"github.com/CodeArena-Org/rating-management/handlers"
	"github.com/CodeArena-Org/rating-management/helpers"
)

type Server struct {
	gm.UnimplementedRatingManagementServer
}

func (s *Server) GetWinnerScore(ctx context.Context, req *gm.GetScoreRequest) (*gm.GetScoreResponse, error) {
	fmt.Println("GetWinnerScore called")
	err := helpers.ValidateScoreRequst(req)
	if err != nil {
		return nil, err
	}
	winnerscore, looserscore := handlers.CalculateScore(ctx, req)
	db.UpdateRating(req.WinnerId, winnerscore)
	db.UpdateRating(req.LooserId, looserscore)
	return &gm.GetScoreResponse{
		WinnerScore: winnerscore,
		LooserScore: looserscore,
	}, err
}

func (s *Server) AssignProblem(ctx context.Context, req *gm.AssignProblemRequest) (*gm.AssignProblemResponse, error) {
	fmt.Println("AssignProblem called")
	err := helpers.ValidateAssignProblemRequest(req)
	if err != nil {
		return nil, err
	}
	problemId, err := handlers.AssignProblem(ctx, req)
	return &gm.AssignProblemResponse{
		ProblemId: problemId,
	}, err
}

func main() {
	// start network listener
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer lis.Close()

	// connect to postgress
	err = db.ConnectToPostGress()
	if err != nil {
		log.Fatalf("failed to connect to postgress: %v", err)
	}

	// create grpc server
	s := grpc.NewServer()
	gm.RegisterRatingManagementServer(s, &Server{})
	log.Printf("gRPC server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
