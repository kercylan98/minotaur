package builtin

import "errors"

var (
	ErrWorldPlayerLimit = errors.New("the number of players in the world has reached the upper limit")
	ErrWorldReleased    = errors.New("the world has been released")
)
