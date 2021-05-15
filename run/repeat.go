package run

import "github.com/Kotodian/floc"

func Repeat(times int, job floc.Job) floc.Job {
	return func(ctx floc.Context, ctrl floc.Control) error {
		for n := 1; n <= times; n++ {
			if ctrl.IsFinished() {
				return nil
			}

			err := job(ctx, ctrl)
			if handledErr := handleResult(ctrl, err); handledErr != nil {
				return handledErr
			}
		}
		return nil
	}
}
