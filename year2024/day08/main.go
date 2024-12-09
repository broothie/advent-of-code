package main

import (
	"strings"

	adventofcode "github.com/broothie/advent-of-code"
	"github.com/samber/lo"
)

type Point struct {
	x, y int
}

type Vector struct {
	x, y int
}

func (p Point) difference(other Point) Vector {
	return Vector{
		x: p.x - other.x,
		y: p.y - other.y,
	}
}

func (p Point) translate(vector Vector) Point {
	return Point{
		x: p.x + vector.x,
		y: p.y + vector.y,
	}
}

func (p Point) isInBounds(width, height int) bool {
	isInXBounds := 0 <= p.x && p.x < width
	isInYBounds := 0 <= p.y && p.y < height

	return isInXBounds && isInYBounds
}

func parse(input string) (int, int, map[rune][]Point) {
	lines := strings.Split(input, "\n")
	width := len(lines[0])
	height := len(lines)

	antennae := make(map[rune][]Point)
	for y, line := range lines {
		for x, character := range line {
			if character == '.' {
				continue
			}

			antennae[character] = append(antennae[character], Point{x: x, y: y})
		}
	}
	return width, height, antennae
}

func main() {
	adventofcode.Part1(2024, 8, 14, func(input string) (int, error) {
		width, height, antennae := parse(input)

		antinodes := make(map[Point]bool)
		for _, antennaPoints := range antennae {
			for index, antennaPoint := range antennaPoints {
				otherPoints := lo.Filter(antennaPoints, func(_ Point, otherPointIndex int) bool { return index != otherPointIndex })
				for _, otherPoint := range otherPoints {
					antinode := otherPoint.translate(otherPoint.difference(antennaPoint))
					if !antinode.isInBounds(width, height) {
						continue
					}

					antinodes[antinode] = true
				}
			}
		}

		return len(antinodes), nil
	})

	adventofcode.Part2(2024, 8, 34, func(input string) (int, error) {
		width, height, antennae := parse(input)

		antinodes := make(map[Point]bool)
		for _, antennaPoints := range antennae {
			for index, antennaPoint := range antennaPoints {
				otherPoints := lo.Filter(antennaPoints, func(_ Point, otherPointIndex int) bool { return index != otherPointIndex })
				for _, otherPoint := range otherPoints {
					difference := otherPoint.difference(antennaPoint)

					for currentAntinode := antennaPoint; currentAntinode.isInBounds(width, height); currentAntinode = currentAntinode.translate(difference) {
						antinodes[currentAntinode] = true
					}
				}
			}
		}

		return len(antinodes), nil
	})
}
