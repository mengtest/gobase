/**************************************************************************************
Code Description    : 会话
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

package tcp

import (
	"net"
	"time"
)

var (
	defautCastQueueSize = 32
)

// Session 用于获取
type Session struct {
	attachData   interface{}
	numberID     int64
	conn         net.Conn
	isQuit       bool
	nextDeadTime time.Time
	reader       *immediateRead
	writer       *immediateWrite
	castQueue    chan []byte
}

// newSession 用于创建一个新的Session
func newSession(numberID int64, c net.Conn) *Session {
	return &Session{
		attachData: nil,
		conn:       c,
		numberID:   numberID,
		isQuit:     false,
		reader:     newImmediateRead(),
		writer:     newImmediateWrite(),
		castQueue:  make(chan []byte, defautCastQueueSize),
	}
}

// SetAttachData 用于设置Session的附加数据
func (s *Session) SetAttachData(data interface{}) {
	s.attachData = data
}

// GetAttachData 用于获取Session的附加数据
func (s *Session) GetAttachData() interface{} {
	return s.attachData
}

// Stop 用于停止会话
func (s *Session) stop() {
	s.isQuit = true
}

func (s *Session) cast(data []byte) {
	if len(s.castQueue) < defautCastQueueSize {
		s.castQueue <- data
	}
}

func (s *Session) closeCastQueue() {
	close(s.castQueue)
}
