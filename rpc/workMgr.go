/**************************************************************************************
Code Description    : 工人的管理者
Code Vesion         :
					|------------------------------------------------------------|
						  Version    					Editor            Time
							1.0        					yuansudong        2016.4.12
					|------------------------------------------------------------|
Version Description	:
                    |------------------------------------------------------------|
						  Version
							1.0
								 ....
					|------------------------------------------------------------|
***************************************************************************************/

package rpc

// workerMgr 工人的管理者
type workerMgr struct {
	length     int
	head, tail *worker
}

func newWokerMgr() *workerMgr {
	ele := newWoker()
	return &workerMgr{
		length: 0,
		head:   ele,
		tail:   ele,
	}
}

func (wm *workerMgr) push(work *worker) {
	wm.tail.next = work
	wm.tail = work
	wm.length++
}

func (wm *workerMgr) pop() *worker {
	var retEle *worker
	retEle = nil
	if wm.length != 0 {
		retEle = wm.head.next
		wm.head.next = retEle.next
		retEle.next = nil
		wm.length--
		if wm.length == 0 {
			wm.tail = wm.head
		}

	}
	return retEle
}
