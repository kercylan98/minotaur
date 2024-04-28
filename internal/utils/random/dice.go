package random

// Dice 掷骰子
//   - 常规掷骰子将返回 1-6 的随机数
func Dice() int {
	return Int(1, 6)
}

// DiceN 掷骰子
//   - 与 Dice 不同的是，将返回 1-N 的随机数
func DiceN(n int) int {
	if n <= 1 {
		return 1
	}
	return Int(1, n)
}
