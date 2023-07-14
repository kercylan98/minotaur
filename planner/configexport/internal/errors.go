package internal

import "errors"

var (
	ErrReadConfigFailedIgnore                     = errors.New("read config skip ignore")
	ErrReadConfigFailedSame                       = errors.New("read config skip, same name")
	ErrReadConfigFailedWithDisplayName            = errors.New("read config display name failed, can not found position")
	ErrReadConfigFailedWithName                   = errors.New("read config name failed, can not found position")
	ErrReadConfigFailedWithIndexCount             = errors.New("read config index count failed, can not found position")
	ErrReadConfigFailedWithIndexCountLessThanZero = errors.New("read config index count failed, value less than zero")
	ErrReadConfigFailedWithNameDuplicate          = errors.New("read config index count failed, duplicate field names")
	ErrReadConfigFailedWithIndexTypeException     = errors.New("read config index count failed, the index type is only allowed to be the basic type")
)
