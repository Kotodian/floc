package run

import "github.com/Kotodian/floc"

func Parallel(jobs ...floc.Job) floc.Job {
	return func(ctx floc.Context, ctrl floc.Control) error {
		if ctrl.IsFinished() {
			return nil
		}

		done := make(chan error, len(jobs))
		defer close(done)

		jobsRunning := 0
		for _, job := range jobs {
			jobsRunning++

			go func(job floc.Job) {
				var err error
				defer func() { done <- err }()
				err = job(ctx, ctrl)
			}(job)
		}
		errs := make([]error, 0, len(jobs))

		for jobsRunning > 0 {
			select {
			case <-ctx.Done():
			case err := <-done:
				if handErr := handleResult(ctrl, err); handErr != nil {
					errs = append(errs, handErr)
				}
				jobsRunning--
			}
		}
		if len(errs) > 0 {
			return floc.NewErrMultiple(errs...)
		}
		return nil
	}
}
