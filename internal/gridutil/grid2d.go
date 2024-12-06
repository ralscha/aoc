package gridutil

import (
	"math"
)

type Coordinate struct {
	Row, Col int
}

type Direction struct {
	Row, Col int
}

var (
	DirectionN  = Direction{Row: -1}
	DirectionS  = Direction{Row: 1}
	DirectionE  = Direction{Col: 1}
	DirectionW  = Direction{Col: -1}
	DirectionNW = Direction{Row: -1, Col: -1}
	DirectionNE = Direction{Row: -1, Col: 1}
	DirectionSW = Direction{Row: 1, Col: -1}
	DirectionSE = Direction{Row: 1, Col: 1}
)

type Grid2D[T comparable] struct {
	grid           map[Coordinate]T
	minCol, maxCol int
	minRow, maxRow int
	wrap           bool
}

func TurnLeft(direction Direction) Direction {
	switch direction {
	case DirectionN:
		return DirectionW
	case DirectionW:
		return DirectionS
	case DirectionS:
		return DirectionE
	case DirectionE:
		return DirectionN
	}
	return direction
}

func TurnRight(direction Direction) Direction {
	switch direction {
	case DirectionN:
		return DirectionE
	case DirectionE:
		return DirectionS
	case DirectionS:
		return DirectionW
	case DirectionW:
		return DirectionN
	}
	return direction
}

func NewGrid2D[T comparable](wrap bool) Grid2D[T] {
	return Grid2D[T]{
		grid:   make(map[Coordinate]T),
		wrap:   wrap,
		maxCol: 0,
		maxRow: 0,
		minCol: math.MaxInt,
		minRow: math.MaxInt,
	}
}

func NewNumberGrid2D(lines []string) Grid2D[int] {
	g := NewGrid2D[int](false)

	for r, line := range lines {
		for c, j := range line {
			g.Set(r, c, int(j-'0'))
		}
	}

	return g
}

func NewCharGrid2D(lines []string) Grid2D[rune] {
	g := NewGrid2D[rune](false)

	for r, line := range lines {
		for c, j := range line {
			g.Set(r, c, j)
		}
	}

	return g
}

func (g *Grid2D[T]) SetMinRow(row int) {
	g.minRow = row
}

func (g *Grid2D[T]) SetMinRowCol(row, col int) {
	g.minRow = row
	g.minCol = col
}

func (g *Grid2D[T]) SetMaxRowCol(row, col int) {
	g.maxRow = row
	g.maxCol = col
}

func (g Grid2D[T]) Width() int {
	return g.maxCol - g.minCol + 1
}

func (g Grid2D[T]) Height() int {
	return g.maxRow - g.minRow + 1
}

func (g Grid2D[T]) GetMinMaxCol() (int, int) {
	return g.minCol, g.maxCol
}

func (g Grid2D[T]) GetMinMaxRow() (int, int) {
	return g.minRow, g.maxRow
}

func (g Grid2D[T]) Get(row, col int) (T, bool) {
	val, ok := g.grid[Coordinate{
		Row: row,
		Col: col,
	}]
	return val, ok
}

func (g Grid2D[T]) GetWithCoordinate(coord Coordinate) (T, bool) {
	val, ok := g.grid[coord]
	return val, ok
}

func (g *Grid2D[T]) Set(row, col int, value T) {
	g.grid[Coordinate{
		Row: row,
		Col: col,
	}] = value
	g.updateMinMax(row, col)
}

func (g *Grid2D[T]) SetWithCoordinate(coord Coordinate, value T) {
	g.grid[coord] = value
	g.updateMinMax(coord.Row, coord.Col)
}

func (g *Grid2D[T]) updateMinMax(row, col int) {
	if col < g.minCol {
		g.minCol = col
	}
	if col > g.maxCol {
		g.maxCol = col
	}
	if row < g.minRow {
		g.minRow = row
	}
	if row > g.maxRow {
		g.maxRow = row
	}
}

func (g *Grid2D[T]) Peek(row, col int, direction Direction) (T, bool) {
	newCoord := Coordinate{
		Row: row + direction.Row,
		Col: col + direction.Col,
	}

	if newCoord.Row >= g.minRow && newCoord.Row <= g.maxRow && newCoord.Col >= g.minCol && newCoord.Col <= g.maxCol {
		if val, ok := g.grid[newCoord]; ok {
			return val, false
		}
	}
	var zero T
	return zero, true
}

func (g Grid2D[T]) Copy() Grid2D[T] {
	newGrid := NewGrid2D[T](g.wrap)
	for coord, val := range g.grid {
		newGrid.SetWithCoordinate(coord, val)
	}
	return newGrid
}

func (g *Grid2D[T]) Move(row, col int, direction Direction) (int, int) {
	newCoord := Coordinate{
		Row: row + direction.Row,
		Col: col + direction.Col,
	}

	if g.wrap {
		height := g.maxRow - g.minRow + 1
		width := g.maxCol - g.minCol + 1
		newCoord.Row = newCoord.Row % height
		newCoord.Col = newCoord.Col % width
	}

	if newCoord.Row >= g.minRow && newCoord.Row <= g.maxRow && newCoord.Col >= g.minCol && newCoord.Col <= g.maxCol {
		oldCoord := Coordinate{
			Row: row,
			Col: col,
		}
		g.grid[newCoord] = g.grid[oldCoord]
		delete(g.grid, oldCoord)
	}

	g.updateMinMax(row, col)

	return newCoord.Row, newCoord.Col
}

func (g Grid2D[T]) GetNeighbours8(row, col int) []T {
	return g.GetNeighbours(row, col, []Direction{
		DirectionN,
		DirectionS,
		DirectionE,
		DirectionW,
		DirectionNW,
		DirectionNE,
		DirectionSW,
		DirectionSE,
	})
}

func (g Grid2D[T]) GetNeighbours4(row, col int) []T {
	return g.GetNeighbours(row, col, []Direction{
		DirectionN,
		DirectionS,
		DirectionE,
		DirectionW,
	})
}

func (g Grid2D[T]) GetNeighbours(row, col int, directions []Direction) []T {
	var neighbours []T
	for _, direction := range directions {
		neighbourCoord := Coordinate{
			Row: row + direction.Row,
			Col: col + direction.Col,
		}

		if g.wrap {
			height := g.maxRow - g.minRow + 1
			width := g.maxCol - g.minCol + 1
			neighbourCoord.Row = neighbourCoord.Row % height
			neighbourCoord.Col = neighbourCoord.Col % width
		}
		if neighbourCoord.Row >= g.minRow && neighbourCoord.Row <= g.maxRow &&
			neighbourCoord.Col >= g.minCol && neighbourCoord.Col <= g.maxCol {
			if neighbour, ok := g.grid[neighbourCoord]; ok {
				neighbours = append(neighbours, neighbour)
			}
		}
	}
	return neighbours
}

func (g *Grid2D[T]) Count() int {
	return len(g.grid)
}
