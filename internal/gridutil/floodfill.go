package gridutil

import (
	"aoc/internal/container"
)

type FloodFillResult struct {
	Count    int
	MinRow   int
	MaxRow   int
	MinCol   int
	MaxCol   int
	Visited  *container.Set[Coordinate]
}

func (g *Grid2D[T]) FloodFill(start Coordinate, shouldFill func(current, neighbor T) bool) FloodFillResult {
	visited := container.NewSet[Coordinate]()
	visited.Add(start)
	
	minRow, maxRow := start.Row, start.Row
	minCol, maxCol := start.Col, start.Col
	
	queue := []Coordinate{start}
	startValue, _ := g.GetC(start)
	
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		
		for _, dir := range Get4Directions() {
			next := Coordinate{current.Row + dir.Row, current.Col + dir.Col}
			if nextValue, ok := g.GetC(next); ok && !visited.Contains(next) {
				if shouldFill(startValue, nextValue) {
					visited.Add(next)
					queue = append(queue, next)
					
					if next.Row < minRow {
						minRow = next.Row
					}
					if next.Row > maxRow {
						maxRow = next.Row
					}
					if next.Col < minCol {
						minCol = next.Col
					}
					if next.Col > maxCol {
						maxCol = next.Col
					}
				}
			}
		}
	}
	
	return FloodFillResult{
		Count:   visited.Len(),
		MinRow:  minRow,
		MaxRow:  maxRow,
		MinCol:  minCol,
		MaxCol:  maxCol,
		Visited: visited,
	}
}

func (g *Grid2D[T]) FloodFill8(start Coordinate, shouldFill func(current, neighbor T) bool) FloodFillResult {
	visited := container.NewSet[Coordinate]()
	visited.Add(start)
	
	minRow, maxRow := start.Row, start.Row
	minCol, maxCol := start.Col, start.Col
	
	queue := []Coordinate{start}
	startValue, _ := g.GetC(start)
	
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		
		for _, dir := range Get8Directions() {
			next := Coordinate{current.Row + dir.Row, current.Col + dir.Col}
			if nextValue, ok := g.GetC(next); ok && !visited.Contains(next) {
				if shouldFill(startValue, nextValue) {
					visited.Add(next)
					queue = append(queue, next)
					
					if next.Row < minRow {
						minRow = next.Row
					}
					if next.Row > maxRow {
						maxRow = next.Row
					}
					if next.Col < minCol {
						minCol = next.Col
					}
					if next.Col > maxCol {
						maxCol = next.Col
					}
				}
			}
		}
	}
	
	return FloodFillResult{
		Count:   visited.Len(),
		MinRow:  minRow,
		MaxRow:  maxRow,
		MinCol:  minCol,
		MaxCol:  maxCol,
		Visited: visited,
	}
}
