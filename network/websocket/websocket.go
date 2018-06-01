package websocket

import (
	"fmt"
	"net/http"
	"runtime"
	"time"

	"golang.org/x/net/websocket"
)

func server(Conn *websocket.Conn) {
	fmt.Println("有新的连接都来")
	fmt.Println("当前协程的运行总量是:", runtime.NumGoroutine())
	Conn.Write([]byte(time.Now().String()))
	time.Sleep(1000 * time.Millisecond)
	Conn.Close()
}

func main() {

	http.Handle("/ws/connect", websocket.Handler(server))
	http.ListenAndServe(":8080", nil)

}
