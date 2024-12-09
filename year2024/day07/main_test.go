package main

import (
	"fmt"
	"slices"
	"testing"

	"github.com/samber/lo"
	"github.com/shoenig/test"
)

func Test_generateOperatorSets(t *testing.T) {
	t.Run("0", func(t *testing.T) {
		test.Eq(t, []string{""}, generateOperatorSets("+*", 0))
	})

	t.Run("2", func(t *testing.T) {
		test.Eq(t,
			[]string{
				"++",
				"+*",
				"*+",
				"**",
			},
			generateOperatorSets("+*", 2),
		)
	})

	t.Run("4", func(t *testing.T) {
		expected := []string{
			"****",
			"***+",
			"**+*",
			"**++",
			"*+**",
			"*+*+",
			"*++*",
			"*+++",
			"+***",
			"+**+",
			"+*+*",
			"+*++",
			"++**",
			"++*+",
			"+++*",
			"++++",
		}

		actual := generateOperatorSets("+*", 4)

		slices.Sort(expected)
		slices.Sort(actual)
		test.Eq(t, expected, actual)

		for _, pair := range lo.Zip2(expected, actual) {
			fmt.Println(pair.A, pair.B)
		}
	})
}
