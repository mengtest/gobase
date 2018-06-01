package idle

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
