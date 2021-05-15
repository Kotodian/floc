package run

import "github.com/Kotodian/floc"

func IfNot(predicate floc.Predicate, jobs ...floc.Job) floc.Job {
	count := len(jobs)
	if count == 1 {
		return func(ctx floc.Context, ctrl floc.Control) error {
			if ctrl.IsFinished() {
				return nil
			}
			if !predicate(ctx) {
				err := jobs[idxTrue](ctx, ctrl)
				if handledErr := handleResult(ctrl, err); handledErr != nil {
					return handledErr
				}
			}
			return nil
		}
	} else if count == 2 {
		return func(ctx floc.Context, ctrl floc.Control) error {
			if ctrl.IsFinished() {
				return nil
			}
			var err error
			if !predicate(ctx) {
				err = jobs[idxTrue](ctx, ctrl)
			} else {
				err = jobs[idxFalse](ctx, ctrl)
			}

			if handledErr := handleResult(ctrl, err); handledErr != nil {
				return handledErr
			}
			return nil
		}
	}
	panic("IfNot requires one or two jobs")
}
