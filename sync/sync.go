package v1

import "sync"

// Counter将会自增1
type Counter struct {
	mu    sync.Mutex
	value int
}

// NewCounter 返回一个新的计数器
func NewCounter() *Counter {
	return &Counter{}
}

// 计数器自增
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// 返回当前计数
func (c *Counter) Value() int {
	return c.value
}
