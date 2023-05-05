package builtin

import "errors"

var (
	ErrWorldPlayerLimit = errors.New("the number of players in the world has reached the upper limit") // 玩家数量达到上限
	ErrWorldReleased    = errors.New("the world has been released")                                    // 世界已被释放
)
