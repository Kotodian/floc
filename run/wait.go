package run

import (
	"github.com/Kotodian/floc"
	"time"
)

func Wait(predicate floc.Predicate, sleep time.Duration) floc.Job {
	return func(ctx floc.Context, ctrl floc.Control) error {
		for !ctrl.IsFinished() && !predicate(ctx) {
			time.Sleep(sleep)
		}
		return nil
	}
}
