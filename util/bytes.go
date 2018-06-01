/**************************************************************************************
Code Description    : 字节工具
Code Vesion         :
					|------------------------------------------------------------|
						  Version    					Editor            Time
							1.0        					yuansudong        2016.4.12
					|------------------------------------------------------------|
Version Description	:
                    |------------------------------------------------------------|
						  Version
							1.0
								 ....
					|------------------------------------------------------------|
***************************************************************************************/

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
