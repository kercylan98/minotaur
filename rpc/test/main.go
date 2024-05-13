package main

import (
	"fmt"
	"github.com/kercylan98/minotaur/rpc"
	"github.com/kercylan98/minotaur/rpc/client"
	"github.com/kercylan98/minotaur/rpc/codec"
	"github.com/kercylan98/minotaur/rpc/transporter"
	"net"
	"time"
)

func main() {
	c := codec.NewJSON()
	r := rpc.NewRouter()
	r.Register("/println", func(ctx rpc.Context) error {
		var str string
		ctx.MustReadTo(&str)
		fmt.Println(str)
		return nil
	})
	r.Register("/echo", func(ctx rpc.Context) error {
		var str string
		ctx.MustReadTo(&str)
		ctx.Reply(str)
		return nil
	})

	srv := rpc.NewServer(
		transporter.NewGoRPC(),
		r,
		c,
	)

	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	go func() {
		if err = srv.Serve(l); err != nil {
			panic(err)
		}
	}()

	cli, err := client.NewGoRPC("tcp", "127.0.0.1:1234", c)
	if err != nil {
		panic(err)
	}

	cli.AsyncTell("/println", "Hello, World!", func(err error) {
		if err != nil {
			panic(err)
		}
	})

	if resp, err := cli.Ask("/echo", "Hello, Echo!"); err != nil {
		panic(err)
	} else {
		var str string
		resp.MustReadTo(&str)
		fmt.Println(str)
	}

	time.Sleep(time.Second * 123123213)
}
