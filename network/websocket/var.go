package websocket

import "time"

const (
	defaultWriteCap      = 1024
	invalidPackageLength = -1
)

var (
	activeTimeoutTime time.Duration
	readTimeout       time.Duration
	writeTimeout      time.Duration
)

func init() {
	activeTimeoutTime = 15 * time.Millisecond * 1000
	readTimeout = time.Millisecond * 1000
	writeTimeout = time.Millisecond * 1000
}
