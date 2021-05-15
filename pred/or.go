package pred

import "github.com/Kotodian/floc"

func Or(predicates ...floc.Predicate) floc.Predicate {
	count := len(predicates)
	if count > 2 {
		return func(ctx floc.Context) bool {
			for _, p := range predicates {
				if p(ctx) {
					return true
				}
			}
			return false
		}
	} else if count == 2 {
		return func(ctx floc.Context) bool {
			return predicates[0](ctx) || predicates[1](ctx)
		}
	}
	panic("Or requires at least 2 predicates")
}
