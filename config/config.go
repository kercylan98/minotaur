package config

import (
	jsonIter "github.com/json-iterator/go"
	"github.com/kercylan98/minotaur/utils/log"
	"github.com/kercylan98/minotaur/utils/timer"
	"go.uber.org/zap"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type LoadHandle func(handle func(filename string, config any) error)
type RefreshHandle func()

const (
	tickerLoadRefresh = "_tickerLoadRefresh"
)

var (
	cLoadDir       string
	cTicker        *timer.Ticker
	cInterval      time.Duration
	cLoadHandle    LoadHandle
	cRefreshHandle RefreshHandle
	json           = jsonIter.ConfigCompatibleWithStandardLibrary
	mutex          sync.Mutex
)

func Init(loadDir string, loadHandle LoadHandle, refreshHandle RefreshHandle) {
	cLoadDir = loadDir
	cLoadHandle = loadHandle
	cRefreshHandle = refreshHandle
}

func Load() {
	mutex.Lock()
	if cTicker != nil {
		WithTickerLoad(cTicker, cInterval)
	} else {
		cLoadHandle(func(filename string, config any) error {
			bytes, err := os.ReadFile(filepath.Join(cLoadDir, filename))
			if err != nil {
				return err
			}
			if err = json.Unmarshal(bytes, &config); err == nil {
				log.Error("Config", zap.String("Name", filename), zap.Bool("LoadSuccess", true))
			}
			return err
		})
	}
	mutex.Unlock()
}

func WithTickerLoad(ticker *timer.Ticker, interval time.Duration) {
	if ticker != cTicker && cTicker != nil {
		cTicker.StopTimer(tickerLoadRefresh)
	}
	cTicker = ticker
	cInterval = interval
	cTicker.Loop(tickerLoadRefresh, timer.Instantly, cInterval, timer.Forever, func() {
		Load()
		Refresh()
	})
}

func StopTickerLoad() {
	if cTicker != nil {
		cTicker.StopTimer(tickerLoadRefresh)
	}
}

func Refresh() {
	mutex.Lock()
	cRefreshHandle()
	OnConfigRefreshEvent()
	mutex.Unlock()
}
