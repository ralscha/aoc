package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"aoc/internal/gridutil"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	partI()
	partII()
	partIII()
}

func partI() {
	lines := readInput("input1")
	grid := gridutil.NewCharGrid2D(lines)
	result := countTrampolinePairs(&grid)
	fmt.Println(result)
}

func partII() {
	lines := readInput("input2")
	grid := gridutil.NewCharGrid2D(lines)
	result := findShortestPath(&grid)
	fmt.Println(result)
}

func partIII() {
	lines := readInput("input3")
	grid := gridutil.NewCharGrid2D(lines)
	result := findShortestPathWithRotation(&grid)
	fmt.Println(result)
}

func findShortestPath(grid *gridutil.Grid2D[rune]) int {
	var start, end gridutil.Coordinate

	for r := range grid.Height() {
		for c := range grid.Width() {
			val, ok := grid.Get(r, c)
			if !ok {
				continue
			}
			switch val {
			case 'S':
				start = gridutil.Coordinate{Row: r, Col: c}
			case 'E':
				end = gridutil.Coordinate{Row: r, Col: c}
			}
		}
	}

	current := container.NewSet[gridutil.Coordinate]()
	current.Add(start)

	visited := container.NewSet[gridutil.Coordinate]()
	visited.Add(start)

	steps := 0

	for current.Len() > 0 {
		steps++
		next := container.NewSet[gridutil.Coordinate]()

		for _, pos := range current.Values() {
			deltas := []gridutil.Coordinate{
				{Row: 0, Col: -1},
				{Row: 0, Col: 1},
			}

			if (pos.Row+pos.Col)%2 == 0 {
				deltas = append(deltas, gridutil.Coordinate{Row: -1, Col: 0})
			} else {
				deltas = append(deltas, gridutil.Coordinate{Row: 1, Col: 0})
			}

			for _, delta := range deltas {
				np := gridutil.Coordinate{Row: pos.Row + delta.Row, Col: pos.Col + delta.Col}

				if np == end {
					return steps
				}

				val, ok := grid.Get(np.Row, np.Col)
				if !ok {
					continue
				}

				if val == 'T' && !visited.Contains(np) {
					visited.Add(np)
					next.Add(np)
				}
			}
		}

		current = next
	}

	return -1
}

type state struct {
	row      int
	col      int
	rotation int
}

func rotateGrid120(grid *gridutil.Grid2D[rune]) gridutil.Grid2D[rune] {
	height := grid.Height()
	width := grid.Width()

	newLines := make([]string, height)
	for i := range height {
		newLines[i] = strings.Repeat(".", width)
	}
	result := gridutil.NewCharGrid2D(newLines)

	for origY := range height {
		for origX := origY; origX < width-origY; origX++ {
			val, ok := grid.Get(origY, origX)
			if !ok {
				continue
			}

			resultX := (width - 1) - ((origX - origY + 1) / 2) - (origY * 2)
			resultY := (origX - origY) / 2

			if resultX >= 0 && resultX < width && resultY >= 0 && resultY < height {
				result.Set(resultY, resultX, val)
			}
		}
	}

	return result
}

func findShortestPathWithRotation(grid *gridutil.Grid2D[rune]) int {
	var start gridutil.Coordinate
	for r := range grid.Height() {
		for c := range grid.Width() {
			val, ok := grid.Get(r, c)
			if ok && val == 'S' {
				start = gridutil.Coordinate{Row: r, Col: c}
				break
			}
		}
	}

	grid2 := rotateGrid120(grid)
	grid3 := rotateGrid120(&grid2)
	grids := [3]*gridutil.Grid2D[rune]{
		grid,
		&grid2,
		&grid3,
	}

	visited := container.NewSet[state]()
	queue := []struct {
		state state
		jumps int
	}{
		{state{start.Row, start.Col, 0}, 0},
	}
	visited.Add(state{start.Row, start.Col, 0})

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		pos := current.state
		jumps := current.jumps

		nextRotation := (pos.rotation + 1) % 3

		neighbors := getNeighbors(pos.row, pos.col)

		for _, next := range neighbors {
			nextState := state{next.Row, next.Col, nextRotation}

			if !visited.Contains(nextState) {
				val, ok := grids[nextRotation].Get(next.Row, next.Col)
				if !ok {
					continue
				}

				if val == 'E' {
					return jumps + 1
				}

				if val == 'T' {
					visited.Add(nextState)
					queue = append(queue, struct {
						state state
						jumps int
					}{nextState, jumps + 1})
				}
			}
		}
	}

	return -1
}

func getNeighbors(row, col int) []gridutil.Coordinate {
	neighbors := []gridutil.Coordinate{
		{Row: row, Col: col - 1},
		{Row: row, Col: col + 1},
	}

	if (row+col)%2 == 0 {
		neighbors = append(neighbors, gridutil.Coordinate{Row: row - 1, Col: col})
	} else {
		neighbors = append(neighbors, gridutil.Coordinate{Row: row + 1, Col: col})
	}

	neighbors = append(neighbors, gridutil.Coordinate{Row: row, Col: col})

	return neighbors
}

func readInput(filename string) []string {
	input, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	inputStr := strings.ReplaceAll(string(input), "\r\n", "\n")
	return conv.SplitNewline(inputStr)
}

func countTrampolinePairs(grid *gridutil.Grid2D[rune]) int {
	count := 0

	for r := 0; r < grid.Height(); r++ {
		for c := 1; c < grid.Width(); c++ {
			curr, currOk := grid.Get(r, c)
			prev, prevOk := grid.Get(r, c-1)
			if currOk && prevOk && curr == 'T' && prev == 'T' {
				count++
			}
		}
	}

	for r := 1; r < grid.Height()-1; r += 2 {
		for c := 0; c < grid.Width(); c++ {
			val, ok := grid.Get(r, c)
			if !ok || val != 'T' {
				continue
			}

			var neighborRow int
			if (r+c)%2 == 1 {
				neighborRow = r + 1
			} else {
				neighborRow = r - 1
			}

			if neighborRow >= 0 && neighborRow < grid.Height() {
				neighbor, neighborOk := grid.Get(neighborRow, c)
				if neighborOk && neighbor == 'T' {
					count++
				}
			}
		}
	}

	return count
}
