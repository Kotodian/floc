package pred

import "github.com/Kotodian/floc"

func Not(predicate floc.Predicate) floc.Predicate {
	return func(ctx floc.Context) bool {
		return !predicate(ctx)
	}
}
