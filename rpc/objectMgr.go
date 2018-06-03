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
	objectMgr.Register(clientModel, createClient)
}

func createClient() interface{} {
	return newClient()
}
func getClient() *client {
	return objectMgr.Get(clientModel).(*client)
}
func putClient(c *client) {
	objectMgr.Put(clientModel, c)
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
