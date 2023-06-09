package ch2

type Int1 interface {
	int16 | int32 | int64
}

func f[T Int1](x T) {
	g(x)
	// h(x)
}

type Int2 interface {
	int16 | int32 | int64
}

func g[T Int2](x T) {
}

type Int3 interface {
	int16 | int32
}

func h[T Int3](x T) {
}

func test() {
	f(int16(0))
}
