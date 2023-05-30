package cross

import "github.com/nats-io/nats.go"

type NatsOption func(n *Nats)

// WithNatsSubject 通过对应的主题名称创建
//   - 默认为：MINOTAUR_CROSS
func WithNatsSubject(subject string) NatsOption {
	return func(n *Nats) {
		n.subject = subject
	}
}

// WithNatsOptions 通过nats自带的可选项创建连接
func WithNatsOptions(options ...nats.Option) NatsOption {
	return func(n *Nats) {
		n.options = options
	}
}

// WithNatsConn 指定通过特定的连接创建
//   - 这将导致 WithNatsOptions 失效
func WithNatsConn(conn *nats.Conn) NatsOption {
	return func(n *Nats) {
		n.conn = conn
	}
}
