package geomutil

import (
	"aoc/internal/gridutil"
	"aoc/internal/mathx"
	"math"
	"sort"
)

// Distance returns the Euclidean distance between two coordinates
func Distance(p1, p2 gridutil.Coordinate) float64 {
	dx := float64(p1.Col - p2.Col)
	dy := float64(p1.Row - p2.Row)
	return math.Sqrt(dx*dx + dy*dy)
}

// ManhattanDistance returns the Manhattan distance between two coordinates
func ManhattanDistance(p1, p2 gridutil.Coordinate) int {
	return mathx.Abs(p1.Col-p2.Col) + mathx.Abs(p1.Row-p2.Row)
}

// Area calculates the area of a polygon defined by a sequence of coordinates
func Area(points []gridutil.Coordinate) float64 {
	n := len(points)
	if n < 3 {
		return 0
	}

	area := 0
	for i := 0; i < n; i++ {
		j := (i + 1) % n
		area += points[i].Col * points[j].Row
		area -= points[j].Col * points[i].Row
	}
	return math.Abs(float64(area)) / 2.0
}

// PointInPolygon determines if a point is inside a polygon using ray casting
func PointInPolygon(point gridutil.Coordinate, polygon []gridutil.Coordinate) bool {
	inside := false
	j := len(polygon) - 1

	for i := 0; i < len(polygon); i++ {
		if ((polygon[i].Row > point.Row) != (polygon[j].Row > point.Row)) &&
			(float64(point.Col) < float64(polygon[j].Col-polygon[i].Col)*float64(point.Row-polygon[i].Row)/float64(polygon[j].Row-polygon[i].Row)+float64(polygon[i].Col)) {
			inside = !inside
		}
		j = i
	}

	return inside
}

// ConvexHull returns the convex hull of a set of points using Graham's scan
func ConvexHull(points []gridutil.Coordinate) []gridutil.Coordinate {
	n := len(points)
	if n < 3 {
		return points
	}

	// Find the bottommost point (and leftmost if tied)
	bottom := 0
	for i := 1; i < n; i++ {
		if points[i].Row < points[bottom].Row ||
			(points[i].Row == points[bottom].Row && points[i].Col < points[bottom].Col) {
			bottom = i
		}
	}

	// Swap the bottommost point to position 0
	points[0], points[bottom] = points[bottom], points[0]

	// Sort points by polar angle with respect to the bottommost point
	p0 := points[0]
	sortedPoints := points[1:]
	sortByPolarAngle(sortedPoints, p0)

	// Build convex hull
	hull := []gridutil.Coordinate{points[0], sortedPoints[0]}
	for i := 1; i < len(sortedPoints); i++ {
		for len(hull) > 1 && !rightTurn(hull[len(hull)-2], hull[len(hull)-1], sortedPoints[i]) {
			hull = hull[:len(hull)-1]
		}
		hull = append(hull, sortedPoints[i])
	}

	return hull
}

// Helper functions

func sortByPolarAngle(points []gridutil.Coordinate, p0 gridutil.Coordinate) {
	// Sort points by polar angle with respect to p0
	// Uses cross product comparison for efficiency (avoids actual angle calculation)
	sort.Slice(points, func(i, j int) bool {
		// Get vectors from p0 to points[i] and points[j]
		a := points[i]
		b := points[j]

		// Calculate cross product
		cross := crossProduct(p0, a, b)
		if cross != 0 {
			return cross > 0
		}

		// If cross product is 0 (points are collinear),
		// sort by distance from p0
		distA := squaredDistance(p0, a)
		distB := squaredDistance(p0, b)
		return distA < distB
	})
}

// crossProduct returns the cross product of vectors p0->p1 and p0->p2
func crossProduct(p0, p1, p2 gridutil.Coordinate) int {
	return (p1.Col-p0.Col)*(p2.Row-p0.Row) - (p1.Row-p0.Row)*(p2.Col-p0.Col)
}

// squaredDistance returns the squared Euclidean distance between two points
// Avoiding square root for efficiency in comparisons
func squaredDistance(p1, p2 gridutil.Coordinate) int {
	dx := p1.Col - p2.Col
	dy := p1.Row - p2.Row
	return dx*dx + dy*dy
}

func rightTurn(p1, p2, p3 gridutil.Coordinate) bool {
	// Returns true if p1->p2->p3 makes a right turn
	return crossProduct(p1, p2, p3) > 0
}
