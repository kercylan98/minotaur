package ecs

import (
	"github.com/kercylan98/minotaur/toolkit/mask"
)

type graphNode struct {
	mask         mask.DynamicMask
	componentIds []ComponentId

	children map[ComponentId]*graphNode
	parents  map[ComponentId]*graphNode
}

func (g *graphNode) AddArchetype(componentIds ...ComponentId) {
	var recursive func(node *graphNode, componentIds []ComponentId)
	recursive = func(node *graphNode, componentIds []ComponentId) {
		if len(componentIds) == 0 {
			return
		}

		if node.children == nil {
			node.children = make(map[ComponentId]*graphNode)
			node.parents = make(map[ComponentId]*graphNode)
		}

		curr := node.mask
		for i := 0; i < len(componentIds); i++ {
			node.mask.Set(componentIds[i])
			node.componentIds = append(node.componentIds, componentIds[i])

			if child, exist := node.children[componentIds[i]]; exist {
				// 忽略当前节点，但是要包含前后节点
				recursive(child, append(componentIds[:i], componentIds[i+1:]...))
			} else {
				child = &graphNode{
					mask:     curr,
					children: map[ComponentId]*graphNode{},
					parents:  map[ComponentId]*graphNode{},
				}
				node.children[componentIds[i]] = child
				child.parents = map[ComponentId]*graphNode{
					componentIds[i]: node,
				}
				recursive(child, append(componentIds[:i], componentIds[i+1:]...))
			}
		}
	}

	recursive(g, componentIds)
}
