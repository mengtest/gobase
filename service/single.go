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
func (s *Single) Execute(functionName string) {
	if function, ok := s.action[functionName]; ok {
		in := make([]reflect.Value, 0)
		function.Call(in)
	}
}
