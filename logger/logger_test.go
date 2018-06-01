package logger

import "testing"

func TestPrint(t *testing.T) {
	Debug("Debug消息:%s", "hhe")
	Info("Info消息")
	Error("Error消息")
}
