package service

type packet struct {
	ServiceName   string `json:"service_name"`
	ServiceMethod string `json:"service_method"`
	Data          []byte `json:"data"`
}

func newPacket() *packet {
	return &packet{}
}
