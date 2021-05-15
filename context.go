package floc

import (
	"context"
	"sync"
)

type Context interface {
	Releaser

	Ctx() context.Context

	UpdateCtx(ctx context.Context)

	Done() <-chan struct{}

	Value(key interface{}) (value interface{})

	AddValue(key, value interface{})
}

type flowContext struct {
	ctx context.Context
	mu  sync.RWMutex
}

func BorrowContext(ctx context.Context) Context {
	if ctx == nil {
		ctx = context.Background()
	}

	return &flowContext{ctx: ctx, mu: sync.RWMutex{}}
}

func NewContext() Context {
	return &flowContext{
		ctx: context.TODO(),
		mu:  sync.RWMutex{},
	}
}

func (f *flowContext) Release() {

}

func (f *flowContext) Ctx() context.Context {
	f.mu.RLock()
	defer f.mu.RUnlock()
	return f.ctx
}

func (f *flowContext) UpdateCtx(ctx context.Context) {
	f.mu.Lock()
	f.ctx = ctx
	f.mu.Unlock()
}

func (f *flowContext) Done() <-chan struct{} {
	f.mu.RLock()
	defer f.mu.RUnlock()
	return f.ctx.Done()
}

func (f *flowContext) Value(key interface{}) (value interface{}) {
	f.mu.RLock()
	ctx := f.ctx
	f.mu.RUnlock()
	return ctx.Value(key)
}

func (f *flowContext) AddValue(key, value interface{}) {
	f.mu.Lock()
	f.ctx = context.WithValue(f.ctx, key, value)
	f.mu.Unlock()
}
