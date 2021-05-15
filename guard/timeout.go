package guard

import (
	"github.com/Kotodian/floc"
	"time"
)

type WhenTimeoutFunc func(ctx floc.Context, id interface{}) time.Duration
type TimeoutTrigger func(ctx floc.Context, ctrl floc.Control, id interface{})

func Timeout(when WhenTimeoutFunc, id interface{}, job floc.Job) floc.Job {
	return OnTimeout(when, id, job, nil)
}
func OnTimeout(when WhenTimeoutFunc, id interface{}, job floc.Job, timeoutTrigger TimeoutTrigger) floc.Job {
	return func(ctx floc.Context, ctrl floc.Control) error {
		done := make(chan error)
		defer close(done)

		timer := time.NewTimer(when(ctx, id))
		defer timer.Stop()

		go func() {
			var err error
			defer func() { done <- err }()
			err = job(ctx, ctrl)
		}()

		select {
		case <-ctx.Done():
		case err := <-done:
			return err
		case <-timer.C:
			if timeoutTrigger != nil {
				timeoutTrigger(ctx, ctrl, id)
			} else {
				ctrl.Fail(id, floc.NewErrTimeout(id, time.Now().Local()))
			}
		}
		return <-done
	}
}
func ConstTimeout(timeout time.Duration) WhenTimeoutFunc {
	return func(ctx floc.Context, id interface{}) time.Duration {
		return timeout
	}
}
