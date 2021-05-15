package run

import (
	"github.com/Kotodian/floc"
	"time"
)

func Delay(delay time.Duration, job floc.Job) floc.Job {
	return func(ctx floc.Context, ctrl floc.Control) error {
		if ctrl.IsFinished() {
			return nil
		}
		timer := time.NewTimer(delay)
		defer timer.Stop()

		select {
		case <-ctx.Done():
		case <-timer.C:
			err := job(ctx, ctrl)
			if handleErr := handleResult(ctrl, err); handleErr != nil {
				return handleErr
			}
		}
		return nil
	}
}
