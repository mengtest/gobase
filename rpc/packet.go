package rpc

import (
	"github.com/go-mangos/mangos"
)

type packet struct {
	msg    *mangos.Message
	socket mangos.Socket
}

func newPakcet() *packet {
	return &packet{}
}
