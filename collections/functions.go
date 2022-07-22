package collections

func SelectMust[TIn any, TOut any](col []TIn, selector func(element TIn) TOut) []TOut {
	result := []TOut{}

	for _, item := range col {
		transformed := selector(item)
		result = append(result, transformed)
	}

	return result
}

func AnyMust[T any](col []T, predicate func(elem T) bool) bool {
	for _, item := range col {
		if predicate(item) {
			return true
		}
	}

	return false
}

func ContainsMust[T comparable](col []T, elem T) bool {
	return AnyMust(col, func(elem2 T) bool {
		return elem == elem2
	})
}
