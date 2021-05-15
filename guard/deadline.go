package guard

import (
	"github.com/Kotodian/floc"
	"time"
)

type WhenDeadlineFunc func(ctx floc.Context, id interface{}) time.Time

func ConstDeadline(deadline time.Time) WhenDeadlineFunc {
	return func(ctx floc.Context, id interface{}) time.Time {
		return deadline
	}
}

func DeadlineIn(delay time.Duration) WhenDeadlineFunc {
	return func(ctx floc.Context, id interface{}) time.Time {
		return time.Now().Add(delay)
	}
}

func Deadline(when WhenDeadlineFunc, id interface{}, job floc.Job) floc.Job {
	return OneDeadline(when, id, job, nil)
}

func OneDeadline(when WhenDeadlineFunc, id interface{}, job floc.Job, timeoutTrigger TimeoutTrigger) floc.Job {
	whenTimeout := func(ctx floc.Context, id interface{}) time.Duration {
		return time.Until(when(ctx, id))
	}
	return OnTimeout(whenTimeout, id, job, timeoutTrigger)
}
