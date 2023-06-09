package ch2

import "cmp"

type Map[K comparable, V any] map[K]V
type List[T any] []T
type Chan[T any] chan T
type TenItems[T any] [10]T
type Func[T any] func() T

type Heap[T cmp.Ordered] []T

func (h Heap[_]) Len() int           { return len(h) }
func (h Heap[_]) Less(i, j int) bool { return h[i] < h[j] }
func (h Heap[_]) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Heap[T]) Push(x T) {
	*h = append(*h, x)
}

func (h *Heap[T]) Pop() T {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
