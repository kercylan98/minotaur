package builtin

import "fmt"

// NewPokerCard 创建一张扑克牌
//   - 当 point 为 PokerPointBlackJoker 或 PokerPointRedJoker 时，color 将没有效果
func NewPokerCard(point PokerPoint, color PokerColor) PokerCard {
	if point == PokerPointRedJoker || point == PokerPointBlackJoker {
		color = PokerColorNone
	}
	card := PokerCard{
		point: point,
		color: color,
	}
	return card
}

// PokerCard 扑克牌
type PokerCard struct {
	point PokerPoint
	color PokerColor
}

// GetPoint 返回扑克牌的点数
func (slf PokerCard) GetPoint() PokerPoint {
	return slf.point
}

// GetColor 返回扑克牌的花色
func (slf PokerCard) GetColor() PokerColor {
	return slf.color
}

// GetPointAndColor 返回扑克牌的点数和花色
func (slf PokerCard) GetPointAndColor() (PokerPoint, PokerColor) {
	return slf.point, slf.color
}

// String 将扑克牌转换为字符串形式
func (slf PokerCard) String() string {
	return fmt.Sprintf("(%s %s)", slf.point, slf.color)
}
