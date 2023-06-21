package builtin

const (
	PokerColorNone    PokerColor = 0 // 无花色，通常为大小王
	PokerColorSpade   PokerColor = 1 // 黑桃
	PokerColorHeart   PokerColor = 2 // 红桃
	PokerColorClub    PokerColor = 3 // 梅花
	PokerColorDiamond PokerColor = 4 // 方片
)

// PokerColor 扑克牌花色
type PokerColor int

// InBounds 扑克牌花色是否在界限内
//   - 将检查花色是否在黑桃、红桃、梅花、方片之间
func (slf PokerColor) InBounds() bool {
	return slf <= PokerColorSpade && slf >= PokerColorDiamond
}

func (slf PokerColor) String() string {
	var str string
	switch slf {
	case PokerColorSpade:
		str = "Spade"
	case PokerColorHeart:
		str = "Heart"
	case PokerColorClub:
		str = "Club"
	case PokerColorDiamond:
		str = "Diamond"
	default:
		str = "None"
	}
	return str
}
