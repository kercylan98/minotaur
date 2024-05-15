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
	srv := server.NewServer(network.Tcp(":1244"))
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
	vivid.RegisterBehavior(ctx, u.onHello)
	vivid.RegisterBehavior(ctx, u.onEcho)
	return nil
}

func (u *UserActor) onHello(ctx vivid.MessageContext, msg string) error {
	fmt.Println(msg)
	return nil
}

func (u *UserActor) onEcho(ctx vivid.MessageContext, msg int) error {
	fmt.Println("ECHO", msg)
	return ctx.Reply(msg)
}

func main() {
	system := vivid.NewActorSystem("User", vivid.NewActorSystemOptions().
		WithAddress(NewServer(), "127.0.0.1", 1244).
		WithClientFactory(NewClient),
	)
	go func() {
		if err := system.Run(); err != nil {
			panic(err)
		}
	}()
	time.Sleep(time.Second)

	localActor, err := vivid.ActorOf[*UserActor](system)
	if err != nil {
		panic(err)
	}

	remoteActor, err := system.GetActor(vivid.NewActorId("tcp", "", "127.0.0.1", 1245, "User", "User1"))
	if err != nil {
		panic(err)
	}

	if err = remoteActor.Tell("Hello, World!"); err != nil {
		panic(err)
	}

	if reply, err := remoteActor.Ask(10086); err != nil {
		panic(err)
	} else {
		fmt.Println("remote reply:", reply)
	}

	if err = localActor.Tell("local: Hello, World!"); err != nil {
		panic(err)
	}

	if reply, err := localActor.Ask(9999); err != nil {
		panic(err)
	} else {
		fmt.Println("local reply:", reply)
	}

	time.Sleep(time.Minute * 10)
}
