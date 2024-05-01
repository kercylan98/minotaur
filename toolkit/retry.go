package toolkit

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"time"
)

// Retry 根据提供的 count 次数尝试执行 f 函数，如果 f 函数返回错误，则在 interval 后重试，直到成功或者达到 count 次数
func Retry(count int, interval time.Duration, f func() error) error {
	var err error
	for i := 0; i < count; i++ {
		if err = f(); err == nil {
			return nil
		}
		time.Sleep(interval)
	}
	return err
}

// RetryByRule 根据提供的规则尝试执行 f 函数，如果 f 函数返回错误，则根据 rule 的返回值进行重试
//   - rule 将包含一个入参，表示第几次重试，返回值表示下一次重试的时间间隔，当返回值为 0 时，表示不再重试
//   - rule 的 count 将在 f 首次失败后变为 1，因此 rule 的入参将从 1 开始
func RetryByRule(f func() error, rule func(count int) time.Duration) error {
	var count int
	var err error
	for {
		if err = f(); err != nil {
			count++
			next := rule(count)
			if next <= 0 {
				break
			}
			time.Sleep(next)
		} else {
			break
		}
	}
	return err
}

// RetryByExponentialBackoff 根据指数退避算法尝试执行 f 函数
//   - maxRetries：最大重试次数
//   - baseDelay：基础延迟
//   - maxDelay：最大延迟
//   - multiplier：延迟时间的乘数，通常为 2
//   - randomization：延迟时间的随机化因子，通常为 0.5
//   - ignoreErrors：忽略的错误，当 f 返回的错误在 ignoreErrors 中时，将不会进行重试
func RetryByExponentialBackoff(f func() error, maxRetries int, baseDelay, maxDelay time.Duration, multiplier, randomization float64, ignoreErrors ...error) error {
	return ConditionalRetryByExponentialBackoff(f, nil, maxRetries, baseDelay, maxDelay, multiplier, randomization, ignoreErrors...)
}

// ConditionalRetryByExponentialBackoff 该函数与 RetryByExponentialBackoff 类似，但是可以被中断
//   - cond 为中断条件，当 cond 返回 false 时，将会中断重试
//
// 该函数通常用于在重试过程中，需要中断重试的场景，例如：
//   - 用户请求开始游戏，由于网络等情况，进入重试状态。此时用户再次发送开始游戏请求，此时需要中断之前的重试，避免重复进入游戏
func ConditionalRetryByExponentialBackoff(f func() error, cond func() bool, maxRetries int, baseDelay, maxDelay time.Duration, multiplier, randomization float64, ignoreErrors ...error) error {
	retry := 0
	for {
		if cond != nil && !cond() {
			return fmt.Errorf("interrupted")
		}
		err := f()
		if err == nil {
			return nil
		}
		for _, ignore := range ignoreErrors {
			if errors.Is(err, ignore) {
				return err
			}
		}

		if retry >= maxRetries {
			return fmt.Errorf("max retries reached: %w", err)
		}

		delay := float64(baseDelay) * math.Pow(multiplier, float64(retry))
		jitter := (rand.Float64() - 0.5) * randomization * float64(baseDelay)
		sleepDuration := time.Duration(delay + jitter)

		if sleepDuration > maxDelay {
			sleepDuration = maxDelay
		}

		time.Sleep(sleepDuration)
		retry++
	}
}

// RetryAsync 与 Retry 类似，但是是异步执行
//   - 传入的 callback 函数会在执行完毕后被调用，如果执行成功，则 err 为 nil，否则为错误信息
//   - 如果 callback 为 nil，则不会在执行完毕后被调用
func RetryAsync(count int, interval time.Duration, f func() error, callback func(err error)) {
	go func() {
		var err error
		for i := 0; i < count; i++ {
			if err = f(); err == nil {
				if callback != nil {
					callback(nil)
				}
				return
			}
			time.Sleep(interval)
		}
		if callback != nil {
			callback(err)
		}
	}()
}

// RetryForever 根据提供的 interval 时间间隔尝试执行 f 函数，如果 f 函数返回错误，则在 interval 后重试，直到成功
func RetryForever(interval time.Duration, f func() error) {
	var err error
	for {
		if err = f(); err == nil {
			return
		}
		time.Sleep(interval)
	}
}
