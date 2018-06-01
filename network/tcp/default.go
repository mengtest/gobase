package tcp

import (
	"base/logger"
)

func defaultDealMessage(data []byte) []byte {
	logger.Debug("接受到了数据,数据的是:" + string(data))
	return []byte("因为没有自动的处理函数,所以是默认的返回值")
}
