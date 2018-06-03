package pull

import (
	"gobase/pool"
)

var (
	objectMgr *pool.Manager
)

func init() {
	objectMgr = pool.NewManager()
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
