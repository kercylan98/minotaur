package internal

import "errors"

var (
	ErrReadConfigFailedWithDisplayName            = errors.New("read config display name failed, can not found position")
	ErrReadConfigFailedWithName                   = errors.New("read config name failed, can not found position")
	ErrReadConfigFailedWithIndexCount             = errors.New("read config index count failed, can not found position")
	ErrReadConfigFailedWithIndexCountLessThanZero = errors.New("read config index count failed, value less than zero")
	ErrReadConfigFailedWithFieldPosition          = errors.New("read config index count failed, field position exception")
	ErrReadConfigFailedWithNameDuplicate          = errors.New("read config index count failed, duplicate field names")
	ErrReadConfigFailedWithExportParamException   = errors.New("read config index count failed, export param must is s or c or sc or cs")
	ErrReadConfigFailedWithIndexTypeException     = errors.New("read config index count failed, the index type is only allowed to be the basic type")
)
