package lockstep

import "errors"

const (
	tickerFrameName = "LOCKSTEP_FRAME"
)

var (
	ErrFrameFactorCanNotIsNull = errors.New("frameFactory can not is nil")
)
