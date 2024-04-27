package helpers

import (
	"math"

	"github.com/CodeArena-Org/rating-management/errors"
	"github.com/CodeArena-Org/rating-management/grpc_mngm"
	"github.com/CodeArena-Org/rating-management/models"
)

func ValidateScoreRequst(req *grpc_mngm.GetScoreRequest) error {
	// validate incoming data
	if req.AcceptedTime <= 0 {
		return errors.ErrInvalidAcceptedTime
	}
	if req.ContestTime != models.TenMinuteSlot && req.ContestTime != models.ThirtyMinuteSlot && req.ContestTime != models.SixtyMinuteSlot {
		return errors.ErrInvalidContestTime
	}
	if req.WinnerCurrentRating <= 0 {
		return errors.ErrInvalidCurrentRating
	}
	if req.LooserCurrentRating <= 0 {
		return errors.ErrInvalidCurrentRating
	}
	if req.WrongAttempts < 0 {
		return errors.ErrInvalidWrongAttempts
	}
	if req.WinnerId == "" {
		return errors.ErrInvalidUserId
	}
	if req.LooserId == "" {
		return errors.ErrInvalidUserId
	}

	return nil
}

// Coefficients of the polynomial equation
var coefficients = []float64{1.203e-05, -0.09414, 199}

// Function to calculate y value for a given x using the polynomial equation
func GetPosScore(x float64) float64 {
	// Evaluate the polynomial equation for the given x
	y := coefficients[0]*math.Pow(x, 2) + coefficients[1]*x + coefficients[2]
	return y
}

// Polynomial coefficients
var coefficients2 = []float64{3.74e-19, -6.245e-15, 4.244e-11, -1.501e-07, 0.0002908, -0.3129, 120}

// Function to evaluate the polynomial equation at a given x value
func GetNegScore(x float64) float64 {
	var y float64
	for i, coef := range coefficients2 {
		y += coef * math.Pow(x, float64(len(coefficients2)-i-1))
	}
	return y
}
