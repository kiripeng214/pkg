package singleflight

import "sync"

type call struct {
	wg  sync.WaitGroup
	val interface{}
	err error
}

type Group struct {
	mu sync.Mutex
	m  map[string]*call
}

func (this *Group) Do(key string, fn func() (interface{}, error)) (interface{}, error) {
	this.mu.Lock()
	if this.m == nil {
		this.m = make(map[string]*call)
	}
	if c, ok := this.m[key]; ok {
		this.mu.Unlock()
		c.wg.Wait()
		return c.val, c.err
	}
	c := new(call)
	c.wg.Add(1)
	this.m[key] = c
	this.mu.Unlock()

	c.val, c.err = fn()
	c.wg.Done()

	this.mu.Lock()
	delete(this.m, key)
	this.mu.Unlock()

	return c.val, c.err
}
