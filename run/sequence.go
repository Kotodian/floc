package run

import "github.com/Kotodian/floc"

func Sequence(jobs ...floc.Job) floc.Job {
	return func(ctx floc.Context, ctrl floc.Control) error {
		for _, job := range jobs {
			if ctrl.IsFinished() {
				return nil
			}

			err := job(ctx, ctrl)
			if handledErr := handleResult(ctrl, err); err != nil {
				return handledErr
			}
		}
		return nil
	}
}
