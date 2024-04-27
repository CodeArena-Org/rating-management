package helpers

import (
	"github.com/CodeArena-Org/rating-management/errors"
	"github.com/CodeArena-Org/rating-management/grpc_mngm"
)

func ValidateAssignProblemRequest(req *grpc_mngm.AssignProblemRequest) error {
	// validate incoming data
	if req.User1Id == "" {
		return errors.ErrInvalidUserId
	}
	if req.User2Id == "" {
		return errors.ErrInvalidUserId
	}
	if req.User1Rating <= 0 {
		return errors.ErrInvalidCurrentRating
	}
	if req.User2Rating <= 0 {
		return errors.ErrInvalidCurrentRating
	}

	return nil
}
