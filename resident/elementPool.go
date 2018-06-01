/**************************************************************************************
Code Description    : 元素池
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
	"sync"
)

var (
	elePool *elementPool
)

func init() {
	elePool = newElementPool()
}

type elementPool struct {
	pool sync.Pool
}

func newElementPool() *elementPool {
	return &elementPool{
		pool: sync.Pool{
			New: func() interface{} {
				return newElement()
			},
		},
	}
}

func (ep *elementPool) get() *element {
	return ep.pool.Get().(*element)
}
func (ep *elementPool) put(ele *element) {
	ep.pool.Put(ele)
}
