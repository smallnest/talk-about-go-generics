package ch2

// func ZeroValue0[T any](v T) bool {
// 	return v == nil
// }

func interfaceIsNil(v any) bool {
	return v == nil
}

// https://github.com/golang/go/issues/56548
// https://github.com/golang/go/issues/52624
func isZero[T comparable](v T) bool {
	var t T
	return v == t //
}

func Zero1[T any]() T {
	return *new(T)
}

func Zero2[T any]() T {
	var t T
	return t
}

func Zero3[T any]() (t T) {
	return
}
