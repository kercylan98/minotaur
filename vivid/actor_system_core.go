package vivid

type ActorSystemCore interface {
	// ModifyGlobalMessageCounter 修改全局消息计数器
	ModifyGlobalMessageCounter(delta int64)
}

type _ActorSystemCore struct {
	system *ActorSystem
}

func (s *_ActorSystemCore) init(system *ActorSystem) *_ActorSystemCore {
	s.system = system
	return s
}

func (s *_ActorSystemCore) ModifyGlobalMessageCounter(delta int64) {
	s.system.waitGroup.Add(delta)
}
