package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/gridutil"
	"fmt"
	"log"
)

func main() {
	input, err := download.ReadInput(2022, 8)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	inputLines := conv.SplitNewline(input)
	grid := gridutil.NewNumberGrid2D(inputLines)

	count := countVisibleTrees(grid)
	fmt.Println(count)
}

func part2(input string) {
	inputLines := conv.SplitNewline(input)
	grid := gridutil.NewNumberGrid2D(inputLines)

	minCol, maxCol := grid.GetMinMaxCol()
	minRow, maxRow := grid.GetMinMaxRow()

	scenicScore := 0
	for r := minRow; r <= maxRow; r++ {
		for c := minCol; c <= maxCol; c++ {
			ss := calculateScenicScore(grid, r, c, maxRow, maxCol)
			if ss > scenicScore {
				scenicScore = ss
			}
		}
	}

	fmt.Println(scenicScore)
}

func countVisibleTrees(grid gridutil.Grid2D[int]) int {
	minCol, maxCol := grid.GetMinMaxCol()
	minRow, maxRow := grid.GetMinMaxRow()
	count := 0

	for r := minRow; r <= maxRow; r++ {
		for c := minCol; c <= maxCol; c++ {
			if r == 0 || r == maxRow || c == 0 || c == maxCol {
				count++
				continue
			}

			treeHeight, _ := grid.Get(r, c)
			if isVisible(grid, r, c, treeHeight) {
				count++
			}
		}
	}

	return count
}

func isVisible(grid gridutil.Grid2D[int], r, c, treeHeight int) bool {
	_, maxCol := grid.GetMinMaxCol()
	_, maxRow := grid.GetMinMaxRow()

	blockingTrees := 0
	for i := r - 1; i >= 0; i-- {
		height, _ := grid.Get(i, c)
		if height >= treeHeight {
			blockingTrees++
			break
		}
	}
	for i := r + 1; i <= maxRow; i++ {
		height, _ := grid.Get(i, c)
		if height >= treeHeight {
			blockingTrees++
			break
		}
	}
	for i := c - 1; i >= 0; i-- {
		height, _ := grid.Get(r, i)
		if height >= treeHeight {
			blockingTrees++
			break
		}
	}
	for i := c + 1; i <= maxCol; i++ {
		height, _ := grid.Get(r, i)
		if height >= treeHeight {
			blockingTrees++
			break
		}
	}

	return blockingTrees < 4
}

func calculateScenicScore(grid gridutil.Grid2D[int], r, c, maxRow, maxCol int) int {
	treeHeight, _ := grid.Get(r, c)

	var scenicScore [4]int
	for i := r - 1; i >= 0; i-- {
		th, _ := grid.Get(i, c)
		if th <= treeHeight {
			scenicScore[0]++
			if th == treeHeight {
				break
			}
		} else {
			break
		}
	}
	for i := r + 1; i <= maxRow; i++ {
		th, _ := grid.Get(i, c)
		if th <= treeHeight {
			scenicScore[1]++
			if th == treeHeight {
				break
			}
		} else {
			break
		}
	}
	for i := c - 1; i >= 0; i-- {
		th, _ := grid.Get(r, i)
		if th <= treeHeight {
			scenicScore[2]++
			if th == treeHeight {
				break
			}
		} else {
			break
		}
	}
	for i := c + 1; i <= maxCol; i++ {
		th, _ := grid.Get(r, i)
		if th <= treeHeight {
			scenicScore[3]++
			if th == treeHeight {
				break
			}
		} else {
			break
		}
	}

	return max(1, scenicScore[0]) * max(1, scenicScore[1]) * max(1, scenicScore[2]) * max(1, scenicScore[3])
}
