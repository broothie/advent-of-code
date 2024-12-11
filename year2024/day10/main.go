package main

import (
	"strconv"
	"strings"

	"github.com/bobg/errors"
	adventofcode "github.com/broothie/advent-of-code"
	"github.com/samber/lo"
)

type TopoMap [][]int

func (m TopoMap) width() int {
	return len(m[0])
}

func (m TopoMap) height() int {
	return len(m)
}

func (m TopoMap) isInBounds(point Point) bool {
	isInXBounds := 0 <= point.x && point.x < m.width()
	isInYBounds := 0 <= point.y && point.y < m.height()

	return isInXBounds && isInYBounds
}

func (m TopoMap) elevationAt(point Point) (int, bool) {
	if !m.isInBounds(point) {
		return 0, false
	}

	return m[point.y][point.x], true
}

func (m TopoMap) trailheads() []Point {
	var trailheads []Point
	for y, row := range m {
		for x, elevation := range row {
			if elevation == 0 {
				trailheads = append(trailheads, Point{x: x, y: y})
			}
		}
	}

	return trailheads
}

type Point struct {
	x, y int
}

type Vector struct {
	x, y int
}

func (p Point) translate(delta Vector) Point {
	return Point{
		x: p.x + delta.x,
		y: p.y + delta.y,
	}
}

func main() {
	adventofcode.Part1(2024, 10, 36, func(input string) (int, error) {
		topo, err := parse(input)
		if err != nil {
			return 0, errors.Wrap(err, "parsing")
		}

		score := 0
		for _, trailhead := range topo.trailheads() {
			score += len(findTrailScores(trailhead, topo))
		}

		return score, nil
	})

	adventofcode.Part2(2024, 10, 81, func(input string) (int, error) {
		topo, err := parse(input)
		if err != nil {
			return 0, errors.Wrap(err, "parsing")
		}

		rating := 0
		for _, trailhead := range topo.trailheads() {
			rating += findTrailRatings(trailhead, topo)
		}

		return rating, nil
	})
}

func parse(input string) (TopoMap, error) {
	var topo TopoMap
	for _, line := range strings.Split(input, "\n") {
		var row []int
		for _, character := range line {
			elevation, err := strconv.Atoi(string(character))
			if err != nil {
				return nil, errors.Wrap(err, "parsing elevation")
			}

			row = append(row, elevation)
		}

		topo = append(topo, row)
	}
	return topo, nil
}

func stepDeltas() []Vector {
	return []Vector{
		{x: 0, y: -1}, // Up
		{x: 0, y: 1},  // Down
		{x: -1, y: 0}, // Left
		{x: 1, y: 0},  // Right
	}
}

func findTrailScores(point Point, topo TopoMap) map[Point]bool {
	elevation, _ := topo.elevationAt(point)
	if elevation == 9 {
		return map[Point]bool{point: true}
	}

	nextPoints := lo.FilterMap(stepDeltas(), func(delta Vector, _ int) (Point, bool) {
		nextPoint := point.translate(delta)
		nextElevation, ok := topo.elevationAt(nextPoint)
		return nextPoint, ok && nextElevation == elevation+1
	})

	topPoints := make(map[Point]bool)
	for _, nextPoint := range nextPoints {
		topPoints = lo.Assign(topPoints, findTrailScores(nextPoint, topo))
	}

	return topPoints
}

func findTrailRatings(point Point, topo TopoMap) int {
	elevation, _ := topo.elevationAt(point)
	if elevation == 9 {
		return 1
	}

	nextPoints := lo.FilterMap(stepDeltas(), func(delta Vector, _ int) (Point, bool) {
		nextPoint := point.translate(delta)
		nextElevation, ok := topo.elevationAt(nextPoint)
		return nextPoint, ok && nextElevation == elevation+1
	})

	rating := 0
	for _, nextPoint := range nextPoints {
		rating += findTrailRatings(nextPoint, topo)
	}

	return rating
}
