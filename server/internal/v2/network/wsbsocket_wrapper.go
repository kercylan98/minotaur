package network

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/kercylan98/minotaur/utils/super"
	"github.com/panjf2000/gnet/v2"
	"io"
	"time"
)

// newWebsocketWrapper 创建 websocket 包装器
func newWebsocketWrapper(conn gnet.Conn) *websocketWrapper {
	wrapper := &websocketWrapper{
		conn:     conn,
		upgraded: false,
		active:   time.Now(),
	}
	return wrapper
}

// websocketWrapper websocket 包装器
type websocketWrapper struct {
	conn     gnet.Conn    // 连接
	upgraded bool         // 是否已经升级
	hs       ws.Handshake // 握手信息
	active   time.Time    // 活跃时间
	buf      bytes.Buffer // 缓冲区

	header *ws.Header   // 当前头部
	cache  bytes.Buffer // 缓存的数据
}

// readToBuffer 将数据读取到缓冲区
func (w *websocketWrapper) readToBuffer() error {
	size := w.conn.InboundBuffered()
	buf := make([]byte, size)
	n, err := w.conn.Read(buf)
	if err != nil {
		return err
	}
	if n < size {
		return fmt.Errorf("incomplete data, read buffer bytes failed! size: %d, read: %d", size, n)
	}
	w.buf.Write(buf)
	return nil
}

// upgrade 升级
func (w *websocketWrapper) upgrade(upgrader ws.Upgrader) (err error) {
	if w.upgraded {
		return
	}

	buf := &w.buf
	reader := bytes.NewReader(buf.Bytes())
	n := reader.Len()

	w.hs, err = upgrader.Upgrade(super.ReadWriter{
		Reader: reader,
		Writer: w.conn,
	})
	skip := n - reader.Len()
	if err != nil {
		if err == io.EOF || errors.Is(err, io.ErrUnexpectedEOF) { //数据不完整，不跳过 buf 中的 skipN 字节（此时 buf 中存放的仅是部分 "handshake data" bytes），下次再尝试读取
			return
		}
		buf.Next(skip)
		return err
	}
	buf.Next(skip)
	w.upgraded = true
	return
}

// decode 解码
func (w *websocketWrapper) decode() (messages []wsutil.Message, err error) {
	if messages, err = w.read(); err != nil {
		return
	}
	var result = make([]wsutil.Message, 0, len(messages))
	for _, message := range messages {
		if message.OpCode.IsControl() {
			err = wsutil.HandleClientControlMessage(w.conn, message)
			if err != nil {
				return
			}
			continue
		}
		if message.OpCode == ws.OpText || message.OpCode == ws.OpBinary {
			result = append(result, message)
		}
	}
	return result, nil
}

// decode 解码
func (w *websocketWrapper) read() (messages []wsutil.Message, err error) {
	var buf = &w.buf
	for {
		// 从缓冲区中读取 header 信息并写入到缓存中
		if w.header == nil {
			if buf.Len() < ws.MinHeaderSize {
				return // 不完整的数据，不做处理
			}
			var header ws.Header
			if buf.Len() >= ws.MaxHeaderSize {
				header, err = ws.ReadHeader(buf)
				if err != nil {
					return
				}
			} else {
				// 使用新的 reader 尝试读取 header，避免 header 不完整
				reader := bytes.NewReader(buf.Bytes())
				currLen := reader.Len()
				header, err = ws.ReadHeader(reader)
				skip := currLen - reader.Len()
				if err != nil {
					if errors.Is(err, io.EOF) || errors.Is(err, io.ErrUnexpectedEOF) {
						return messages, nil
					}
					buf.Next(skip)
					return nil, err
				}
				buf.Next(skip)
			}

			w.header = &header
			if err = ws.WriteHeader(&w.cache, header); err != nil {
				return nil, err
			}
		}

		// 将缓冲区数据读出并写入缓存
		if n := int(w.header.Length); n > 0 {
			if buf.Len() < n {
				return // 不完整的数据，不做处理
			}

			if _, err = io.CopyN(&w.cache, buf, int64(n)); err != nil {
				return
			}
		}

		// 消息已完整，读取数据帧，否则数据将被分割为多个数据帧
		if w.header.Fin {
			messages, err = wsutil.ReadClientMessage(&w.cache, messages)
			if err != nil {
				return
			}
			w.cache.Reset()
		}
		w.header = nil
	}
}
