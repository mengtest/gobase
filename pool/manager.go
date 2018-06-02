package pool

// 非协程安全
import (
	"sync"
)

// Manager 用于管理池
type Manager struct {
	pools map[string]sync.Pool
}

// NewManager 用于新建一个Manager
func NewManager() *Manager {
	pManager := &Manager{
		pools: make(map[string]sync.Pool),
	}
	return pManager
}

// Register 用于注册新的对象池
func (m *Manager) Register(name string,
	create func() interface{}) {
	m.pools[name] = sync.Pool{
		New: create,
	}
}

// UnRegister 用于注销已有的对象池
func (m *Manager) UnRegister(name string) {
	delete(m.pools, name)
}

// Get 用于获取指定模块的对象
func (m *Manager) Get(name string) interface{} {
	model, ok := m.pools[name]
	if ok {
		return model.Get()
	}
	return nil
}

// Put 用于将对象放回到池子里
func (m *Manager) Put(name string, data interface{}) {
	if model, ok := m.pools[name]; ok {
		model.Put(data)
	}
}
