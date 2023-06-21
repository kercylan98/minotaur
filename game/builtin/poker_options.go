package builtin

type PokerOption[PlayerID comparable] func(poker *Poker[PlayerID])

// WithPokerPointOrderOfSize 通过特定的点数大小顺序创建扑克玩法
//   - 顺序由大到小，points 数组必须包含每一个点数
func WithPokerPointOrderOfSize[PlayerID comparable](points [15]PokerPoint) PokerOption[PlayerID] {
	return func(poker *Poker[PlayerID]) {
		var compare = make(map[PokerPoint]int)
		for i, point := range points {
			compare[point] = len(points) - i
		}
		if len(compare) != len(points) {
			panic("not every point included")
		}
		poker.comparePoint = compare
	}
}

// WithPokerColorOrderOfSize 通过特定的花色大小顺序创建扑克玩法
//   - 顺序由大到小，colors 数组必须包含每一种花色
func WithPokerColorOrderOfSize[PlayerID comparable](colors [4]PokerColor) PokerOption[PlayerID] {
	return func(poker *Poker[PlayerID]) {
		var compare = make(map[PokerColor]int)
		for i, color := range colors {
			compare[color] = len(colors) - i
		}
		if len(compare) != len(colors) {
			panic("not every color included")
		}
		poker.compareColor = compare
	}
}
