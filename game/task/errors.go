package task

import "errors"

var (
	// ErrTaskNotFinish 任务未完成
	ErrTaskNotFinish = errors.New("task not finish")
	// ErrTaskRewardReceived 任务奖励已领取
	ErrTaskRewardReceived = errors.New("task reward received")
	// ErrTaskNotStart 任务未开始
	ErrTaskNotStart = errors.New("task not start")
	// ErrTaskFail 任务失败
	ErrTaskFail = errors.New("task fail")
)
