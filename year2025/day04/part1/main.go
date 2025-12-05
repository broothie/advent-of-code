package main

import (
	"strings"

	"github.com/broothie/advent-of-code"
)

const (
	rollRune                     = '@'
	inaccessibleNumAdjacentRolls = 4
)

func main() {
	adventofcode.Run(13, func(input string) (int, error) {
		var grid [][]rune
		for _, line := range strings.Split(input, "\n") {
			grid = append(grid, []rune(line))
		}

		numAccessibleRolls := 0
		for y, row := range grid {
			for x, cell := range row {
				if cell != rollRune {
					continue
				}

				numAdjacentRolls := 0
				for _, offset := range offsets() {
					if offset.X == 0 && offset.Y == 0 {
						continue
					}

					lookupPair := Pair{
						X: x + offset.X,
						Y: y + offset.Y,
					}

					xInBounds := 0 <= lookupPair.X && lookupPair.X < len(grid[0])
					yInBounds := 0 <= lookupPair.Y && lookupPair.Y < len(grid)
					inBounds := xInBounds && yInBounds
					if !inBounds {
						continue
					}

					cellAtOffset := grid[lookupPair.Y][lookupPair.X]
					if cellAtOffset == rollRune {
						numAdjacentRolls++
					}
				}

				if numAdjacentRolls < inaccessibleNumAdjacentRolls {
					numAccessibleRolls++
				}
			}
		}

		return numAccessibleRolls, nil
	})
}

type Pair struct {
	X int
	Y int
}

func offsets() []Pair {
	return []Pair{
		{-1, -1}, {0, -1}, {1, -1},
		{-1, 0}, {0, 0}, {1, 0},
		{-1, 1}, {0, 1}, {1, 1},
	}
}
