package rpc

import "gobase/service"

var (
	packetModel string
)

func init() {
	packetModel = "packet"
}

var (
	servicesMgr *service.Service
)

func init() {
	servicesMgr = service.NewService()
}
