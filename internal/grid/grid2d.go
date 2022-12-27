package grid

type Coordinates struct {
	row, col int
}

type Grid2D[T any] struct {
	grid          map[Coordinates]T
	width, height int
	wrap          bool
}

func NewGrid2D[T any](width, height int, wrap bool) Grid2D[T] {
	return Grid2D[T]{
		grid:   make(map[Coordinates]T),
		width:  width,
		height: height,
		wrap:   wrap,
	}
}

func (g Grid2D[T]) Width() int {
	return g.width
}

func (g Grid2D[T]) Height() int {
	return g.height
}

func (g Grid2D[T]) Get(coord Coordinates) T {
	return g.grid[coord]
}

func (g Grid2D[T]) Set(coord Coordinates, value T) {
	g.grid[coord] = value
}

func (g Grid2D[T]) Move(coord, direction Coordinates) {
	newCoord := Coordinates{
		row: coord.row + direction.row,
		col: coord.col + direction.col,
	}
	if g.wrap {
		newCoord.row = (newCoord.row + g.height) % g.height
		newCoord.col = (newCoord.col + g.width) % g.width
	}

	if newCoord.row >= 0 && newCoord.row < g.height && newCoord.col >= 0 && newCoord.col < g.width {
		g.grid[newCoord] = g.grid[coord]
		delete(g.grid, coord)
	}
}

func (g Grid2D[T]) GetNeighbours8(coord Coordinates) []T {
	var directions = []Coordinates{
		{row: -1, col: -1},
		{row: -1, col: 0},
		{row: -1, col: 1},
		{row: 0, col: -1},
		{row: 0, col: 1},
		{row: 1, col: -1},
		{row: 1, col: 0},
		{row: 1, col: 1},
	}
	return g.GetNeighbours(coord, directions)
}

func (g Grid2D[T]) GetNeighbours4(coord Coordinates) []T {
	var directions = []Coordinates{
		{row: -1, col: 0},
		{row: 0, col: -1},
		{row: 0, col: 1},
		{row: 1, col: 0},
	}
	return g.GetNeighbours(coord, directions)
}

func (g Grid2D[T]) GetNeighbours(coord Coordinates, directions []Coordinates) []T {
	var neighbours []T
	for _, direction := range directions {
		neighbourCoord := Coordinates{
			row: coord.row + direction.row,
			col: coord.col + direction.col,
		}
		if g.wrap {
			neighbourCoord.row = (neighbourCoord.row + g.height) % g.height
			neighbourCoord.col = (neighbourCoord.col + g.width) % g.width
		}
		if neighbourCoord.row >= 0 && neighbourCoord.row < g.height && neighbourCoord.col >= 0 && neighbourCoord.col < g.width {
			if neighbour, ok := g.grid[neighbourCoord]; ok {
				neighbours = append(neighbours, neighbour)
			}
		}
	}
	return neighbours
}
