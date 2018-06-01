/**************************************************************************************
Code Description    : 会话管理者
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

package tcp

import (
	"sync"
)

var (
	defaultCapMask = uint64(0xff)
)

// SessionMgr 用于管理Session会话
type SessionMgr struct {
	value  int64
	store  []map[int64]*Session
	rwlock sync.RWMutex
}

// notifyStop 用于通知会话停止
func (sm *SessionMgr) notifyStop() {
	for _, slotMap := range sm.store {
		for _, s := range slotMap {
			s.stop()
		}
	}
}

// newSessionMgr 用于创建一个会话管理者
func newSessionMgr() *SessionMgr {
	pSessionMgr := &SessionMgr{
		value: 0x1,
		store: make([]map[int64]*Session, 0x1),
	}
	for i := int64(0x1); i < pSessionMgr.value; i++ {
		pSessionMgr.store[i] = make(map[int64]*Session, defaultCapMask)
	}
	return pSessionMgr
}

// addSession 用于增加会话
func (sm *SessionMgr) addSession(s *Session) {
	sm.rwlock.Lock()
	defer sm.rwlock.Unlock()
	sm.store[s.numberID%sm.value][s.numberID] = s
	if len(sm.store[s.numberID%sm.value]) == 0xff {
		sm.value++
		newStore := make([]map[int64]*Session, sm.value)
		for i := int64(0); i < sm.value; i++ {
			newStore[i] = map[int64]*Session{}
		}
		for _, slotMap := range sm.store {
			for slotKey, slotValue := range slotMap {
				newStore[slotKey%sm.value][slotKey] = slotValue
			}

		}
		sm.store = newStore
	}
}

// deleteSession 用于删除会话
func (sm *SessionMgr) deleteSession(s *Session) {
	sm.rwlock.Lock()
	defer sm.rwlock.Unlock()
	delete(sm.store[s.numberID%sm.value], s.numberID)
	if len(sm.store[s.numberID%sm.value]) == 0x0 && sm.value > 1 {
		sm.value--
		newStore := make([]map[int64]*Session, sm.value)
		for i := int64(0); i < sm.value; i++ {
			newStore[i] = map[int64]*Session{}
		}
		for _, slotMap := range sm.store {
			for slotKey, slotValue := range slotMap {
				newStore[slotKey%sm.value][slotKey] = slotValue
			}

		}
		sm.store = newStore
	}
}

func (sm *SessionMgr) broadCast(data []byte) {
	for _, slotMap := range sm.store {
		for _, session := range slotMap {
			session.cast(data)
		}
	}
}
