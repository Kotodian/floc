package guard

import "github.com/Kotodian/floc"

type PanicTrigger func(ctx floc.Context, ctrl floc.Control, v interface{})

func Panic(job floc.Job) floc.Job {
	return OnPanic(job, nil)
}

func IgnorePanic(job floc.Job) floc.Job {
	return OnPanic(job, func(ctx floc.Context, ctrl floc.Control, v interface{}) {})
}

func OnPanic(job floc.Job, panicTrigger PanicTrigger) floc.Job {
	return func(ctx floc.Context, ctrl floc.Control) error {
		defer func() {
			if r := recover(); r != nil {
				if panicTrigger != nil {
					panicTrigger(ctx, ctrl, r)
				} else {
					ctrl.Fail(r, floc.NewErrPanic(r))
				}
			}
		}()

		return job(ctx, ctrl)
	}
}
