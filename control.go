package floc

import (
	"context"
	"sync/atomic"
)

type Control interface {
	Releaser

	Complete(data interface{})

	Cancel(data interface{})

	Fail(data interface{}, err error)

	IsFinished() bool

	Result() (result Result, data interface{}, err error)
}

const (
	statusRunning  = 0
	statusFinished = 1
)

type flowControl struct {
	ctx    Context
	cancel context.CancelFunc
	status int32
	result int32
	data   interface{}
	err    error
}

func (f *flowControl) Release() {
	f.Cancel(nil)
}

func (f *flowControl) Complete(data interface{}) {
	f.finish(Completed, data, nil)
}

func (f *flowControl) Cancel(data interface{}) {
	f.finish(Canceled, data, nil)
}

func (f *flowControl) Fail(data interface{}, err error) {
	f.finish(Failed, data, err)
}

func (f *flowControl) IsFinished() bool {
	r := atomic.LoadInt32(&f.result)
	return Result(r).IsFinished()
}

func (f *flowControl) Result() (result Result, data interface{}, err error) {
	r := atomic.LoadInt32(&f.result)
	result = Result(r)

	if result.IsFinished() {
		return result, f.data, f.err
	}
	return result, nil, nil
}

func NewControl(ctx Context) Control {
	if ctx == nil {
		panic("context is nil")
	}
	oldCtx := ctx.Ctx()
	cancelCtx, cancelFunc := context.WithCancel(oldCtx)
	ctx.UpdateCtx(cancelCtx)
	return &flowControl{
		ctx:    ctx,
		cancel: cancelFunc,
		status: statusRunning,
		result: None.i32(),
	}
}

func (f *flowControl) finish(result Result, data interface{}, err error) {
	if atomic.CompareAndSwapInt32(&f.status, statusRunning, statusFinished) {
		f.data = data
		f.err = err

		atomic.StoreInt32(&f.result, result.i32())
		f.cancel()
	}
}
