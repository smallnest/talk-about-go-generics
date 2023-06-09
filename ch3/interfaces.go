package ch3

import (
	"io"

	"github.com/golang/snappy"
)

func ReadZippedData(r io.Reader) ([]byte, error) {
	r = snappy.NewReader(r)

	return io.ReadAll(r)
}

// func ReadZippedData2[T io.Reader](r T) ([]byte, error) {
// 	r = snappy.NewReader(r)

// 	return io.ReadAll(r)
// }
