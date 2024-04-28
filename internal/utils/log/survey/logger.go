package survey

import (
	"sync"
	"time"
)

// logger 用于埋点数据的运营日志记录器
type logger struct {
	bl       sync.Mutex // writer lock
	wl       sync.Mutex // flush lock
	bs       []string
	interval time.Duration
	flusher  Flusher
}

// flush 将记录器缓冲区的数据写入到文件
func (slf *logger) flush() {
	slf.bl.Lock()
	count := len(slf.bs)
	if count == 0 {
		slf.bl.Unlock()
		return
	}
	ds := slf.bs[:]
	slf.bs = slf.bs[count:]
	slf.bl.Unlock()

	slf.wl.Lock()
	defer slf.wl.Unlock()
	slf.flusher.Flush(ds)
}

// writer 写入数据到记录器缓冲区
func (slf *logger) writer(d string) {
	slf.bl.Lock()
	slf.bs = append(slf.bs, d)
	slf.bl.Unlock()
	if slf.interval <= 0 {
		slf.flush()
	}
}
