/**************************************************************************************
Code Description    : 路由
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

package router

var (
	slotSize = 64
)

// FunctionCallBack 用于与行为码对应的执行函数
type FunctionCallBack func([]byte) []byte

// Router 用于描述一个路由
type Router struct {
	value int
	store []map[int]FunctionCallBack
}

// NewRouter 用于新建一个路由
func NewRouter() *Router {
	pRouter := &Router{
		value: 1,
		store: make([]map[int]FunctionCallBack, 1),
	}
	for i := 0; i < pRouter.value; i++ {
		pRouter.store[i] = make(map[int]FunctionCallBack, slotSize)
	}
	return pRouter
}

// AddAction 用于添加行为码
func (r *Router) AddAction(code int, function FunctionCallBack) {
	r.store[code%r.value][code] = function
	if len(r.store[code%r.value]) >= slotSize {
		r.value++
		newStore := make([]map[int]FunctionCallBack, r.value)
		for i := 0; i < r.value; i++ {
			newStore[i] = make(map[int]FunctionCallBack)
		}

		for _, slotMap := range r.store {
			for actionCode, actionFunction := range slotMap {
				newStore[actionCode%r.value][actionCode] = actionFunction
			}
		}
		r.store = newStore
	}
}
