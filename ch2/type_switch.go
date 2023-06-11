package ch2

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// func tswitch[T any](v T) {
// 	switch v := v.(type) {
// 	case int:
// 		println(v)
// 	case string:
// 		println(v)
// 	default:
// 		println(v)
// 	}
// }

func tswitch2[T any](v T) {
	switch (any)(v).(type) {
	case int:
		println(v)
	case string:
		println(v)
	default:
		println(v)
	}
}

// func tassert[T any](v T) {
// 	i := v.(int)
// 	fmt.Println(i)
// }

// func tassert2[T any](v T) {
// 	i := int(v)
// 	fmt.Println(i)
// }

func tassert3[T any](v T) {
	i := (any)(v).(int)
	fmt.Println(i)
}

func tassert4[T constraints.Integer](v T) {
	i := int(v)
	fmt.Println(i)
}
