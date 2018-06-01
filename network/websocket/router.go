package websocket

import (
	"encoding/json"
	"fmt"
)

const (
	monoBucketSize = 32
)

type (
	// FuncRouterCallBack 用于表示路由的回调函数
	FuncRouterCallBack func(session *Session, request map[string]interface{})
	// FuncInnerCallBack 用于超时时,要调用的函数
	FuncInnerCallBack func(session *Session)
	// Router 用于描述一个路由
	router struct {
		hashValue                 int
		logicNavigate             []map[int]FuncRouterCallBack
		unknownNavigate           FuncRouterCallBack
		pingNavigate              FuncInnerCallBack
		openConnectNavigate       FuncInnerCallBack
		closeConnectNavigate      FuncInnerCallBack
		maxConnectArrivedNavigate FuncInnerCallBack
	}
	// ExpectionRouter 这个是出现异常码要执行的函数
)

// newRouter 用于新建并实例化
func newRouter() *router {
	pRouter := &router{
		hashValue:                 1,
		unknownNavigate:           defaultUnknownCodeCallBack,
		pingNavigate:              defaultPing,
		openConnectNavigate:       defaultConnectOpenCallBack,
		closeConnectNavigate:      defaultConnectCloseCallBack,
		maxConnectArrivedNavigate: defaultMaxConnectArrivedCallBack,
	}
	pRouter.logicNavigate = createSliceMap(pRouter.hashValue)
	return pRouter
}
func createSliceMap(count int) []map[int]FuncRouterCallBack {
	slicemap := make([]map[int]FuncRouterCallBack, count)
	for i := 0; i < len(slicemap); i++ {
		slicemap[i] = map[int]FuncRouterCallBack{}
	}
	return slicemap
}

// AddFunction code表示行为码,callBack 表示与行为码对应的回调函数
func (r *router) AddFunction(code int, callBack FuncRouterCallBack) {
	r.logicNavigate[code%r.hashValue][code] = callBack
	if len(r.logicNavigate[code%r.hashValue]) > monoBucketSize {
		r.hashValue++
		reHash(r)
	}
}

// Execute 用于执行相应的code
func (r *router) Execute(session *Session, request []byte) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	m := map[string]interface{}{}
	if err := json.Unmarshal(request, &m); err != nil {
		r.unknownNavigate(session, m)
		return
	}

	if codeInterface, ok := m["code"]; ok {
		code := int(codeInterface.(float64))
		if Func, ok := r.logicNavigate[code%r.hashValue][code]; ok {
			Func(session, m)
		} else {
			r.unknownNavigate(session, m)
		}
	}
}

func reHash(r *router) {
	newLogicNavigate := createSliceMap(r.hashValue)
	length := len(r.logicNavigate)
	for i := 0; i < length; i++ {
		for code, callBack := range r.logicNavigate[i] {
			newLogicNavigate[code%r.hashValue][code] = callBack
		}
	}
	r.logicNavigate = newLogicNavigate
}
