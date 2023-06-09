package ch2

type A[T any] struct {
	V T
}

func (a *A[T]) Value() T {
	return a.V
}

// https://github.com/golang/go/issues/52654
// https://github.com/golang/go/issues/46477
// type B = A
// type B[T any] = A[T]

// type C A
type CC[T any] A[T]

func TestB() {
	// var b B[int]
	// b.V = 1

	var c CC[int]
	c.V = 1
}
