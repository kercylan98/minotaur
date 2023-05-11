package builtin

import "errors"

var (
	ErrRoomPlayerLimit = errors.New("the number of players in the room has reached the upper limit") // 玩家数量达到上限
	ErrRoomNoHasMaster = errors.New("room not has master, can't kick player")
	ErrRoomNotIsOwner  = errors.New("not is room owner, can't kick player")
)
