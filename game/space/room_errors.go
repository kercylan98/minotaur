package space

import "errors"

var (
	// ErrRoomFull 房间已满
	ErrRoomFull = errors.New("room is full")
	// ErrSeatNotEmpty 座位上已经有实体
	ErrSeatNotEmpty = errors.New("seat is not empty")
	// ErrNotInRoom 实体不在房间中
	ErrNotInRoom = errors.New("not in room")
	// ErrRoomPasswordNotMatch 房间密码不匹配
	ErrRoomPasswordNotMatch = errors.New("room password not match")
	// ErrPermissionDenied 权限不足
	ErrPermissionDenied = errors.New("permission denied")
)
