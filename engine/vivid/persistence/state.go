package persistence

func NewState(name Name, configurator ...StateConfigurator) *State {
	state := &State{
		config: newStateConfiguration(),
		name:   name,
	}

	for _, c := range configurator {
		c.Configure(state.config)
	}

	return state
}

// State 是持久化对象的运行时状态，其中包含的记录的事件及快照信息
type State struct {
	config   *StateConfiguration
	events   []Event
	snapshot Snapshot
	name     Name
}

func (s *State) StateChanged(event Event) int {
	s.events = append(s.events, event)
	return s.EventCount()
}

func (s *State) SaveSnapshot(snapshot Snapshot) {
	s.snapshot = snapshot
	s.events = s.events[:0]
}

func (s *State) EventCount() int {
	return len(s.events)
}

func (s *State) Persist() error {
	if s.snapshot == nil && len(s.events) == 0 {
		return nil
	}
	return s.config.storage.Save(s.name, s.snapshot, s.events)
}
func (s *State) Load() (snapshot Snapshot, events []Event, err error) {
	return s.config.storage.Load(s.name)
}

func (s *State) Clear() error {
	return s.config.storage.Clear(s.name)
}
