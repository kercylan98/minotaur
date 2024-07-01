package cluster

import (
	"github.com/hashicorp/memberlist"
)

type Cluster struct {
	memberlist.Memberlist
}
