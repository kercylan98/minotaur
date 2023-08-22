package survey

import (
	"bufio"
	"fmt"
	"github.com/kercylan98/minotaur/utils/log"
	"os"
	"path/filepath"
	"sync"
)

const (
	DATETIME_FORMAT = "2006-01-02 15:04:05"
	DATE_FORMAT     = "2006-01-02"
	dateLen         = len(DATE_FORMAT)
	datetimeLen     = len(DATETIME_FORMAT)
)

// logger 用于埋点数据的运营日志记录器
type logger struct {
	bl            sync.Mutex // writer lock
	wl            sync.Mutex // flush lock
	dir           string
	fn            string
	fe            string
	bs            []string
	layout        string
	layoutLen     int
	dataLayout    string
	dataLayoutLen int
}

// flush 将记录器缓冲区的数据写入到文件
func (slf *logger) flush() {
	slf.bl.Lock()
	count := len(slf.bs)
	if count == 0 {
		slf.bl.Unlock()
		return
	}
	ds := slf.bs[:]
	slf.bs = slf.bs[count:]
	slf.bl.Unlock()

	slf.wl.Lock()
	defer slf.wl.Unlock()

	var (
		file   *os.File
		writer *bufio.Writer
		err    error
		last   string
	)
	for _, data := range ds {
		tick := data[0:slf.layoutLen]
		if tick != last {
			if file != nil {
				_ = writer.Flush()
				_ = file.Close()
			}
			fp := slf.filePath(tick)
			file, err = os.OpenFile(fp, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
			if err != nil {
				log.Fatal("Survey", log.String("Action", "DateSwitch"), log.String("FilePath", fp), log.Err(err))
				return
			}
			writer = bufio.NewWriterSize(file, 1024*10240)
			last = tick
		}
		_, _ = writer.WriteString(data)
	}
	_ = writer.Flush()
	_ = file.Close()
}

// writer 写入数据到记录器缓冲区
func (slf *logger) writer(d string) {
	slf.bl.Lock()
	slf.bs = append(slf.bs, d)
	slf.bl.Unlock()
}

// filePath 获取文件路径
func (slf *logger) filePath(t string) string {
	return filepath.Join(slf.dir, fmt.Sprintf("%s.%s%s", slf.fn, t, slf.fe))
}
