package client

import "net"

func dial(network string, addr string, runState chan<- error, receive func(wst int, packet []byte), setConn func(conn net.Conn), isClosed func() bool) {
	c, err := net.Dial(network, addr)
	if err != nil {
		runState <- err
		return
	}
	setConn(c)
	runState <- nil
	packet := make([]byte, 1024)
	for !isClosed() {
		n, readErr := c.Read(packet)
		if readErr != nil {
			panic(readErr)
		}
		receive(0, packet[:n])
	}
}
