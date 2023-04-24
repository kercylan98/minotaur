package builtin

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
	x, y, z float64
}

// GetX 返回Position对象的X坐标。
func (p *Position) GetX() float64 {
	return p.x
}

// GetY 返回Position对象的Y坐标。
func (p *Position) GetY() float64 {
	return p.y
}

// GetZ 返回Position对象的Z坐标。
func (p *Position) GetZ() float64 {
	return p.z
}

// GetXY 返回Position对象的X和Y坐标。
func (p *Position) GetXY() (float64, float64) {
	return p.x, p.y
}

// GetXYZ 返回Position对象的X、Y和Z坐标。
func (p *Position) GetXYZ() (float64, float64, float64) {
	return p.x, p.y, p.z
}

// SetX 设置Position对象的X坐标。
func (p *Position) SetX(x float64) {
	p.x = x
}

// SetY 设置Position对象的Y坐标。
func (p *Position) SetY(y float64) {
	p.y = y
}

// SetZ 设置Position对象的Z坐标。
func (p *Position) SetZ(z float64) {
	p.z = z
}

// SetXY 设置Position对象的X和Y坐标。
func (p *Position) SetXY(x, y float64) {
	p.x = x
	p.y = y
}

// SetXYZ 设置Position对象的X、Y和Z坐标。
func (p *Position) SetXYZ(x, y, z float64) {
	p.x = x
	p.y = y
	p.z = z
}
