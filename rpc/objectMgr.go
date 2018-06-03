package rpc

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
	return newPakcet()
}
func getPacket() *packet {
	return objectMgr.Get(packetModel).(*packet)
}
func putPacket(p *packet) {
	objectMgr.Put(packetModel, p)
}
