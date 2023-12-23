# Space

[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/kercylan98/minotaur/game/space)

计划提供游戏中常见的空间设计，例如房间、地图等。开发者可以使用它来快速构建游戏中的常见空间，例如多人房间、地图等。

## Room [`房间`]((https://pkg.go.dev/github.com/kercylan98/minotaur/game/space#RoomManager))
房间在 `Minotaur` 中仅仅只是一个可以为任意可比较类型的 `ID`，当需要将现有或新设计的房间纳入 [`RoomManager`](https://pkg.go.dev/github.com/kercylan98/minotaur/game/space#RoomManager) 管理时，需要实现 [`Room`](https://pkg.go.dev/github.com/kercylan98/minotaur/game/space#RoomManager) 管理时，仅需要实现 [`generic.IdR`](https://pkg.go.dev/github.com/kercylan98/minotaur/utils/generic#IdR) 接口即可。

该功能由 
[`RoomManager`](https://pkg.go.dev/github.com/kercylan98/minotaur/game/space#RoomManager)、
[`RoomController`](https://pkg.go.dev/github.com/kercylan98/minotaur/game/space#RoomController)
组成。

当创建一个新的房间并纳入 [`RoomManager`](https://pkg.go.dev/github.com/kercylan98/minotaur/game/space#RoomManager) 管理后，将会得到一个 [`RoomController`](https://pkg.go.dev/github.com/kercylan98/minotaur/game/space#RoomController)。通过 [`RoomController`](https://pkg.go.dev/github.com/kercylan98/minotaur/game/space#RoomController) 可以对房间进行管理，例如：获取房间信息、加入房间、退出房间等。

### 使用示例
```go
package main

import (
    "fmt"
    "github.com/kercylan98/minotaur/game/space"
)

type Room struct {
	Id int64
}

func (r *Room) GetId() int64 {
	return r.Id
}

type Player struct {
	Id string
}

func (p *Player) GetId() string {
	return p.Id
}

func main() {
	var rm = space.NewRoomManager[string, int64, *Player, *Room]()
	var room = &Room{Id: 1}
	var controller = rm.AssumeControl(room)

	if err := controller.AddEntity(&Player{Id: "1"}); err != nil {
		// 房间密码不匹配或者房间已满
		panic(err)
	}

	fmt.Println(controller.GetEntityCount()) // 1
}
```