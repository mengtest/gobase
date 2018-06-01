/**************************************************************************************
Code Description    : 该文件主要用于封装原子锁,其实可以封装成读写锁,但是项目中不需要这个
Code Vesion         :
					|------------------------------------------------------------|
						  Version    					Editor            Time
							1.0        					yuansudong        2018.4.12
					|------------------------------------------------------------|
Version Description	:
                    |------------------------------------------------------------|
						  Version
							1.0
								  (1) 主要封装相关的原子锁操作
								      SyncExec,Create
					|------------------------------------------------------------|
***************************************************************************************/

package alock

import (
	"runtime"
	"sync/atomic"
	"time"
)

const (
	// lock 代表着处于加锁状态
	lock = 1
	// free 代表处于自由状态
	free = 0
)

type (
	// CallBack 用于原子回调函数
	CallBack func(...interface{}) ([]interface{}, error)
	// Synchro 用于描述一个同步原子锁
	Synchro struct {
		// lock 用于保持加锁的状态
		lock int32
	}
)

// Create 用于创建一个原子锁
func Create() *Synchro {
	return &Synchro{
		lock: free,
	}
}

// SyncExec 用于同步执行某个函数
func (S *Synchro) SyncExec(ExecFunc CallBack, Args ...interface{}) ([]interface{}, error) {
	var Result []interface{}
	var Err error
	for {
		OldLock := atomic.LoadInt32(&S.lock)
		if OldLock == lock {
			// 如果满足,则代表此时该队列是一个上了锁的队列
			runtime.Gosched()
			time.Sleep(time.Millisecond * 100)
			continue
		}
		// 如果进入了这里,那么此时代表OldLock 是处于自由状态
		// 所以下面要对其进行上锁
		if !atomic.CompareAndSwapInt32(&S.lock, OldLock, lock) {
			// 如果进入了这里,代表着上锁失败
			runtime.Gosched()
			time.Sleep(time.Millisecond * 100)
			continue
		}
		// 在这里执行函数的结果
		Result, Err = ExecFunc(Args...)
		// 此时应该进行解锁
		atomic.CompareAndSwapInt32(&S.lock, lock, free)
		break
	}
	return Result, Err
}
