package poker

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/hash"
	"github.com/kercylan98/minotaur/utils/random"
	"github.com/kercylan98/minotaur/utils/slice"
	"sort"
)

// NewCardPile 返回一个新的牌堆，其中 size 表示了该牌堆由多少副牌组成
//   - 在不洗牌的情况下，默认牌堆顶部到底部为从大到小排列
func NewCardPile(size int, options ...CardPileOption) *CardPile {
	pile := &CardPile{
		size: size,
		pile: make([]Card, 0, size*54),
	}
	pile.shuffleHandle = func(cards []Card) []Card {
		sort.Slice(cards, func(i, j int) bool {
			return random.Float64() >= 0.5
		})
		return cards
	}
	for _, option := range options {
		option(pile)
	}
	pile.Reset()
	return pile
}

// CardPile 扑克牌堆
type CardPile struct {
	pile          []Card
	size          int
	shuffleHandle func(cards []Card) []Card
	excludeColor  map[Color]struct{}
	excludePoint  map[Point]struct{}
	excludeCard   map[Point]map[Color]struct{}
}

// Reset 重置牌堆的扑克牌数量及顺序
func (slf *CardPile) Reset() {
	var cards = make([]Card, 0, 54)
	if !slf.IsExclude(PointRedJoker, ColorNone) {
		cards = append(cards, NewCard(PointRedJoker, ColorNone))
	}
	if !slf.IsExclude(PointBlackJoker, ColorNone) {
		cards = append(cards, NewCard(PointBlackJoker, ColorNone))
	}
	for point := PointK; point >= PointA; point-- {
		for color := ColorSpade; color <= ColorDiamond; color++ {
			if !slf.IsExclude(point, color) {
				cards = append(cards, NewCard(point, color))
			}
		}
	}
	slf.pile = slf.pile[0:0]
	for i := 0; i < slf.size; i++ {
		slf.pile = append(slf.pile, cards...)
	}
}

// IsExclude 检查特定点数和花色是否被排除在外
func (slf *CardPile) IsExclude(point Point, color Color) bool {
	if point == PointRedJoker || point == PointBlackJoker {
		color = ColorNone
	}
	return hash.Exist(slf.excludePoint, point) || hash.Exist(slf.excludeColor, color) || hash.Exist(slf.excludeCard[point], color)
}

// IsExcludeWithCard 检查特定扑克牌是否被排除在外
func (slf *CardPile) IsExcludeWithCard(card Card) bool {
	point, color := card.GetPointAndColor()
	return hash.Exist(slf.excludePoint, point) || hash.Exist(slf.excludeColor, color) || hash.Exist(slf.excludeCard[point], color)
}

// Shuffle 洗牌
func (slf *CardPile) Shuffle() {
	before := slf.Count()
	cards := slf.shuffleHandle(slf.Cards())
	if len(cards) != before {
		panic("the count after shuffling does not match the count before shuffling")
	}
	slf.pile = cards
}

// Cards 获取当前牌堆的所有扑克牌
func (slf *CardPile) Cards() []Card {
	return slf.pile
}

// IsFree 返回牌堆是否没有扑克牌了
func (slf *CardPile) IsFree() bool {
	return len(slf.pile) == 0
}

// Count 获取牌堆剩余牌量
func (slf *CardPile) Count() int {
	return len(slf.pile)
}

// Pull 从牌堆特定位置抽出一张牌
func (slf *CardPile) Pull(index int) Card {
	if index >= slf.Count() || index < 0 {
		panic(fmt.Errorf("failed to pull a poker card from the pile, the index is less than 0 or exceeds the remaining number of cards in the pile. count: %d, index: %d", slf.Count(), index))
	}
	pc := slf.pile[index]
	slice.Del(&slf.pile, index)
	return pc
}

// PullTop 从牌堆顶部抽出一张牌
func (slf *CardPile) PullTop() Card {
	if slf.IsFree() {
		panic("empty poker cards pile")
	}
	pc := slf.pile[0]
	slice.Del(&slf.pile, 0)
	return pc
}

// PullBottom 从牌堆底部抽出一张牌
func (slf *CardPile) PullBottom() Card {
	if slf.IsFree() {
		panic("empty poker cards pile")
	}
	i := len(slf.pile) - 1
	pc := slf.pile[i]
	slice.Del(&slf.pile, i)
	return pc
}

// Push 将扑克牌插入到牌堆特定位置
func (slf *CardPile) Push(index int, card Card) {
	slice.Insert(&slf.pile, index, card)
	return
}

// PushTop 将扑克牌插入到牌堆顶部
func (slf *CardPile) PushTop(card Card) {
	slf.pile = append([]Card{card}, slf.pile...)
}

// PushBottom 将扑克牌插入到牌堆底部
func (slf *CardPile) PushBottom(card Card) {
	slf.pile = append(slf.pile, card)
}
