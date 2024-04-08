package server

func newProducer(srv Server, conn Conn) Producer {
	return Producer{
		srv:  srv,
		conn: conn,
	}
}

type Producer struct {
	srv  Server
	conn Conn
}

func (p Producer) GetServer() Server {
	return p.srv
}

func (p Producer) GetConn() (conn Conn, exist bool) {
	return p.conn, p.conn != nil
}
