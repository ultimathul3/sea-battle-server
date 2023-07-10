package domain

import "errors"

var (
	ErrGameDoesNotExist        = errors.New("game does not exist")
	ErrInvalidHostNickname     = errors.New("invalid host nickname")
	ErrInvalidOpponentNickname = errors.New("invalid opponent nickname")
	ErrInvalidHostField        = errors.New("invalid host field")
	ErrInvalidOpponentField    = errors.New("invalid opponent field")
	ErrInvalidField            = errors.New("invalid field")
	ErrInvalidHitCell          = errors.New("invalid hit cell")
	ErrInvalidAction           = errors.New("invalid action")
)
