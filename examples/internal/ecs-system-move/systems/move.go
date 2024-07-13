package systems

import (
	"github.com/kercylan98/minotaur/core/ecs"
	cmps2 "github.com/kercylan98/minotaur/examples/internal/ecs-system-move/cmps"
	"math"
)

type Move struct {
	positionId   ecs.ComponentId
	velocityId   ecs.ComponentId
	collider2DId ecs.ComponentId

	queryWithCollider    ecs.Query
	queryWithoutCollider ecs.Query
}

func (m *Move) OnLifecycle(world ecs.World, lifecycle ecs.Lifecycle) {
	switch lifecycle {
	case ecs.OnInit:
		m.positionId = world.RegComponent(new(cmps2.Position2d))
		m.velocityId = world.RegComponent(new(cmps2.Velocity))
		m.collider2DId = world.RegComponent(new(cmps2.Collider2D))

		m.queryWithCollider = ecs.In(m.positionId, m.velocityId, m.collider2DId)
		m.queryWithoutCollider = ecs.And(ecs.In(m.positionId, m.velocityId), ecs.NotIn(m.collider2DId))
	default:
	}
}

func (m *Move) OnUpdate(world ecs.World) {
	// 遍历具有 Position2d、Velocity 和 Collider2D 组件的实体
	resultWithCollider := world.Query(m.queryWithCollider)
	iterWithCollider := resultWithCollider.Iterator()
	iterWithColliderCheck := resultWithCollider.Iterator()
	for iterWithCollider.Next() {
		position := iterWithCollider.Get(m.positionId).(*cmps2.Position2d)
		velocity := iterWithCollider.Get(m.velocityId).(*cmps2.Velocity)
		collider := iterWithCollider.Get(m.collider2DId).(*cmps2.Collider2D)

		// 更新位置
		newPos := cmps2.Vector2{
			X: position.X + velocity.X*float64(world.DeltaTime()),
			Y: position.Y + velocity.Y*float64(world.DeltaTime()),
		}

		// 碰撞检测
		collisionDetected := false
		for iterWithColliderCheck.Reset(); iterWithColliderCheck.Next(); {
			otherEntity := iterWithCollider.Entity()
			if otherEntity == iterWithCollider.Entity() {
				continue
			}

			otherPosition := iterWithCollider.Get(m.positionId).(*cmps2.Position2d)
			otherCollider := iterWithCollider.Get(m.collider2DId).(*cmps2.Collider2D)

			if checkCollision(newPos, *position, *collider, *otherPosition, *otherCollider) {
				// 处理碰撞（简单示例：停止运动）
				velocity.X = 0
				velocity.Y = 0
				collisionDetected = true
				break
			}
		}

		// 如果没有碰撞，更新位置
		if !collisionDetected {
			position.X = newPos.X
			position.Y = newPos.Y
		}
	}

	// 遍历具有 Position2d 和 Velocity 但不包含 Collider2D 组件的实体
	resultWithoutCollider := world.Query(m.queryWithoutCollider)
	iterWithoutCollider := resultWithoutCollider.Iterator()
	for iterWithoutCollider.Next() {
		position := iterWithoutCollider.Get(m.positionId).(*cmps2.Position2d)
		velocity := iterWithoutCollider.Get(m.velocityId).(*cmps2.Velocity)

		// 更新位置
		position.X += velocity.X * float64(world.DeltaTime())
		position.Y += velocity.Y * float64(world.DeltaTime())
	}
}

func checkCollision(newPos, oldPos cmps2.Vector2, collider cmps2.Collider2D, otherPos cmps2.Vector2, otherCollider cmps2.Collider2D) bool {
	// 简单圆形碰撞检测
	dx := newPos.X - otherPos.X
	dy := newPos.Y - otherPos.Y
	distance := math.Sqrt(dx*dx + dy*dy)
	return distance < (collider.Radius + otherCollider.Radius)
}
