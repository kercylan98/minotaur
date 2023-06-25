package poker

import "fmt"

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

// String 将扑克牌转换为字符串形式
func (slf Card) String() string {
	return fmt.Sprintf("(%s %s)", slf.point, slf.color)
}
