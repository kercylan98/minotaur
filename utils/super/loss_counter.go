package super

import (
	"fmt"
	"strings"
	"time"
)

// StartLossCounter 开始损耗计数
func StartLossCounter() *LossCounter {
	return &LossCounter{curr: time.Now()}
}

type LossCounter struct {
	curr    time.Time
	loss    []time.Duration
	lossKey []string
}

// Record 记录一次损耗
func (slf *LossCounter) Record(name string) {
	slf.loss = append(slf.loss, time.Since(slf.curr))
	slf.lossKey = append(slf.lossKey, name)
	slf.curr = time.Now()
}

// GetLoss 获取损耗
func (slf *LossCounter) GetLoss(handler func(step int, name string, loss time.Duration)) {
	for i, loss := range slf.loss {
		handler(i, slf.lossKey[i], loss)
	}
}

func (slf *LossCounter) String() string {
	var lines []string
	slf.GetLoss(func(step int, name string, loss time.Duration) {
		lines = append(lines, fmt.Sprintf("%d. %s: %s", step, name, loss.String()))
	})
	return strings.Join(lines, "\n")
}

// StopWatch 计时器，返回 fn 执行耗时
func StopWatch(fn func()) time.Duration {
	start := time.Now()
	fn()
	return time.Since(start)
}

// StopWatchAndPrintln 计时器，返回 fn 执行耗时，并打印耗时
func StopWatchAndPrintln(name string, fn func()) time.Duration {
	loss := StopWatch(fn)
	fmt.Println(fmt.Sprintf("%s cost: %s", name, loss.String()))
	return loss
}
