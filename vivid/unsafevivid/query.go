package unsafevivid

import (
	"github.com/kercylan98/minotaur/toolkit/collection"
	vivid "github.com/kercylan98/minotaur/vivid/vivids"
	"path"
	"path/filepath"
)

// NewQuery 创建一个 Actor 查询器
func NewQuery(system *ActorSystem, core *ActorCore) *Query {
	return &Query{
		system:    system,
		core:      core,
		resultIds: map[vivid.ActorId]struct{}{},
		lock:      true,
	}
}

// Query 是 Actor 查询器
type Query struct {
	system *ActorSystem
	core   *ActorCore

	results   []*ActorCore               // 查询结果
	resultIds map[vivid.ActorId]struct{} // 去重查询结果 ID

	actions []func()
	lock    bool
}

func (q *Query) ActorId(actorIds ...vivid.ActorId) vivid.Query {
	q.actions = append(q.actions, func() {
		if q.resultIds == nil {
			q.resultIds = make(map[vivid.ActorId]struct{})
		}
		for _, actorId := range actorIds {
			if _, exist := q.resultIds[actorId]; !exist {
				if ref, exist := q.system.actors[actorId]; exist {
					q.results = append(q.results, ref)
				} else {
					q.results = append(q.results, &ActorCore{ActorRef: NewActorRef(q.system, actorId)})
				}
				q.resultIds[actorId] = struct{}{}
			}
		}
	})
	return q

}

func (q *Query) MustActorId(actorIds ...vivid.ActorId) vivid.Query {
	q.actions = append(q.actions, func() {
		if q.resultIds == nil {
			q.resultIds = make(map[vivid.ActorId]struct{})
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

func (q *Query) ActorName(names ...vivid.ActorName) vivid.Query {
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

func (q *Query) ActorPath(actorPath vivid.ActorPath) vivid.Query {
	findPath := path.Clean(path.Join(q.core.Id.Path(), actorPath))
	q.actions = append(q.actions, func() {
		node := q.core
		if path.IsAbs(findPath) {
			node = q.system.user
		}
		for _, name := range filepath.SplitList(filepath.Clean(findPath)) {
			if name == "*" {
				for _, actor := range node.Children {
					q.results = append(q.results, actor)
				}
				break
			}
			if len(name) == 0 {
				continue
			}
			if actor, exist := node.Children[name]; exist {
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
func (q *Query) Many() []vivid.ActorRef {
	return collection.MappingFromSlice(q.internalMany(), func(value *ActorCore) vivid.ActorRef {
		return value
	})
}

// First 获取第一个响应结果
func (q *Query) First() (vivid.ActorRef, error) {
	return q.internalFirst()
}

// One 获取唯一的响应结果
func (q *Query) One() (vivid.ActorRef, error) {
	return q.internalOne()
}

// exec 执行查询
func (q *Query) exec() {
	if q.lock {
		q.system.actorsRW.RLock()
		defer q.system.actorsRW.RUnlock()
	}

	for _, action := range q.actions {
		action()
	}
}

// internalOne 获取唯一的响应结果
func (q *Query) internalOne() (*ActorCore, error) {
	q.exec()
	if len(q.results) == 0 {
		return nil, vivid.ErrActorNotFound
	}
	if len(q.results) > 1 {
		return nil, vivid.ErrActorNotUnique
	}
	return q.results[0], nil
}

// internalMany 获取多个响应结果
func (q *Query) internalMany() []*ActorCore {
	q.exec()
	return q.results
}

// internalFirst 获取第一个响应结果
func (q *Query) internalFirst() (*ActorCore, error) {
	q.exec()
	if len(q.results) == 0 {
		return nil, vivid.ErrActorNotFound
	}
	return q.results[0], nil
}

// withLock 禁用锁
func (q *Query) withLock(isLock bool) *Query {
	q.lock = isLock
	return q
}
