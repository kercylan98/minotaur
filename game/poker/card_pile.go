package poker

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/generic"
	"github.com/kercylan98/minotaur/utils/hash"
	"github.com/kercylan98/minotaur/utils/random"
	"github.com/kercylan98/minotaur/utils/slice"
	"sort"
)

// NewCardPile 返回一个新的牌堆，其中 size 表示了该牌堆由多少副牌组成
//   - 在不洗牌的情况下，默认牌堆顶部到底部为从大到小排列
//
// Deprecated: 从 Minotaur 0.0.25 开始，由于设计原因已弃用，请尝试考虑使用 deck.Deck 或 deck.Group 代替，构建函数为 deck.NewDeck 或 deck.NewGroup
func NewCardPile[P, C generic.Number, T Card[P, C]](size int, jokers [2]P, points [13]P, colors [4]C, generateCard func(guid int64, point P, color C) T, options ...CardPileOption[P, C, T]) *CardPile[P, C, T] {
	pile := &CardPile[P, C, T]{
		size:         size,
		pile:         make([]T, 0, size*54),
		generateCard: generateCard,
		cards:        map[int64]T{},
		jokers:       jokers,
		points:       points,
		colors:       colors,
	}
	pile.shuffleHandle = func(cards []T) []T {
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
type CardPile[P, C generic.Number, T Card[P, C]] struct {
	pile          []T
	size          int
	shuffleHandle func(cards []T) []T
	excludeColor  map[C]struct{}
	excludePoint  map[P]struct{}
	excludeCard   map[P]map[C]struct{}
	generateCard  func(guid int64, point P, color C) T
	guid          int64
	cards         map[int64]T
	jokers        [2]P
	points        [13]P
	colors        [4]C
}

// GetCard 通过 guid 获取一张牌
func (slf *CardPile[P, C, T]) GetCard(guid int64) T {
	return slf.cards[guid]
}

// Reset 重置牌堆的扑克牌数量及顺序
func (slf *CardPile[P, C, T]) Reset() {
	slf.guid = 0
	var cards = make([]T, 0, 54*slf.size)
	for i := 0; i < slf.size; i++ {
		for _, joker := range slf.jokers {
			if !slf.IsExclude(joker, C(0)) {
				slf.guid++
				card := slf.generateCard(slf.guid, joker, C(0))
				slf.cards[slf.guid] = card
				cards = append(cards, card)
			}
		}

		for p := 0; p < len(slf.points); p++ {
			for c := 0; c < len(slf.colors); c++ {
				point, color := slf.points[p], slf.colors[c]
				if !slf.IsExclude(point, color) {
					slf.guid++
					card := slf.generateCard(slf.guid, point, color)
					slf.cards[slf.guid] = card
				}
			}
		}
	}
	slf.pile = slf.pile[0:0]
	for i := 0; i < slf.size; i++ {
		slf.pile = append(slf.pile, cards...)
	}
}

// IsExclude 检查特定点数和花色是否被排除在外
func (slf *CardPile[P, C, T]) IsExclude(point P, color C) bool {
	return hash.Exist(slf.excludePoint, point) || hash.Exist(slf.excludeColor, color) || hash.Exist(slf.excludeCard[point], color)
}

// IsExcludeWithCard 检查特定扑克牌是否被排除在外
func (slf *CardPile[P, C, T]) IsExcludeWithCard(card T) bool {
	point, color := GetPointAndColor[P, C, T](card)
	return hash.Exist(slf.excludePoint, point) || hash.Exist(slf.excludeColor, color) || hash.Exist(slf.excludeCard[point], color)
}

// Shuffle 洗牌
func (slf *CardPile[P, C, T]) Shuffle() {
	before := slf.Count()
	cards := slf.shuffleHandle(slf.Cards())
	if len(cards) != before {
		panic("the count after shuffling does not match the count before shuffling")
	}
	slf.pile = cards
}

// Cards 获取当前牌堆的所有扑克牌
func (slf *CardPile[P, C, T]) Cards() []T {
	return slf.pile
}

// IsFree 返回牌堆是否没有扑克牌了
func (slf *CardPile[P, C, T]) IsFree() bool {
	return len(slf.pile) == 0
}

// Count 获取牌堆剩余牌量
func (slf *CardPile[P, C, T]) Count() int {
	return len(slf.pile)
}

// Pull 从牌堆特定位置抽出一张牌
func (slf *CardPile[P, C, T]) Pull(index int) T {
	if index >= slf.Count() || index < 0 {
		panic(fmt.Errorf("failed to pull a poker card from the pile, the index is less than 0 or exceeds the remaining number of cards in the pile. count: %d, index: %d", slf.Count(), index))
	}
	pc := slf.pile[index]
	slice.Del(&slf.pile, index)
	return pc
}

// PullTop 从牌堆顶部抽出一张牌
func (slf *CardPile[P, C, T]) PullTop() T {
	if slf.IsFree() {
		panic("empty poker cards pile")
	}
	pc := slf.pile[0]
	slice.Del(&slf.pile, 0)
	return pc
}

// PullBottom 从牌堆底部抽出一张牌
func (slf *CardPile[P, C, T]) PullBottom() T {
	if slf.IsFree() {
		panic("empty poker cards pile")
	}
	i := len(slf.pile) - 1
	pc := slf.pile[i]
	slice.Del(&slf.pile, i)
	return pc
}

// Push 将扑克牌插入到牌堆特定位置
func (slf *CardPile[P, C, T]) Push(index int, card T) {
	slice.Insert(&slf.pile, index, card)
	return
}

// PushTop 将扑克牌插入到牌堆顶部
func (slf *CardPile[P, C, T]) PushTop(card T) {
	slf.pile = append([]T{card}, slf.pile...)
}

// PushBottom 将扑克牌插入到牌堆底部
func (slf *CardPile[P, C, T]) PushBottom(card T) {
	slf.pile = append(slf.pile, card)
}
