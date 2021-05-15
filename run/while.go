package run

import "github.com/Kotodian/floc"

func While(predicate floc.Predicate, job floc.Job) floc.Job {
	return func(ctx floc.Context, ctrl floc.Control) error {
		for !ctrl.IsFinished() && predicate(ctx) {
			err := job(ctx, ctrl)
			if handledErr := handleResult(ctrl, err); handledErr != nil {
				return handledErr
			}
		}
		return nil
	}
}
