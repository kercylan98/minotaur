package room

import "errors"

var (
	// ErrRoomNotExist 房间不存在
	ErrRoomNotExist = errors.New("room not exist")
	// ErrRoomPlayerFull 房间人数已满
	ErrRoomPlayerFull = errors.New("room player full")
	// ErrPlayerNotInRoom 玩家不在房间中
	ErrPlayerNotInRoom = errors.New("player not in room")
	// ErrRoomOrPlayerNotExist 房间不存在或玩家不在房间中
	ErrRoomOrPlayerNotExist = errors.New("room or player not exist")
)
