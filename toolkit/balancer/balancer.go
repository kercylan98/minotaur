package balancer

import (
	"errors"
	"github.com/kercylan98/minotaur/toolkit/constraints"
)

var (
	ErrNoInstance = errors.New("no instance") // 没有实例
)

type Balancer[I constraints.Ordered, T Item[I]] interface {
	// Select 选择一个实例，如果没有实例则返回 ErrNoInstance
	Select(opts ...*SelectOptions) (T, error)

	// Add 添加一个实例
	Add(instance T)

	// Remove 移除一个实例
	Remove(instance T)

	// GetInstances 获取所有实例
	GetInstances() []T
}
