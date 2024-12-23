package gridutil

import (
	"maps"
	"math"
	"slices"
)

// Coordinate represents a position in a 2D grid with row and column values.
type Coordinate struct {
	Row, Col int
}

// Direction represents a movement vector in a 2D grid with row and column deltas.
type Direction struct {
	Row, Col int
}

var (
	// DirectionN Predefined cardinal directions
	DirectionN  = Direction{Row: -1}          // North: Move up one row
	DirectionS  = Direction{Row: 1}           // South: Move down one row
	DirectionE  = Direction{Col: 1}           // East: Move right one column
	DirectionW  = Direction{Col: -1}          // West: Move left one column
	DirectionNW = Direction{Row: -1, Col: -1} // Northwest: Move up and left
	DirectionNE = Direction{Row: -1, Col: 1}  // Northeast: Move up and right
	DirectionSW = Direction{Row: 1, Col: -1}  // Southwest: Move down and left
	DirectionSE = Direction{Row: 1, Col: 1}   // Southeast: Move down and right

	// Slice containing all 8 possible directions (cardinal and diagonal)
	directions8 = []Direction{
		DirectionN,
		DirectionS,
		DirectionE,
		DirectionW,
		DirectionNW,
		DirectionNE,
		DirectionSW,
		DirectionSE,
	}

	// Slice containing the 4 cardinal directions
	directions4 = []Direction{
		DirectionN,
		DirectionS,
		DirectionE,
		DirectionW,
	}
)

// Grid2D represents a two-dimensional grid with generic type T.
// It supports wrapping around edges if wrap is set to true.
type Grid2D[T comparable] struct {
	grid           map[Coordinate]T
	minCol, maxCol int
	minRow, maxRow int
	wrap           bool
}

// Get8Directions returns a slice containing all 8 possible directions (cardinal and diagonal).
func Get8Directions() []Direction {
	return slices.Clone(directions8)
}

// Get4Directions returns a slice containing the 4 cardinal directions.
func Get4Directions() []Direction {
	return slices.Clone(directions4)
}

// TurnLeft returns a new direction after rotating the given direction 90 degrees counterclockwise.
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

// TurnRight returns a new direction after rotating the given direction 90 degrees clockwise.
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

// NewGrid2D creates a new empty grid with the specified wrap behavior.
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

// NewNumberGrid2D creates a new grid from a slice of strings, converting each character to its numeric value.
// For example, '0' becomes 0, '1' becomes 1, etc.
func NewNumberGrid2D(lines []string) Grid2D[int] {
	g := NewGrid2D[int](false)

	for r, line := range lines {
		for c, j := range line {
			g.Set(r, c, int(j-'0'))
		}
	}

	return g
}

// NewCharGrid2D creates a new grid from a slice of strings, storing each character as a rune.
func NewCharGrid2D(lines []string) Grid2D[rune] {
	g := NewGrid2D[rune](false)

	for r, line := range lines {
		for c, j := range line {
			g.Set(r, c, j)
		}
	}

	return g
}

// SetMinRow sets the minimum row value for the grid.
func (g *Grid2D[T]) SetMinRow(row int) {
	g.minRow = row
}

// SetMinRowCol sets both the minimum row and column values for the grid.
func (g *Grid2D[T]) SetMinRowCol(row, col int) {
	g.minRow = row
	g.minCol = col
}

// SetMaxRowCol sets both the maximum row and column values for the grid.
func (g *Grid2D[T]) SetMaxRowCol(row, col int) {
	g.maxRow = row
	g.maxCol = col
}

// SetMaxRowColC sets the maximum row and column values from a Coordinate.
func (g *Grid2D[T]) SetMaxRowColC(coord Coordinate) {
	g.maxRow = coord.Row
	g.maxCol = coord.Col
}

// Width returns the width of the grid (number of columns).
func (g Grid2D[T]) Width() int {
	return g.maxCol - g.minCol + 1
}

// Height returns the height of the grid (number of rows).
func (g Grid2D[T]) Height() int {
	return g.maxRow - g.minRow + 1
}

// Count returns the total number of elements in the grid.
func (g *Grid2D[T]) Count() int {
	return len(g.grid)
}

// GetMinMaxCol returns the minimum and maximum column values of the grid.
func (g Grid2D[T]) GetMinMaxCol() (int, int) {
	return g.minCol, g.maxCol
}

// GetMinMaxRow returns the minimum and maximum row values of the grid.
func (g Grid2D[T]) GetMinMaxRow() (int, int) {
	return g.minRow, g.maxRow
}

// Get retrieves the value at the specified row and column.
// Returns the value and true if found, zero value and false if not found.
func (g Grid2D[T]) Get(row, col int) (T, bool) {
	val, ok := g.grid[Coordinate{
		Row: row,
		Col: col,
	}]
	return val, ok
}

// GetC retrieves the value at the specified coordinate.
// Returns the value and true if found, zero value and false if not found.
func (g Grid2D[T]) GetC(coord Coordinate) (T, bool) {
	val, ok := g.grid[coord]
	return val, ok
}

// Set stores a value at the specified row and column.
func (g *Grid2D[T]) Set(row, col int, value T) {
	g.grid[Coordinate{
		Row: row,
		Col: col,
	}] = value
	g.updateMinMax(row, col)
}

// SetC stores a value at the specified coordinate.
func (g *Grid2D[T]) SetC(coord Coordinate, value T) {
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

// Peek looks at the value in the specified direction from the given position.
// Returns the value and true if out of bounds, value and false if within bounds.
func (g *Grid2D[T]) Peek(row, col int, direction Direction) (T, bool) {
	return g.PeekC(Coordinate{Row: row, Col: col}, direction)
}

// PeekC looks at the value in the specified direction from the given coordinate.
// Returns the value and true if out of bounds, value and false if within bounds.
func (g *Grid2D[T]) PeekC(coord Coordinate, direction Direction) (T, bool) {
	newCoord := Coordinate{
		Row: coord.Row + direction.Row,
		Col: coord.Col + direction.Col,
	}

	if newCoord.Row >= g.minRow && newCoord.Row <= g.maxRow && newCoord.Col >= g.minCol && newCoord.Col <= g.maxCol {
		if val, ok := g.grid[newCoord]; ok {
			return val, false
		}
	}
	var zero T
	return zero, true
}

// Copy creates a deep copy of the grid.
func (g Grid2D[T]) Copy() Grid2D[T] {
	newGrid := NewGrid2D[T](g.wrap)
	newGrid.grid = maps.Clone(g.grid)
	newGrid.minCol = g.minCol
	newGrid.maxCol = g.maxCol
	newGrid.minRow = g.minRow
	newGrid.maxRow = g.maxRow
	return newGrid
}

// GetNeighbours8 returns all 8 neighboring values (including diagonals) for the given position.
func (g Grid2D[T]) GetNeighbours8(row, col int) []T {
	return g.GetNeighboursC(Coordinate{Row: row, Col: col}, directions8)
}

// GetNeighbours8C returns all 8 neighboring values (including diagonals) for the given coordinate.
func (g Grid2D[T]) GetNeighbours8C(coord Coordinate) []T {
	return g.GetNeighboursC(coord, directions8)
}

// GetNeighbours4 returns the 4 neighboring values (cardinal directions only) for the given position.
func (g Grid2D[T]) GetNeighbours4(row, col int) []T {
	return g.GetNeighboursC(Coordinate{Row: row, Col: col}, directions4)
}

// GetNeighbours4C returns the 4 neighboring values (cardinal directions only) for the given coordinate.
func (g Grid2D[T]) GetNeighbours4C(coord Coordinate) []T {
	return g.GetNeighboursC(coord, directions4)
}

// GetNeighbours returns neighboring values for the given position in the specified directions.
func (g Grid2D[T]) GetNeighbours(row, col int, directions []Direction) []T {
	return g.GetNeighboursC(Coordinate{Row: row, Col: col}, directions)
}

// GetNeighboursC returns neighboring values for the given coordinate in the specified directions.
func (g Grid2D[T]) GetNeighboursC(coord Coordinate, directions []Direction) []T {
	var neighbours []T
	for _, direction := range directions {
		neighbourCoord := Coordinate{
			Row: coord.Row + direction.Row,
			Col: coord.Col + direction.Col,
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

// RotateRow rotates the values in the specified row by the given amount.
// Positive amount rotates right, negative amount rotates left.
func (g *Grid2D[T]) RotateRow(row, amount int) {
	width := g.Width()
	if width == 0 {
		return
	}

	// Normalize amount to be within grid width
	amount = ((amount % width) + width) % width
	if amount == 0 {
		return
	}

	// Create a temporary copy of the row
	temp := make([]T, width)
	for i := range width {
		if val, ok := g.Get(row, g.minCol+i); ok {
			temp[i] = val
		}
	}

	// Rotate and set values back
	for i := range width {
		newPos := (i + amount) % width
		g.Set(row, g.minCol+newPos, temp[i])
	}
}

// RotateColumn rotates the values in the specified column by the given amount.
// Positive amount rotates down, negative amount rotates up.
func (g *Grid2D[T]) RotateColumn(col, amount int) {
	height := g.Height()
	if height == 0 {
		return
	}

	// Normalize amount to be within grid height
	amount = ((amount % height) + height) % height
	if amount == 0 {
		return
	}

	// Create a temporary copy of the column
	temp := make([]T, height)
	for i := range height {
		if val, ok := g.Get(g.minRow+i, col); ok {
			temp[i] = val
		}
	}

	// Rotate and set values back
	for i := range height {
		newPos := (i + amount) % height
		g.Set(g.minRow+newPos, col, temp[i])
	}
}
