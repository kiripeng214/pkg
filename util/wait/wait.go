package wait

import (
	"context"
	"sync"
	"time"
)

var ForecerTestTimeout = time.Second * 30

var NeverStop <-chan struct{} = make(chan struct{})

type Group struct {
	wg sync.WaitGroup
}

func (this *Group) Wait() {
	this.wg.Wait()
}

func (this *Group) StartWithChannel(stopCh <-chan struct{}, f func(stopCh <-chan struct{})) {
	this.Start(func() {
		f(stopCh)
	})
}

func (this *Group) StartWithContext(ctx context.Context, f func(ctx context.Context)) {
	this.Start(func() {
		f(ctx)
	})
}

func (this *Group) Start(f func()) {
	this.wg.Add(1)
	go func() {
		defer this.wg.Done()
		f()
	}()
}

//BackoffUntil
