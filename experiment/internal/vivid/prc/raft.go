package prc

import (
	"github.com/hashicorp/raft"
	raftboltdb "github.com/hashicorp/raft-boltdb"
	"io"
	"net"
	"os"
	"path/filepath"
	"time"
)

// https://zhuanlan.zhihu.com/p/58048906
type raftInstance struct {
}

func newRaft() {
	raftConfig := raft.DefaultConfig()
	raftConfig.LocalID = raft.ServerID(opts.raftTCPAddress) // 或许应该使用物理地址作为集群标识

	logStore, err := raftboltdb.NewBoltStore(filepath.Join(opts.dataDir, "raft-log.bolt"))
	stableStore, err := raftboltdb.NewBoltStore(filepath.Join(opts.dataDir, "raft-stable.bolt"))

	snapshotStore, err := raft.NewFileSnapshotStore(opts.dataDir, 1, os.Stderr)
}

func newRaftTransport(opts *options) (*raft.NetworkTransport, error) {
	address, err := net.ResolveTCPAddr("tcp", opts.raftTCPAddress)
	if err != nil {
		return nil, err
	}
	transport, err := raft.NewTCPTransport(address.String(), address, 3, 10*time.Second, os.Stderr)
	if err != nil {
		return nil, err
	}
	return transport, nil
}

type fms struct {
}

func (s *fms) Apply(entry *raft.Log) any {
	entry.Type
	return nil
}

func (s *fms) Snapshot() (raft.FSMSnapshot, error) {

}

func (s *fms) Restore(rc io.ReadCloser) error {

}
