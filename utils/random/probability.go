package random

import (
	"math/rand"
	"time"
)

// ProbabilitySlice 按概率随机从切片中产生一个数据并返回命中的对象及是否未命中
//   - 当总概率小于 1 将会发生未命中的情况
func ProbabilitySlice[T any](getProbabilityHandle func(data T) float64, data ...T) (hit T, miss bool) {
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))

	var total float64
	for _, d := range data {
		total += getProbabilityHandle(d)
	}

	scaleFactor := 1.0
	var indexLimit = len(data)
	var limitP = 0.0
	if total > 1.0 {
		scaleFactor = 1.0 / total
	} else if total < 1.0 {
		indexLimit++
		limitP = 1.0 - total
	}

	var overlayProbability []float64
	cumulativeProbability := 0.0
	for i := 0; i < indexLimit; i++ {
		if i < len(data) {
			cumulativeProbability += getProbabilityHandle(data[i]) * scaleFactor
			overlayProbability = append(overlayProbability, cumulativeProbability)
		} else {
			cumulativeProbability += limitP * scaleFactor
			overlayProbability = append(overlayProbability, cumulativeProbability)
		}

	}

	if total < 1.0 {
		overlayProbability[len(overlayProbability)-1] = 1.0
	}

	var r = rd.Float64()
	var i, count = 0, len(overlayProbability)
	for i < count {
		h := int(uint(i+count) >> 1)
		if overlayProbability[h] <= r {
			i = h + 1
		} else {
			count = h
		}
	}
	if i >= len(data) {
		return hit, true
	}
	hit = data[i]
	return hit, false
}

// Probability 输入一个概率，返回是否命中
//   - 当 full 不为空时，将以 full 为基数，p 为分子，计算命中概率
func Probability(p int, full ...int) bool {
	var f = 100
	if len(full) > 0 {
		f = full[0]
		if f <= 0 {
			f = 100
		} else if p > f {
			return true
		}
	}
	r := Int(1, f)
	return r <= p
}

// ProbabilityChooseOne 输入一组概率，返回命中的索引
func ProbabilityChooseOne(ps ...int) int {
	var f int
	for _, p := range ps {
		f += p
	}
	if f <= 0 {
		panic("total probability less than or equal to 0")
	}
	r := Int(1, f)
	for i, p := range ps {
		if r <= p {
			return i
		}
		r -= p
	}
	panic("probability choose one error")
}
