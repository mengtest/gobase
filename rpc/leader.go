/**************************************************************************************
Code Description    : 领导工人模式中领导者
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

package rpc

import (
	"gobase/logger"
	"runtime"
	"sync"
	"sync/atomic"
	"time"

	"nanomsg.org/go-mangos"
)

// leader 用于描述领导者
type leader struct {
	wmIdle   *workerMgr
	wmDead   *workerMgr
	wmlock   sync.Mutex
	idleTime time.Duration
	isQuit   bool
}

func newLeader(maxWorkerCount int,
	idleTime time.Duration) *leader {
	pLeader := &leader{
		wmIdle:   newWokerMgr(),
		wmDead:   newWokerMgr(),
		idleTime: idleTime,
		isQuit:   false,
	}
	for i := 0; i < maxWorkerCount; i++ {
		pLeader.wmDead.push(newWoker())
	}
	return pLeader
}

func (l *leader) check() {
	l.wmlock.Lock()
	defer l.wmlock.Unlock()
	if w := l.wmIdle.pop(); w != nil {
		if time.Now().UnixNano() > w.nextDeadTime.UnixNano() {
			w.reset()
			l.wmDead.push(w)
		} else {
			l.wmIdle.push(w)
		}
	}
}

func (l *leader) put(w *worker) {
	l.wmlock.Lock()
	defer l.wmlock.Unlock()
	w.nextDeadTime = time.Now().Add(l.idleTime)
	l.wmIdle.push(w)
}

func (l *leader) get() *worker {
	var w *worker
	var isGet bool
	for {
		w, isGet = getLogic(l)
		if isGet {
			break
		}
		time.Sleep(time.Nanosecond)
		runtime.Gosched()
	}
	w.run(l)
	return w
}
func getLogic(l *leader) (*worker, bool) {
	var w *worker
	isGet := false
	l.wmlock.Lock()
	defer l.wmlock.Unlock()
	if w = l.wmIdle.pop(); w != nil {
		isGet = true
	} else {
		if w = l.wmDead.pop(); w != nil {
			isGet = true
		}
	}
	return w, isGet
}

func (l *leader) Run(server *Server) {
	go runLogic(server, l)
}
func runLogic(server *Server, l *leader) {
	atomic.AddInt32(&server.leaderCount, 1)
	var work *worker
	for {
		if l.isQuit {
			goto end
		}
		message, err := server.socket.RecvMsg()
		if err != nil {
			if err == mangos.ErrRecvTimeout {
			} else {
				logger.Debug("发生了位置的错误:" + err.Error())
			}
		} else {
			work = l.get()
			work.dispatch(message, server.socket)
		}
		l.check()
		time.Sleep(100)
	}
end:
	atomic.AddInt32(&server.leaderCount, -1)
}
