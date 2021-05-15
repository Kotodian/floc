package guard

import "github.com/Kotodian/floc"

func Complete(data interface{}) floc.Job {
	return func(ctx floc.Context, ctrl floc.Control) error {
		ctrl.Complete(data)
		return nil
	}
}
