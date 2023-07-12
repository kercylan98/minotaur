package log

const (
	// RunModeDev 开发模式是默认的运行模式，同时也是最基础的运行模式
	//   - 开发模式下，将会输出所有级别的日志到控制台
	//   - 默认不再输出日志到文件
	RunModeDev RunMode = iota
	// RunModeTest 测试模式是一种特殊的运行模式，用于测试
	//   - 测试模式下，将会输出所有级别的日志到控制台和文件
	RunModeTest
	// RunModeProd 生产模式是一种特殊的运行模式，用于生产
	//   - 生产模式下，将会输出 InfoLevel 及以上级别的日志到控制台和文件
	RunModeProd
)

type RunMode uint8
