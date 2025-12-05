package main

import (
	"strconv"
	"strings"

	"github.com/bobg/errors"
	"github.com/broothie/advent-of-code"
	"github.com/samber/lo"
)

func main() {
	adventofcode.Run(3, func(input string) (int, error) {
		chunks := strings.Split(input, "\n\n")
		rangesString, idsString := chunks[0], chunks[1]

		var idRanges []Range
		for _, rangeString := range strings.Split(rangesString, "\n") {
			idRange, err := parseRange(rangeString)
			if err != nil {
				return 0, errors.Wrap(err, "parsing range")
			}

			idRanges = append(idRanges, idRange)
		}

		numFreshIDs := 0
		for _, idString := range strings.Split(idsString, "\n") {
			id, err := strconv.Atoi(idString)
			if err != nil {
				return 0, errors.Wrap(err, "parsing id")
			}

			if lo.SomeBy(idRanges, func(idRange Range) bool { return idRange.contains(id) }) {
				numFreshIDs++
			}
		}

		return numFreshIDs, nil
	})
}

type Range struct {
	Low, High int
}

func parseRange(rangeString string) (Range, error) {
	split := strings.SplitN(rangeString, "-", 2)

	low, err := strconv.Atoi(split[0])
	if err != nil {
		return Range{}, errors.Wrap(err, "parsing range low")
	}

	high, err := strconv.Atoi(split[1])
	if err != nil {
		return Range{}, errors.Wrap(err, "parsing range high")
	}

	return Range{Low: low, High: high}, nil
}

func (r Range) contains(id int) bool {
	return r.Low <= id && id <= r.High
}
