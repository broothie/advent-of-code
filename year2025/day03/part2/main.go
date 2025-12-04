package main

import (
	"strconv"
	"strings"

	"github.com/bobg/errors"
	"github.com/broothie/advent-of-code"
)

const numBatteries = 12

func main() {
	adventofcode.Run(3121910778619, func(input string) (int, error) {
		sum := 0

		lines := strings.Split(input, "\n")
		for _, line := range lines {
			if line == "" {
				continue
			}

			bank, err := parseBank(line)
			if err != nil {
				return 0, errors.Wrap(err, "parsing bank")
			}

			joltage, err := joltageForBank(bank)
			if err != nil {
				return 0, errors.Wrap(err, "calculating joltage")
			}

			sum += joltage
		}

		return sum, nil
	})
}

func parseBank(line string) ([]int, error) {
	var bank []int
	for _, joltageChar := range line {
		joltage, err := strconv.Atoi(string(joltageChar))
		if err != nil {
			return nil, errors.Wrap(err, "parsing digit")
		}

		bank = append(bank, joltage)
	}

	return bank, nil
}

func joltageForBank(bank []int) (int, error) {
	startWindowIndex := 0
	joltageString := ""
	for len(joltageString) < numBatteries {
		numJoltagesFound := len(joltageString)
		numBatteriesToSave := numBatteries - numJoltagesFound - 1
		endWindowIndex := len(bank) - numBatteriesToSave
		window := bank[startWindowIndex:endWindowIndex]

		enabledBatteryJoltage, foundIndex := maxWithIndex(window)
		startWindowIndex = startWindowIndex + foundIndex + 1
		joltageString += strconv.Itoa(enabledBatteryJoltage)
	}

	joltage, err := strconv.Atoi(joltageString)
	if err != nil {
		return 0, errors.Wrap(err, "parsing joltage")
	}

	return joltage, nil
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
