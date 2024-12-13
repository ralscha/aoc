package geomutil

import (
	"aoc/internal/gridutil"
	"aoc/internal/mathx"
)

// ManhattanDistance returns the Manhattan distance between two coordinates
func ManhattanDistance(p1, p2 gridutil.Coordinate) int {
	return mathx.Abs(p1.Col-p2.Col) + mathx.Abs(p1.Row-p2.Row)
}
