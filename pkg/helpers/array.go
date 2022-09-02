package helpers

func Map[TIn, TOut any](s []TIn, f func(TIn) TOut) []TOut {
	r := make([]TOut, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

func Fileter[T any](s []T, f func(T) bool) []T {
	r := make([]T, 0)
	for _, v := range s {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}

func Contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}

	return false
}

func FindIndex[T comparable](s []T, e T) int {
	for i, v := range s {
		if v == e {
			return i
		}
	}

	return -1
}
