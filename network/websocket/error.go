package websocket

import "errors"

var (
	// ErrPipeFull 管道满了
	ErrPipeFull error
)

func init() {
	ErrPipeFull = errors.New("Pipe is Full")
}
