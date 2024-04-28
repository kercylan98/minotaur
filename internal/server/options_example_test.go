package server_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/server"
	"github.com/kercylan98/minotaur/utils/times"
	"time"
)

// 服务器在启动时将阻塞 1s，模拟了慢消息的过程，这时候如果通过 RegMessageLowExecEvent 函数注册过慢消息事件，将会收到该事件的消息
//   - 该示例中，将在收到慢消息时关闭服务器
func ExampleWithLowMessageDuration() {
	srv := server.New(server.NetworkNone,
		server.WithLowMessageDuration(time.Second),
	)
	srv.RegStartFinishEvent(func(srv *server.Server) {
		time.Sleep(time.Second)
	})
	srv.RegMessageLowExecEvent(func(srv *server.Server, message *server.Message, cost time.Duration) {
		srv.Shutdown()
		fmt.Println(times.GetSecond(cost))
	})
	if err := srv.RunNone(); err != nil {
		panic(err)
	}
	// Output:
	// 1
}

// 服务器在启动时将发布一条阻塞 1s 的异步消息，模拟了慢消息的过程，这时候如果通过 RegMessageLowExecEvent 函数注册过慢消息事件，将会收到该事件的消息
//   - 该示例中，将在收到慢消息时关闭服务器
func ExampleWithAsyncLowMessageDuration() {
	srv := server.New(server.NetworkNone,
		server.WithAsyncLowMessageDuration(time.Second),
	)
	srv.RegStartFinishEvent(func(srv *server.Server) {
		srv.PushAsyncMessage(func() error {
			time.Sleep(time.Second)
			return nil
		}, nil)
	})
	srv.RegMessageLowExecEvent(func(srv *server.Server, message *server.Message, cost time.Duration) {
		srv.Shutdown()
		fmt.Println(times.GetSecond(cost))
	})
	if err := srv.RunNone(); err != nil {
		panic(err)
	}
	// Output:
	// 1
}
