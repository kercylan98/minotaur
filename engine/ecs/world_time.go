package ecs

import "time"

func newWorldTime() *worldTime {
	return &worldTime{
		prevTime:  time.Now(),
		timeScale: 1,
	}
}

type worldTime struct {
	timeScale    float64       // 控制时间流逝的速度
	deltaTime    time.Duration // 上一次更新和当前更新之间的时间
	sleep        time.Duration // 休眠时间
	prevTime     time.Time     // 上一次更新的时间
	isPaused     bool          // 是否暂停
	maxDeltaTime time.Duration // 最大时间间隔
}

// Update 更新时间
func (t *worldTime) update() {
	now := time.Now()
	if t.isPaused {
		t.prevTime = now
	}

	t.deltaTime = now.Sub(t.prevTime)

	if t.timeScale != 1 {
		ms := float64(t.deltaTime.Milliseconds())
		ms *= t.timeScale
		t.deltaTime = time.Duration(ms) * time.Millisecond
	}

	if t.deltaTime > t.maxDeltaTime {
		t.deltaTime = t.maxDeltaTime
	}

	if t.sleep > 0 {
		d := t.deltaTime
		t.deltaTime -= t.sleep
		if t.deltaTime < 0 {
			t.deltaTime = 0
		}
		t.sleep -= d
		if t.sleep < 0 {
			t.sleep = 0
		}
	}

	t.prevTime = now
}

// SetSleep 设置休眠时间
func (t *worldTime) SetSleep(d time.Duration) {
	t.sleep = d
}

// SetTimeScale 设置时间流逝的速度
func (t *worldTime) SetTimeScale(scale float64) {
	t.timeScale = scale
}

// TimeScale 返回时间流逝的速度
func (t *worldTime) TimeScale() float64 {
	return t.timeScale
}

// DeltaTime 返回上一次更新和当前更新之间的时间
func (t *worldTime) DeltaTime() time.Duration {
	return t.deltaTime
}

// Pause 暂停时间
func (t *worldTime) Pause() {
	t.isPaused = true
}

// Resume 恢复时间
func (t *worldTime) Resume() {
	t.isPaused = false
}
