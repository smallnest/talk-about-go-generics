package ch2

import (
	"bytes"
	"fmt"
)

func Stringify[T fmt.Stringer](s []T) (ret []string) {
	for i := range s {
		ret = append(ret, s[i].String())
	}
	return ret
}

func Stringify2[T fmt.Stringer](s []T) (ret []string) {

	for i := range s {
		ret = append(ret, s[i].String())
	}
	return ret
}

func TestStringify() {
	var s []bytes.Buffer

	var ret []string
	for i := range s {
		ret = append(ret, s[i].String())
	}

	var result = Stringify(s)
	fmt.Println(result)
}
