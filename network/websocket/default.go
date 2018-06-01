package websocket

import (
	"fmt"
)

func defaultMaxConnectArrivedCallBack(session *Session) {
	fmt.Println("达到了最大的连接数量")
}

func defaultConnectCloseCallBack(session *Session) {
	fmt.Println("有新的连接关闭")
}

func defaultConnectOpenCallBack(session *Session) {
	fmt.Println("有新的连接开启")
}

func defaultRepeatedLoginCallBack(session *Session) {
	fmt.Println("重复登录")
}

func defaultPing(session *Session) {
	session.SendMessage([]byte("我希望你能回答我一下,如果一段时间没有回应我要关掉你了哦"))
}
func defaultUnknownCodeCallBack(session *Session, m map[string]interface{}) {
	session.SendMessage([]byte("我不认识你啊"))
}
