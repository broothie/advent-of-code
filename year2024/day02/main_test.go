package main

import "testing"
import "github.com/shoenig/test"

func TestSlicing(t *testing.T) {
	a := []int{1, 2, 3}
	test.Eq(t, a, a[:len(a)-1])
}
