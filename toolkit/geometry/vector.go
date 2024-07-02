package geometry

import (
	"github.com/kercylan98/minotaur/toolkit/constraints"
	"github.com/kercylan98/minotaur/toolkit/ident"
	"math"
)

var (
	vector2Zero = NewVector(0, 0)
	vector3Zero = NewVector(0, 0, 0)
	vector4Zero = NewVector(0, 0, 0, 0)
	vector5Zero = NewVector(0, 0, 0, 0, 0)
	vector6Zero = NewVector(0, 0, 0, 0, 0, 0)
	vector7Zero = NewVector(0, 0, 0, 0, 0, 0, 0)
	vector8Zero = NewVector(0, 0, 0, 0, 0, 0, 0, 0)
	vector9Zero = NewVector(0, 0, 0, 0, 0, 0, 0, 0, 0)
)

type Vector []float64
type Vector2 = Vector
type Vector3 = Vector
type Vector4 = Vector
type Vector5 = Vector
type Vector6 = Vector
type Vector7 = Vector
type Vector8 = Vector
type Vector9 = Vector

func Vector2Zero() Vector2 {
	return vector2Zero
}

func Vector3Zero() Vector3 {
	return vector3Zero
}

func Vector4Zero() Vector4 {
	return vector4Zero
}

func Vector5Zero() Vector5 {
	return vector5Zero
}

func Vector6Zero() Vector6 {
	return vector6Zero
}

func Vector7Zero() Vector7 {
	return vector7Zero
}

func Vector8Zero() Vector8 {
	return vector8Zero
}

func Vector9Zero() Vector9 {
	return vector9Zero
}

// NewVector 创建一个向量
func NewVector[T constraints.Number](v ...T) Vector {
	vec := make(Vector, len(v))
	for i, val := range v {
		vec[i] = float64(val)
	}
	return vec
}

// NewVector2 创建一个二维向量
func NewVector2[T constraints.Number](x, y T) Vector2 {
	return NewVector(x, y)
}

// NewVector3 创建一个三维向量
func NewVector3[T constraints.Number](x, y, z T) Vector3 {
	return NewVector(x, y, z)
}

// Add 向量相加
func (v Vector) Add(v2 Vector) Vector {
	if len(v) != len(v2) {
		panic("vector size mismatch")
	}
	vec := make(Vector, len(v))
	for i := range v {
		vec[i] = v[i] + v2[i]
	}
	return vec
}

// Sub 向量相减
func (v Vector) Sub(v2 Vector) Vector {
	if len(v) != len(v2) {
		panic("vector size mismatch")
	}
	vec := make(Vector, len(v))
	for i := range v {
		vec[i] = v[i] - v2[i]
	}
	return vec
}

// Mul 向量数乘
func (v Vector) Mul(scalar float64) Vector {
	vec := make(Vector, len(v))
	for i := range v {
		vec[i] = v[i] * scalar
	}
	return vec
}

// Div 向量除法
func (v Vector) Div(scalar float64) Vector {
	vec := make(Vector, len(v))
	for i := range v {
		vec[i] = v[i] / scalar
	}
	return vec
}

// Dot 向量点乘
func (v Vector) Dot(v2 Vector) float64 {
	if len(v) != len(v2) {
		panic("vector size mismatch")
	}
	var result float64
	for i := range v {
		result += v[i] * v2[i]
	}
	return result
}

// Cross3D 对三维向量进行叉乘
func (v Vector) Cross3D(v2 Vector3) Vector {
	if len(v) != 3 || len(v2) != 3 {
		panic("vector size mismatch")
	}

	return NewVector(
		v[1]*v2[2]-v[2]*v2[1],
		v[2]*v2[0]-v[0]*v2[2],
		v[0]*v2[1]-v[1]*v2[0],
	)
}

// Cross2D 对二维向量进行叉乘
func (v Vector2) Cross2D(v2 Vector2) float64 {
	if len(v) != 2 || len(v2) != 2 {
		panic("vector size mismatch")
	}

	return v[0]*v2[1] - v[1]*v2[0]
}

// Length 向量长度
func (v Vector) Length() float64 {
	return math.Sqrt(v.Dot(v))
}

// Normalize 向量归一化
func (v Vector) Normalize() Vector {
	return v.Div(v.Length())
}

// Angle 向量夹角
func (v Vector) Angle(v2 Vector) float64 {
	return math.Acos(v.Dot(v2) / (v.Length() * v2.Length()))
}

// PolarAngle 极坐标角度，即点在极坐标系中的角度
func (v Vector2) PolarAngle(v2 Vector2) float64 {
	return math.Atan2(v2.GetY()-v.GetY(), v2.GetX()-v.GetX())
}

// Equal 判断两个向量是否相等
func (v Vector) Equal(v2 Vector) bool {
	if len(v) != len(v2) {
		return false
	}
	for i := range v {
		if v[i] != v2[i] {
			return false
		}
	}
	return true
}

// Clone 复制一个向量
func (v Vector) Clone() Vector {
	vec := make(Vector, len(v))
	copy(vec, v)
	return vec
}

// IsZero 判断向量是否为零向量
func (v Vector) IsZero() bool {
	for _, val := range v {
		if val != 0 {
			return false
		}
	}
	return true
}

// Key 返回向量的键值
func (v Vector) Key() string {
	return ident.GenerateOrderedUniqueIdentStringWithUint64()
}

// IsParallel 判断两个向量是否平行
func (v Vector3) IsParallel(v2 Vector3) bool {
	return v.Cross3D(v2).IsZero()
}

// IsOrthogonal 判断两个向量是否垂直
func (v Vector) IsOrthogonal(v2 Vector) bool {
	return v.Dot(v2) == 0
}

// ProjectTo 向量投影
func (v Vector) ProjectTo(v2 Vector) Vector {
	return v2.Mul(v.Dot(v2) / v2.Dot(v2))
}

// RejectFrom 向量拒投影
func (v Vector) RejectFrom(v2 Vector) Vector {
	return v.Sub(v.ProjectTo(v2))
}

// ReflectBy 向量反射
func (v Vector) ReflectBy(normal Vector) Vector {
	return v.Sub(normal.Mul(2 * v.Dot(normal) / normal.Dot(normal)))
}

// AngleBetweenVectors 计算两个向量之间的夹角
func AngleBetweenVectors(v1, v2 Vector) float64 {
	return v1.Angle(v2)
}

// GetX 获取 x 坐标
func (v Vector2) GetX() float64 {
	return v[0]
}

// GetY 获取 y 坐标
func (v Vector2) GetY() float64 {
	return v[1]
}

// GetZ 获取 z 坐标
func (v Vector3) GetZ() float64 {
	return v[2]
}

// Quadrant 获取向量所在象限
func (v Vector2) Quadrant() int {
	AssertVector2Valid(v)
	x, y := v.GetXY()
	switch {
	case x > 0 && y > 0:
		return 1
	case x < 0 && y > 0:
		return 2
	case x < 0 && y < 0:
		return 3
	case x > 0 && y < 0:
		return 4
	default:
		return 0
	}
}

// Distance2D 计算二维向量之间的距离
func (v Vector2) Distance2D(v2 Vector2) float64 {
	x1, y1 := v.GetXY()
	x2, y2 := v2.GetXY()
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}

// DistanceSquared2D 计算二维向量之间的距离的平方
//   - 用于比较距离，避免开方运算
func (v Vector2) DistanceSquared2D(v2 Vector2) float64 {
	x1, y1 := v.GetXY()
	x2, y2 := v2.GetXY()
	return math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2)
}

// GetXY 返回该点的 x、y 坐标
func (v Vector2) GetXY() (x, y float64) {
	return v.GetX(), v.GetY()
}

// GetXYZ 返回该点的 x、y、z 坐标
func (v Vector3) GetXYZ() (x, y, z float64) {
	return v.GetX(), v.GetY(), v.GetZ()
}
