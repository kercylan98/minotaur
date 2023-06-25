package poker

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
