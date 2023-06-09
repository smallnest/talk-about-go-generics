package ch1

import (
	"io"

	"golang.org/x/exp/constraints"
)

// ---------------------------
// 基本定义

// 泛型类型
type List[T any] struct {
	next  *List[T]
	value T
}

// 泛型类型的方法
func (l *List[T]) Len() int {
	ll := 0
	for l.next != nil {
		ll++
	}
	return ll
}

func (l *List[V]) Size() int {
	ll := 0
	for l.next != nil {
		ll++
	}
	return ll
}

// 泛型函数
func min[T ~int | ~float64](x, y T) T {
	if x < y {
		return x
	}
	return y
}

func testArray[T any](a [10]T) T {
	return a[9]
}

// ---------------------------
// 泛型类型的扩展
type (
	S1[P any]                           struct{}
	S2[S interface{ ~[]byte | string }] struct{}
	S3[S ~[]E, E any]                   struct{}
	S4[P S1[int]]                       struct{}
	S5[_ any]                           struct{}
	S6[_ any, _ io.Reader]              struct{}
	S7[K any, _ io.Reader, V any]       struct{}
	S8[K, _ any, S, T S4[S1[int]]]      struct{}
	S9[K int, T map[K]any]              struct{}
)
type MyList[T any] List[T]
type MyList2 List[int]
type MyMap[K comparable, V any] map[K]V
type MyMap2[K comparable] map[K]int

// ---------------------------
// 泛型方法的扩展
type MyFunc[T constraints.Ordered, V any] func(T) V
