package gridutil

import "aoc/internal/mathx"

// Coordinate3D represents a point in 3D space
type Coordinate3D struct {
	X, Y, Z int
}

// Add returns a new coordinate that is the sum of this coordinate and another
func (c Coordinate3D) Add(other Coordinate3D) Coordinate3D {
	return Coordinate3D{
		X: c.X + other.X,
		Y: c.Y + other.Y,
		Z: c.Z + other.Z,
	}
}

// Sub returns a new coordinate that is this coordinate minus another
func (c Coordinate3D) Sub(other Coordinate3D) Coordinate3D {
	return Coordinate3D{
		X: c.X - other.X,
		Y: c.Y - other.Y,
		Z: c.Z - other.Z,
	}
}

// ManhattanDistance returns the Manhattan distance between this coordinate and another
func (c Coordinate3D) ManhattanDistance(other Coordinate3D) int {
	return mathx.Abs(c.X-other.X) + mathx.Abs(c.Y-other.Y) + mathx.Abs(c.Z-other.Z)
}
