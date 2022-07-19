package collections

func SelectMust[TIn any, TOut any](col []TIn, selector func(element TIn) TOut) []TOut {
	result := []TOut{}

	for _, item := range col {
		transformed := selector(item)
		result = append(result, transformed)
	}

	return result
}
