package builtin

import "errors"

var (
	ErrRoomNoHasMaster = errors.New("room not has master, can't kick player")
	ErrRoomNotIsOwner  = errors.New("not is room owner, can't kick player")
)
