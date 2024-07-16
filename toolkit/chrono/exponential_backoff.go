package chrono

import (
	"math"
	"math/rand/v2"
	"time"
)

// StandardExponentialBackoff 退避指数函数用于计算下一次重试将在多长时间后发生，当返回 -1 时表示不再重试
//   - count：当前重试次数
//   - maxRetries：最大重试次数
//
// 该函数将以 2 为基数，0.5 为随机化因子进行计算
func StandardExponentialBackoff(count, maxRetries int, baseDelay, maxDelay time.Duration) time.Duration {
	return ExponentialBackoff(count, maxRetries, baseDelay, maxDelay, 2, 0.5)
}

// ExponentialBackoff 退避指数函数用于计算下一次重试将在多长时间后发生，当返回 -1 时表示不再重试
//   - count：当前重试次数
//   - maxRetries：最大重试次数
//   - baseDelay：基础延迟
//   - maxDelay：最大延迟
//   - multiplier：延迟时间的乘数，通常为 2
//   - randomization：延迟时间的随机化因子，通常为 0.5
func ExponentialBackoff(count, maxRetries int, baseDelay, maxDelay time.Duration, multiplier, randomization float64) time.Duration {
	for {
		if count >= maxRetries {
			return -1
		}

		delay := float64(baseDelay) * math.Pow(multiplier, float64(count))
		jitter := (rand.Float64() - 0.5) * randomization * float64(baseDelay)
		sleepDuration := time.Duration(delay + jitter)

		if sleepDuration > maxDelay {
			sleepDuration = maxDelay
		}

		return sleepDuration
	}
}
