package run

import "github.com/Kotodian/floc"

func handleResult(ctrl floc.Control, err error) error {
	if err != nil {
		ctrl.Fail(nil, err)
		return err
	}
	return nil
}
