package pull

import "gobase/service"

var (
	packetModel string
	clientModel string
)

func init() {
	clientModel = "client"
}

var (
	servicesMgr *service.Service
)

func init() {
	servicesMgr = service.NewService()
}
