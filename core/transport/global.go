package transport

import (
	"sync"
	"sync/atomic"
)

var externalNetworkNum atomic.Int32
var externalNetworkLaunchedNum atomic.Int32
var externalNetworkOnceLaunchInfo sync.Once
