package memory

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/collection"
	"github.com/kercylan98/minotaur/utils/random"
	"github.com/kercylan98/minotaur/utils/super"
	"reflect"
	"strings"
	"sync"
)

var (
	running            = false
	runLock            sync.Mutex
	persistTickerNames = make(map[string]map[string]struct{})
	persistRun         = make([]func(), 0)
	actionOutput       = make(map[string][]reflect.Type)
	caches             = make(map[string]map[string][]reflect.Value)
	cachesRWMutex      sync.RWMutex
)

// Run 运行持久化缓存程序
func Run() {
	runLock.Lock()
	defer runLock.Unlock()
	if running {
		panic(fmt.Errorf("can not run persist cache program twice"))
	}
	running = true
	for _, f := range persistRun {
		f()
	}
	persistRun = nil
}

// BindPersistCacheProgram 绑定持久化缓存程序
//   - name 持久化缓存程序名称
//   - handler 持久化缓存程序处理函数
//   - option 持久化缓存程序选项
//
// 注意事项：
//   - 持久化程序建议声明为全局变量进行使用
//   - 持久化程序处理函数参数类型必须与绑定的缓存程序输出参数类型一致，并且相同 name 的持久化程序必须在 BindAction 之后进行绑定
//   - 默认情况下只有执行该函数返回的函数才会进行持久化，如果需要持久化策略，可以设置 option 参数或者自行实现策略调用返回的函数
//   - 所有持久化程序绑定完成后，应该主动调用 Run 函数运行
func BindPersistCacheProgram[OutputParamHandlerFunc any](name string, handler OutputParamHandlerFunc, option ...*Option) func() {
	runLock.Lock()
	defer runLock.Unlock()
	if running {
		panic(fmt.Errorf("can not bind persist cache program after running"))
	}
	v := reflect.ValueOf(handler)
	if v.Kind() != reflect.Func {
		panic("handle is not a function")
	}

	outputs, exist := actionOutput[name]
	if !exist {
		panic(fmt.Errorf("action %s not exist", name))
	}

	if len(outputs) != v.Type().NumIn() {
		panic(fmt.Errorf("action %s output params count %d not equal handler input params count %d", name, len(outputs), v.Type().NumIn()))
	}

	for i := 0; i < v.Type().NumIn(); i++ {
		if outputs[i] != v.Type().In(i) {
			panic(fmt.Errorf("action %s output param %d type %s not equal handler input param %d type %s", name, i, outputs[i].String(), i, v.Type().In(i).String()))
		}
	}

	persist := reflect.MakeFunc(v.Type(), func(args []reflect.Value) []reflect.Value {
		results := v.Call(args)
		return results
	})
	executor := func() {
		cachesRWMutex.RLock()
		funcCache, exist := caches[name]
		if !exist {
			cachesRWMutex.RUnlock()
			return
		}
		funcCache = collection.CloneMap(funcCache)
		cachesRWMutex.RUnlock()
		for _, results := range funcCache {
			persist.Call(results)
		}
	}

	var opt *Option
	if len(option) > 0 {
		opt = option[0]
	}
	if opt != nil {
		if opt.ticker == nil {
			panic(fmt.Errorf("option ticker is nil"))
		}
		var loopName = fmt.Sprintf("periodic_persistence:%d:%s:%s", len(persistTickerNames[name]), name, random.HostName())
		if _, exist := persistTickerNames[name]; !exist {
			persistTickerNames[name] = make(map[string]struct{})
		}
		persistTickerNames[name][loopName] = struct{}{}

		var after = super.If(opt.firstDelay == 0, opt.interval, opt.firstDelay)
		if opt.delay > 0 {
			executor = func() {
				cachesRWMutex.RLock()
				funcCache, exist := caches[name]
				if !exist {
					cachesRWMutex.RUnlock()
					return
				}
				funcCache = collection.CloneMap(funcCache)
				cachesRWMutex.RUnlock()
				delay := opt.delay
				tick := delay
				for actionId, c := range funcCache {
					opt.ticker.After(fmt.Sprintf("%s:%v", loopName, actionId), tick, func(c []reflect.Value) {
						persist.Call(c)
					}, c)
					tick += delay
				}
			}
		}
		persistRun = append(persistRun, func() {
			opt.ticker.Loop("periodic_persistence", after, opt.interval, -1, executor)
		})
	}

	return executor
}

// BindAction 绑定需要缓存的操作函数
//   - name 缓存操作名称
//   - handler 缓存操作处理函数
//
// 注意事项：
//   - 关于持久化缓存程序的绑定请参考 BindPersistCacheProgram
//   - handler 函数的返回值将被作为缓存目标，如果返回值为非指针类型，将可能会发生意外的情况
//   - 当传入的 handler 没有任何返回值时，将不会被缓存，并且不会占用缓存操作名称
//
// 使用场景：
//   - 例如在游戏中，需要根据玩家 ID 查询玩家信息，可以使用该函数进行绑定，当查询玩家信息时，如果缓存中存在该玩家信息，将直接返回缓存中的数据，否则将执行 handler 函数进行查询并缓存
func BindAction[Func any](name string, handler Func) Func {
	v := reflect.ValueOf(handler)
	if v.Kind() != reflect.Func {
		panic(fmt.Errorf("handle is not a function"))
	}
	if v.Type().NumOut() == 0 {
		return handler
	}

	if _, exist := actionOutput[name]; exist {
		panic(fmt.Errorf("action %s already exist", name))
	}

	outputs := make([]reflect.Type, 0, v.Type().NumOut())
	for i := 0; i < v.Type().NumOut(); i++ {
		outputs = append(outputs, v.Type().Out(i))
	}
	actionOutput[name] = outputs

	return reflect.MakeFunc(v.Type(), func(args []reflect.Value) []reflect.Value {
		argsKeys := make([]string, 0, len(args))
		for i, arg := range args {
			argsKeys = append(argsKeys, fmt.Sprintf("%d:%s:%v", i, arg.Type().String(), arg))
		}
		argsKey := strings.Join(argsKeys, ":")

		cachesRWMutex.RLock()
		cache, exist := caches[name][argsKey]
		if exist {
			cachesRWMutex.RUnlock()
			return cache
		}
		cachesRWMutex.RUnlock()

		results := v.Call(args)

		cachesRWMutex.Lock()
		defer cachesRWMutex.Unlock()
		funcCache, exist := caches[name]
		if !exist {
			funcCache = make(map[string][]reflect.Value)
			caches[name] = funcCache
		}
		cache, exist = funcCache[argsKey]
		if !exist {
			funcCache[argsKey] = results
		}
		return results
	}).Interface().(Func)
}
