package builtin

import (
	"github.com/kercylan98/minotaur/game"
	"github.com/kercylan98/minotaur/utils/hash"
	"github.com/kercylan98/minotaur/utils/log"
	"github.com/kercylan98/minotaur/utils/synchronization"
	"go.uber.org/zap"
	"sync/atomic"
)

// NewWorld 创建一个内置的游戏世界
func NewWorld[PlayerID comparable, Player game.Player[PlayerID]](guid int64, options ...WorldOption[PlayerID, Player]) *World[PlayerID, Player] {
	world := &World[PlayerID, Player]{
		guid:         guid,
		players:      synchronization.NewMap[PlayerID, Player](),
		playerActors: synchronization.NewMap[PlayerID, hash.Map[int64, game.Actor]](),
		owners:       synchronization.NewMap[int64, PlayerID](),
		actors:       synchronization.NewMap[int64, game.Actor](),
	}
	for _, option := range options {
		option(world)
	}
	return world
}

// World 游戏世界的内置实现，实现了基本的游戏世界接口
type World[PlayerID comparable, Player game.Player[PlayerID]] struct {
	guid         int64
	actorGuid    atomic.Int64
	playerLimit  int
	players      hash.Map[PlayerID, Player]
	playerActors hash.Map[PlayerID, hash.Map[int64, game.Actor]]
	owners       hash.Map[int64, PlayerID]
	actors       hash.Map[int64, game.Actor]

	playerJoinWorldEventHandles   []game.PlayerJoinWorldEventHandle[PlayerID, Player]
	playerLeaveWorldEventHandles  []game.PlayerLeaveWorldEventHandle[PlayerID, Player]
	actorGeneratedEventHandles    []game.ActorGeneratedEventHandle[PlayerID, Player]
	actorAnnihilationEventHandles []game.ActorAnnihilationEventHandle[PlayerID, Player]
	actorOwnerChangeEventHandles  []game.ActorOwnerChangeEventHandle[PlayerID, Player]
	worldResetEventHandles        []game.WorldResetEventHandle[PlayerID, Player]
	worldReleasedEventHandles     []game.WorldReleaseEventHandle[PlayerID, Player]

	released atomic.Bool
}

func (slf *World[PlayerID, Player]) GetGuid() int64 {
	return slf.guid
}

func (slf *World[PlayerID, Player]) GetPlayerLimit() int {
	return slf.playerLimit
}

func (slf *World[PlayerID, Player]) GetPlayer(id PlayerID) Player {
	return slf.players.Get(id)
}

func (slf *World[PlayerID, Player]) GetPlayers() hash.MapReadonly[PlayerID, Player] {
	return slf.players.(hash.MapReadonly[PlayerID, Player])
}

func (slf *World[PlayerID, Player]) GetActor(guid int64) game.Actor {
	return slf.actors.Get(guid)
}

func (slf *World[PlayerID, Player]) GetActors() hash.MapReadonly[int64, game.Actor] {
	return slf.actors.(hash.MapReadonly[int64, game.Actor])
}

func (slf *World[PlayerID, Player]) GetPlayerActor(id PlayerID, guid int64) game.Actor {
	if actors := slf.playerActors.Get(id); actors != nil {
		return actors.Get(guid)
	}
	return nil
}

func (slf *World[PlayerID, Player]) GetPlayerActors(id PlayerID) hash.MapReadonly[int64, game.Actor] {
	return slf.playerActors.Get(id).(hash.MapReadonly[int64, game.Actor])
}

func (slf *World[PlayerID, Player]) IsExistPlayer(id PlayerID) bool {
	return slf.players.Exist(id)
}

func (slf *World[PlayerID, Player]) IsExistActor(guid int64) bool {
	return slf.actors.Exist(guid)
}

func (slf *World[PlayerID, Player]) IsOwner(id PlayerID, guid int64) bool {
	actors := slf.playerActors.Get(id)
	if actors != nil {
		return actors.Exist(guid)
	}
	return false
}

func (slf *World[PlayerID, Player]) Join(player Player) error {
	if slf.released.Load() {
		return ErrWorldReleased
	}
	if slf.players.Size() >= slf.playerLimit && slf.playerLimit > 0 {
		return ErrWorldPlayerLimit
	}
	log.Debug("World.Join", zap.Int64("guid", slf.GetGuid()), zap.Any("player", player.GetID()))
	slf.players.Set(player.GetID(), player)
	if actors := slf.playerActors.Get(player.GetID()); actors == nil {
		actors = synchronization.NewMap[int64, game.Actor]()
		slf.playerActors.Set(player.GetID(), actors)
	}
	slf.OnPlayerJoinWorldEvent(player)
	return nil
}

func (slf *World[PlayerID, Player]) Leave(id PlayerID) {
	player, exist := slf.players.GetExist(id)
	if !exist {
		return
	}
	log.Debug("World.Leave", zap.Int64("guid", slf.GetGuid()), zap.Any("player", player.GetID()))
	slf.OnPlayerLeaveWorldEvent(player)
	slf.playerActors.Get(player.GetID()).Range(func(guid int64, actor game.Actor) {
		slf.OnActorAnnihilationEvent(actor)
		slf.owners.Delete(guid)
	})
	slf.playerActors.Delete(player.GetID())
	slf.players.Delete(player.GetID())
}

func (slf *World[PlayerID, Player]) AddActor(actor game.Actor) {
	guid := slf.actorGuid.Add(1)
	actor.SetGuid(guid)
	slf.actors.Set(actor.GetGuid(), actor)
	slf.OnActorGeneratedEvent(actor)
}

func (slf *World[PlayerID, Player]) RemoveActor(guid int64) {
	if actor, exist := slf.actors.GetExist(guid); exist {
		slf.OnActorAnnihilationEvent(actor)
		if id, exist := slf.owners.DeleteGetExist(guid); exist {
			slf.playerActors.Get(id).Delete(guid)
		}
		slf.actors.Delete(guid)
	}
}

func (slf *World[PlayerID, Player]) SetActorOwner(guid int64, id PlayerID) {
	oldId, exist := slf.owners.GetExist(guid)
	if exist && oldId == id {
		return
	}
	actor := slf.GetActor(guid)
	if actor == nil {
		return
	}
	slf.owners.Set(guid, id)
	slf.playerActors.Get(id).Set(guid, actor)
	slf.OnActorOwnerChangeEvent(actor, oldId, id, false)
}

func (slf *World[PlayerID, Player]) RemoveActorOwner(guid int64) {
	id, exist := slf.owners.GetExist(guid)
	if !exist {
		return
	}
	slf.owners.Delete(guid)
	slf.playerActors.Get(id).Delete(guid)
	slf.OnActorOwnerChangeEvent(slf.GetActor(guid), id, id, true)
}

func (slf *World[PlayerID, Player]) Reset() {
	log.Debug("World", zap.Int64("Reset", slf.guid))
	slf.players.Clear()
	slf.playerActors.Range(func(id PlayerID, actors hash.Map[int64, game.Actor]) {
		actors.Clear()
	})
	slf.playerActors.Clear()
	slf.owners.Clear()
	slf.actors.Clear()
	slf.actorGuid.Store(0)
	slf.OnWorldResetEvent()
}

func (slf *World[PlayerID, Player]) Release() {
	if !slf.released.Swap(true) {
		log.Debug("World", zap.Int64("Release", slf.guid))
		slf.OnWorldReleaseEvent()
		slf.Reset()
		slf.players = nil
		slf.playerActors = nil
		slf.owners = nil
		slf.actors = nil

		slf.playerJoinWorldEventHandles = nil
		slf.playerLeaveWorldEventHandles = nil
		slf.actorGeneratedEventHandles = nil
		slf.actorAnnihilationEventHandles = nil
		slf.actorOwnerChangeEventHandles = nil
		slf.worldResetEventHandles = nil
		slf.worldReleasedEventHandles = nil
	}
}

func (slf *World[PlayerID, Player]) RegWorldResetEvent(handle game.WorldResetEventHandle[PlayerID, Player]) {
	slf.worldResetEventHandles = append(slf.worldResetEventHandles, handle)
}

func (slf *World[PlayerID, Player]) OnWorldResetEvent() {
	for _, handle := range slf.worldResetEventHandles {
		handle(slf)
	}
}

func (slf *World[PlayerID, Player]) RegWorldReleaseEvent(handle game.WorldReleaseEventHandle[PlayerID, Player]) {
	slf.worldReleasedEventHandles = append(slf.worldReleasedEventHandles, handle)
}

func (slf *World[PlayerID, Player]) OnWorldReleaseEvent() {
	for _, handle := range slf.worldReleasedEventHandles {
		handle(slf)
	}
}

func (slf *World[PlayerID, Player]) RegPlayerJoinWorldEvent(handle game.PlayerJoinWorldEventHandle[PlayerID, Player]) {
	slf.playerJoinWorldEventHandles = append(slf.playerJoinWorldEventHandles, handle)
}

func (slf *World[PlayerID, Player]) OnPlayerJoinWorldEvent(player Player) {
	for _, handle := range slf.playerJoinWorldEventHandles {
		handle(slf, player)
	}
}

func (slf *World[PlayerID, Player]) RegPlayerLeaveWorldEvent(handle game.PlayerLeaveWorldEventHandle[PlayerID, Player]) {
	slf.playerLeaveWorldEventHandles = append(slf.playerLeaveWorldEventHandles, handle)
}

func (slf *World[PlayerID, Player]) OnPlayerLeaveWorldEvent(player Player) {
	for _, handle := range slf.playerLeaveWorldEventHandles {
		handle(slf, player)
	}
}

func (slf *World[PlayerID, Player]) RegActorGeneratedEvent(handle game.ActorGeneratedEventHandle[PlayerID, Player]) {
	slf.actorGeneratedEventHandles = append(slf.actorGeneratedEventHandles, handle)
}

func (slf *World[PlayerID, Player]) OnActorGeneratedEvent(actor game.Actor) {
	for _, handle := range slf.actorGeneratedEventHandles {
		handle(slf, actor)
	}
}

func (slf *World[PlayerID, Player]) RegActorAnnihilationEvent(handle game.ActorAnnihilationEventHandle[PlayerID, Player]) {
	slf.actorAnnihilationEventHandles = append(slf.actorAnnihilationEventHandles, handle)
}

func (slf *World[PlayerID, Player]) OnActorAnnihilationEvent(actor game.Actor) {
	for _, handle := range slf.actorAnnihilationEventHandles {
		handle(slf, actor)
	}
}

func (slf *World[PlayerID, Player]) RegActorOwnerChangeEvent(handle game.ActorOwnerChangeEventHandle[PlayerID, Player]) {
	slf.actorOwnerChangeEventHandles = append(slf.actorOwnerChangeEventHandles, handle)
}

func (slf *World[PlayerID, Player]) OnActorOwnerChangeEvent(actor game.Actor, old, new PlayerID, isolated bool) {
	for _, handle := range slf.actorOwnerChangeEventHandles {
		handle(slf, actor, old, new, isolated)
	}
}
