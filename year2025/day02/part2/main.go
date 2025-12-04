package main

import (
	"strconv"
	"strings"

	"github.com/broothie/advent-of-code"
	"github.com/samber/lo"
)

func main() {
	adventofcode.Run(4174379265, func(input string) (int, error) {
		sum := 0
		for _, idRange := range strings.Split(input, ",") {
			split := strings.SplitN(idRange, "-", 2)
			low, _ := strconv.Atoi(split[0])
			high, _ := strconv.Atoi(split[1])

			for id := low; id <= high; id++ {
				if !isValidID(id) {
					sum += id
				}
			}
		}

		return sum, nil
	})
}

func isValidID(id int) bool {
	idString := strconv.Itoa(id)
	idStringLength := len(idString)
	for length := 1; length <= idStringLength/2; length++ {
		if idStringLength%length != 0 {
			continue
		}

		sequences := lo.Map(lo.Chunk([]rune(idString), length), func(sequence []rune, _ int) string { return string(sequence) })
		if lo.EveryBy(sequences, func(sequence string) bool { return sequence == sequences[0] }) {
			return false
		}
	}

	return true
}
