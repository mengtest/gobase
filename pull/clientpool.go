package pull

import (
	"runtime"
	"sync"
	"time"
)

// ClientPool 用于描述一个连接池
type ClientPool struct {
	lock           sync.Mutex
	poolSize       int
	cm             *clientMgr
	MaxConnectsNum int
	IdleTime       time.Duration
	addr           string
}

// NewClientPool 用于新建一个池子
func NewClientPool(addr string, maxConnectsNum int, idleTime time.Duration) *ClientPool {
	return &ClientPool{
		poolSize:       0,
		cm:             newClientMgr(),
		MaxConnectsNum: maxConnectsNum,
		IdleTime:       idleTime,
		addr:           addr,
	}

}

// Release 用于释放驻留式池
func (cp *ClientPool) Release() {
	cp.lock.Lock()
	defer cp.lock.Unlock()
	var c *client
	for {
		c = cp.cm.pop()
		if c == nil {
			break
		}
		c.reset()
		putClient(c)
	}
}

// Put 用于释放一个连接
func (cp *ClientPool) put(c *client) {
	cp.lock.Lock()
	defer cp.lock.Unlock()
	c.nextDeadTime = time.Now().Add(cp.IdleTime)
	cp.cm.push(c)
}

// Get 用于获得一个连接
func (cp *ClientPool) get() *client {
	var connect *client
	var isGet bool
	for {
		connect, isGet = getClientLogic(cp)
		if isGet {
			break
		}
		time.Sleep(time.Nanosecond)
		runtime.Gosched()
	}
	return connect
}
func getClientLogic(cp *ClientPool) (*client, bool) {
	var connect *client
	isHave := false
	cp.lock.Lock()
	defer cp.lock.Unlock()
	if cp.poolSize < cp.MaxConnectsNum {
		connect = cp.cm.pop()
		if connect != nil {
			if connect.nextDeadTime.Unix() > time.Now().Unix() {
				isHave = true
			} else {
				cp.poolSize--
				connect.reset()
				putClient(connect)
			}
		} else {
			connect = getClient()
			connect.build(cp)
			isHave = true
			cp.poolSize++
		}
	}
	return connect, isHave
}

// Cast 用于进行调用
func (cp *ClientPool) Cast(data []byte) error {
	c := cp.get()
	defer cp.put(c)
	return c.socket.Send(data)
}
