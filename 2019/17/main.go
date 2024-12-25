package main

import (
	"aoc/2019/intcomputer"
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/gridutil"
	"fmt"
	"log"
)

func main() {
	input, err := download.ReadInput(2019, 17)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	program := conv.ToIntSliceComma(input)
	computer := intcomputer.NewIntcodeComputer(program)
	scaffold := make(map[gridutil.Coordinate]rune)
	var x, y int

	for !computer.Halted {
		output := computer.Run()
		if output == -1 {
			break
		}
		char := rune(output)
		if char == '\n' {
			y++
			x = 0
		} else {
			scaffold[gridutil.Coordinate{Row: y, Col: x}] = char
			x++
		}
	}

	alignmentSum := 0
	for point, char := range scaffold {
		if char == '#' && isIntersection(scaffold, point) {
			alignmentSum += point.Row * point.Col
		}
	}

	fmt.Println("Part 1", alignmentSum)
}

func isIntersection(scaffold map[gridutil.Coordinate]rune, point gridutil.Coordinate) bool {
	directions := gridutil.Get4Directions()
	for _, dir := range directions {
		neighbor := gridutil.Coordinate{Row: point.Row + dir.Row, Col: point.Col + dir.Col}
		if scaffold[neighbor] != '#' {
			return false
		}
	}
	return true
}

func part2(input string) {
	fmt.Println("Part 2", "TODO")
}
