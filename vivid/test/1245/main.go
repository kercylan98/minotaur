package main

import (
	"fmt"
	"github.com/kercylan98/minotaur/server"
	"github.com/kercylan98/minotaur/server/network"
	"github.com/kercylan98/minotaur/vivid"
	"github.com/kercylan98/minotaur/vivid/components"
	"net"
	"time"
)

type UserActor struct {
	components.BehaviorAutoExecutor
}

func NewServer() *Server {
	return &Server{
		c: make(chan []byte, 1024),
	}
}

type Server struct {
	c chan []byte
}

func (s *Server) Run() error {
	srv := server.NewServer(network.Tcp(":1245"))
	srv.RegisterConnectionReceivePacketEvent(func(srv server.Server, conn server.Conn, packet server.Packet) {
		s.c <- packet.GetBytes()
	})
	return srv.Run()
}

func (s *Server) Shutdown() error {
	return nil
}

func (s *Server) C() <-chan []byte {
	return s.c
}

func NewClient(network, host string, port uint16) (vivid.Client, error) {
	tcp, err := net.Dial(network, fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		return nil, err
	}
	return &Client{
		tcp: tcp,
	}, nil
}

type Client struct {
	tcp net.Conn
}

func (c Client) Exec(data []byte) error {
	_, err := c.tcp.Write(data)
	return err
}

func (u *UserActor) OnPreStart(ctx vivid.ActorContext) error {
	u.BehaviorAutoExecutor.Init(ctx)
	vivid.RegisterBehavior[string](ctx, u.onHello)
	vivid.RegisterBehavior[int](ctx, u.onEcho)
	return nil
}

func (u *UserActor) onHello(ctx vivid.MessageContext) error {
	fmt.Println(ctx.GetMessage())
	return nil
}

func (u *UserActor) onEcho(ctx vivid.MessageContext) error {
	return ctx.Reply(ctx.GetMessage())
}

func main() {
	system := vivid.NewActorSystem("User", vivid.NewActorSystemOptions().
		WithAddress(NewServer(), "127.0.0.1", 1245).
		WithClientFactory(NewClient),
	)
	go func() {
		if err := system.Run(); err != nil {
			panic(err)
		}
	}()
	time.Sleep(time.Second)

	localActor, err := vivid.ActorOf[*UserActor](system, vivid.NewActorOptions().WithName("User1"))
	if err != nil {
		panic(err)
	}

	_ = localActor.Tell("Hello, World!")
	time.Sleep(time.Minute * 10)
}
