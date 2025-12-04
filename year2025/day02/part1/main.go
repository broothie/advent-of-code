package main

import (
	"strconv"
	"strings"

	"github.com/broothie/advent-of-code"
)

func main() {
	adventofcode.Run(1227775554, func(input string) (int, error) {
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
	length := len(idString)
	if length%2 != 0 {
		return true
	}

	firstHalf, secondHalf := idString[:length/2], idString[length/2:]
	return firstHalf != secondHalf
}
