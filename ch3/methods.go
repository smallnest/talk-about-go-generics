package ch3

// func Index[T map[int]string | []string](m T, k int) string {
// 	return m[k]
// }

func IndexMap[T map[int]string](m T, k int) string {
	return m[k] // map的查找算法
}

func IndexSlice[T []string](m T, k int) string {
	return m[k] // slice查找算法
}
