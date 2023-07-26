package room

import "errors"

var (
	// ErrRoomNotExist 房间不存在
	ErrRoomNotExist = errors.New("room not exist")
	// ErrRoomPlayerFull 房间人数已满
	ErrRoomPlayerFull = errors.New("room player full")
	// ErrPlayerNotExist 玩家不存在
	ErrPlayerNotExist = errors.New("player not exist")
)
