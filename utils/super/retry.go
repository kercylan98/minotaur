package super

import "time"

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

// RetryAsync 与 Retry 类似，但是是异步执行
//   - 传入的 callback 函数会在执行完毕后被调用，如果执行成功，则 err 为 nil，否则为错误信息
//   - 如果 callback 为 nil，则不会在执行完毕后被调用
func RetryAsync(count int, interval time.Duration, f func() error, callback func(err error)) {
	go func() {
		var err error
		for i := 0; i < count; i++ {
			if err = f(); err == nil {
				HandleV(nil, callback)
				return
			}
			time.Sleep(interval)
		}
		HandleV(err, callback)
	}()
}
