package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/mathx"
	"fmt"
	"log"
)

func main() {
	inputFile := "./2022/08/input.txt"
	input, err := download.ReadInput(inputFile, 2022, 8)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	inputLines := conv.SplitNewline(input)
	grid := conv.CreateNumberGrid(inputLines)

	count := countVisibleTrees(grid)
	fmt.Println(count)
}

func part2(input string) {
	inputLines := conv.SplitNewline(input)
	grid := conv.CreateNumberGrid(inputLines)

	nRows := len(grid)
	nColumns := len(grid[0])
	scenicScore := 0
	for r := 0; r < nRows; r++ {
		for c := 0; c < nColumns; c++ {
			ss := calculateScenicScore(grid, r, c)
			if ss > scenicScore {
				scenicScore = ss
			}
		}
	}

	fmt.Println(scenicScore)
}

func countVisibleTrees(grid [][]int) int {
	nRows := len(grid)
	nColumns := len(grid[0])
	count := 0

	for r := 0; r < nRows; r++ {
		for c := 0; c < nColumns; c++ {
			if r == 0 || r == nRows-1 || c == 0 || c == nColumns-1 {
				count++
				continue
			}

			treeHeight := grid[r][c]
			if isVisible(grid, r, c, treeHeight) {
				count++
			}
		}
	}

	return count
}

func isVisible(grid [][]int, r, c, treeHeight int) bool {
	nRows := len(grid)
	nColumns := len(grid[0])

	blockingTrees := 0
	for i := r - 1; i >= 0; i-- {
		if grid[i][c] >= treeHeight {
			blockingTrees++
			break
		}
	}
	for i := r + 1; i < nRows; i++ {
		if grid[i][c] >= treeHeight {
			blockingTrees++
			break
		}
	}
	for i := c - 1; i >= 0; i-- {
		if grid[r][i] >= treeHeight {
			blockingTrees++
			break
		}
	}
	for i := c + 1; i < nColumns; i++ {
		if grid[r][i] >= treeHeight {
			blockingTrees++
			break
		}
	}

	return blockingTrees < 4
}

func calculateScenicScore(grid [][]int, r, c int) int {
	nRows := len(grid)
	nColumns := len(grid[0])
	treeHeight := grid[r][c]

	var scenicScore [4]int
	for i := r - 1; i >= 0; i-- {
		if grid[i][c] <= treeHeight {
			scenicScore[0]++
			if grid[i][c] == treeHeight {
				break
			}
		} else {
			break
		}
	}
	for i := r + 1; i < nRows; i++ {
		if grid[i][c] <= treeHeight {
			scenicScore[1]++
			if grid[i][c] == treeHeight {
				break
			}
		} else {
			break
		}
	}
	for i := c - 1; i >= 0; i-- {
		if grid[r][i] <= treeHeight {
			scenicScore[2]++
			if grid[r][i] == treeHeight {
				break
			}
		} else {
			break
		}
	}
	for i := c + 1; i < nColumns; i++ {
		if grid[r][i] <= treeHeight {
			scenicScore[3]++
			if grid[r][i] == treeHeight {
				break
			}
		} else {
			break
		}
	}

	return mathx.Max(1, scenicScore[0]) * mathx.Max(1, scenicScore[1]) * mathx.Max(1, scenicScore[2]) * mathx.Max(1, scenicScore[3])
}
