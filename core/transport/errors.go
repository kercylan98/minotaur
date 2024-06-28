package transport

import "errors"

var (
	ErrorMessageMustIsProtoMessage = errors.New("message must be a proto.Message")
)
