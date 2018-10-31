package HTTPGateway

import (
	"net/http"
)

// httpService 用于描述HTTPGateway的一个实例
type httpService struct {
}

// newHttpService 用于实例化httpService
func newHTTPService() *httpService {
	return &httpService{}
}

func (hs *httpService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Content-Type", "application/json")

}

// panicDealWith 用于处理异常
func panicDealWith() {

}
