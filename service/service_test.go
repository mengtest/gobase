package service

import (
	"fmt"
	"testing"
)

type Test1 struct {
}

// Hello 用于测试
func (t *Test1) Hello() {
	fmt.Println("hello world")
}

func TestSingle(t *testing.T) {
	h := newSingleWith(&Test1{})
	h.execute("Hello")
}
