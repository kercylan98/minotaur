package geometry

import (
	"github.com/kercylan98/minotaur/toolkit/constraints"
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

// Cross 对三维向量进行叉乘
func (v Vector) Cross(v2 Vector3) Vector {
	if len(v) != 3 || len(v2) != 3 {
		panic("vector size mismatch")
	}
	return NewVector(
		v[1]*v2[2]-v[2]*v2[1],
		v[2]*v2[0]-v[0]*v2[2],
		v[0]*v2[1]-v[1]*v2[0],
	)
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

// IsParallel 判断两个向量是否平行
func (v Vector) IsParallel(v2 Vector3) bool {
	return v.Cross(v2).IsZero()
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
