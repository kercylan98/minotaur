package builtin

import (
	"github.com/kercylan98/minotaur/game"
)

type ActorMove struct {
	game.Actor
	game.Position
	speed float64 // 移动速度
}

func (slf *ActorMove) MoveTo2D(x, y float64) {
	slf.SetXY(x, y)
}

func (slf *ActorMove) MoveBy2D(dx, dy float64) {
	slf.SetXY(slf.GetX()+dx, slf.GetY()+dy)
}

func (slf *ActorMove) MoveTo3D(x, y, z float64) {
	slf.SetXYZ(x, y, z)
}

func (slf *ActorMove) MoveBy3D(dx, dy, dz float64) {
	slf.SetXYZ(slf.GetX()+dx, slf.GetY()+dy, slf.GetZ()+dz)
}

func (slf *ActorMove) MoveToX(x float64) {
	slf.SetX(x)
}

func (slf *ActorMove) MoveByX(dx float64) {
	slf.SetX(slf.GetX() + dx)
}

func (slf *ActorMove) MoveToY(y float64) {
	slf.SetY(y)
}

func (slf *ActorMove) MoveByY(dy float64) {
	slf.SetY(slf.GetY() + dy)
}

func (slf *ActorMove) MoveToZ(z float64) {
	slf.SetZ(z)
}

func (slf *ActorMove) MoveByZ(dz float64) {
	slf.SetZ(slf.GetZ() + dz)
}

func (slf *ActorMove) GetSpeed() float64 {
	return slf.speed
}

func (slf *ActorMove) SetSpeed(speed float64) {
	slf.speed = speed
}
