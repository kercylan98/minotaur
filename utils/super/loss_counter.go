package super

import "time"

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
