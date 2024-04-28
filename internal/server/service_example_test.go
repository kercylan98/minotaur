package server_test

import (
	"github.com/kercylan98/minotaur/server"
	"time"
)

// 这个案例中我们将 `TestService` 绑定到了 `srv` 服务器中，当服务器启动时，将会对 `TestService` 进行初始化
//
// 其中 `TestService` 的定义如下：
// ```go
//
//	type TestService struct{}
//
//	func (ts *TestService) OnInit(srv *server.Server) {
//		srv.RegStartFinishEvent(onStartFinish)
//
//		srv.RegStopEvent(func(srv *server.Server) {
//			fmt.Println("server stop")
//		})
//	}
//
//	func (ts *TestService) onStartFinish(srv *server.Server) {
//		fmt.Println("server start finish")
//	}
//
// ```
//
// 可以看出，在服务初始化时，该服务向服务器注册了启动完成事件及停止事件。这是我们推荐的编码方式，这样编码有以下好处：
//   - 具备可控制的初始化顺序，避免 init 产生的各种顺序导致的问题，如配置还未加载完成，即开始进行数据库连接等操作
//   - 可以方便的将不同的服务拆分到不同的包中进行管理
//   - 当不需要某个服务时，可以直接删除该服务的绑定，而不需要修改其他代码
//   - ...
func ExampleBindService() {
	srv := server.New(server.NetworkNone, server.WithLimitLife(time.Second))
	server.BindService(srv, new(TestService))

	if err := srv.RunNone(); err != nil {
		panic(err)
	}
	// Output:
	// server start finish
	// server stop
}
