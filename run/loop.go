package run

import "github.com/Kotodian/floc"

func Loop(job floc.Job) floc.Job {
	return func(ctx floc.Context, ctrl floc.Control) error {
		for {
			if ctrl.IsFinished() {
				return nil
			}

			err := job(ctx, ctrl)
			if handledErr := handleResult(ctrl, err); handledErr != nil {
				return handledErr
			}
		}
	}
}
