/**************************************************************************************
Code Description    : 连接管理者
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

// clientMgr 工人的管理者
type clientMgr struct {
	length     int
	head, tail *client
}

func newClientMgr() *clientMgr {
	ele := newClient()
	return &clientMgr{
		length: 0,
		head:   ele,
		tail:   ele,
	}
}

func (cm *clientMgr) push(client *client) {
	cm.tail.next = client
	cm.tail = client
	cm.length++
}

func (cm *clientMgr) pop() *client {
	var retEle *client
	retEle = nil
	if cm.length != 0 {
		retEle = cm.head.next
		cm.head.next = retEle.next
		retEle.next = nil
		cm.length--
		if cm.length == 0 {
			cm.tail = cm.head
		}

	}
	return retEle
}
