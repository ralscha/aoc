package grid

import "math"

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

type Grid2D[T any] struct {
	grid           map[Coordinate]T
	minCol, maxCol int
	minRow, maxRow int
	wrap           bool
}

func NewGrid2D[T any](wrap bool) Grid2D[T] {
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

func (g *Grid2D[T]) Set(row, col int, value T) {
	g.grid[Coordinate{
		Row: row,
		Col: col,
	}] = value
	g.updateMinMax(row, col)
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

func (g *Grid2D[T]) Move(row, col int, direction Direction) {
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
