package geometry

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/generic"
	"github.com/kercylan98/minotaur/utils/random"
	"math"
)

// NewSimpleCircle 通过传入圆的半径和圆心位置，生成一个圆
func NewSimpleCircle[V generic.SignedNumber](radius V, centroid Point[V]) SimpleCircle[V] {
	if radius <= 0 {
		panic(fmt.Errorf("radius must be greater than 0, but got %v", radius))
	}
	return SimpleCircle[V]{
		centroid: centroid,
		radius:   radius,
	}
}

// SimpleCircle 仅由位置和半径组成的圆形数据结构
type SimpleCircle[V generic.SignedNumber] struct {
	centroid Point[V] // 圆心位置
	radius   V        // 半径
}

// String 获取圆形的字符串表示
func (sc SimpleCircle[V]) String() string {
	return fmt.Sprintf("SimpleCircle{centroid: %v, %v, radius: %v}", sc.centroid.GetX(), sc.centroid.GetY(), sc.radius)
}

// Centroid 获取圆形质心位置
func (sc SimpleCircle[V]) Centroid() Point[V] {
	return sc.centroid
}

// CentroidX 获取圆形质心位置的 X 坐标
func (sc SimpleCircle[V]) CentroidX() V {
	return sc.centroid.GetX()
}

// CentroidY 获取圆形质心位置的 Y 坐标
func (sc SimpleCircle[V]) CentroidY() V {
	return sc.centroid.GetY()
}

// CentroidXY 获取圆形质心位置的 X、Y 坐标
func (sc SimpleCircle[V]) CentroidXY() (V, V) {
	return sc.centroid.GetXY()
}

// PointIsIn 检查特定点是否位于圆内
func (sc SimpleCircle[V]) PointIsIn(pos Point[V]) bool {
	return V(pos.Distance(sc.centroid)) <= sc.radius
}

// CentroidDistance 计算与另一个圆的质心距离
func (sc SimpleCircle[V]) CentroidDistance(circle SimpleCircle[V]) float64 {
	return sc.centroid.Distance(circle.centroid)
}

// Radius 获取圆形半径
func (sc SimpleCircle[V]) Radius() V {
	return sc.radius
}

// ZoomRadius 获取缩放后的半径
func (sc SimpleCircle[V]) ZoomRadius(zoom float64) V {
	return V(float64(sc.radius) * zoom)
}

// Area 获取圆形面积
func (sc SimpleCircle[V]) Area() V {
	return sc.radius * sc.radius
}

// Projection 获取圆形投影到另一个圆形的特定比例下的位置和半径
func (sc SimpleCircle[V]) Projection(circle SimpleCircle[V], ratio float64) SimpleCircle[V] {
	// 计算圆心朝目标按比例移动后的位置
	distance := float64(sc.Centroid().Distance(circle.centroid))
	moveDistance := distance * ratio
	newX := float64(sc.CentroidX()) + moveDistance*(float64(circle.CentroidX())-float64(sc.CentroidX()))/distance
	newY := float64(sc.CentroidY()) + moveDistance*(float64(circle.CentroidY())-float64(sc.CentroidY()))/distance

	return NewSimpleCircle(V(float64(sc.radius)*ratio), NewPoint(V(newX), V(newY)))
}

// Length 获取圆的周长
func (sc SimpleCircle[V]) Length() V {
	return 2 * sc.radius
}

// Overlap 与另一个圆是否发生重叠
func (sc SimpleCircle[V]) Overlap(circle SimpleCircle[V]) bool {
	return sc.centroid.Distance(circle.centroid) < float64(sc.radius+circle.radius)
}

// RandomPoint 获取圆内随机点
func (sc SimpleCircle[V]) RandomPoint() Point[V] {
	rx := V(random.Float64() * float64(sc.radius))
	ry := V(random.Float64() * float64(sc.radius))
	if random.Bool() {
		rx = -rx
	}
	if random.Bool() {
		ry = -ry
	}
	return sc.centroid.GetOffset(rx, ry)
}

// RandomPointWithinCircle 获取圆内随机点，且该圆在 radius 小于父圆时不会超出父圆
func (sc SimpleCircle[V]) RandomPointWithinCircle(radius V) Point[V] {
	// 生成随机角度
	angle := random.Float64() * 2 * math.Pi

	// 限制坐标随机范围
	var rx, ry float64
	if radius < sc.radius {
		r := float64(sc.radius - radius)
		rx = random.Float64() * r
		ry = random.Float64() * r
	} else {
		r := float64(sc.radius)
		rx = random.Float64() * r
		ry = random.Float64() * r
	}

	// 生成随机点
	return sc.centroid.GetOffset(V(rx*math.Cos(angle)), V(ry*math.Sin(angle)))
}

// RandomPointWithinRadius 获取圆内随机点，且该点与圆心的距离小于等于指定半径
func (sc SimpleCircle[V]) RandomPointWithinRadius(radius V) Point[V] {
	if radius > sc.radius {
		panic("radius must be less than or equal to the circle radius")
	}
	rx := V(random.Float64() * float64(radius))
	ry := V(random.Float64() * float64(radius))
	if random.Bool() {
		rx = -rx
	}
	if random.Bool() {
		ry = -ry
	}
	return sc.centroid.GetOffset(rx, ry)
}

// RandomPointWithinRadiusAndSector 获取圆内随机点，且距离圆心小于等于指定半径，且角度在指定范围内
//   - startAngle: 起始角度，取值范围为 0 到 360 度
//   - endAngle: 结束角度，取值范围为 0 到 360 度，且大于起始角度
func (sc SimpleCircle[V]) RandomPointWithinRadiusAndSector(radius, startAngle, endAngle V) Point[V] {
	var full = 360
	if radius > sc.radius {
		panic("radius must be less than or equal to the circle radius")
	}
	if startAngle < 0 || startAngle > V(full) {
		panic("startAngle must be in the range 0 to 360 degrees")
	}
	if endAngle < 0 || endAngle > V(full) {
		panic("endAngle must be in the range 0 to 360 degrees")
	}
	if startAngle > endAngle {
		panic("startAngle must be less than or equal to endAngle")
	}
	angle := V(random.Float64() * float64(endAngle-startAngle))
	return sc.centroid.GetOffset(radius*V(math.Cos(float64(angle))), radius*V(math.Sin(float64(angle))))
}

// RandomCircleWithinParent 根据指定半径，生成一个圆内随机子圆，该圆不会超出父圆
func (sc SimpleCircle[V]) RandomCircleWithinParent(radius V) SimpleCircle[V] {
	if radius > sc.radius {
		panic("radius must be less than or equal to the circle radius")
	}
	return NewSimpleCircle(radius, sc.RandomPointWithinCircle(radius))
}
