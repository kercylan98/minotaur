package poker

// IsContainJoker 检查扑克牌是否包含大小王
func IsContainJoker(cards ...Card) bool {
	for _, card := range cards {
		if card.IsJoker() {
			return true
		}
	}
	return false
}

// GroupByPoint 将扑克牌按照点数分组
func GroupByPoint(cards ...Card) map[Point][]Card {
	group := map[Point][]Card{}
	for _, card := range cards {
		group[card.GetPoint()] = append(group[card.GetPoint()], card)
	}
	return group
}

// GroupByColor 将扑克牌按照花色分组
func GroupByColor(cards ...Card) map[Color][]Card {
	group := map[Color][]Card{}
	for _, card := range cards {
		group[card.GetColor()] = append(group[card.GetColor()], card)
	}
	return group
}

// IsRocket 两张牌能否组成红黑 Joker
func IsRocket(cardA, cardB Card) bool {
	return cardA.GetPoint() == PointRedJoker && cardB.GetPoint() == PointBlackJoker || cardA.GetPoint() == PointBlackJoker && cardB.GetPoint() == PointRedJoker
}

// IsFlush 判断是否是同花
func IsFlush(cards ...Card) bool {
	if len(cards) == 0 {
		return false
	}
	if len(cards) == 1 {
		return true
	}

	color := cards[0].GetColor()
	for i := 1; i < len(cards); i++ {
		if cards[i].GetColor() != color {
			return false
		}
	}
	return true
}

// GetCardsPoint 获取一组扑克牌的点数
func GetCardsPoint(cards ...Card) []Point {
	var points = make([]Point, len(cards))
	for i, card := range cards {
		points[i] = card.GetPoint()
	}
	return points
}

// GetCardsColor 获取一组扑克牌的花色
func GetCardsColor(cards ...Card) []Color {
	var colors = make([]Color, len(cards))
	for i, card := range cards {
		colors[i] = card.GetColor()
	}
	return colors
}
