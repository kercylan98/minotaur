package config

import (
	jsonIter "github.com/json-iterator/go"
	"github.com/kercylan98/minotaur/utils/log"
	"github.com/kercylan98/minotaur/utils/timer"
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
	isInit         = true
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
	cLoadHandle(func(filename string, config any) error {
		bytes, err := os.ReadFile(filepath.Join(cLoadDir, filename))
		if err != nil {
			return err
		}
		if err = json.Unmarshal(bytes, &config); err == nil && isInit {
			log.Info("Config", log.String("Name", filename), log.Bool("LoadSuccess", true))
		}
		return err
	})
	isInit = false
	mutex.Unlock()
}

// WithTickerLoad 通过定时器加载配置
//   - 通过定时器加载配置后，会自动刷新线上配置
//   - 调用该函数后将会立即加载并刷新一次配置，随后每隔 interval 时间加载并刷新一次配置
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
