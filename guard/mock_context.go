package guard

import "github.com/Kotodian/floc"

type mockContext struct {
	floc.Context
	mock floc.Context
}

func (ctx mockContext) Release() {
	ctx.mock.Release()
}

func (ctx mockContext) Done() <-chan struct{} {
	return ctx.mock.Done()
}
