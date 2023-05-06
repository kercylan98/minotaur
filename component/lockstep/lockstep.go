package lockstep

import (
	"github.com/kercylan98/minotaur/utils/timer"
	"sync"
	"time"
)

func New[WriterID comparable, FrameCommand any](frameFactory func(frameIndex uint32) Frame[FrameCommand]) *Lockstep[WriterID, FrameCommand] {
	lockstep := &Lockstep[WriterID, FrameCommand]{
		frameFactory:       frameFactory,
		ticker:             timer.GetTicker(30),
		writers:            map[WriterID]Writer[WriterID, FrameCommand]{},
		writerCurrentFrame: map[WriterID]uint32{},
		frames:             map[uint32]Frame[FrameCommand]{},
	}
	return lockstep
}

type Lockstep[WriterID comparable, FrameCommand any] struct {
	frameFactory func(frameIndex uint32) Frame[FrameCommand]

	FrameLimit             uint32 // 帧数上限
	FrameRate              uint32 // 帧率(每秒的帧数)
	FrameBroadcastInterval uint32 // 帧数广播间隔帧数
	FrameOnceLimit         uint32 // 每次消息最大容纳帧数

	ticker                *timer.Ticker                               // 定时器
	writers               map[WriterID]Writer[WriterID, FrameCommand] // 被广播的对象
	writerCurrentFrame    map[WriterID]uint32                         // 被广播的对象当前帧
	currentFrame          uint32                                      // 当前帧
	writerMaxCurrentFrame uint32                                      // 最大写入器当前帧
	frames                map[uint32]Frame[FrameCommand]              // 所有帧
	framesRWMutes         sync.RWMutex                                // 所有帧读写锁
}

// SetWriter 设置需要被广播的 Writer
func (slf *Lockstep[WriterID, FrameCommand]) SetWriter(writer ...Writer[WriterID, FrameCommand]) {
	for _, w := range writer {
		slf.writers[w.GetID()] = w
	}
}

func (slf *Lockstep[WriterID, FrameCommand]) Run() error {
	if slf.frameFactory == nil {
		return ErrFrameFactorCanNotIsNull
	}
	slf.Release()
	slf.ticker.Loop(tickerFrameName, timer.Instantly, time.Second/time.Duration(slf.FrameRate), timer.Forever, slf.tick)
	return nil
}

func (slf *Lockstep[WriterID, FrameCommand]) Record(command FrameCommand) {
	slf.framesRWMutes.RLock()
	frame, exist := slf.frames[slf.currentFrame]
	slf.framesRWMutes.RUnlock()
	if !exist {
		frame = slf.frameFactory(slf.currentFrame)
		slf.framesRWMutes.Lock()
		slf.frames[slf.currentFrame] = frame
		slf.framesRWMutes.Unlock()
	}
	frame.AddCommand(command)
}

func (slf *Lockstep[WriterID, FrameCommand]) Release() {
	slf.ticker.StopTimer(tickerFrameName)
	for k := range slf.writers {
		delete(slf.writers, k)
	}
	for k := range slf.writerCurrentFrame {
		delete(slf.writers, k)
	}
	slf.currentFrame = 0
	slf.framesRWMutes.Lock()
	for k := range slf.frames {
		delete(slf.frames, k)
	}
	slf.framesRWMutes.Unlock()
}

// ReWrite 重写
func (slf *Lockstep[WriterID, FrameCommand]) ReWrite(writer Writer[WriterID, FrameCommand]) {
	if !writer.Healthy() {
		return
	}

	var writerCurrentFrame uint32
	var frameCounter uint32
	var frames = make(map[uint32]Frame[FrameCommand])
	for ; writerCurrentFrame < slf.currentFrame; writerCurrentFrame++ {
		slf.framesRWMutes.RLock()
		var frame = slf.frames[writerCurrentFrame]
		slf.framesRWMutes.RUnlock()
		if frame == nil && writerCurrentFrame != (slf.currentFrame-1) {
			continue
		}

		frames[frame.GetIndex()] = frame
		frameCounter++
		if writerCurrentFrame == slf.currentFrame-1 || (slf.FrameOnceLimit > 0 && frameCounter >= slf.FrameOnceLimit) {
			data := writer.Marshal(frames)
			// TODO: writer.Write error not handle
			_ = writer.Write(data)
			frameCounter = 0
			for k := range frames {
				delete(frames, k)
			}
		}
	}
	slf.writerCurrentFrame[writer.GetID()] = writerCurrentFrame
	if writerCurrentFrame > slf.writerMaxCurrentFrame {
		slf.writerMaxCurrentFrame = writerCurrentFrame
	}
}

func (slf *Lockstep[WriterID, FrameCommand]) tick() {
	if slf.FrameLimit > 0 && slf.currentFrame >= slf.FrameLimit {
		slf.ticker.StopTimer(tickerFrameName)
		return
	}

	slf.currentFrame++

	if slf.currentFrame-slf.writerMaxCurrentFrame < slf.FrameBroadcastInterval {
		return
	}

	for id, writer := range slf.writers {

		if !writer.Healthy() {
			continue
		}

		var writerCurrentFrame uint32
		var frameCounter uint32
		var frames = make(map[uint32]Frame[FrameCommand])
		for writerCurrentFrame = slf.writerCurrentFrame[id]; writerCurrentFrame < slf.currentFrame; writerCurrentFrame++ {
			slf.framesRWMutes.RLock()
			var frame = slf.frames[writerCurrentFrame]
			slf.framesRWMutes.RUnlock()
			if frame == nil && writerCurrentFrame != (slf.currentFrame-1) {
				continue
			}

			frames[frame.GetIndex()] = frame
			frameCounter++
			if writerCurrentFrame == slf.currentFrame-1 || (slf.FrameOnceLimit > 0 && frameCounter >= slf.FrameOnceLimit) {
				data := writer.Marshal(frames)
				// TODO: writer.Write error not handle
				_ = writer.Write(data)
				frameCounter = 0
				for k := range frames {
					delete(frames, k)
				}
			}
		}
		slf.writerCurrentFrame[id] = writerCurrentFrame
		if writerCurrentFrame > slf.writerMaxCurrentFrame {
			slf.writerMaxCurrentFrame = writerCurrentFrame
		}
	}
}
