package main

import (
	"testing"

	"github.com/shoenig/test"
)

func Test_copy(t *testing.T) {
	slice := []int{1, 2, 3}

	copy(slice, []int{10, 20})

	test.Eq(t, []int{10, 20, 3}, slice)
}
