package HTTPGateway

import (
	"net"
	"sync"
	"time"
)

var defaultMaxPoolSize int = 512
var defaultWriteTimeout time.Duration = time.Second * 10

var connDailNotDesc []byte = []byte("连接不通")

// connPool 用于描述一个 连接池
type connPool struct {
	idleHeader *connection // header 为头
	idleTailer *connection // tail 为尾部
	idleNum    int         // 驻留池子里拥有的连接数
	connNum    int         // 已经开启的连接数
	tcpAddr    *net.TCPAddr
	poolMutex  sync.Mutex
}

// newConnPool 用于创建一个连接池
func newConnPool(addr string) *connPool {
	monitor := connMemMgrInstance.Get()
	ta, err := net.ResolveTCPAddr(connProtocol, addr)
	if err != nil {
		panic(err)
	}
	return &connPool{
		idleHeader: monitor,
		idleTailer: monitor,
		idleNum:    0,
		connNum:    0,
		tcpAddr:    ta,
	}
}

// dispath 用于派发数据下去
func (cp *connPool) dispatch(dispatchData []byte) (returnData []byte) {
	var conn *connection
	var isDailNot bool
	for {
		conn, isDailNot = cp.getConn()
		if isDailNot {
			returnData = connDailNotDesc
			break
		}
		if conn != nil {
			conn.socket.Write(dispatchData)
		}
	}
	return returnData
}

func (cp *connPool) getConn() (conn *connection, isDailNot bool) {
	cp.poolMutex.Lock()
	defer cp.poolMutex.Unlock()
	if cp.idleNum == 0 {
		if cp.connNum < defaultMaxPoolSize {
			conn = connMemMgrInstance.Get()
			if conn.dialConnection(cp.tcpAddr) != nil {
				isDailNot = true
			}
		}
	} else {
		conn = cp.idleHeader.next
		cp.idleHeader.next = conn.next
		cp.idleNum--
		if cp.idleNum == 0 {
			cp.idleTailer = cp.idleHeader
		}
	}
	return conn, isDailNot
}
