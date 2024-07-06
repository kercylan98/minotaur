package ecs

func newEntityManager() *entityManager {
	return &entityManager{}
}

type entityManager struct {
	entities []EntityId
}

func (em *entityManager) createEntity() EntityId {
	entityId := newEntityId(uint32(len(em.entities))+1, 0)
	em.entities = append(em.entities, entityId)
	return entityId
}
