package ecs

func newSystems(world *world) *systems {
	s := &systems{
		world: world,
	}
	return s
}

type systems struct {
	world     *world
	lifecycle Lifecycle
	systems   []System
}

func (s *systems) setSystems(systems ...System) {
	s.systems = systems
}

func (s *systems) update() {
	if s.lifecycle != OnRunning {
		return
	}
	s.world.worldTime.update()
	for _, system := range s.systems {
		system.OnUpdate(s.world)
	}
}

func (s *systems) onLifecycle(lifecycle Lifecycle) {
	s.lifecycle = lifecycle
	for _, system := range s.systems {
		system.OnLifecycle(s.world, lifecycle)
	}

	switch lifecycle {
	case OnInit:
		s.lifecycle = OnRunning
	default:
	}
}
