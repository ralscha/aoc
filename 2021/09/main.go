package main

import (
	"aoc/internal/download"
	"bufio"
	"fmt"
	"log"
	"sort"
	"strings"
)

type point struct {
	row int
	col int
}

func main() {
	input, err := download.ReadInput(2021, 9)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	var grid [][]int32

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		var lg []int32
		for _, c := range line {
			lg = append(lg, c-'0')
		}
		grid = append(grid, lg)
	}

	var lowPoints []int32
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			currentValue := grid[row][col]
			isLower := true
			//up
			if row > 0 {
				isLower = isLower && currentValue < grid[row-1][col]
			}
			//down
			if row < len(grid)-1 {
				isLower = isLower && currentValue < grid[row+1][col]
			}
			//left
			if col > 0 {
				isLower = isLower && currentValue < grid[row][col-1]
			}
			//right
			if col < len(grid[row])-1 {
				isLower = isLower && currentValue < grid[row][col+1]
			}

			if isLower {
				lowPoints = append(lowPoints, currentValue)
			}
		}
	}

	var total int32
	for _, lp := range lowPoints {
		total += lp + 1
	}

	fmt.Println(total)
}

func part2(input string) {
	var grid [][]int32

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		var lg []int32
		for _, c := range line {
			lg = append(lg, c-'0')
		}
		grid = append(grid, lg)
	}

	var basinSize []int
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			islp := isLowPoint(row, col, grid)
			if islp {
				size := crawl(row, col, grid, map[point]struct{}{}, 1)
				if size > 0 {
					basinSize = append(basinSize, size)
				}
			}
		}
	}
	sort.Slice(basinSize, func(i, j int) bool {
		return basinSize[i] > basinSize[j]
	})

	result := basinSize[0] * basinSize[1] * basinSize[2]
	fmt.Println(result)
}

func crawl(row, col int, grid [][]int32, visited map[point]struct{}, size int) int {
	key := point{row: row, col: col}
	if _, ok := visited[key]; ok {
		return size - 1
	} else {
		visited[key] = struct{}{}
	}

	if grid[row][col] == 9 {
		return size - 1
	}

	if row > 0 {
		up := row - 1
		size = crawl(up, col, grid, visited, size+1)
	}

	if row < len(grid)-1 {
		down := row + 1
		size = crawl(down, col, grid, visited, size+1)
	}

	if col > 0 {
		left := col - 1
		size = crawl(row, left, grid, visited, size+1)
	}

	if col < len(grid[row])-1 {
		right := col + 1
		size = crawl(row, right, grid, visited, size+1)
	}

	return size
}

func isLowPoint(row, col int, grid [][]int32) bool {
	currentValue := grid[row][col]
	//up
	if row > 0 && currentValue >= grid[row-1][col] {
		return false
	}
	//down
	if row < len(grid)-1 && currentValue >= grid[row+1][col] {
		return false
	}
	//left
	if col > 0 && currentValue >= grid[row][col-1] {
		return false
	}
	//right
	if col < len(grid[row])-1 && currentValue >= grid[row][col+1] {
		return false
	}
	return true
}
