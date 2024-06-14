package minotaur

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/pprof"
	"sync"
)

var (
	httpPProf      map[string]*http.Server // HTTP PProf 服务器
	httpPProfMutex sync.Mutex              // HTTP PProf 服务器互斥锁
)

// EnableHttpPProf 设置启用 HTTP PProf
//   - 该函数支持运行时调用且支持重复调用，重复调用不会重复开启
func EnableHttpPProf(addr, prefix string, errorHandler func(err error)) {
	httpPProfMutex.Lock()
	defer httpPProfMutex.Unlock()

	_, exist := httpPProf[addr]
	if exist {
		return
	}

	if httpPProf == nil {
		httpPProf = make(map[string]*http.Server)
	}

	mux := http.NewServeMux()
	mux.HandleFunc(fmt.Sprintf("GET %s/", prefix), pprof.Index)
	mux.HandleFunc(fmt.Sprintf("GET %s/heap", prefix), pprof.Handler("heap").ServeHTTP)
	mux.HandleFunc(fmt.Sprintf("GET %s/goroutine", prefix), pprof.Handler("goroutine").ServeHTTP)
	mux.HandleFunc(fmt.Sprintf("GET %s/block", prefix), pprof.Handler("block").ServeHTTP)
	mux.HandleFunc(fmt.Sprintf("GET %s/threadcreate", prefix), pprof.Handler("threadcreate").ServeHTTP)
	mux.HandleFunc(fmt.Sprintf("GET %s/cmdline", prefix), pprof.Cmdline)
	mux.HandleFunc(fmt.Sprintf("GET %s/profile", prefix), pprof.Profile)
	mux.HandleFunc(fmt.Sprintf("GET %s/symbol", prefix), pprof.Symbol)
	mux.HandleFunc(fmt.Sprintf("POST %s/symbol", prefix), pprof.Symbol)
	mux.HandleFunc(fmt.Sprintf("GET %s/trace", prefix), pprof.Trace)
	mux.HandleFunc(fmt.Sprintf("GET %s/mutex", prefix), pprof.Handler("mutex").ServeHTTP)
	srv := &http.Server{Addr: addr, Handler: mux}
	httpPProf[addr] = srv
	go func(srv *http.Server, errHandler func(err error)) {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errorHandler(err)
		}
	}(srv, errorHandler)
}

// DisableHttpPProf 设置禁用 HTTP PProf
//   - 当 HTTP PProf 未启用时不会产生任何效果
//   - 该函数支持运行时调用且支持重复调用，重复调用不会重复禁用
func DisableHttpPProf(addr string) {
	httpPProfMutex.Lock()
	defer httpPProfMutex.Unlock()
	p, exist := httpPProf[addr]
	if !exist {
		return
	}
	delete(httpPProf, addr)

	_ = p.Close()
	httpPProf = nil
}
