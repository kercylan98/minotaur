package modular

import (
	"github.com/kercylan98/minotaur/toolkit/log"
	"sync"
)

// startLifecycle 开始生命周期
func startLifecycle(services []GlobalService) *lifecycle {
	l := &lifecycle{services: services, wait: new(sync.WaitGroup)}
	l.root = l
	return l
}

// lifecycle 生命周期函数
type lifecycle struct {
	name     string                           // 生命周期名称
	services []GlobalService                  // 生命周期服务
	handler  func(service GlobalService) bool // 生命周期处理函数
	n        *lifecycle                       // 下一个生命周期
	wait     *sync.WaitGroup                  // 等待组
	last     bool                             // 是否是最后一个生命周期
	root     *lifecycle                       // 根生命周期
	running  bool                             // 是否正在运行
}

// next 设置下一个生命周期
func (l *lifecycle) next(name string, handler func(service GlobalService) bool) *lifecycle {
	l.last = false
	l.n = &lifecycle{name: name, services: l.services, handler: handler, wait: l.wait, last: true, root: l.root}
	return l.n
}

// run 运行生命周期
func (l *lifecycle) run() {
	if !l.root.running {
		l.root.running = true
		l.root.run()
		return
	}
	var num int
	if l.handler != nil {
		for _, service := range l.services {
			l.wait.Add(1)
			if l.handler(service) {
				num++
			}
			l.wait.Done()
		}
		log.Info("modular", log.String("lifecycle", l.name), log.Int("num", num))
	}
	if l.n != nil {
		l.n.run()
	}
	if l.last {
		l.wait.Wait()
	}
}
