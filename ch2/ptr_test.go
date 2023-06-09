package ch2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func ptr[T any](v T) *T {
	return &v
}

func TestPtr(t *testing.T) {
	var x int
	var y = ptr(x) // 生成一个指向 x类型 的指针，不是指向x的指针
	*y = 10

	assert.Equal(t, 0, x)
}
