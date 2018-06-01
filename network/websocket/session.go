package websocket

/**************************************************************************************
Code Description    : session
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

import (
	"io"
	"net"
	"runtime"
	"time"

	"golang.org/x/net/websocket"
)

var (
	nextDeadIntervalTime = time.Millisecond * 1000 * 15
	castQueueCap         = 16
	exitConn             int32
)

type (
	// Session 用于描述会话
	Session struct {
		AttachData   *message.HotelSessionData
		castQueue    chan []byte
		nextDeadTime time.Time
		conn         *websocket.Conn
		isExit       bool
		isMaxConnect bool
	}
)

func newSessionWith(c *websocket.Conn) *Session {
	pSession := &Session{
		AttachData:   nil,
		castQueue:    make(chan []byte, castQueueCap),
		nextDeadTime: time.Now().Add(nextDeadIntervalTime),
		conn:         c,
		isExit:       false,
		isMaxConnect: false,
	}
	return pSession
}

// GetRemoteAddr 用于获取远程连接的地址
func (s *Session) GetRemoteAddr() string {
	return s.conn.RemoteAddr().String()
}

// Logout 登出
func (s *Session) Logout() {
	s.isExit = true
}

// SendMessage 用于发送消息
func (s *Session) SendMessage(data []byte) {
	if _, err := s.conn.Write(data); err != nil {
		s.isExit = true
		return
	}
	s.nextDeadTime = time.Now().Add(GetServer().GetEchoConnectMaxIdleTimeout())
}

// CastMessage 用于广播消息
func (s *Session) CastMessage(data []byte) {
	if len(s.castQueue) == castQueueCap {
		return
	}
	s.castQueue <- data
}

// runSession 用于运行回话
func runSession(c *websocket.Conn) {
	session := newSessionWith(c)
	buffer := make([]byte, 1024)
	var length int
	var err error
	var data []byte
	if !GetServer().IncConnectCount() {
		GetServer().navigate.maxConnectArrivedNavigate(session)
		session.isMaxConnect = true
		goto end
	}
	GetServer().navigate.openConnectNavigate(session)
	for {
		if session.isExit {
			goto end
		}
		select {
		case data = <-session.castQueue:
			session.SendMessage(data)
		default:
			session.conn.SetReadDeadline(time.Now().Add(GetServer().GetEchoConnectReadTimeout()))
			length, err = session.conn.Read(buffer)
			if err != nil {
				if err == io.EOF {
					goto end
				}
				if nErr, ok := err.(net.Error); ok && nErr.Timeout() {
					GetServer().navigate.pingNavigate(session)
				}
			} else {
				GetServer().navigate.Execute(session, buffer[0:length])
				time.Now().Add(GetServer().GetEchoConnectMaxIdleTimeout())
			}
		}
		if time.Now().Unix() > session.nextDeadTime.Unix() {
			goto end
		}
		time.Sleep(time.Millisecond)
		runtime.Gosched()
	}
end:
	if session.isMaxConnect {
		GetServer().IncConnectCount()
	}
	GetServer().navigate.closeConnectNavigate(session)
	session.conn.Close()
}
