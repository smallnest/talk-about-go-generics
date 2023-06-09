package ch2

import "cmp"

// func compare[T comparable](a, b T) int {
// 	if a < b {
// 		return -1
// 	} else if a == b {
// 		return 0
// 	} else {
// 		return 1
// 	}
// }

func compare2[T cmp.Ordered](a, b T) int {
	if a < b {
		return -1
	} else if a == b {
		return 0
	} else {
		return 1
	}
}
