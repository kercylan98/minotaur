package poker

import (
	"fmt"
	"math"
)

// NewCard 创建一张扑克牌
//   - 当 point 为 PointBlackJoker 或 PointRedJoker 时，color 将没有效果
func NewCard(point Point, color Color) Card {
	if point == PointRedJoker || point == PointBlackJoker {
		color = ColorNone
	}
	card := Card{
		point: point,
		color: color,
	}
	return card
}

// Card 扑克牌
type Card struct {
	point Point
	color Color
}

// GetPoint 返回扑克牌的点数
func (slf Card) GetPoint() Point {
	return slf.point
}

// GetColor 返回扑克牌的花色
func (slf Card) GetColor() Color {
	if slf.point == PointRedJoker || slf.point == PointBlackJoker {
		return ColorNone
	}
	return slf.color
}

// GetPointAndColor 返回扑克牌的点数和花色
func (slf Card) GetPointAndColor() (Point, Color) {
	return slf.GetPoint(), slf.GetColor()
}

// EqualPoint 比较与另一张扑克牌的点数是否相同
func (slf Card) EqualPoint(card Card) bool {
	return slf.GetPoint() == card.GetPoint()
}

// EqualColor 比较与另一张扑克牌的花色是否相同
func (slf Card) EqualColor(card Card) bool {
	return slf.GetColor() == card.GetColor()
}

// Equal 比较与另一张扑克牌的点数和花色是否相同
func (slf Card) Equal(card Card) bool {
	return slf.GetPoint() == card.GetPoint() && slf.GetColor() == card.GetColor()
}

// MaxPoint 返回两张扑克牌中点数较大的一张
func (slf Card) MaxPoint(card Card) Card {
	if slf.GetPoint() > card.GetPoint() {
		return slf
	}
	return card
}

// MinPoint 返回两张扑克牌中点数较小的一张
func (slf Card) MinPoint(card Card) Card {
	if slf.GetPoint() < card.GetPoint() {
		return slf
	}
	return card
}

// MaxColor 返回两张扑克牌中花色较大的一张
func (slf Card) MaxColor(card Card) Card {
	if slf.GetColor() > card.GetColor() {
		return slf
	}
	return card
}

// MinColor 返回两张扑克牌中花色较小的一张
func (slf Card) MinColor(card Card) Card {
	if slf.GetColor() < card.GetColor() {
		return slf
	}
	return card
}

// Max 返回两张扑克牌中点数和花色较大的一张
func (slf Card) Max(card Card) Card {
	if slf.GetPoint() > card.GetPoint() {
		return slf
	} else if slf.GetPoint() < card.GetPoint() {
		return card
	} else {
		if slf.GetColor() > card.GetColor() {
			return slf
		}
		return card
	}
}

// Min 返回两张扑克牌中点数和花色较小的一张
func (slf Card) Min(card Card) Card {
	if slf.GetPoint() < card.GetPoint() {
		return slf
	} else if slf.GetPoint() > card.GetPoint() {
		return card
	} else {
		if slf.GetColor() < card.GetColor() {
			return slf
		}
		return card
	}
}

// IsJoker 判断是否为大小王
func (slf Card) IsJoker() bool {
	point := slf.GetPoint()
	return point == PointRedJoker || point == PointBlackJoker
}

// CalcPointDifference 计算两张扑克牌的点数差
func (slf Card) CalcPointDifference(card Card) int {
	return int(slf.GetPoint()) - int(card.GetPoint())
}

// CalcPointDifferenceAbs 计算两张扑克牌的点数差的绝对值
func (slf Card) CalcPointDifferenceAbs(card Card) int {
	return int(math.Abs(float64(slf.CalcPointDifference(card))))
}

// CalcColorDifference 计算两张扑克牌的花色差
func (slf Card) CalcColorDifference(card Card) int {
	return int(slf.GetColor()) - int(card.GetColor())
}

// CalcColorDifferenceAbs 计算两张扑克牌的花色差的绝对值
func (slf Card) CalcColorDifferenceAbs(card Card) int {
	return int(math.Abs(float64(slf.CalcColorDifference(card))))
}

// IsNeighborPoint 判断两张扑克牌是否为相邻的点数
func (slf Card) IsNeighborPoint(card Card) bool {
	return slf.CalcPointDifferenceAbs(card) == 1
}

// IsNeighborColor 判断两张扑克牌是否为相邻的花色
func (slf Card) IsNeighborColor(card Card) bool {
	return slf.CalcColorDifferenceAbs(card) == 1
}

// String 将扑克牌转换为字符串形式
func (slf Card) String() string {
	return fmt.Sprintf("(%s %s)", slf.point, slf.color)
}
