/**************************************************************************************
Code Description    : 该文件主要用于封装原子管道,但是由于时间紧迫,工作中要用.
					  但是按照reactor的模式来的管道,下面代码还要很多事情没有做
Code Vesion         :
					|------------------------------------------------------------|
						  Version    					Editor            Time
							1.0        					yuansudong        2018.4.12
					|------------------------------------------------------------|
Version Description	:
                    |------------------------------------------------------------|
						  Version
							1.0
								  (1) 主要封装相关的原子管道操作
									  Cap,Length,Read,Write
								  (2) MT开头的为消息类型
								  (3) AE开头的为错误类型
					|------------------------------------------------------------|
***************************************************************************************/

package achannel

import (
	"errors"
	"gobase/alock"
	"runtime"
	"time"
)

/*
  File Description: 不够完善,为了工作
  File Editor     : yuansudong
*/

var (
	// readFunc 读回调函数
	readFunc alock.CallBack
	// writeFunc 写回调函数
	writeFunc alock.CallBack
	// lengthFunc 长度回调函数
	lengthFunc alock.CallBack
)

var (
	// AETimeout 超时错误
	AETimeout error
)

// 包初始化函数
func init() {

	// 回调函数初始化
	readFunc = func(Args ...interface{}) ([]interface{}, error) {
		ChannelArg := Args[0].(*Channel)
		if ChannelArg.len == 0 {
			return []interface{}{}, errors.New("管道为空")
		}
		ReturnNode := ChannelArg.head.next
		ChannelArg.head.next = ReturnNode.next
		ChannelArg.len--
		if ChannelArg.len == 0 {
			ChannelArg.tail = ChannelArg.head
		}
		return []interface{}{ReturnNode.msg}, nil
	}

	writeFunc = func(Args ...interface{}) ([]interface{}, error) {
		ChannelArg := Args[0].(*Channel)
		if ChannelArg.len == ChannelArg.cap {
			return nil, errors.New("管道满了")
		}
		InsertNode := Args[1].(*channelMsgNode)
		ChannelArg.tail.next = InsertNode
		ChannelArg.tail = InsertNode
		ChannelArg.len++
		return nil, nil
	}
	lengthFunc = func(Args ...interface{}) ([]interface{}, error) {
		ChannelArg := Args[0].(*Channel)
		return []interface{}{ChannelArg.len}, nil
	}

	// 错误类型初始化
	AETimeout = errors.New("Achannel Timeout")

}

type (

	// Channel 用于描述一个双向的原子管道
	Channel struct {
		cap  int             // 管道总容量
		len  int             // 大小
		head *channelMsgNode // 头结点
		tail *channelMsgNode // 尾结点
		sync *alock.Synchro
	}
	// channelMsgNode 用于描述一个管道消息
	channelMsgNode struct {
		msg  interface{}
		next *channelMsgNode
	}
)

// GetMessage 用于获取管道中的消息点
func (CM *channelMsgNode) GetMessage() interface{} {
	return CM.msg
}

// createchannelMsgNode 用于创建一个管道结点
func createchannelMsgNode(D interface{}) *channelMsgNode {
	return &channelMsgNode{
		msg: D,
	}
}

// Create 用于创建一个管道
func Create(Cap int) *Channel {
	SentinelNode := createchannelMsgNode(nil)
	return &Channel{
		cap:  Cap,
		len:  0,
		head: SentinelNode,
		tail: SentinelNode,
		sync: alock.Create(),
	}
}

// Length 用于返回管道的长度
func (C *Channel) Length() int {
	Result, _ := C.sync.SyncExec(lengthFunc, C)
	return Result[0].(int)
}

// Cap 用于返回管道的总容量
func (C *Channel) Cap() int {
	return C.cap
}

// Read 用于读取管道中的数据,该函数如果一直娶不到数据会阻塞的
func (C *Channel) Read() interface{} {
	for {
		Result, Err := C.sync.SyncExec(readFunc, C)
		if Err != nil {
			//fmt.Println("read", Err)
			runtime.Gosched()
			time.Sleep(time.Microsecond)
			continue
		}
		return Result[0]

	}
	// 永远都不可能会执行到这里
}

// ReadTimeout 用于读取管道中的数据,只不过到达了指定的时间内,如果还没有取到数据,那么就会返回
func (C *Channel) ReadTimeout(Timeout time.Time) (interface{}, error) {
	for {
		Result, Err := C.sync.SyncExec(readFunc, C)
		if Err != nil {
			if time.Now().Sub(Timeout) > 0 {
				// 进入了这里就代表着超时了
				return nil, errors.New("超时")
			}
			runtime.Gosched()
			time.Sleep(time.Microsecond)
			continue
		}
		return Result[0], nil
	}
}

// Write 用于向管道中写数据,直到成功为止
func (C *Channel) Write(Msg interface{}) {
	PchannelMsgNode := createchannelMsgNode(Msg)
	for {
		_, Err := C.sync.SyncExec(writeFunc, C, PchannelMsgNode)
		if Err != nil {
			//fmt.Println("write", Err)
			runtime.Gosched()
			time.Sleep(time.Microsecond)
			continue
		}
		return
	}
	// 永远都不会执行到这里的
}

// WriteTimeout 用于向管道中写数据,只不过经过一段时间后,队列还是满的话, 就会返回超时错误
func (C *Channel) WriteTimeout(Msg interface{}, Timeout time.Time) error {
	PchannelMsgNode := createchannelMsgNode(Msg)
	for {
		_, Err := C.sync.SyncExec(writeFunc, C, PchannelMsgNode)
		if Err != nil {
			if time.Now().Sub(Timeout) > 0 {
				return AETimeout
			}
			runtime.Gosched()
			time.Sleep(time.Microsecond)
			continue
		}
		return nil
	}
}
