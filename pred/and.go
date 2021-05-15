package pred

import "github.com/Kotodian/floc"

func And(predicates ...floc.Predicate) floc.Predicate {
	count := len(predicates)
	if count > 2 {
		return func(ctx floc.Context) bool {
			for _, p := range predicates {
				if !p(ctx) {
					return false
				}
			}
			return true
		}
	} else if count == 2 {
		return func(ctx floc.Context) bool {
			return predicates[0](ctx) && predicates[1](ctx)
		}
	}
	panic("And requries at least predicates")
}
