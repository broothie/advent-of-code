package main

import (
	"regexp"
	"strconv"
	"strings"

	adventofcode "github.com/broothie/advent-of-code"
)

var mulRegexp = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

func main() {
	adventofcode.Part1(2024, 3, 161, func(input string) (int, error) {
		sum := 0
		for _, captures := range mulRegexp.FindAllStringSubmatch(input, -1) {
			left, err := strconv.Atoi(captures[1])
			if err != nil {
				return 0, err
			}

			right, err := strconv.Atoi(captures[2])
			if err != nil {
				return 0, err
			}

			sum += left * right
		}

		return sum, nil
	})

	adventofcode.Part2(2024, 3, 48, func(input string) (int, error) {
		sum := 0

		enabled := true
		for index := range input {
			if strings.HasPrefix(input[index:], "do()") {
				enabled = true
				continue
			}

			if strings.HasPrefix(input[index:], "don't()") {
				enabled = false
				continue
			}

			if enabled {
				if loc := mulRegexp.FindStringIndex(input[index:]); loc != nil && loc[0] == 0 {
					captures := mulRegexp.FindStringSubmatch(input[index:][loc[0]:loc[1]])
					left, err := strconv.Atoi(captures[1])
					if err != nil {
						return 0, err
					}

					right, err := strconv.Atoi(captures[2])
					if err != nil {
						return 0, err
					}

					sum += left * right
				}
			}
		}

		return sum, nil
	})
}
