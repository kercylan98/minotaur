package storage

import "time"

var (
	// globalDataSaveHandles 全局数据保存句柄
	globalDataSaveHandles []func() error
)

// SaveAll 保存所有数据
//   - errorHandle 错误处理中如果返回 false 将重试，否则跳过当前保存下一个
func SaveAll(errorHandle func(err error) bool, retryInterval time.Duration) {
	var err error
	for _, handle := range globalDataSaveHandles {
		for {
			if err = handle(); err != nil {
				if !errorHandle(err) {
					time.Sleep(retryInterval)
					continue
				}
				break
			}
			break
		}
	}
}
