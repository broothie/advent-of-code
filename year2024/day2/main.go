package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	adventofcode "github.com/broothie/advent-of-code"
	"github.com/samber/lo"
)

func main() {
	adventofcode.Part1(2024, 2, 2, func(input string) (int, error) {
		var reports [][]int
		for _, line := range strings.Split(input, "\n") {
			var report []int
			for _, rawLevel := range strings.Fields(line) {
				level, err := strconv.Atoi(rawLevel)
				if err != nil {
					return 0, err
				}

				report = append(report, level)
			}

			reports = append(reports, report)
		}

		return lo.CountBy(reports, isSafeReport), nil
	})

	adventofcode.Part2(2024, 2, 4, func(input string) (int, error) {
		var reports [][]int
		for _, line := range strings.Split(input, "\n") {
			var report []int
			for _, rawLevel := range strings.Fields(line) {
				level, err := strconv.Atoi(rawLevel)
				if err != nil {
					return 0, err
				}

				report = append(report, level)
			}

			reports = append(reports, report)
		}

		return lo.CountBy(reports, func(report []int) bool {
			reportVersions := [][]int{report}
			for index := range report {
				reportVersions = append(reportVersions, lo.DropByIndex(report, index))
			}

			return lo.SomeBy(reportVersions, isSafeReport)
		}), nil
	})
}

func isSafeReport(report []int) bool {
	sorted := slices.Sorted(slices.Values(report))
	allIncreasing := fmt.Sprint(report) == fmt.Sprint(sorted)
	slices.Reverse(sorted)
	allDecreasing := fmt.Sprint(report) == fmt.Sprint(sorted)
	if !allIncreasing && !allDecreasing {
		return false
	}

	for index := range report[:len(report)-1] {
		a, b := report[index], report[index+1]
		difference := a - b
		if difference < 0 {
			difference = -difference
		}

		if !lo.Contains([]int{1, 2, 3}, difference) {
			return false
		}
	}

	return true
}
