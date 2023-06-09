package ch2

type Bytes []byte

var x []byte
var y Bytes

func init() {
	x = y // okay
	y = x // okay
}

func f[T Bytes](v T) {
	x = v // okay
	y = v // error
	v = x // okay
	v = y // error
}

func g[T []byte](v T) {
	x = v // okay
	y = v // error
	v = x // okay
	v = y // error
}

func h[T Bytes | []byte](v T) {
	x = v // okay
	y = v // error
	v = x // okay
	v = y // error
}
