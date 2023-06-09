package ch2

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// https://go.dev/blog/intro-generics

func Scale[E constraints.Integer](s []E, c E) []E {
	r := make([]E, len(s))
	for i, v := range s {
		r[i] = v * c
	}
	return r
}

// Scale returns a copy of s with each element multiplied by c.
func Scale1[S ~[]E, E constraints.Integer](s S, c E) S {
	r := make(S, len(s))
	for i, v := range s {
		r[i] = v * c
	}
	return r
}

type Point []int32

func (p Point) String() string {
	return ""
}

func ScaleAndPrint(p Point) {
	r := Scale1(p, 2)
	fmt.Println(r.String())
}
