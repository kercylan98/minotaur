package sorts

import "errors"

var (
	ErrCircularDependencyDetected = errors.New("circular dependency detected")
)
