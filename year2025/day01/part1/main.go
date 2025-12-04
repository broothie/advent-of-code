package main

import (
	"strconv"
	"strings"

	"github.com/bobg/errors"
	"github.com/broothie/advent-of-code"
)

const (
	dialStart = 50
	dialSize  = 100
)

func main() {
	adventofcode.Run(3, func(input string) (int, error) {
		lines := strings.Split(input, "\n")

		zeroes := 0
		dial := dialStart
		for _, line := range lines {
			if line == "" {
				continue
			}

			direction := line[0]
			distance, err := strconv.Atoi(line[1:])
			if err != nil {
				return 0, errors.Wrap(err, "parsing distance")
			}

			switch direction {
			case 'R':
				dial += distance

			case 'L':
				dial -= distance

			default:
				return 0, errors.Errorf("unknown direction %q", direction)
			}

			dial = dial % dialSize
			if dial < 0 {
				dial += dialSize
			}

			if dial == 0 {
				zeroes++
			}
		}

		return zeroes, nil
	})
}
