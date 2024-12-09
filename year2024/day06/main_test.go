package main

import (
	"testing"

	"github.com/shoenig/test"
)

func Test_Grid_setObstacleAt(t *testing.T) {
	grid := Grid{
		[]rune(".."),
		[]rune(".."),
	}

	grid.setObstacleAt(Point{x: 0, y: 0})
	test.Eq(t, '#', grid.at(Point{x: 0, y: 0}))
}
