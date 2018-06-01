package util

import (
	"bytes"
	"encoding/binary"
)

// BytesToInt 用于 字节装换成整形
func BytesToInt(data []byte) int {
	var x int32
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &x)
	return int(x)
}

// IntToBytes 将int转换为字节数组
func IntToBytes(number int) []byte {
	x := int32(number)
	buf := bytes.NewBuffer(make([]byte, 4))
	binary.Write(buf, binary.BigEndian, x)
	return buf.Bytes()
}
