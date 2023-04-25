package builtin

import (
	"minotaur/game"
	"minotaur/utils/synchronization"
	"sync/atomic"
)

func NewWorld[PlayerID comparable](guid int64, options ...WorldOption[PlayerID]) *World[PlayerID] {
	world := &World[PlayerID]{
		guid:         guid,
		players:      synchronization.NewMap[PlayerID, game.Player[PlayerID]](),
		playerActors: synchronization.NewMap[PlayerID, *synchronization.Map[int64, game.Actor]](),
		owners:       synchronization.NewMap[int64, PlayerID](),
		actors:       synchronization.NewMap[int64, game.Actor](),
	}
	for _, option := range options {
		option(world)
	}
	return world
}

type World[PlayerID comparable] struct {
	guid         int64
	actorGuid    atomic.Int64
	playerLimit  int
	players      *synchronization.Map[PlayerID, game.Player[PlayerID]]
	playerActors *synchronization.Map[PlayerID, *synchronization.Map[int64, game.Actor]]
	owners       *synchronization.Map[int64, PlayerID]
	actors       *synchronization.Map[int64, game.Actor]

	playerJoinWorldEventHandles   []game.PlayerJoinWorldEventHandle[PlayerID]
	playerLeaveWorldEventHandles  []game.PlayerLeaveWorldEventHandle[PlayerID]
	actorGeneratedEventHandles    []game.ActorGeneratedEventHandle
	actorAnnihilationEventHandles []game.ActorAnnihilationEventHandle
	actorOwnerChangeEventHandles  []game.ActorOwnerChangeEventHandle[PlayerID]

	released atomic.Bool
}

func (slf *World[PlayerID]) GetGuid() int64 {
	return slf.guid
}

func (slf *World[PlayerID]) GetPlayerLimit() int {
	return slf.playerLimit
}

func (slf *World[PlayerID]) GetPlayer(id PlayerID) game.Player[PlayerID] {
	return slf.players.Get(id)
}

func (slf *World[PlayerID]) GetPlayers() map[PlayerID]game.Player[PlayerID] {
	return slf.players.Map()
}

func (slf *World[PlayerID]) GetActor(guid int64) game.Actor {
	return slf.actors.Get(guid)
}

func (slf *World[PlayerID]) GetActors() map[int64]game.Actor {
	return slf.actors.Map()
}

func (slf *World[PlayerID]) GetPlayerActor(id PlayerID, guid int64) game.Actor {
	if actors := slf.playerActors.Get(id); actors != nil {
		return actors.Get(guid)
	}
	return nil
}

func (slf *World[PlayerID]) GetPlayerActors(id PlayerID) map[int64]game.Actor {
	return slf.playerActors.Get(id).Map()
}

func (slf *World[PlayerID]) IsExistPlayer(id PlayerID) bool {
	return slf.players.Exist(id)
}

func (slf *World[PlayerID]) IsExistActor(guid int64) bool {
	return slf.actors.Exist(guid)
}

func (slf *World[PlayerID]) IsOwner(id PlayerID, guid int64) bool {
	actors := slf.playerActors.Get(id)
	if actors != nil {
		return actors.Exist(guid)
	}
	return false
}

func (slf *World[PlayerID]) Join(player game.Player[PlayerID]) error {
	if slf.released.Load() {
		return ErrWorldReleased
	}
	if slf.players.Size() >= slf.playerLimit && slf.playerLimit > 0 {
		return ErrWorldPlayerLimit
	}
	slf.players.Set(player.GetID(), player)
	if actors := slf.playerActors.Get(player.GetID()); actors == nil {
		actors = synchronization.NewMap[int64, game.Actor]()
		slf.playerActors.Set(player.GetID(), actors)
	}
	slf.OnPlayerJoinWorldEvent(player)
	return nil
}

func (slf *World[PlayerID]) Leave(player game.Player[PlayerID]) {
	if !slf.players.Exist(player.GetID()) {
		return
	}
	slf.OnPlayerLeaveWorldEvent(player)
	slf.playerActors.Get(player.GetID()).Range(func(guid int64, actor game.Actor) {
		slf.OnActorAnnihilationEvent(actor)
		slf.owners.Delete(guid)
	})
	slf.playerActors.Delete(player.GetID())
	slf.players.Delete(player.GetID())
}

func (slf *World[PlayerID]) AddActor(actor game.Actor) {
	guid := slf.actorGuid.Add(1)
	actor.SetGuid(guid)
	slf.actors.Set(actor.GetGuid(), actor)
	slf.OnActorGeneratedEvent(actor)
}

func (slf *World[PlayerID]) RemoveActor(guid int64) {
	if actor, exist := slf.actors.GetExist(guid); exist {
		slf.OnActorAnnihilationEvent(actor)
		if id, exist := slf.owners.DeleteGetExist(guid); exist {
			slf.playerActors.Get(id).Delete(guid)
		}
		slf.actors.Delete(guid)
	}
}

func (slf *World[PlayerID]) SetActorOwner(guid int64, id PlayerID) {
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

func (slf *World[PlayerID]) RemoveActorOwner(guid int64) {
	id, exist := slf.owners.GetExist(guid)
	if !exist {
		return
	}
	slf.owners.Delete(guid)
	slf.playerActors.Get(id).Delete(guid)
	slf.OnActorOwnerChangeEvent(slf.GetActor(guid), id, id, true)
}

func (slf *World[PlayerID]) Reset() {
	slf.players.Clear()
	slf.playerActors.Range(func(id PlayerID, actors *synchronization.Map[int64, game.Actor]) {
		actors.Clear()
	})
	slf.playerActors.Clear()
	slf.owners.Clear()
	slf.actors.Clear()
	slf.actorGuid.Store(0)
}

func (slf *World[PlayerID]) Release() {
	if !slf.released.Swap(true) {
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
	}
}

func (slf *World[PlayerID]) RegPlayerJoinWorldEvent(handle game.PlayerJoinWorldEventHandle[PlayerID]) {
	slf.playerJoinWorldEventHandles = append(slf.playerJoinWorldEventHandles, handle)
}

func (slf *World[PlayerID]) OnPlayerJoinWorldEvent(player game.Player[PlayerID]) {
	for _, handle := range slf.playerJoinWorldEventHandles {
		handle(player)
	}
}

func (slf *World[PlayerID]) RegPlayerLeaveWorldEvent(handle game.PlayerLeaveWorldEventHandle[PlayerID]) {
	slf.playerLeaveWorldEventHandles = append(slf.playerLeaveWorldEventHandles, handle)
}

func (slf *World[PlayerID]) OnPlayerLeaveWorldEvent(player game.Player[PlayerID]) {
	for _, handle := range slf.playerLeaveWorldEventHandles {
		handle(player)
	}
}

func (slf *World[PlayerID]) RegActorGeneratedEvent(handle game.ActorGeneratedEventHandle) {
	slf.actorGeneratedEventHandles = append(slf.actorGeneratedEventHandles, handle)
}

func (slf *World[PlayerID]) OnActorGeneratedEvent(actor game.Actor) {
	for _, handle := range slf.actorGeneratedEventHandles {
		handle(actor)
	}
}

func (slf *World[PlayerID]) RegActorAnnihilationEvent(handle game.ActorAnnihilationEventHandle) {
	slf.actorAnnihilationEventHandles = append(slf.actorAnnihilationEventHandles, handle)
}

func (slf *World[PlayerID]) OnActorAnnihilationEvent(actor game.Actor) {
	for _, handle := range slf.actorAnnihilationEventHandles {
		handle(actor)
	}
}

func (slf *World[PlayerID]) RegActorOwnerChangeEvent(handle game.ActorOwnerChangeEventHandle[PlayerID]) {
	slf.actorOwnerChangeEventHandles = append(slf.actorOwnerChangeEventHandles, handle)
}

func (slf *World[PlayerID]) OnActorOwnerChangeEvent(actor game.Actor, old, new PlayerID, isolated bool) {
	for _, handle := range slf.actorOwnerChangeEventHandles {
		handle(actor, old, new, isolated)
	}
}
