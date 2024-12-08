package main

import (
	"strings"

	adventofcode "github.com/broothie/advent-of-code"
	"github.com/samber/lo"
)

var directions = []string{"up", "right", "down", "left"}

type Point struct {
	x, y int
}

func (p Point) inBounds(width, height int) bool {
	inXBounds := 0 <= p.x && p.x < width
	inYBounds := 0 <= p.y && p.y < height

	return inXBounds && inYBounds
}

type Guard struct {
	position  Point
	direction string
}

func parse(input string) (int, int, Guard, map[Point]bool) {
	guard := Guard{direction: "up"}
	obstacles := make(map[Point]bool)

	lines := strings.Split(input, "\n")
	width := len(lines[0])
	height := len(lines)
	for y, line := range lines {
		for x, character := range line {
			if character == '#' {
				obstacles[Point{x: x, y: y}] = true
			}

			if lo.Contains([]rune{'^', '>', 'v', '<'}, character) {
				guard.position = Point{x: x, y: y}

				switch character {
				case '^':
					guard.direction = "up"
				case '>':
					guard.direction = "right"
				case 'v':
					guard.direction = "down"
				case '<':
					guard.direction = "left"
				}
			}
		}
	}
	return width, height, guard, obstacles
}

func simulateMap(width, height int, guard Guard, obstacles map[Point]bool) (_ map[Guard]bool, looped bool) {
	guards := make(map[Guard]bool)

	for guard.position.inBounds(width, height) {
		if guards[guard] {
			return guards, true
		}

		guards[guard] = true

		for i := 0; i < 3; i++ {
			var nextPosition Point
			switch guard.direction {
			case "up":
				nextPosition = Point{
					x: guard.position.x,
					y: guard.position.y - 1,
				}

			case "right":
				nextPosition = Point{
					x: guard.position.x + 1,
					y: guard.position.y,
				}

			case "down":
				nextPosition = Point{
					x: guard.position.x,
					y: guard.position.y + 1,
				}

			case "left":
				nextPosition = Point{
					x: guard.position.x - 1,
					y: guard.position.y,
				}
			}

			if !obstacles[nextPosition] {
				guard.position = nextPosition
				break
			}

			directionIndex := lo.IndexOf(directions, guard.direction)
			guard.direction = directions[(directionIndex+1)%len(directions)]
		}
	}

	return guards, false
}

func main() {
	adventofcode.Part1(2024, 6, 41, func(input string) (int, error) {
		guards, _ := simulateMap(parse(input))
		positions := lo.Map(lo.Keys(guards), func(guard Guard, _ int) Point { return guard.position })
		return len(lo.Uniq(positions)), nil
	})

	adventofcode.Part2(2024, 6, 6, func(input string) (int, error) {
		width, height, guard, obstacles := parse(input)

		sum := 0
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				position := Point{x: x, y: y}
				if guard.position == position {
					continue
				}

				if obstacles[position] {
					continue
				}

				obstaclesDup := lo.SliceToMap(lo.Keys(obstacles), func(point Point) (Point, bool) { return point, true })
				obstaclesDup[position] = true
				if _, looped := simulateMap(width, height, guard, obstaclesDup); looped {
					sum += 1
				}
			}
		}

		return sum, nil
	})
}
