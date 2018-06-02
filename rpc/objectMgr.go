package rpc

import (
	"gobase/pool"
)

var (
	objectMgr *pool.Manager
)

func init() {
	objectMgr = pool.NewManager()
}
