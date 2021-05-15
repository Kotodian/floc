package guard

import "github.com/Kotodian/floc"

func Cancel(data interface{}) floc.Job {
	return func(ctx floc.Context, ctrl floc.Control) error {
		ctrl.Cancel(data)
		return nil
	}
}
