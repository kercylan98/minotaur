package ecs

import (
	"bytes"
	"fmt"
	"github.com/kercylan98/minotaur/core/vivid"
	"github.com/kercylan98/minotaur/toolkit/log"
	"os"
	"os/exec"
)

func NewWorld() *World {
	w := &World{
		entities:        newEntities(256),
		archetypes:      map[DynamicBitSetKey]*archetype{},
		entityArchetype: make(map[EntityId]*archetype),
		componentIds:    make(map[ComponentId]*componentInfo),
		componentTypes:  make(map[ComponentType]*componentInfo),
		logger: func() *log.Logger {
			return log.GetDefault()
		},
	}
	w.root = newRootArchetype(w)
	return w
}

type World struct {
	root            *archetype
	archetypes      map[DynamicBitSetKey]*archetype
	entities        *entities
	entityArchetype map[EntityId]*archetype
	componentIds    map[ComponentId]*componentInfo
	componentTypes  map[ComponentType]*componentInfo
	logger          vivid.LoggerProvider
}

func (w *World) Query(query QueryCondition) *Result {
	var result = &Result{
		archetypes: make(map[DynamicBitSetKey]*archetype),
	}

	for _, arch := range w.archetypes {
		if !query.Evaluate(arch.mask) {
			continue
		}

		result.archetypes[arch.mask.Key()] = arch
	}

	result.expansion()
	return result
}

func (w *World) RegisterComponent(componentType ComponentType) ComponentId {
	info, exists := w.componentTypes[componentType]
	if exists {
		return info.id
	}
	info = newComponentInfo(ComponentId(len(w.componentTypes)+1), componentType)

	w.componentIds[info.id] = info
	w.componentTypes[info.typ] = info

	w.logger().Debug("ECS", log.String("info", "register component"), log.Uint32("id", info.id), log.String("type", info.typ.String()))

	return info.id
}

func (w *World) Spawn(componentIds ...ComponentId) EntityId {
	art := w.root.mutation(componentIds, nil)

	eid := w.entities.get()
	art.addEntity(eid, nil)
	w.entityArchetype[eid] = art

	return eid
}

func (w *World) Annihilate(id EntityId) {
	w.entities.recycle(id)
}

func (w *World) AddComponent(entityId EntityId, componentIds ...ComponentId) {
	curr := w.entityArchetype[entityId]
	if curr == nil {
		panic("entity not found")
	}

	art := curr.mutation(componentIds, nil)
	curr.migrate(art, entityId)
}

func (w *World) DelComponent(entityId EntityId, componentIds ...ComponentId) {
	curr := w.entityArchetype[entityId]
	if curr == nil {
		panic("entity not found")
	}

	art := curr.mutation(nil, componentIds)
	curr.migrate(art, entityId)
}

func (w *World) getEntityComponentData(id EntityId, cid ComponentId) any {
	if !w.entities.alive(id) {
		return nil
	}

	at := w.entityArchetype[id]
	if at == nil {
		return nil
	}

	return at.getEntityComponentData(id, cid)
}

func (w *World) getComponentInfoById(id ComponentId) *componentInfo {
	info, exists := w.componentIds[id]
	if !exists {
		return nil
	}
	return info
}

func (w *World) GenerateDotFile(path string) {
	var buffer bytes.Buffer
	buffer.WriteString("digraph G {\n")

	for _, art := range w.archetypes {
		currKey := fmt.Sprint(parseDynamicBitSetKey(art.mask.Key()).Bits())
		for key := range art.delEdges {
			buffer.WriteString(fmt.Sprintf("  \"%v\" -> \"%v\" [label=\"del\"];\n", currKey, parseDynamicBitSetKey(key).Bits()))
		}
		for key := range art.addEdges {
			buffer.WriteString(fmt.Sprintf("  \"%v\" -> \"%v\" [label=\"add\"];\n", currKey, parseDynamicBitSetKey(key).Bits()))
		}
	}

	buffer.WriteString("}\n")

	dotFileName := path
	err := os.WriteFile(dotFileName, buffer.Bytes(), 0644)
	if err != nil {
		fmt.Println("Error writing dot file:", err)
		return
	}

	// 生成图形文件
	outFileName := dotFileName + ".png"
	cmd := exec.Command("dot", "-Tpng", dotFileName, "-o", outFileName)
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error generating graph:", err)
		return
	}

	fmt.Println("Graph generated and saved as", outFileName)
}
