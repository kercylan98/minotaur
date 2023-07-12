package log

import (
	rotateLogs "github.com/lestrrat-go/file-rotatelogs"
)

type Option func(log *Log)

// WithRunMode 设置运行模式
//   - 默认的运行模式为: RunModeDev
//   - 当 handle 不为空时，将会调用 handle()，并将返回值添加到日志记录器中，同时将会抑制默认的日志记录器
func WithRunMode(mode RunMode, handle func() Core) Option {
	return func(log *Log) {
		log.mode = mode
		if handle != nil {
			log.cores = append(log.cores, handle())
		}
	}
}

// WithFilename 设置日志文件名
//   - 默认的日志文件名为: {level}.log
func WithFilename(filename func(level Level) string) Option {
	return func(log *Log) {
		log.filename = filename
	}
}

// WithRotateFilename 设置日志分割文件名
//   - 默认的日志分割文件名为: {level}.%Y%m%d.log
func WithRotateFilename(filename func(level Level) string) Option {
	return func(log *Log) {
		log.rotateFilename = filename
	}
}

// WithRotateOption 设置日志分割选项
//   - 默认的日志分割选项为: WithMaxAge(7天), WithRotationTime(1天)
func WithRotateOption(options ...rotateLogs.Option) Option {
	return func(log *Log) {
		log.rotateOptions = options
	}
}

// WithLogDir 设置日志文件夹
//   - 默认情况下不会设置日志文件夹，日志将不会被文件存储
func WithLogDir(logDir, rotateLogDir string) Option {
	return func(log *Log) {
		log.logDir = logDir
		log.rotateLogDir = rotateLogDir
	}
}
