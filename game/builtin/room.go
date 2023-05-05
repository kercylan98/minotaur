package builtin

import (
	"go.uber.org/zap"
	"minotaur/game"
	"minotaur/utils/log"
	"minotaur/utils/synchronization"
)

func NewRoom[PlayerID comparable, Player game.Player[PlayerID]](guid int64) *Room[PlayerID, Player] {
	return &Room[PlayerID, Player]{
		guid:        guid,
		playersConn: synchronization.NewMap[string, Player](),
		players:     synchronization.NewMap[PlayerID, Player](),
	}
}

type Room[PlayerID comparable, Player game.Player[PlayerID]] struct {
	guid            int64
	owner           PlayerID
	noMaster        bool
	playerLimit     int
	playersConn     *synchronization.Map[string, Player]
	players         *synchronization.Map[PlayerID, Player]
	kickCheckHandle func(id, target PlayerID) error

	playerJoinRoomEventHandles  []game.PlayerJoinRoomEventHandle[PlayerID, Player]
	playerLeaveRoomEventHandles []game.PlayerLeaveRoomEventHandle[PlayerID, Player]
	playerKickedOutEventHandles []game.PlayerKickedOutEventHandle[PlayerID, Player]
}

func (slf *Room[PlayerID, Player]) GetGuid() int64 {
	return slf.guid
}

func (slf *Room[PlayerID, Player]) GetPlayerLimit() int {
	return slf.playerLimit
}

func (slf *Room[PlayerID, Player]) GetPlayerWithConnID(id string) Player {
	return slf.playersConn.Get(id)
}

func (slf *Room[PlayerID, Player]) GetPlayer(id PlayerID) Player {
	return slf.players.Get(id)
}

func (slf *Room[PlayerID, Player]) GetPlayers() synchronization.MapReadonly[PlayerID, Player] {
	return slf.players
}

func (slf *Room[PlayerID, Player]) IsExistPlayer(id PlayerID) bool {
	return slf.players.Exist(id)
}

func (slf *Room[PlayerID, Player]) IsOwner(id PlayerID) bool {
	return !slf.noMaster && slf.owner == id
}

func (slf *Room[PlayerID, Player]) Join(player Player) error {
	if slf.players.Size() >= slf.playerLimit && slf.playerLimit > 0 {
		return ErrWorldPlayerLimit
	}
	log.Debug("Room.Join", zap.Any("guid", slf.GetGuid()), zap.Any("player", player.GetID()), zap.String("conn", player.GetConnID()))
	slf.players.Set(player.GetID(), player)
	slf.playersConn.Set(player.GetConnID(), player)
	if slf.players.Size() == 1 && !slf.noMaster {
		slf.owner = player.GetID()
	}
	slf.OnPlayerJoinRoomEvent(player)
	return nil
}

func (slf *Room[PlayerID, Player]) Leave(id PlayerID) {
	player, exist := slf.players.GetExist(id)
	if !exist {
		return
	}
	log.Debug("Room.Leave", zap.Any("guid", slf.GetGuid()), zap.Any("player", player.GetID()), zap.String("conn", player.GetConnID()))
	slf.OnPlayerLeaveRoomEvent(player)
	slf.players.Delete(player.GetID())
	slf.playersConn.Delete(player.GetConnID())
}

func (slf *Room[PlayerID, Player]) KickOut(id, target PlayerID, reason string) error {
	player, exist := slf.players.GetExist(target)
	if !exist {
		return nil
	}
	if slf.kickCheckHandle != nil {
		if err := slf.kickCheckHandle(id, target); err != nil {
			return err
		}
	} else if slf.noMaster {
		return ErrRoomNoHasMaster
	} else if slf.owner != id {
		return ErrRoomNotIsOwner
	}

	slf.OnPlayerKickedOutEvent(id, target, reason)
	slf.Leave(player.GetID())
	return nil
}

func (slf *Room[PlayerID, Player]) RegPlayerJoinRoomEvent(handle game.PlayerJoinRoomEventHandle[PlayerID, Player]) {
	slf.playerJoinRoomEventHandles = append(slf.playerJoinRoomEventHandles, handle)
}

func (slf *Room[PlayerID, Player]) OnPlayerJoinRoomEvent(player Player) {
	for _, handle := range slf.playerJoinRoomEventHandles {
		handle(slf, player)
	}
}

func (slf *Room[PlayerID, Player]) RegPlayerLeaveRoomEvent(handle game.PlayerLeaveRoomEventHandle[PlayerID, Player]) {
	slf.playerLeaveRoomEventHandles = append(slf.playerLeaveRoomEventHandles, handle)
}

func (slf *Room[PlayerID, Player]) OnPlayerLeaveRoomEvent(player Player) {
	for _, handle := range slf.playerLeaveRoomEventHandles {
		handle(slf, player)
	}
}

func (slf *Room[PlayerID, Player]) RegPlayerKickedOutEvent(handle game.PlayerKickedOutEventHandle[PlayerID, Player]) {
	slf.playerKickedOutEventHandles = append(slf.playerKickedOutEventHandles, handle)
}

func (slf *Room[PlayerID, Player]) OnPlayerKickedOutEvent(executor, kicked PlayerID, reason string) {
	for _, handle := range slf.playerKickedOutEventHandles {
		handle(slf, executor, kicked, reason)
	}
}
