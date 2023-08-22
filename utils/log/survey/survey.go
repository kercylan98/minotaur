package survey

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/log"
	"github.com/kercylan98/minotaur/utils/super"
	"path/filepath"
	"strings"
	"time"
)

var (
	survey = make(map[string]*logger)
)

// RegSurvey 注册一个运营日志记录器
func RegSurvey(name, filePath string, options ...Option) {
	fn := filepath.Base(filePath)
	ext := filepath.Ext(fn)
	fn = strings.TrimSuffix(fn, ext)
	_, exist := survey[name]
	if exist {
		panic(fmt.Errorf("survey %s already exist", name))
	}
	dir := filepath.Dir(filePath)
	logger := &logger{
		dir:           dir,
		fn:            fn,
		fe:            ext,
		layout:        time.DateOnly,
		layoutLen:     len(time.DateOnly),
		dataLayout:    time.DateTime,
		dataLayoutLen: len(time.DateTime) + 3,
	}
	for _, option := range options {
		option(logger)
	}
	survey[name] = logger
	log.Info("Survey", log.String("Action", "RegSurvey"), log.String("Name", name), log.String("FilePath", dir+"/"+fn+".${DATE}"+ext))
}

// Record 记录一条运营日志
func Record(name string, data map[string]any) {
	logger, exist := survey[name]
	if !exist {
		panic(fmt.Errorf("survey %s not exist", name))
	}
	logger.writer(fmt.Sprintf("%s - %s\n", time.Now().Format(time.DateTime), super.MarshalJSON(data)))
}

// Flush 将所有运营日志记录器的缓冲区数据写入到文件
func Flush() {
	for _, logger := range survey {
		logger.flush()
	}
}
