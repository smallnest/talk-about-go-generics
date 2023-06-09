package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func min[T constraints.Ordered](x, y T) T {
	if x < y {
		return x
	}
	return y
}

func main() {
	fmt.Println(min(1, 2))
	fmt.Println(min(1.0, 2.0))
	fmt.Print(min("a", "b"))
}
