package pred

import "github.com/Kotodian/floc"

func Xor(predicates ...floc.Predicate) floc.Predicate {
	count := len(predicates)
	if count > 2 {
		return func(ctx floc.Context) bool {
			result := predicates[0](ctx) != predicates[1](ctx)

			for i := 2; i < count; i++ {
				result = result != predicates[i](ctx)
			}
			return result
		}
	} else if count == 2 {
		return func(ctx floc.Context) bool {
			return predicates[0](ctx) != predicates[1](ctx)
		}
	}

	panic("Xor requires at least 2 predicates")
}
