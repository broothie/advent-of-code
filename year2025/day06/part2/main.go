package main

import (
	"strconv"
	"strings"

	"github.com/bobg/errors"
	"github.com/broothie/advent-of-code"
	"github.com/samber/lo"
)

func main() {
	adventofcode.Run(3263827, func(input string) (int, error) {
		lines := strings.Split(input, "\n")

		// Last line is signs, signs position indicates column start
		signsLine, lines := lines[len(lines)-1], lines[:len(lines)-1]
		signs := strings.Fields(signsLine)

		var signPositions []int
		for index, char := range signsLine {
			if char != ' ' {
				signPositions = append(signPositions, index)
			}
		}

		sum := 0
		for columnIndex, start := range signPositions {
			var end int
			if columnIndex == len(signPositions)-1 {
				end = len(lines[0])
			} else {
				end = signPositions[columnIndex+1]
			}

			var transposedColumn [][]rune
			for _, line := range lines {
				transposedColumn = append(transposedColumn, []rune(line[start:end]))
			}

			column := lo.Map(transpose(transposedColumn), func(row []rune, _ int) string { return string(row) })

			var intColumn []int
			for _, row := range column {
				cleanedRow := string(lo.Filter([]rune(row), func(char rune, _ int) bool { return char != ' ' }))
				if cleanedRow == "" {
					continue
				}

				value, err := strconv.Atoi(cleanedRow)
				if err != nil {
					return 0, errors.Wrap(err, "parsing int")
				}

				intColumn = append(intColumn, value)
			}

			if sign := signs[columnIndex]; sign == "+" {
				sum += lo.Reduce(intColumn, func(a int, b int, _ int) int { return a + b }, 0)
			} else if sign == "*" {
				sum += lo.Reduce(intColumn, func(a int, b int, _ int) int { return a * b }, 1)
			} else {
				return 0, errors.Errorf("unknown sign: %q", sign)
			}
		}

		return sum, nil
	})
}

type Range struct {
	Start int // Inclusive
	End   int // Exclusive
}

func transpose[T any](matrix [][]T) [][]T {
	height, width := len(matrix), len(matrix[0])

	// Transpose dimensions
	newHeight, newWidth := width, height

	newMatrix := make([][]T, newHeight)
	for y := range newHeight {
		newMatrix[y] = make([]T, newWidth)
	}

	for y, row := range matrix {
		for x, value := range row {
			newMatrix[x][y] = value
		}
	}

	return newMatrix
}
