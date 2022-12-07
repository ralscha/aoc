package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"bufio"
	"fmt"
	"log"
	"strings"
)

func main() {
	inputFile := "./2021/05/input.txt"
	input, err := download.ReadInput(inputFile, 2021, 5)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	var grid [1000][1000]int
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()

		splitted := strings.Split(line, " -> ")

		startSplitted := strings.Split(splitted[0], ",")
		endSplitted := strings.Split(splitted[1], ",")
		startX := conv.MustAtoi(startSplitted[0])
		startY := conv.MustAtoi(startSplitted[1])
		endX := conv.MustAtoi(endSplitted[0])
		endY := conv.MustAtoi(endSplitted[1])

		if startX == endX {
			start, end := startY, endY
			if start > end {
				start, end = end, start
			}
			for y := start; y <= end; y++ {
				grid[y][startX] += 1
			}
		} else if startY == endY {
			start, end := startX, endX
			if start > end {
				start, end = end, start
			}
			for x := start; x <= end; x++ {
				grid[startY][x] += 1
			}
		}
	}

	count := 0
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			if grid[x][y] >= 2 {
				count++
			}
		}
	}

	fmt.Println("Result: ", count)
}

func part2(input string) {
	var grid [1000][1000]int
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()

		splitted := strings.Split(line, " -> ")

		startSplitted := strings.Split(splitted[0], ",")
		endSplitted := strings.Split(splitted[1], ",")
		startX := conv.MustAtoi(startSplitted[0])
		startY := conv.MustAtoi(startSplitted[1])
		endX := conv.MustAtoi(endSplitted[0])
		endY := conv.MustAtoi(endSplitted[1])

		if startX == endX {
			start, end := startY, endY
			if start > end {
				start, end = end, start
			}
			for y := start; y <= end; y++ {
				grid[y][startX] += 1
			}
		} else if startY == endY {
			start, end := startX, endX
			if start > end {
				start, end = end, start
			}
			for x := start; x <= end; x++ {
				grid[startY][x] += 1
			}
		} else {
			incrX := 1
			if startX > endX {
				incrX = -1
			}
			incrY := 1
			if startY > endY {
				incrY = -1
			}
			for x, y := startX, startY; ; x, y = x+incrX, y+incrY {
				grid[y][x] += 1
				if y == endY {
					break
				}
			}
		}
	}

	count := 0
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			if grid[x][y] >= 2 {
				count++
			}
		}
	}

	fmt.Println("Result: ", count)
}
