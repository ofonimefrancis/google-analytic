package counter

import "sync"

//Counter is a go routine safe counter used to count events.
// We add Read/Write Mutex to prevent race conditions.
type Counter struct {
	counter map[string]uint64
	sync.RWMutex
}

func NewCounter() *Counter {
	return &Counter{counter: make(map[string]uint64)}
}

func (c *Counter) Incr(k string) {
	c.Lock()
	c.counter[k]++
	c.Unlock()
}

func (c *Counter) Val(k string) uint64 {
	var v uint64
	c.RLock()
	v = c.counter[k]
	c.RUnlock()
	return v
}
