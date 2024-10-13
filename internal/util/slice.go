package util

func SliceFilter[T any](ss []T, test func(T) bool) []T {
	var ret []T
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}

	return ret
}

func SliceFind[T any](ss []T, test func(T) bool) *T {
	for _, s := range ss {
		if test(s) {
			return &s
		}
	}

	return nil
}
