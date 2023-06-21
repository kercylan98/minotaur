package builtin

import "strconv"

const (
	PokerPointA          PokerPoint = 1
	PokerPoint2          PokerPoint = 2
	PokerPoint3          PokerPoint = 3
	PokerPoint4          PokerPoint = 4
	PokerPoint5          PokerPoint = 5
	PokerPoint6          PokerPoint = 6
	PokerPoint7          PokerPoint = 7
	PokerPoint8          PokerPoint = 8
	PokerPoint9          PokerPoint = 9
	PokerPoint10         PokerPoint = 10
	PokerPointJ          PokerPoint = 11
	PokerPointQ          PokerPoint = 12
	PokerPointK          PokerPoint = 13
	PokerPointBlackJoker PokerPoint = 14
	PokerPointRedJoker   PokerPoint = 15
)

// PokerPoint 扑克点数
type PokerPoint int

func (slf PokerPoint) String() string {
	var str string
	switch slf {
	case PokerPointA:
		str = "A"
	case PokerPointJ:
		str = "J"
	case PokerPointQ:
		str = "Q"
	case PokerPointK:
		str = "K"
	case PokerPointBlackJoker:
		str = "B"
	case PokerPointRedJoker:
		str = "R"
	default:
		str = strconv.Itoa(int(slf))
	}
	return str
}
