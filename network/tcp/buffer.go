package tcp

import (
	"net"
	"time"
)

type reader struct {
	start, end int
	buffer     []byte
}

func newReaderWith(size int) *reader {
	return &reader{
		start:  0,
		end:    0,
		buffer: make([]byte, size),
	}
}

func (r *reader) tryUpdate(session *Session) {
	if r.end-r.start == 0 {
		session.conn.SetReadDeadline(time.Now().Add(readTimeout))
		n, err := session.conn.Read(r.buffer)
		if err != nil {
			if nErr, ok := err.(net.Error); ok && nErr.Timeout() {
			} else {
				session.isQuit = true
			}
		} else {
			r.end = n
			r.start = 0
		}
	}
}

// read 用于填充缓冲区
func (r *reader) read(start int, end int, container []byte) int {
	readLength := 0
	dataLength := r.end - r.start
	containerLength := end - start
	if dataLength > 0 {
		if dataLength < containerLength {
			readLength = dataLength
		} else {
			readLength = containerLength
		}
		for i := 0; i < readLength; i++ {
			container[start] = r.buffer[r.start]
			start++
			r.start++
		}
	}
	return readLength
}
