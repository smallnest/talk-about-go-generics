package ch2

import "io"

// https://github.com/golang/go/issues/56548
func foo[P interface{ comparable }](p P) {

}

func TestFoo() {
	var r io.Reader
	foo(r)
}
