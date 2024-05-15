package common

import (
	"fmt"
	"github.com/kercylan98/minotaur/vivid"
	"net"
)

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
