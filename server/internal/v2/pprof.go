package server

import (
	"fmt"
	"net/http"
	"net/http/pprof"
	"sync"
)

var (
	httpPProf      *http.Server // HTTP PProf 服务器
	httpPProfMutex sync.Mutex   // HTTP PProf 服务器互斥锁
)

// EnableHttpPProf 设置启用 http pprof
//   - 该函数支持运行时调用
func EnableHttpPProf(addr, prefix string, errorHandler func(err error)) {
	httpPProfMutex.Lock()
	defer httpPProfMutex.Unlock()
	if httpPProf != nil {
		return
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
	httpPProf = srv
	go func(srv *http.Server, errHandler func(err error)) {
		if err := srv.ListenAndServe(); err != nil {
			errorHandler(err)
		}
	}(srv, errorHandler)
}

// DisableHttpPProf 设置禁用 http pprof
//   - 该函数支持运行时调用
func DisableHttpPProf() {
	httpPProfMutex.Lock()
	defer httpPProfMutex.Unlock()
	if httpPProf == nil {
		return
	}

	_ = httpPProf.Close()
	httpPProf = nil
}
