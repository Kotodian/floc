package run

import "github.com/Kotodian/floc"

const (
	idxTrue  = 0
	idxFalse = 1
)

func If(predicate floc.Predicate, jobs ...floc.Job) floc.Job {
	count := len(jobs)
	if count == 1 {
		return func(ctx floc.Context, ctrl floc.Control) error {
			if ctrl.IsFinished() {
				return nil
			}

			if predicate(ctx) {
				err := jobs[idxTrue](ctx, ctrl)
				if handleErr := handleResult(ctrl, err); handleErr != nil {
					return handleErr
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
			if predicate(ctx) {
				err = jobs[idxTrue](ctx, ctrl)
			} else {
				err = jobs[idxFalse](ctx, ctrl)
			}

			if handleErr := handleResult(ctrl, err); handleErr != nil {
				return handleErr
			}
			return nil
		}
	}
	panic("If requires one or tow jobs")
}
