package server_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/server"
	"time"
)

// 该案例将创建一个简单的 WebSocket 服务器，如果需要更多的服务器类型可参考 [` Network `](#struct_Network) 部分
//   - server.WithLimitLife(time.Millisecond) 通常不是在正常开发应该使用的，在这里只是为了让服务器在启动完成后的 1 毫秒后自动关闭
//
// 该案例的输出结果为 true
func ExampleNew() {
	srv := server.New(server.NetworkWebsocket, server.WithLimitLife(time.Millisecond))
	fmt.Println(srv != nil)
	// Output:
	// true
}

// 该案例将创建两个不同类型的服务器，其中 WebSocket 是一个 Socket 服务器，而 Http 是一个非 Socket 服务器
//
// 可知案例输出结果为：
//   - true
//   - false
func ExampleServer_IsSocket() {
	srv1 := server.New(server.NetworkWebsocket)
	fmt.Println(srv1.IsSocket())
	srv2 := server.New(server.NetworkHttp)
	fmt.Println(srv2.IsSocket())
	// Output:
	// true
	// false
}

// 该案例将创建一个简单的 WebSocket 服务器并启动监听 `:9999/` 作为 WebSocket 监听地址，如果需要更多的服务器类型可参考 [` Network `](#struct_Network) 部分
//   - 当服务器启动失败后，将会返回错误信息并触发 panic
//   - server.WithLimitLife(time.Millisecond) 通常不是在正常开发应该使用的，在这里只是为了让服务器在启动完成后的 1 毫秒后自动关闭
func ExampleServer_Run() {
	srv := server.New(server.NetworkWebsocket, server.WithLimitLife(time.Millisecond))
	if err := srv.Run(":9999"); err != nil {
		panic(err)
	}
	// Output:
}

// RunNone 函数并没有特殊的意义，该函数内部调用了 `srv.run("")` 函数，仅是一个语法糖，用来表示服务器不需要监听任何地址
func ExampleServer_RunNone() {
	srv := server.New(server.NetworkNone)
	if err := srv.RunNone(); err != nil {
		panic(err)
	}
	// Output:
}
