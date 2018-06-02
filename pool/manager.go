package pool

import (
	"sync"
)

// Manager 用于管理池
type Manager struct {
	value int
	pools []map[string]sync.Pool
}

// NewManager 用于新建一个Manager
func NewManager() *Manager {

}
