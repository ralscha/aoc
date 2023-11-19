package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"strings"
)

func main() {
	inputFile := "./2016/08/input.txt"
	input, err := download.ReadInputAuto(inputFile)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}
	lines := conv.SplitNewline(input)
	part1(lines)
}

func part1(lines []string) {
	var grid [50][6]bool

	for _, line := range lines {
		spaceIndex := strings.Index(line, " ")
		command := line[0:spaceIndex]
		switch command {
		case "rect":
			rect := line[spaceIndex+1:]
			xIndex := strings.Index(rect, "x")
			col := conv.MustAtoi(rect[0:xIndex])
			row := conv.MustAtoi(rect[xIndex+1:])
			for y := 0; y < row; y++ {
				for x := 0; x < col; x++ {
					grid[x][y] = true
				}
			}
		case "rotate":
			rotate := line[spaceIndex+1:]
			equalsIndex := strings.Index(rotate, "=")
			byIndex := strings.Index(rotate, " by ")
			axis := rotate[0 : equalsIndex-2]
			index := conv.MustAtoi(rotate[equalsIndex+1 : byIndex])
			amount := conv.MustAtoi(rotate[byIndex+4:])
			switch axis {
			case "row":
				gridCopy := grid
				for x := 0; x < 50; x++ {
					grid[(x+amount)%50][index] = gridCopy[x][index]
				}
			case "column":
				gridCopy := grid
				for y := 0; y < 6; y++ {
					grid[index][(y+amount)%6] = gridCopy[index][y]
				}
			}
		}
	}
	count := 0
	for x := 0; x < 50; x++ {
		for y := 0; y < 6; y++ {
			if grid[x][y] {
				count++
			}
		}
	}

	for y := 0; y < 6; y++ {
		for x := 0; x < 50; x++ {
			if grid[x][y] {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

	fmt.Println(count)
}
