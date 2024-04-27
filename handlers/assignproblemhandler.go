package handlers

import (
	"context"
	"fmt"

	"github.com/CodeArena-Org/rating-management/db"
	gm "github.com/CodeArena-Org/rating-management/grpc_mngm"
)

func AssignProblem(ctx context.Context, req *gm.AssignProblemRequest) (int32, error) {
	// Assign a problem to the user
	avgRating := (req.User1Rating + req.User2Rating) / 2

	// find problem with rating closest to avgRating
	problemId, err := findProblem(avgRating)
	if err != nil {
		return 0, fmt.Errorf("unable to assign problem: %v", err)
	}

	// insert into problemsolved table
	_, err = db.DB.Exec("INSERT INTO problemsolved (userid, problemid) VALUES ($1, $2) ON CONFLICT DO NOTHING", req.User1Id, problemId)
	if err != nil {
		return 0, fmt.Errorf("unable to assign problem: %v", err)
	}
	_, err = db.DB.Exec("INSERT INTO problemsolved (userid, problemid) VALUES ($1, $2) ON CONFLICT DO NOTHING", req.User2Id, problemId)
	if err != nil {
		return 0, fmt.Errorf("unable to assign problem: %v", err)
	}
	return problemId, nil
}

func findProblem(avgRating int32) (int32, error) {
	// find problem with rating closest to avgRating , also exculde the prev solved problems
	rows, err := db.DB.Query("SELECT problemid, problemname, ratingrange FROM problems WHERE ratingrange[1] <= $1 AND ratingrange[2] >= $1 AND problemid NOT IN (SELECT problemid FROM problemsolved)", avgRating)
	if err != nil {
		return 0, fmt.Errorf("unable to assign problem: %v", err)
	}
	defer rows.Close()
	var problemId int32
	var problemName string
	var ratingRange []uint8
	for rows.Next() {
		err = rows.Scan(&problemId, &problemName, &ratingRange)
		if err != nil {
			return 0, fmt.Errorf("unable to assign problem: %v", err)
		}
		// print problem in json form
		fmt.Printf("Problem: {\"problemId\": %d, \"problemName\": %s, \"ratingRange\": %s}\n", problemId, problemName, ratingRange)

		if problemId != 0 {
			return problemId, nil
		}
	}

	// No problem found, return default value
	return 0, fmt.Errorf("unable to assign problem: no problem found")

}
