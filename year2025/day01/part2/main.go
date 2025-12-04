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
	adventofcode.Run(6, func(input string) (int, error) {
		lines := strings.Split(input, "\n")

		zeroCrosses := 0
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

			delta := 1
			if direction == 'L' {
				delta = -1
			}

			for i := 0; i < distance; i++ {
				dial += delta

				dial = dial % dialSize
				if dial < 0 {
					dial += dialSize
				}

				if dial == 0 {
					zeroCrosses++
				}
			}
		}

		return zeroCrosses, nil
	})
}
