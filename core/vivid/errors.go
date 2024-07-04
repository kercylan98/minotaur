package vivid

import "errors"

var (
	ErrFutureTimeout     = errors.New("vivid: future timeout")
	ErrActorAlreadyExist = errors.New("vivid: actor already exist")
)
