package poker

const (
	ColorNone    Color = 0 // 无花色，通常为大小王
	ColorSpade   Color = 4 // 黑桃
	ColorHeart   Color = 3 // 红桃
	ColorClub    Color = 2 // 梅花
	ColorDiamond Color = 1 // 方片
)

var defaultColorSort = map[Color]int{
	ColorSpade:   int(ColorSpade),
	ColorHeart:   int(ColorHeart),
	ColorClub:    int(ColorClub),
	ColorDiamond: int(ColorDiamond),
	ColorNone:    int(ColorDiamond + 1),
}

// Color 扑克牌花色
type Color int

// InBounds 扑克牌花色是否在界限内
//   - 将检查花色是否在黑桃、红桃、梅花、方片之间
func (slf Color) InBounds() bool {
	return slf <= ColorSpade && slf >= ColorDiamond
}

func (slf Color) String() string {
	var str string
	switch slf {
	case ColorSpade:
		str = "Spade"
	case ColorHeart:
		str = "Heart"
	case ColorClub:
		str = "Club"
	case ColorDiamond:
		str = "Diamond"
	default:
		str = "None"
	}
	return str
}
