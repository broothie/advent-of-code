package main

import (
	"math"
	"sort"
	"strconv"
	"strings"

	adventofcode "github.com/broothie/advent-of-code"
)

func main() {
	adventofcode.Part1(2024, 1, 11, func(input string) (int, error) {
		var leftIDs, rightIDs []int
		for _, line := range strings.Split(input, "\n") {
			ids := strings.Split(line, "   ")

			leftID, err := strconv.Atoi(ids[0])
			if err != nil {
				return 0, err
			}

			rightID, err := strconv.Atoi(ids[1])
			if err != nil {
				return 0, err
			}

			leftIDs = append(leftIDs, leftID)
			rightIDs = append(rightIDs, rightID)
		}

		sort.Ints(leftIDs)
		sort.Ints(rightIDs)

		sum := 0
		for index := range leftIDs {
			leftID := leftIDs[index]
			rightID := rightIDs[index]

			sum += int(math.Abs(float64(leftID - rightID)))
		}

		return sum, nil
	})

	adventofcode.Part2(2024, 1, 31, func(input string) (int, error) {
		var leftIDs []int
		var rightIDs []int
		for _, line := range strings.Split(input, "\n") {
			ids := strings.Split(line, "   ")

			leftID, err := strconv.Atoi(ids[0])
			if err != nil {
				return 0, err
			}

			rightID, err := strconv.Atoi(ids[1])
			if err != nil {
				return 0, err
			}

			leftIDs = append(leftIDs, leftID)
			rightIDs = append(rightIDs, rightID)
		}

		counts := make(map[int]int)
		for _, rightID := range rightIDs {
			if _, found := counts[rightID]; !found {
				counts[rightID] = 0
			}

			counts[rightID] += 1
		}

		sum := 0
		for _, leftID := range leftIDs {
			sum += leftID * counts[leftID]
		}

		return sum, nil
	})
}
