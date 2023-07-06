package domain

import "errors"

var (
	ErrGameDoesNotExist        = errors.New("game does not exist")
	ErrInvalidHostNickname     = errors.New("invalid host nickname")
	ErrInvalidOpponentNickname = errors.New("invalid opponent nickname")
	ErrInvalidFieldIndex       = errors.New("invalid field index")
	ErrInvalidHostField        = errors.New("invalid host field")
	ErrInvalidOpponentField    = errors.New("invalid opponent field")
)
