package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isValidID(t *testing.T) {
	assert.False(t, isValidID(11))
	assert.False(t, isValidID(22))
	assert.False(t, isValidID(99))
	assert.False(t, isValidID(111))
	assert.False(t, isValidID(1188511885))
	assert.False(t, isValidID(222222))
	assert.False(t, isValidID(446446))
	assert.False(t, isValidID(38593859))
	assert.False(t, isValidID(565656))
	assert.False(t, isValidID(824824824))
	assert.False(t, isValidID(2121212121))
}
