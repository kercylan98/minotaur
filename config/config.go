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

// LoadHandle 配置加载处理函数
type LoadHandle func(handle func(filename string, config any) error)

// RefreshHandle 配置刷新处理函数
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

// Init 配置初始化
func Init(loadDir string, loadHandle LoadHandle, refreshHandle RefreshHandle) {
	cLoadDir = loadDir
	cLoadHandle = loadHandle
	cRefreshHandle = refreshHandle
	Load()
	Refresh()
}

// Load 加载配置
//   - 加载后并不会刷新线上配置，需要执行 Refresh 函数对线上配置进行刷新
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
				log.Info("Config", zap.String("Name", filename), zap.Bool("LoadSuccess", true))
			}
			return err
		})
	}
	mutex.Unlock()
}

// WithTickerLoad 通过定时器加载配置
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

// StopTickerLoad 停止通过定时器加载配置
func StopTickerLoad() {
	if cTicker != nil {
		cTicker.StopTimer(tickerLoadRefresh)
	}
}

// Refresh 刷新配置
func Refresh() {
	mutex.Lock()
	cRefreshHandle()
	OnConfigRefreshEvent()
	mutex.Unlock()
}
