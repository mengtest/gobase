package rpc

import "gobase/service"

var (
	packetModel string
	clientModel string
)

func init() {
	packetModel = "packet"
	clientModel = "client"
}

var (
	servicesMgr *service.Service
)

func init() {
	servicesMgr = service.NewService()
}
