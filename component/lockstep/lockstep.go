package lockstep

import (
	"minotaur/utils/synchronization"
	"minotaur/utils/timer"
	"time"
)

func New[WriterID comparable, FrameCommand any]() *Lockstep[WriterID, FrameCommand] {
	lockstep := &Lockstep[WriterID, FrameCommand]{
		ticker:             timer.GetTicker(30),
		writers:            synchronization.NewMap[WriterID, Writer[WriterID, FrameCommand]](),
		writerCurrentFrame: synchronization.NewMap[WriterID, uint32](),
		frames:             map[uint32]Frame[FrameCommand]{},
	}
	return lockstep
}

type Lockstep[WriterID comparable, FrameCommand any] struct {
	FrameLimit             uint32 // 帧数上限
	FrameRate              uint32 // 帧率(每秒的帧数)
	FrameBroadcastInterval uint32 // 帧数广播间隔帧数
	FrameOnceLimit         uint32 // 每次消息最大容纳帧数

	ticker             *timer.Ticker                                                  // 定时器
	writers            *synchronization.Map[WriterID, Writer[WriterID, FrameCommand]] // 被广播的对象
	writerCurrentFrame *synchronization.Map[WriterID, uint32]                         // 被广播的对象当前帧
	currentFrame       uint32                                                         // 当前帧
	currentClientFrame uint32                                                         // 当前客户端帧数
	frames             map[uint32]Frame[FrameCommand]                                 // 所有帧
}

// SetWriter 设置需要被广播的 Writer
func (slf *Lockstep[WriterID, FrameCommand]) SetWriter(writer ...Writer[WriterID, FrameCommand]) {
	for _, w := range writer {
		slf.writers.Set(w.GetID(), w)
	}
}

func (slf *Lockstep[WriterID, FrameCommand]) Run() {
	slf.Release()
	slf.ticker.Loop(tickerFrameName, timer.Instantly, time.Second/time.Duration(slf.FrameRate), timer.Forever, slf.tick)
}

func (slf *Lockstep[WriterID, FrameCommand]) Release() {
	slf.ticker.StopTimer(tickerFrameName)
	slf.writers.Clear()
	slf.writerCurrentFrame.Clear()
	slf.currentFrame = 0
	slf.currentClientFrame = 0
	for k := range slf.frames {
		delete(slf.frames, k)
	}
}

func (slf *Lockstep[WriterID, FrameCommand]) tick() {
	if slf.FrameLimit > 0 && slf.currentFrame >= slf.FrameLimit {
		slf.ticker.StopTimer(tickerFrameName)
		return
	}

	slf.currentFrame++

	if slf.currentFrame-slf.currentClientFrame < slf.FrameBroadcastInterval {
		return
	}

	slf.writers.RangeSkip(func(id WriterID, writer Writer[WriterID, FrameCommand]) bool {

		if !writer.Healthy() {
			return false
		}

		var frameCounter uint32
		var frames = make(map[uint32]Frame[FrameCommand])
		for i := slf.writerCurrentFrame.Get(id); i < slf.currentFrame; i++ {
			var frame = slf.frames[i]
			if frame == nil && i != (slf.currentFrame-1) {
				continue
			}

			frames[frame.GetIndex()] = frame
			frameCounter++
			if i == slf.currentFrame-1 || frameCounter >= slf.FrameOnceLimit {
				data := writer.Marshal(frames)
				writer.Write(data)
				frameCounter = 0
				for k := range frames {
					delete(frames, k)
				}
			}
		}
		slf.currentClientFrame = slf.currentFrame

		return true
	})
}
