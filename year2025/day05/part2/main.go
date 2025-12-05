package main

import (
	"sort"
	"strconv"
	"strings"

	"github.com/bobg/errors"
	"github.com/broothie/advent-of-code"
	"github.com/samber/lo"
)

// 307_661_172_093_586: too low
// 406_078_764_982_245: too high

func main() {
	adventofcode.Run(14, func(input string) (int, error) {
		chunks := strings.Split(input, "\n\n")
		rangesString := chunks[0]

		var idRanges []Range
		for _, rangeString := range strings.Split(rangesString, "\n") {
			idRange, err := parseRange(rangeString)
			if err != nil {
				return 0, errors.Wrap(err, "parsing range")
			}

			idRanges = append(idRanges, idRange)
		}

		sort.Slice(idRanges, func(i, j int) bool { return idRanges[i].Low < idRanges[j].Low })

		firstIDRange := idRanges[0].duplicate()
		combinedRanges := []*Range{&firstIDRange}
		for _, idRange := range idRanges[1:] {
			lastRange := combinedRanges[len(combinedRanges)-1]
			if idRange.Low <= lastRange.High+1 {
				if idRange.High > lastRange.High {
					lastRange.High = idRange.High
				}
			} else {
				duplicate := idRange.duplicate()
				combinedRanges = append(combinedRanges, &duplicate)
			}
		}

		return lo.SumBy(combinedRanges, func(idRange *Range) int { return idRange.size() }), nil
	})
}

type Range struct{ Low, High int }

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

func (r Range) duplicate() Range {
	return Range{Low: r.Low, High: r.High}
}

func (r Range) contains(id int) bool {
	return r.Low <= id && id <= r.High
}

func (r Range) size() int {
	return r.High - r.Low + 1
}
