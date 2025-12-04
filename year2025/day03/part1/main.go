package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bobg/errors"
	"github.com/broothie/advent-of-code"
)

func main() {
	adventofcode.Run(357, func(input string) (int, error) {
		sum := 0

		lines := strings.Split(input, "\n")
		for _, line := range lines {
			if line == "" {
				continue
			}

			var digits []int
			for _, char := range line {
				digit, err := strconv.Atoi(string(char))
				if err != nil {
					return 0, errors.Wrap(err, "parsing digit")
				}

				digits = append(digits, digit)
			}

			firstDigit, firstDigitIndex := maxWithIndex(digits[:len(digits)-1])
			secondDigit, _ := maxWithIndex(digits[firstDigitIndex+1:])

			joltage, err := strconv.Atoi(fmt.Sprintf("%d%d", firstDigit, secondDigit))
			if err != nil {
				return 0, errors.Wrap(err, "parsing joltage")
			}

			sum += joltage
		}

		return sum, nil
	})
}

func maxWithIndex(digits []int) (int, int) {
	resultDigit := digits[0]
	resultIndex := 0
	for index, digit := range digits {
		if digit > resultDigit {
			resultDigit = digit
			resultIndex = index
		}
	}

	return resultDigit, resultIndex
}
