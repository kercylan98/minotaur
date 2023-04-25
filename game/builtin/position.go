package builtin

import "minotaur/game"

// NewPosition 创建一个新的Position对象。
func NewPosition(x, y, z float64) *Position {
	return &Position{
		x: x,
		y: y,
		z: z,
	}
}

// Position 是一个具有位置信息的对象。
type Position struct {
	x, y, z                    float64
	positionChangeEventHandles []game.PositionChangeEventHandle
}

// GetX 返回Position对象的X坐标。
func (slf *Position) GetX() float64 {
	return slf.x
}

// GetY 返回Position对象的Y坐标。
func (slf *Position) GetY() float64 {
	return slf.y
}

// GetZ 返回Position对象的Z坐标。
func (slf *Position) GetZ() float64 {
	return slf.z
}

// GetXY 返回Position对象的X和Y坐标。
func (slf *Position) GetXY() (float64, float64) {
	return slf.x, slf.y
}

// GetXYZ 返回Position对象的X、Y和Z坐标。
func (slf *Position) GetXYZ() (float64, float64, float64) {
	return slf.x, slf.y, slf.z
}

// SetX 设置Position对象的X坐标。
func (slf *Position) SetX(x float64) {
	old := slf.Clone()
	defer slf.OnPositionChangeEvent(old, slf)
	slf.x = x
}

// SetY 设置Position对象的Y坐标。
func (slf *Position) SetY(y float64) {
	old := slf.Clone()
	defer slf.OnPositionChangeEvent(old, slf)
	slf.y = y
}

// SetZ 设置Position对象的Z坐标。
func (slf *Position) SetZ(z float64) {
	old := slf.Clone()
	defer slf.OnPositionChangeEvent(old, slf)
	slf.z = z
}

// SetXY 设置Position对象的X和Y坐标。
func (slf *Position) SetXY(x, y float64) {
	old := slf.Clone()
	defer slf.OnPositionChangeEvent(old, slf)
	slf.x = x
	slf.y = y
}

// SetXYZ 设置Position对象的X、Y和Z坐标。
func (slf *Position) SetXYZ(x, y, z float64) {
	old := slf.Clone()
	defer slf.OnPositionChangeEvent(old, slf)
	slf.x = x
	slf.y = y
	slf.z = z
}

func (slf *Position) Clone() game.Position {
	return &Position{
		x: slf.x,
		y: slf.y,
		z: slf.z,
	}
}

func (slf *Position) Compare(position game.Position) bool {
	return slf.x == position.GetX() && slf.y == position.GetY() && slf.z == position.GetZ()
}

func (slf *Position) RegPositionChangeEvent(handle game.PositionChangeEventHandle) {
	slf.positionChangeEventHandles = append(slf.positionChangeEventHandles, handle)
}

func (slf *Position) OnPositionChangeEvent(old, new game.Position) {
	if !old.Compare(new) {
		for _, handle := range slf.positionChangeEventHandles {
			handle(old, new)
		}
	}
}
