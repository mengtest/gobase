package common

import (
	"gobase/pool"
)

var objectPoolInstance *pool.Manager
var rpFormatObjectModel string

func init() {
	objectPoolInstance = pool.NewManager()
	rpFormatObjectModel = "RP_FORMAT"
	objectPoolInstance.Register(rpFormatObjectModel, createRPFormatObject)
}

func createRPFormatObject() interface{} {
	return newRPFormat()
}

// GetRPFormatObject 用于获取RPFormat的实例
func GetRPFormatObject() *RPFormat {
	return objectPoolInstance.Get(rpFormatObjectModel).(*RPFormat)
}

// PutRPFormatObject 用于释放RPFormat的实例
func PutRPFormatObject(obj *RPFormat) {
	objectPoolInstance.Put(rpFormatObjectModel, obj)
}
