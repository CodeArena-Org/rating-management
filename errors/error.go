package errors

import "errors"

var ErrInvalidAcceptedTime = errors.New("accepted time cannot be negative or zero")
var ErrInvalidContestTime = errors.New("contest time cannot be negative or zero")
var ErrInvalidCurrentRating = errors.New("current rating cannot be negative or zero")
var ErrInvalidProblemLevel = errors.New("problem level cannot be empty")
var ErrInvalidWrongAttempts = errors.New("wrong attempts cannot be negative")
var ErrInvalidUserId = errors.New("user id cannot be empty")