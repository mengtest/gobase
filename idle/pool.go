/**************************************************************************************
Code Description    : 驻留式池
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

package idle

import (
	"runtime"
	"sync"
	"time"
)

// Pool 用于描述一个连接池
type Pool struct {
	lock           sync.Mutex
	poolSize       int
	queue          *linkedList
	Create         func() interface{}
	Close          func(interface{})
	MaxConnectsNum int
	IdleTime       time.Duration
}

// NewPool 用于新建一个池子
func NewPool(create func() interface{},
	close func(interface{}),
	maxConnectsNum int,
	idleTime time.Duration) *Pool {
	return &Pool{
		poolSize:       0,
		queue:          newLinkedList(),
		Create:         create,
		Close:          close,
		MaxConnectsNum: maxConnectsNum,
		IdleTime:       idleTime,
	}

}

// Release 用于释放驻留式池
func (p *Pool) Release() {
	p.lock.Lock()
	defer p.lock.Unlock()
	var ele *element
	for {
		ele = p.queue.Pop()
		if ele == nil {
			break
		}
		p.Close(ele.data)
	}
}

// Put 用于释放一个连接
func (p *Pool) Put(connect interface{}) {
	p.lock.Lock()
	defer p.lock.Unlock()
	ele := elePool.get()
	ele.data = connect
	ele.nextDeadTime = time.Now().Add(p.IdleTime)
	ele.next = nil
	p.queue.Push(ele)
}

// Get 用于获得一个连接
func (p *Pool) Get() interface{} {
	var connect interface{}
	var isGet bool
	for {
		connect, isGet = get(p)
		if isGet {
			break
		}
		time.Sleep(time.Nanosecond)
		runtime.Gosched()
	}
	return connect
}
func get(p *Pool) (interface{}, bool) {
	var connect interface{}
	isHave := false
	p.lock.Lock()
	defer p.lock.Unlock()
	if p.poolSize < p.MaxConnectsNum {
		ele := p.queue.Pop()
		if ele != nil {
			if ele.nextDeadTime.Unix() > time.Now().Unix() {
				connect = ele.data
				isHave = true
			} else {
				p.Close(ele.data)
				p.poolSize--
			}
			ele.Reset()
			elePool.put(ele)
		} else {
			connect = p.Create()
			isHave = true
			p.poolSize++
		}
	}
	return connect, isHave
}
