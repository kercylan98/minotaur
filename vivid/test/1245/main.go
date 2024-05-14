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
		c: make(chan vivid.RemoteMessageEvent, 1024),
	}
}

type Server struct {
	c chan vivid.RemoteMessageEvent
}

func (s *Server) Run() error {
	srv := server.NewServer(network.Tcp(":1245"))
	srv.RegisterConnectionReceivePacketEvent(func(srv server.Server, conn server.Conn, packet server.Packet) {
		message, err := vivid.ParseRemoteMessage(packet.GetBytes(), func(err error) {
			fmt.Println("remote handle", err)
		})
		if err != nil {
			return
		}

		s.c <- message
	})
	return srv.Run()
}

func (s *Server) Shutdown() error {
	return nil
}

func (s *Server) C() <-chan vivid.RemoteMessageEvent {
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

func (c Client) Exec(data []byte) ([]byte, error) {
	_, err := c.tcp.Write(data)
	if err != nil {
		return nil, err
	}

	var buf = make([]byte, 1024)
	n, err := c.tcp.Read(buf)
	if err != nil {
		return nil, err
	}

	return buf[:n], nil
}

func (c Client) AsyncExec(data []byte, callback func([]byte, error)) error {
	_, err := c.tcp.Write(data)
	if err != nil {
		return err
	}

	go func() {
		var buf = make([]byte, 1024)
		n, err := c.tcp.Read(buf)
		if err != nil {
			callback(nil, err)
			return
		}

		callback(buf[:n], nil)
	}()
	return nil
}

func (u *UserActor) OnPreStart(ctx *vivid.ActorContext) error {
	u.BehaviorAutoExecutor.Init(ctx)
	vivid.RegisterBehavior(ctx, u.onHello)
	return nil
}

func (u *UserActor) onHello(ctx *vivid.ActorContext, msg string) error {
	fmt.Println(msg)
	return nil
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
