package service

import (
	"gobase/pool"
)

var (
	objectMgr *pool.Manager
)

func init() {
	objectMgr = pool.NewManager()
	objectMgr.Register(packetModel, createPacket)
}
func createPacket() interface{} {
	return NewPacket()

}

// GetPacket 用于获取一个Packet 包
func GetPacket() *Packet {
	return objectMgr.Get(packetModel).(*Packet)
}

// PutPacket 用于释放一个包
func PutPacket(p *Packet) {
	objectMgr.Put(packetModel, p)
}
