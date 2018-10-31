package HTTPGateway

import (
	"io"
	"net"
	"time"
)

var connProtocol string = "tcp"

// connection 用于描述一个连接
type connection struct {
	socket net.Conn
	next   *connection
}

// newConnection 用于创建一个连接
func newConnection() *connection {
	return &connection{
		socket: nil,
		next:   nil,
	}
}

// resetConnection 用于重置连接
func (c *connection) resetConnection() {
	c.socket = nil
	c.next = nil
}

// dialConnection 用于建立连接 tcpAddr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8000")
func (c *connection) dialConnection(tcpAddr *net.TCPAddr) (err error) {
	c.socket, err = net.DialTCP(connProtocol, nil, tcpAddr)
	return err
}

// write 用于写数据
func (c *connection) write(tcpAddr net.TCPAddr, data []byte) (err error) {
	c.socket.SetWriteDeadline(time.Now().Add(defaultWriteTimeout))
	var current, n, total int
	total = len(data)
	if n, err = c.socket.Write(data); err != nil {
		if nErr, ok := err.(net.Error); ok && nErr.Timeout() {
			goto end
		}
		if err == io.EOF {
           
		}
	}
	current += n
	if current != total {
		for {

		}
	}
end:
	return err

}
