package ecs

func NewWorld() World {
	w := &world{
		worldTime:  newWorldTime(),
		entities:   newEntities(32),
		components: newComponents(),
	}
	w.archetypes = newArchetypes(w)
	return w
}

type World interface {
	Query(query Query) *Result

	RegComponent(component any) ComponentId

	Spawn(componentIds ...ComponentId) Entity

	Spawns(count int, componentIds ...ComponentId) []Entity

	Annihilate(entity Entity)

	Annihilates(entities []Entity)

	//Update()
	//
	//SetSleep(d time.Duration)
	//
	//SetTimeScale(scale float64)
	//
	//TimeScale() float64
	//
	//DeltaTime() time.Duration
	//
	//Pause()
	//
	//Resume()
}

type world struct {
	*worldTime
	archetypes *archetypes
	entities   *entities
	components *components
}

func (w *world) Query(query Query) *Result {
	var result = &Result{world: w}

	for _, arch := range w.archetypes.arts {
		if !query.Evaluate(arch.mask) {
			continue
		}

		result.archetypes = append(result.archetypes, arch)
	}

	result.expansion()

	return result
}

func (w *world) Update() {
	w.worldTime.update()
}

func (w *world) RegComponent(component any) ComponentId {
	return w.components.reg(component)
}

func (w *world) Spawn(componentIds ...ComponentId) Entity {
	arch := w.archetypes.get(componentIds...)
	entity := w.entities.get()
	arch.bind(entity)
	return entity
}

func (w *world) Spawns(count int, componentIds ...ComponentId) []Entity {
	arch := w.archetypes.get(componentIds...)
	entities := w.entities.getMany(count)
	arch.bindMany(entities)
	return entities
}

func (w *world) Annihilate(entity Entity) {
	w.entities.recycle(entity)

	w.archetypes.unbind(entity)
}

func (w *world) Annihilates(entities []Entity) {
	for _, entity := range entities {
		w.entities.recycle(entity)
	}

	w.archetypes.unBindMany(entities)
}
