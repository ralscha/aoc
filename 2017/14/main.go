package main

import (
	"fmt"
	"log"
	"strings"

	"aoc/internal/container"
	"aoc/internal/download"
)

const gridSize = 128

func main() {
	input, err := download.ReadInput(2017, 14)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	usedSquares := 0
	for i := range gridSize {
		hashInput := fmt.Sprintf("%s-%d", input, i)
		binaryString := calculateKnotHashBinary(hashInput)
		usedSquares += strings.Count(binaryString, "1")
	}
	fmt.Println("Part 1", usedSquares)
}

func part2(input string) {
	grid := make([][]int, gridSize)
	for i := range gridSize {
		grid[i] = make([]int, gridSize)
		hashInput := fmt.Sprintf("%s-%d", input, i)
		binaryString := calculateKnotHashBinary(hashInput)
		for j, bit := range binaryString {
			if bit == '1' {
				grid[i][j] = 1
			}
		}
	}

	regions := 0
	visited := container.NewSet[[2]int]()

	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || r >= gridSize || c < 0 || c >= gridSize || visited.Contains([2]int{r, c}) || grid[r][c] == 0 {
			return
		}
		visited.Add([2]int{r, c})
		dfs(r+1, c)
		dfs(r-1, c)
		dfs(r, c+1)
		dfs(r, c-1)
	}

	for i := range gridSize {
		for j := range gridSize {
			if grid[i][j] == 1 && !visited.Contains([2]int{i, j}) {
				regions++
				dfs(i, j)
			}
		}
	}

	fmt.Println("Part 2", regions)
}

func calculateKnotHashBinary(input string) string {
	list := make([]int, 256)
	for i := range 256 {
		list[i] = i
	}

	lengths := make([]int, len(input))
	for i, char := range input {
		lengths[i] = int(char)
	}
	lengths = append(lengths, 17, 31, 73, 47, 23)

	currentPosition := 0
	skipSize := 0

	for range 64 {
		for _, length := range lengths {
			sublist := make([]int, length)
			for i := range length {
				sublist[i] = list[(currentPosition+i)%256]
			}
			for i := range length / 2 {
				a := (currentPosition + i) % 256
				b := (currentPosition + length - 1 - i) % 256
				list[a], list[b] = list[b], list[a]
			}
			currentPosition = (currentPosition + length + skipSize) % 256
			skipSize++
		}
	}

	denseHash := make([]int, 16)
	for i := range 16 {
		xor := 0
		for j := range 16 {
			xor ^= list[i*16+j]
		}
		denseHash[i] = xor
	}

	binaryString := ""
	for _, val := range denseHash {
		binaryString += fmt.Sprintf("%08b", val)
	}

	return binaryString
}
