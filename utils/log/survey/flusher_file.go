package survey

import (
	"bufio"
	"fmt"
	"github.com/kercylan98/minotaur/utils/log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// NewFileFlusher 创建一个文件刷新器
//   - layout 为日志文件名的时间戳格式 (默认为 time.DateOnly)
func NewFileFlusher(filePath string, layout ...string) *FileFlusher {
	fn := filepath.Base(filePath)
	ext := filepath.Ext(fn)
	fn = strings.TrimSuffix(fn, ext)
	dir := filepath.Dir(filePath)
	fl := &FileFlusher{
		dir:       dir,
		fn:        fn,
		fe:        ext,
		layout:    time.DateOnly,
		layoutLen: len(time.DateOnly),
	}
	if len(layout) > 0 {
		fl.layout = layout[0]
		fl.layoutLen = len(fl.layout)
	}
	return fl
}

type FileFlusher struct {
	dir       string
	fn        string
	fe        string
	layout    string
	layoutLen int
}

func (slf *FileFlusher) Flush(records []string) {
	var (
		file   *os.File
		writer *bufio.Writer
		err    error
		last   string
	)
	for _, data := range records {
		tick := data[0:slf.layoutLen]
		if tick != last {
			if file != nil {
				_ = writer.Flush()
				_ = file.Close()
			}
			fp := filepath.Join(slf.dir, fmt.Sprintf("%s.%s%s", slf.fn, tick, slf.fe))
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

func (slf *FileFlusher) Info() string {
	return fmt.Sprintf("%s/%s.${DATE}%s", slf.dir, slf.fn, slf.fe)
}
