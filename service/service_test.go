package service

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJson(t *testing.T) {
	p := newPacket()
	p.ServiceName = "yuansudong"
	data, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
}
