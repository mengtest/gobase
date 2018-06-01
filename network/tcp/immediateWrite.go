package tcp

import (
	"base/util"
	"net"
	"time"
)

var (
	singlePacketSize = 1024
)

// immediateWrite 用于写缓冲区
type immediateWrite struct {
	start, end int
	data       []byte
}

// newImmediateWrite 用于实例化一个新的即使写缓冲区
func newImmediateWrite() *immediateWrite {
	return &immediateWrite{
		start: 0,
		end:   0,
		data:  nil,
	}
}

func (iw *immediateWrite) reset() {
	if iw.start == iw.end && iw.start != 0 {
		iw.start = 0
		iw.end = 0
		iw.data = nil
	}
}

// isCanUse 用于检测当前是否可以写入新的数据
func (iw *immediateWrite) isCanUse() bool {
	return iw.data == nil
}

func (iw *immediateWrite) write(d []byte) {
	headBytes := util.IntToBytes(len(d))
	packetBytes := make([]byte, 0, len(headBytes)+len(d))
	packetBytes = append(append(packetBytes, headBytes...), d...)
	iw.start = 0
	iw.end = len(packetBytes)
	iw.data = packetBytes
}

func (iw *immediateWrite) update(session *Session) {
	if iw.data != nil {
		session.conn.SetWriteDeadline(time.Now().Add(writeTimeout))
		n, err := session.conn.Write(iw.data[iw.start:iw.end])
		if err != nil {
			if nErr, ok := err.(net.Error); ok && nErr.Timeout() {
			} else {
				session.isQuit = true
			}
		} else {
			iw.start += n
			iw.reset()
			session.nextDeadTime = time.Now().Add(sessionAliveTime)
		}
	}
}
