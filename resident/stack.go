/**************************************************************************************
Code Description    : 单列表
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

package resident

// linkedList 为一个单向列表
type stack struct {
	length     int
	head, tail *element
}

func newStack() *stack {
	ele := elePool.get()
	return &stack{
		length: 0,
		head:   ele,
		tail:   ele,
	}
}

// element 为单向列表中的元素
type element struct {
	data interface{}
	next *element
}

func newElement() *element {
	return &element{
		next: nil,
	}
}

// SetData 用于设置元素中的数据
func (ele *element) SetData(data interface{}) {
	ele.data = data
}

// Reset 用于重置元素
func (ele *element) Reset() {
	ele.data = nil
	ele.next = nil
}

// Push 用于向linkedList的末尾追加元素
func (ll *stack) Push(ele *element) {
	ll.tail.next = ele
	ll.tail = ele
	ll.length++
}

// Pop 用于从头节点删除一个元素
func (ll *stack) Pop() *element {
	var retEle *element
	retEle = nil
	if ll.length != 0 {
		retEle = ll.head.next
		ll.head.next = retEle.next
		ll.length--
		if ll.length == 0 {
			ll.tail = ll.head
		}
	}
	return retEle
}
