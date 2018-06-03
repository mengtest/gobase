/**************************************************************************************
Code Description    : 领导工人模式中工人
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
	"runtime"
	"time"

	"github.com/go-mangos/mangos"
)

// worker 工人
type worker struct {
	channel      chan *packet
	nextDeadTime time.Time
	next         *worker
}

func newWoker() *worker {
	pWorker := &worker{
		channel: nil,
	}
	return pWorker
}

func (w *worker) reset() {
	close(w.channel)
	w.channel = nil
}

func (w *worker) run(l *leader) {
	if w.channel == nil {
		w.channel = make(chan *packet, 1)
		go work(l, w)
	}
}

func (w *worker) dispatch(message *mangos.Message, socket mangos.Socket) {
	p := getPacket()
	p.msg = message
	p.socket = socket
	w.channel <- p
}

func work(l *leader, w *worker) {
	var task *packet
	var isClose bool
	for {
		select {
		case task, isClose = <-w.channel:
			if isClose {
				goto end
			} else {
				task.msg.Body = servicesMgr.Execute(task.msg.Body)
				task.socket.SendMsg(task.msg)
				putPacket(task)
				l.put(w)
			}
			// 处理完成后需要将该工人放回去
		default:
			time.Sleep(time.Nanosecond)
			runtime.Gosched()
		}
	}
end:
}
