package socket

import "errors"

const (
	internalWebSocketCloseMessageType int = 8
	readDeadlineTaskName                  = "socket_internal_read_deadline"
	writeDeadlineTaskName                 = "socket_internal_write_deadline"
)

var (
	readDeadlineError  = errors.New("read deadline timeout")
	writeDeadlineError = errors.New("write deadline timeout")
)
