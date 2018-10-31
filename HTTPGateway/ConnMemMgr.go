package HTTPGateway

import (
	"sync"
)

var connMemMgrInstance *connMemMgr

// connMemMgr 用于管理连接所分配的内存
type connMemMgr struct {
	memPool sync.Pool
}

func init() {
	connMemMgrInstance = newConnMemMgr()
}

// newConnMemMgr 用于创建
func newConnMemMgr() *connMemMgr {
	return &connMemMgr{
		memPool: sync.Pool{
			New: func() interface{} {
				return newConnection()
			},
		},
	}
}

// Get 用于获取一个连接
func (cmm *connMemMgr) Get() *connection {
	return cmm.memPool.Get().(*connection)
}

// Put 用于释放一个连接
func (cmm *connMemMgr) Put(c *connection) {
	cmm.memPool.Put(c)
}
