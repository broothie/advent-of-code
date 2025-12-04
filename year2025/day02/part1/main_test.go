package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isValidID(t *testing.T) {
	assert.True(t, isValidID(100))

	assert.False(t, isValidID(11))
	assert.False(t, isValidID(22))
	assert.False(t, isValidID(99))
	assert.False(t, isValidID(1010))
	assert.False(t, isValidID(1188511885))
	assert.False(t, isValidID(222222))
	assert.False(t, isValidID(446446))
	assert.False(t, isValidID(38593859))
}
