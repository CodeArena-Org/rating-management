package handlers

import (
	"context"

	gm "github.com/CodeArena-Org/rating-management/grpc_mngm"
	"github.com/CodeArena-Org/rating-management/helpers"
)

func CalculateScore(ctx context.Context, req *gm.GetScoreRequest) (int32, int32) {
	wrong_attempts_score := req.WrongAttempts * 5
	time_elapsed_score := (req.AcceptedTime * 50) / req.ContestTime
	max_score := int32(helpers.GetPosScore(float64(req.WinnerCurrentRating)))
	winner_score := max_score - wrong_attempts_score - time_elapsed_score
	looser_score := int32(helpers.GetNegScore(float64(req.LooserCurrentRating)))
	looser_score = looser_score - time_elapsed_score/2

	return winner_score, looser_score
}
