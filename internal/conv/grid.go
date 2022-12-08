package conv

func CreateNumberGrid(lines []string) [][]int {
	grid := make([][]int, len(lines))
	for i := range grid {
		grid[i] = make([]int, len(lines[i]))
	}

	for i := range lines {
		for j := range lines[i] {
			grid[i][j] = int(lines[i][j] - '0')
		}
	}

	return grid
}
