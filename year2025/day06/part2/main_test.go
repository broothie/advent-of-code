package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_transpose(t *testing.T) {
	t.Run("wide to tall", func(t *testing.T) {
		matrix := [][]int{
			{1, 2, 3},
			{4, 5, 6},
		}

		expected := [][]int{
			{1, 4},
			{2, 5},
			{3, 6},
		}

		actual := transpose(matrix)
		assert.Equal(t, expected, actual)
	})

	t.Run("tall to wide", func(t *testing.T) {
		matrix := [][]int{
			{1, 4},
			{2, 5},
			{3, 6},
		}

		expected := [][]int{
			{1, 2, 3},
			{4, 5, 6},
		}

		actual := transpose(matrix)
		assert.Equal(t, expected, actual)
	})
}
