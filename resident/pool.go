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

// Pool 用于一个池子
type Pool struct {
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
