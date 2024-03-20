package server

import (
	"context"
	"github.com/kercylan98/minotaur/utils/log"
	"net"
	"time"
)

// connections 结构体用于管理连接
type connections struct {
	ctx   context.Context // 上下文对象，用于取消连接管理器
	ch    chan any        // 事件通道，用于接收连接管理器的操作事件
	items []*conn         // 连接列表，存储所有打开的连接
	gap   []int           // 连接空隙，记录已关闭的连接索引，用于重用索引
}

// 初始化连接管理器
func (cs *connections) init(ctx context.Context) *connections {
	cs.ctx = ctx
	cs.ch = make(chan any, 1024)
	cs.items = make([]*conn, 0, 128)
	go cs.awaitRun()
	return cs
}

// 清理连接列表中的空隙
func (cs *connections) clearGap() {
	cs.gap = cs.gap[:0]
	var gap = make([]int, 0, len(cs.items))
	for i, c := range cs.items {
		if c == nil {
			continue
		}
		c.idx = i
		gap = append(gap, i)
	}

	cs.gap = gap
}

// 打开新连接
func (cs *connections) open(c net.Conn) error {
	// 如果存在连接空隙，则重用连接空隙中的索引，否则分配新的索引
	var idx int
	var reuse bool
	if len(cs.gap) > 0 {
		idx = cs.gap[0]
		cs.gap = cs.gap[1:]
		reuse = true
	} else {
		idx = len(cs.items)
	}
	conn := new(conn).init(cs.ctx, cs, c, idx)
	if reuse {
		cs.items[idx] = conn
	} else {
		cs.items = append(cs.items, conn)
	}
	go conn.awaitRead()
	return nil
}

// 关闭连接
func (cs *connections) close(c *conn) error {
	if c == nil {
		return nil
	}
	defer c.cancel()
	// 如果连接索引是连接列表的最后一个索引，则直接删除连接对象，否则将连接对象置空，并将索引添加到连接空隙中
	if c.idx == len(cs.items)-1 {
		cs.items = cs.items[:c.idx]
	} else {
		cs.items[c.idx] = nil
		cs.gap = append(cs.gap, c.idx)
	}
	return c.Conn.Close()
}

// 等待连接管理器的事件并处理
func (cs *connections) awaitRun() {
	clearGapTicker := time.NewTicker(time.Second * 30)
	defer clearGapTicker.Stop()

	for {
		select {
		case <-cs.ctx.Done():
			return
		case <-clearGapTicker.C:
			cs.clearGap()
		case a := <-cs.ch:
			var err error

			switch v := a.(type) {
			case *conn:
				err = cs.close(v)
			case net.Conn:
				err = cs.open(v)
			}

			if err != nil {
				log.Error("connections.awaitRun", log.Any("err", err))
			}
		}
	}
}

// Event 获取连接管理器的事件通道
func (cs *connections) Event() chan<- any {
	return cs.ch
}
