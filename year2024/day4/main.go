package main

import (
	"strings"

	adventofcode "github.com/broothie/advent-of-code"
	"github.com/samber/lo"
)

const (
	XMAS = "XMAS"

	Up    = -1
	Down  = 1
	Left  = -1
	Right = 1
)

type Point struct {
	X, Y int
}

func (p Point) Translate(vector Vector) Point {
	return Point{
		X: p.X + vector.X,
		Y: p.Y + vector.Y,
	}
}

type Vector Point

func (v Vector) Scale(magnitude int) Vector {
	return Vector{
		X: v.X * magnitude,
		Y: v.Y * magnitude,
	}
}

func (v Vector) Normalize() Vector {
	x := v.X
	if x > 0 {
		x = x / x
	}

	y := v.Y
	if y > 0 {
		y = y / y
	}

	return Vector{
		X: x,
		Y: y,
	}
}

type WordSearch [][]rune

func (w WordSearch) width() int {
	return len(w[0])
}

func (w WordSearch) height() int {
	return len(w)
}

func (w WordSearch) pointIsInBounds(point Point) bool {
	xInBounds := 0 <= point.X && point.X < w.width()
	yInBounds := 0 <= point.Y && point.Y < w.height()

	return xInBounds && yInBounds
}

func (w WordSearch) get(point Point) (rune, bool) {
	if !w.pointIsInBounds(point) {
		return 0, false
	}

	return w[point.Y][point.X], true
}

func main() {
	adventofcode.Part1(2024, 4, 18, func(input string) (int, error) {
		var wordSearch WordSearch
		for _, line := range strings.Split(input, "\n") {
			wordSearch = append(wordSearch, []rune(line))
		}

		sum := 0
		for y, line := range wordSearch {
			for x := range line {
				sum += checkForXmasAtPoint(wordSearch, Point{X: x, Y: y})
			}
		}

		return sum, nil
	})

	adventofcode.Part2(2024, 4, 9, func(input string) (int, error) {
		var wordSearch WordSearch
		for _, line := range strings.Split(input, "\n") {
			wordSearch = append(wordSearch, []rune(line))
		}

		sum := 0
		for y, line := range wordSearch {
			for x := range line {
				if checkForCrossmasesAtPoint(wordSearch, Point{X: x, Y: y}) {
					sum += 1
				}
			}
		}

		return sum, nil
	})
}

var CrossmasTemplates = [][]string{
	{
		"M.M",
		".A.",
		"S.S",
	},
	{
		"S.M",
		".A.",
		"S.M",
	},
	{
		"M.S",
		".A.",
		"M.S",
	},
	{
		"S.S",
		".A.",
		"M.M",
	},
}

func checkForCrossmasAtPoint(wordSearch WordSearch, start Point, crossmasTemplate []string) bool {
	for yOffset, crossmasLine := range crossmasTemplate {
		for xOffset, crossmasCharacter := range crossmasLine {
			if crossmasCharacter == '.' {
				continue
			}

			wordSearchCharacter, isInBounds := wordSearch.get(start.Translate(Vector{X: xOffset, Y: yOffset}))
			if !isInBounds {
				return false
			}

			if crossmasCharacter != wordSearchCharacter {
				return false
			}
		}
	}

	return true
}

func checkForCrossmasesAtPoint(wordSearch WordSearch, start Point) bool {
	return lo.SomeBy(CrossmasTemplates, func(crossmasTemplate []string) bool {
		return checkForCrossmasAtPoint(wordSearch, start, crossmasTemplate)
	})
}

func checkForXmasInDirection(wordSearch WordSearch, start Point, direction Vector) bool {
	for index, xmasCharacter := range XMAS {
		point := start.Translate(direction.Normalize().Scale(index))
		wordSearchCharacter, isInBounds := wordSearch.get(point)
		if !isInBounds {
			return false
		}

		if xmasCharacter != wordSearchCharacter {
			return false
		}
	}

	return true
}

func checkForXmasAtPoint(wordSearch WordSearch, start Point) int {
	directions := []Vector{
		{Up, Left}, {Up, 0}, {Up, Right},
		{0, Left}, {0, 0}, {0, Right},
		{Down, Left}, {Down, 0}, {Down, Right},
	}

	return lo.CountBy(directions, func(direction Vector) bool {
		return checkForXmasInDirection(wordSearch, start, direction)
	})
}
