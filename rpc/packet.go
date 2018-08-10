package rpc

import "nanomsg.org/go-mangos"

type packet struct {
	msg    *mangos.Message
	socket mangos.Socket
}

func newPakcet() *packet {
	return &packet{}
}
