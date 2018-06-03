package service

var (
	slotMask       = 0xff
	packetModel    string
	exceptionModel string
	errMethod      string
	unknownMethod  string
)

func init() {
	packetModel = "packet"
	exceptionModel = "exception"
	errMethod = "error"
	unknownMethod = "unknown"
}

var (
	unknownMethodDesc  []byte
	unknownServiceDesc []byte
)

func init() {
	unknownMethodDesc = []byte("未知的方法名")
	unknownServiceDesc = []byte("未知的服务名")
}
