package pull

import (
	"time"

	"github.com/go-mangos/mangos"
	"github.com/go-mangos/mangos/protocol/push"
	"github.com/go-mangos/mangos/transport/tcp"
)

type client struct {
	socket       mangos.Socket
	nextDeadTime time.Time
	next         *client
}

// newClient 用于新建一个客户端
func newClient() *client {
	pClient := &client{next: nil}
	pClient.socket, _ = push.NewSocket()
	pClient.socket.AddTransport(tcp.NewTransport())
	return pClient
}

func (c *client) build(cp *ClientPool) error {
	c.socket.SetOption(mangos.OptionReconnectTime, time.Millisecond)
	c.socket.SetOption(mangos.OptionSendDeadline, time.Second)
	return c.socket.Dial(cp.addr)
}

func (c *client) reset() {
	c.socket.Close()
}
