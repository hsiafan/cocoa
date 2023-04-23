package internal

import "sync"

type SyncCache[K comparable, V any] struct {
	m           sync.Map
	computeLock sync.Mutex
	Compute     func(key K) V
}

func (c *SyncCache[K, V]) Load(key K) V {
	v, ok := c.m.Load(key)
	if ok {
		return v.(V)
	}
	c.computeLock.Lock()
	defer c.computeLock.Unlock()
	v, ok = c.m.Load(key)
	if ok {
		return v.(V)
	}

	v = c.Compute(key)
	c.m.Store(key, v)
	return v.(V)
}
