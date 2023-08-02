package random

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
