package vivid

import (
	"github.com/kercylan98/minotaur/toolkit/collection"
	"path"
	"path/filepath"
)

// newQuery 创建一个 Actor 查询器
func newQuery(system *ActorSystem, core *actorCore) *query {
	return &query{
		system:    system,
		core:      core,
		resultIds: map[ActorId]struct{}{},
	}
}

type Query interface {
	// ActorId 使用特定的 ActorId，该查询可能包含不存在的 ActorRef
	ActorId(actorIds ...ActorId) *query

	// MustActorId 使用特定的 ActorId，该查询包含必须存在的 ActorRef
	MustActorId(actorIds ...ActorId) *query

	// ActorName 通过 ActorName 进行查询匹配
	ActorName(names ...ActorName) *query

	// ActorPath 通过 ActorPath 进行查询匹配，支持格式如下：
	// - "./user/actor1"
	// - "/user/actor1"
	// - "user/actor1"
	// - "actor1"
	// - "actor1/actor2/*"
	ActorPath(actorPath ActorPath) *query

	// Many 获取多个响应结果
	Many() []ActorRef

	// First 获取第一个响应结果
	First() (ActorRef, error)

	// One 获取唯一的响应结果
	One() (ActorRef, error)
}

// query 是 Actor 查询器
type query struct {
	system *ActorSystem
	core   *actorCore

	results   []*actorCore         // 查询结果
	resultIds map[ActorId]struct{} // 去重查询结果 ID

	actions []func()
}

func (q *query) ActorId(actorIds ...ActorId) *query {
	q.actions = append(q.actions, func() {
		if q.resultIds == nil {
			q.resultIds = make(map[ActorId]struct{})
		}
		for _, actorId := range actorIds {
			if _, exist := q.resultIds[actorId]; !exist {
				if ref, exist := q.system.actors[actorId]; exist {
					q.results = append(q.results, ref)
				} else {
					q.results = append(q.results, &actorCore{ActorRef: newLocalActorRef(q.system, actorId)})
				}
				q.resultIds[actorId] = struct{}{}
			}
		}
	})
	return q

}

func (q *query) MustActorId(actorIds ...ActorId) *query {
	q.actions = append(q.actions, func() {
		if q.resultIds == nil {
			q.resultIds = make(map[ActorId]struct{})
		}
		for _, actorId := range actorIds {
			if _, exist := q.resultIds[actorId]; !exist {
				if ref, exist := q.system.actors[actorId]; exist {
					q.results = append(q.results, ref)
					q.resultIds[actorId] = struct{}{}
				}
			}
		}
	})
	return q
}

func (q *query) ActorName(names ...ActorName) *query {
	q.actions = append(q.actions, func() {
		for _, actor := range q.system.actors {
			for _, actorName := range names {
				if actor.GetId().Name() == actorName {
					q.results = append(q.results, actor)
				}
			}
		}
	})
	return q
}

func (q *query) ActorPath(actorPath ActorPath) *query {
	findPath := path.Clean(path.Join(q.core.id.Path(), actorPath))
	q.actions = append(q.actions, func() {
		node := q.core
		if path.IsAbs(findPath) {
			node = q.system.user
		}
		for _, name := range filepath.SplitList(filepath.Clean(findPath)) {
			if name == "*" {
				for _, actor := range node.children {
					q.results = append(q.results, actor)
				}
				break
			}
			if len(name) == 0 {
				continue
			}
			if actor, exist := node.children[name]; exist {
				node = actor
			} else {
				return
			}
		}
		q.results = append(q.results, node)
	})
	return q
}

// Many 获取多个响应结果
func (q *query) Many() []ActorRef {
	return collection.MappingFromSlice(q.internalMany(), func(value *actorCore) ActorRef {
		return value
	})
}

// First 获取第一个响应结果
func (q *query) First() (ActorRef, error) {
	return q.internalFirst()
}

// One 获取唯一的响应结果
func (q *query) One() (ActorRef, error) {
	return q.internalOne()
}

// exec 执行查询
func (q *query) exec() {
	q.system.actorsRW.RLock()
	defer q.system.actorsRW.RUnlock()

	for _, action := range q.actions {
		action()
	}
}

// internalOne 获取唯一的响应结果
func (q *query) internalOne() (*actorCore, error) {
	q.exec()
	if len(q.results) == 0 {
		return nil, ErrActorNotFound
	}
	if len(q.results) > 1 {
		return nil, ErrActorNotUnique
	}
	return q.results[0], nil
}

// internalMany 获取多个响应结果
func (q *query) internalMany() []*actorCore {
	q.exec()
	return q.results
}

// internalFirst 获取第一个响应结果
func (q *query) internalFirst() (*actorCore, error) {
	q.exec()
	if len(q.results) == 0 {
		return nil, ErrActorNotFound
	}
	return q.results[0], nil
}
