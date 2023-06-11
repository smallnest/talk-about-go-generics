package ch2

import (
	"io"
	"os"
)

// 底层类型不能是struct InvalidUnion
type Error struct {
	Code  int
	Error string
}

// type I0 interface {
// 	~Error
// 	~error
// }

type I000 interface {
	~[]byte
	~struct{ f int }
}

type Error2 = struct {
	Code  int
	Error string
}

type I00 interface {
	~Error2
}

func I0_1[K interface{ io.Reader }]() {
}

// // MisplacedTypeParam
// type T1[P any] P

// type T2[P any] struct{ *P }

// func I1[K any, V interface{ K }]() {
// }

// func I2[K any, V interface{ int | K }]() {
// }

// func I4[K any, V interface{ int | ~int }]() {
// }

type MyInt int

func I5[K any, V interface{ int | MyInt }]() {
}

type MyInt2 = int

// func I6[K any, V interface{ int | MyInt2 }]() {
// }

func I7[K any, V interface{ int | any }]() {
}

type I8 interface {
	~int
}

func I9[K any, V interface{ int | I8 }]() {
}

func I10[K interface{ io.Reader }]() {
}

// MisplacedConstraintIface

type IntStr interface {
	int | string
}

func add[T IntStr](a, b T) T {
	return a + b
}

// func add1(a, b IntStr) IntStr {
// 	return a + b
// }

func I11[K interface {
	io.Reader
	io.Writer
}]() {
}

// func I12[K interface {
// 	io.Reader | io.Writer
// }]() {
// }

// func I12_1[K interface {
// 	interface{ foo() } | int
// }]() {
// }

// func I13[K interface {
// 	io.Reader | os.File
// }]() {
// }

func I14[K interface {
	any | os.File
}]() {
}

func I15[K interface {
	comparable
}]() {
}

func I16[K interface {
	error
	comparable
}]() {
}

// func I17[K interface {
// 	comparable | os.File
// }]() {
// }

// func bar[T any](a T) {
// 	if reflect.TypeOf(a).Kind() == reflect.Ptr {
// 		fmt.Println("a is a ptr")
// 		return
// 	}
// 	bar[*T](&a)
// }

type Tree[T any] struct {
	Left  *Tree[T]
	V     Value[T, int]
	Right *Tree[T]
}

type Value[T any, V any] struct {
	Hodler *Tree[T]
	Value  V
}

// type Tree1 struct {
// 	Left  *Tree1
// 	V     Value2
// 	Right *Tree1
// }

// type Value2 struct {
// 	Hodler *Tree1
// 	Value  int
// }

type ListHead[T any] struct {
	head *ListElement[T]
}

type ListElement[T any] struct {
	next *ListElement[T]
	val  T
	head *ListHead[T]
}

type (
	P struct{}
	C struct{}

	// TT0 [P * C]struct{} // P*C
	// TT1 [P(C)]struct{}  // P(C)

	TT2[P interface{ *C }] struct{}
	TT3[P interface{ C }]  struct{}
)

type Animal[T any] struct {
	Val T
}

// type Animal2[T any] struct {
// 	T
// }

// type Animal3[T any] struct {
// 	*T
// }
