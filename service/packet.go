package service

type Packet struct {
	ServiceName   string `json:"service_name"`
	ServiceMethod string `json:"service_method"`
	Data          []byte `json:"data"`
}

func NewPacket() *Packet {
	return &Packet{}
}
