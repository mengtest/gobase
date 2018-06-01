package tcp

import (
	"base/logger"
	"net"
	"runtime"
	"sync/atomic"
	"time"
)

var (
	readTimeout      time.Duration
	writeTimeout     time.Duration
	sessionAliveTime time.Duration
	packetHeadSize   int
	readBufferSize   int
)

func init() {
	readTimeout = time.Millisecond
	writeTimeout = time.Millisecond
	sessionAliveTime = time.Second * 60
	packetHeadSize = 4
	readBufferSize = 1024
}

// SetReadBufferSize 用于设置读缓冲区的大小
func (s *Server) SetReadBufferSize(size int) {
	readBufferSize = size
}

// SetPacketHeadSize 用于设置包头的大小
func (s *Server) SetPacketHeadSize(size int) {
	packetHeadSize = size
}

// SetSessionAliveTime 用于设置每个会话的最大存活时间
func (s *Server) SetSessionAliveTime(sat time.Duration) {
	sessionAliveTime = sat
}

// SetReadTimeout 用于设置读取超时时间
func (s *Server) SetReadTimeout(rt time.Duration) {
	readTimeout = rt
}

// SetWirteTimeout 用于设置写入的超时时间
func (s *Server) SetWirteTimeout(wt time.Duration) {
	writeTimeout = wt
}

// Server 用于描述Server的数据
type Server struct {
	isQuit      bool
	sessionMgr  *SessionMgr
	sessionNum  int64
	dealMessage func([]byte) []byte
}

// SetDealMessageCallBack 用于设置当消息到来的时候的回调函数
func (s *Server) SetDealMessageCallBack(function func([]byte) []byte) {
	s.dealMessage = function
}

// NewWith 用于创建一个Server 实例
func NewWith(addr string) *Server {
	return &Server{
		isQuit:      false,
		sessionMgr:  newSessionMgr(),
		dealMessage: defaultDealMessage,
		sessionNum:  0,
	}
}

// Stop 用于停止当前的会话
func (s *Server) Stop() {
	s.sessionMgr.notifyStop()
	s.isQuit = true
}

// Run 用于运行函数
func (s *Server) Run(addr string) error {
	runListener(s, addr)
	exit(s)
	return nil
}

func exit(s *Server) {
	for {
		if atomic.LoadInt64(&s.sessionNum) == 0 {
			break
		}
		time.Sleep(time.Millisecond * 1000)
		runtime.Gosched()
	}
}
func runListener(s *Server, addr string) error {
	listener, listenErr := net.Listen("tcp", addr)
	if listenErr != nil {
		logger.Debug("listen addr happen error:%s\n", listenErr)
		return listenErr
	}
	var acceptErr error
	var conn net.Conn
	for {
		if s.isQuit {
			goto end
		}
		conn, acceptErr = listener.Accept()
		if acceptErr != nil {
			logger.Error("接受套接字发生了错误," + acceptErr.Error())
		} else {
			go newConnect(s, conn)
		}
		runtime.Gosched()
	}
end:
	listener.Close()
	return nil
}

// newConnect 用于处理新的连接
func newConnect(server *Server, conn net.Conn) {
	session := newSession(atomic.AddInt64(&server.sessionNum, 1), conn)
	var castData, retData, contentData []byte
	var isHave bool
	for {
		if session.isQuit {
			goto end
		}
		if session.writer.isCanUse() {
			select {
			case castData = <-session.castQueue:
				session.writer.write(castData)
			default:
				contentData, isHave = session.reader.read()
				if isHave {
					// 此时需要去执行路由
					retData = server.dealMessage(contentData)
					session.writer.write(retData)
				}
			}
		}
		session.reader.update(session)
		session.writer.update(session)
		if time.Now().Unix() > session.nextDeadTime.Unix() {
			session.isQuit = true
		}
		time.Sleep(time.Nanosecond)
		runtime.Gosched()
	}
end:
	atomic.AddInt64(&server.sessionNum, -1)
	server.sessionMgr.deleteSession(session)
	session.closeCastQueue()
	session.conn.Close()
}
