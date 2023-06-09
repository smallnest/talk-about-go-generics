package ch1

import (
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
)

// ---------------------------
// 接口和类型约束 Type constraints

// 类型约束
type Int interface {
	int
}
type AllInt interface {
	int | int8 | int16 | int32 | int64
}

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type SignedToString interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64

	ToString() string
}

// 类型约束不能再当做普通的接口，只能用作类型约束
type IntToString int

func (it IntToString) ToString() string {
	return strconv.Itoa(int(it))
}

// var _ SignedToString = IntToString(0)

func ToString[T SignedToString](t T) string {
	return t.ToString()
}
func TestToString() {
	println(ToString(IntToString(123)))
}

// 嵌入接口
type ReadWriterCloser interface {
	*os.File | *net.TCPConn | *net.UDPConn | *net.UnixConn | *net.IPConn

	io.Reader
	io.Writer
	io.Closer
}

func TestReadWriterCloser[RWC ReadWriterCloser](rw RWC) {
	file, _ := os.Open("interfaces.go")
	// var _ ReadWriterCloser = file

	TestReadWriterCloser(file)
}

// 空类型
type Empty interface {
	int
	float64
}

type IntOrFloat64 interface {
	~int | ~float64
}

// 迭代
func _[S []byte | string](s S, i, j int) S {
	return s[i:j]
}

// func rangeTheValue[R map[int]string | []string](r R) {
// 	for i, s := range r {
// 		fmt.Println(i, s)
// 	}
// }

// func Entry[T []int | map[int]int](c T, i int) int {
// 	return c[i]
// }

func builtin[S map[int]int | []byte](s *S) {
	s1 := new(S)
	fmt.Println(len(*s1))
}
