/**************************************************************************************
Code Description    : rpc 服务
Code Vesion         :
					|------------------------------------------------------------|
						  Version    					Editor            Time
							1.0        					yuansudong        2016.4.12
					|------------------------------------------------------------|
Version Description	:
                    |------------------------------------------------------------|
						  Version
							1.0
								 ....
					|------------------------------------------------------------|
***************************************************************************************/

package pull

import (
	"runtime"
	"sync/atomic"
	"time"

	"github.com/go-mangos/mangos"
	"github.com/go-mangos/mangos/protocol/pull"
	"github.com/go-mangos/mangos/transport/tcp"
)

var (
	workIdleTime time.Duration
)

// Server 用于描述一个RPC服务
type Server struct {
	socket      mangos.Socket
	leaderCount int32
	leaders     []*leader
}

// NewServerWith 用于新建一个server
func NewServerWith(url string) *Server {
	pServer := &Server{}
	pServer.leaderCount = 0
	pServer.leaders = make([]*leader, runtime.NumCPU())
	for i := 0; i < runtime.NumCPU(); i++ {
		pServer.leaders[i] = newLeader(512, time.Millisecond)
	}
	var err error
	pServer.socket, err = pull.NewSocket()
	if err != nil {
		panic("NewSocket happen error , because of " + err.Error())
	}
	if err = pServer.socket.SetOption(mangos.OptionRaw, true); err != nil {
		panic("SetOption raw happen error , because of " + err.Error())
	}
	if err = pServer.socket.SetOption(mangos.OptionRecvDeadline, time.Millisecond); err != nil {
		panic("SetOption recvdeadline happen error , because of " + err.Error())
	}
	pServer.socket.AddTransport(tcp.NewTransport())
	if err = pServer.socket.Listen(url); err != nil {
		panic("listen happen error , because of " + err.Error())
	}
	return pServer
}

// RegisterService 用于注册服务
func (s *Server) RegisterService(i interface{}) {
	servicesMgr.Register(i)
}

// Run 用于运行
func (s *Server) Run() {
	for i := 0; i < runtime.NumCPU(); i++ {
		s.leaders[i].Run(s)
	}
}

// Close 用于关闭套接字
func (s *Server) Close() {
	s.socket.Close()
	for i := 0; i < runtime.NumCPU(); i++ {
		s.leaders[i].isQuit = true
	}
	for {
		if atomic.LoadInt32(&s.leaderCount) == 0 {
			break
		}
		time.Sleep(time.Millisecond)
		runtime.Gosched()
	}
}
