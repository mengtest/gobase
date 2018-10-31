package HTTPGateway

import (
	"net/http"
	"time"
)

// Instance 用于描述HTTPGateway的一个实例
type Instance struct {
	server *http.Server
}

// newInstance 用于实例化HTTPGateway Instance
func newInstance() *Instance {
	s := &http.Server{
		Addr:           "",
		Handler:        &httpService{},
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.SetKeepAlivesEnabled(false)

	return &Instance{}
}

// RunHTTPGateway 用于运行HTTP
func (i *Instance) RunHTTPGateway(addr string) error {
	i.server.Addr = addr
	return i.server.ListenAndServe()
}

// RunHTTPSGateway 用于运行HTTPS方式的服务器
func (i *Instance) RunHTTPSGateway(addr, certFile, keyFile string) error {
	i.server.Addr = addr
	return i.server.ListenAndServeTLS(certFile, keyFile)
}
