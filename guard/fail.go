package guard

import "github.com/Kotodian/floc"

func Fail(data interface{}, err error) floc.Job {
	return func(ctx floc.Context, ctrl floc.Control) error {
		ctrl.Fail(data, err)
		return nil
	}
}
