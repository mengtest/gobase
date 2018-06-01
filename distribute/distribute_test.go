package distribute

import (
	"fmt"
	"testing"
)

func TestID(t *testing.T) {
	InitDistribute(127)
	count := 10000000
	store := make(map[int64]int, count)
	var num int
	var ok bool
	var curr int64
	for i := 0; i < count; i++ {
		curr = GetUniqueID()
		if num, ok = store[curr]; ok {
			store[curr] = num + 1
		} else {
			store[curr] = 1
		}
	}
	for id, count := range store {
		if count > 1 {
			fmt.Println("有重复的数据,", id)
		}
	}

}
