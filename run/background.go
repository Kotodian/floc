package run

import "github.com/Kotodian/floc"

func Background(job floc.Job) floc.Job {
	return func(ctx floc.Context, ctrl floc.Control) error {
		if ctrl.IsFinished() {
			return nil
		}

		go func(job floc.Job) {
			err := job(ctx, ctrl)
			_ = handleResult(ctrl, err)
		}(job)
		return nil
	}
}
