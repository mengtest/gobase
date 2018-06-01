package websocket

import (
	"bytes"
	"encoding/binary"
)

// packet 用于打包
func packet(code int, response []byte) []byte {
	codeBytes := intToBytes(code)
	lengthBytes := intToBytes(len(codeBytes) + len(response))
	bytesBuffer := bytes.NewBuffer([]byte{})
	bytesBuffer.Write(lengthBytes)
	bytesBuffer.Write(codeBytes)
	bytesBuffer.Write(response)
	return bytesBuffer.Bytes()
}

func unPacket(request []byte) (int, []byte) {
	code := bytesToInt(request[0:4])
	return code, request[4:]
}

//字节转换成整形
func bytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)
	var tmp int32
	binary.Read(bytesBuffer, binary.BigEndian, &tmp)
	return int(tmp)
}

func intToBytes(n int) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, n)
	return bytesBuffer.Bytes()
}
