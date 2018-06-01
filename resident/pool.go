/**************************************************************************************
Code Description    : 常驻式池
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

package resident

import (
	"runtime"
	"sync"
	"time"
)

// Pool 用于一个池子
type Pool struct {
	lock  sync.Mutex
	queue *stack
	New   func() interface{}
	Close func(interface{})
}

// NewPoolWith 用于实例化一个新的池子
func NewPoolWith(capacity int,
	new func() interface{},
	close func(interface{})) *Pool {
	pPool := &Pool{
		queue: newStack(),
		New:   new,
		Close: close,
	}
	var ele *element
	for i := 0; i < capacity; i++ {
		ele = elePool.get()
		ele.data = pPool.New()
		pPool.queue.Push(ele)
	}
	return pPool
}

// Release 用于释放池中的元素
func (p *Pool) Release() {
	var ele *element
	for {
		ele = p.queue.Pop()
		if ele == nil {
			break
		}
		p.Close(ele.data)
	}
}

// Get 用于获取一个对象
func (p *Pool) Get() interface{} {
	var data interface{}
	var isGet bool
	for {
		data, isGet = get(p)
		if isGet {
			break
		}
		time.Sleep(time.Millisecond)
		runtime.Gosched()
	}

	return data
}
func get(p *Pool) (interface{}, bool) {
	var data interface{}
	isGet := false
	p.lock.Lock()
	defer p.lock.Unlock()
	ele := p.queue.Pop()
	if ele != nil {
		data = ele.data
		isGet = true
		ele.Reset()
		elePool.put(ele)
	}
	return data, isGet
}

// Put 用于还回元素
func (p *Pool) Put(data interface{}) {
	p.lock.Lock()
	p.lock.Unlock()
	ele := elePool.get()
	ele.data = data
	p.queue.Push(ele)
}
