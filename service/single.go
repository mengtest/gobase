/**************************************************************************************
Code Description    : 单服务
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

package service

import (
	"reflect"
)

// Single 是服务的一个实例
type Single struct {
	action map[string]reflect.Value
}

// NewSingleWith 用于新建一个单服务
func NewSingleWith(service interface{}) (string, *Single) {
	pSingle := &Single{
		action: make(map[string]reflect.Value),
	}
	value := reflect.ValueOf(service)
	if value.Kind() == reflect.Ptr {
		ValueType := value.Type()
		Num := ValueType.NumMethod()
		for i := 0; i < Num; i++ {
			FuncMethod := ValueType.Method(i)
			FuncMethodValue := value.MethodByName(FuncMethod.Name)
			c := FuncMethod.Name[0]
			if c >= 'A' && c <= 'Z' {
				pSingle.action[FuncMethod.Name] = FuncMethodValue
			}
		}
	} else {
		panic("your register service is not pointer")
	}
	return value.Elem().Type().Name(), pSingle
}

// Execute 用于执行函数
func (s *Single) Execute(functionName string, data []byte) (retData []byte) {
	defer func() {
		if err := recover(); err != nil {
			retData = err.([]byte)
		}
	}()
	retData = unknownMethodDesc
	if function, ok := s.action[functionName]; ok {
		in := make([]reflect.Value, 1)
		in[0] = reflect.ValueOf(data)
		retData = function.Call(in)[0].Bytes()
	}
	return retData
}
