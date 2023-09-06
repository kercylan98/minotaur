package survey

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/file"
	"github.com/kercylan98/minotaur/utils/log"
	"github.com/kercylan98/minotaur/utils/super"
	"sync"
	"time"
)

var (
	survey          = make(map[string]*logger)
	timers          = make(map[time.Duration]*time.Timer)
	timerSurvey     = make(map[time.Duration]map[string]struct{})
	timerSurveyLock sync.Mutex
)

// Reg 注册一个运营日志记录器
func Reg(name string, flusher Flusher, options ...Option) {

	timerSurveyLock.Lock()
	defer timerSurveyLock.Unlock()

	_, exist := survey[name]
	if exist {
		panic(fmt.Errorf("survey %s already exist", name))
	}
	logger := &logger{
		flusher:  flusher,
		interval: time.Second * 3,
	}
	for _, option := range options {
		option(logger)
	}

	if logger.interval > 0 {
		_, exist = timers[logger.interval]
		if !exist {
			t := time.NewTimer(logger.interval)
			timers[logger.interval] = t
			timerSurvey[logger.interval] = make(map[string]struct{})
			go func(interval time.Duration) {
				for {
					<-t.C
					timerSurveyLock.Lock()
					for n := range timerSurvey[interval] {
						survey[n].flush()
					}
					timerSurveyLock.Unlock()
					t.Reset(interval)
				}
			}(logger.interval)
		}
		timerSurvey[logger.interval][name] = struct{}{}
	}

	survey[name] = logger
	log.Info("Survey", log.String("Action", "Reg"), log.String("Name", name), log.String("Info", logger.flusher.Info()))
}

// Record 记录一条运营日志
func Record(name string, data map[string]any) {
	logger, exist := survey[name]
	if !exist {
		panic(fmt.Errorf("survey %s not exist", name))
	}
	logger.writer(fmt.Sprintf("%s - %s\n", time.Now().Format(time.DateTime), super.MarshalJSON(data)))
}

// Flush 将运营日志记录器的缓冲区数据写入到文件
//   - name 为空时，将所有记录器的缓冲区数据写入到文件
func Flush(names ...string) {
	timerSurveyLock.Lock()
	defer timerSurveyLock.Unlock()
	if len(names) == 0 {
		for _, logger := range survey {
			logger.flush()
		}
		return
	}
	for _, n := range names {
		l, e := survey[n]
		if e {
			l.flush()
		}
	}
}

// Close 关闭运营日志记录器
func Close(names ...string) {
	timerSurveyLock.Lock()
	defer timerSurveyLock.Unlock()
	if len(names) == 0 {
		for _, timer := range timers {
			timer.Stop()
		}
		Flush()
		return
	}
	for _, name := range names {
		l, e := survey[name]
		if e {
			delete(survey, name)
			delete(timerSurvey[l.interval], name)
			if len(timerSurvey[l.interval]) == 0 {
				delete(timerSurvey, l.interval)
				timers[l.interval].Stop()
				delete(timers, l.interval)
			}
			l.flush()
		}
	}
}

// Analyze 分析特定文件的记录，当发生错误时，会发生 panic
//   - handle 为并行执行的，需要自行处理并发安全
//   - 适用于外部进程对于日志文件的读取，但是需要注意的是，此时日志文件可能正在被写入，所以可能会读取到错误的数据
func Analyze(filePath string, handle func(analyzer *Analyzer, record R)) *Report {
	analyzer := new(Analyzer)
	err := file.ReadLineWithParallel(filePath, 1*1024*1024*1024, func(s string) {
		handle(analyzer, R(s))
	})
	if err != nil {
		panic(err)
	}

	return newReport(analyzer)
}
