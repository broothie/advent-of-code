package main

import (
	"strconv"
	"strings"

	"github.com/bobg/errors"
	"github.com/broothie/advent-of-code"
	"github.com/samber/lo"
)

func main() {
	adventofcode.Run(4277556, func(input string) (int, error) {
		var fieldsGrid [][]string
		for _, line := range strings.Split(input, "\n") {
			fieldsGrid = append(fieldsGrid, strings.Fields(line))
		}

		// Last row is signs
		signs, valueStringsGrid := fieldsGrid[len(fieldsGrid)-1], fieldsGrid[:len(fieldsGrid)-1]

		// Parse valueStrings
		var valuesGrid [][]int
		for _, valueStrings := range valueStringsGrid {
			var values []int
			for _, valueString := range valueStrings {
				value, err := strconv.Atoi(valueString)
				if err != nil {
					return 0, errors.Wrap(err, "parsing value")
				}

				values = append(values, value)
			}

			valuesGrid = append(valuesGrid, values)
		}

		// Set first row to results
		results, valuesGrid := valuesGrid[0], valuesGrid[1:]

		// Run math
		for _, valuesRow := range valuesGrid {
			for index, value := range valuesRow {
				sign := signs[index]
				if sign == "+" {
					results[index] += value
				} else if sign == "*" {
					results[index] *= value
				} else {
					return 0, errors.Errorf("unknown sign: %q", sign)
				}
			}
		}

		return lo.Sum(results), nil
	})
}
