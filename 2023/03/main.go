package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/gridutil"
	"fmt"
	"log"
	"unicode"
)

func main() {
	input, err := download.ReadInput(2023, 3)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1and2(input)
}

type position struct {
	row int
	col int
}

func part1and2(input string) {
	lines := conv.SplitNewline(input)
	grid := gridutil.NewCharGrid2D(lines)

	sum := 0
	starNumbers := make(map[position][]int)

	minRow, maxRow := grid.GetMinMaxRow()
	minCol, maxCol := grid.GetMinMaxCol()

	for row := minRow; row <= maxRow; row++ {
		for col := minCol; col <= maxCol; col++ {
			curr, exists := grid.Get(row, col)
			if !exists || !unicode.IsDigit(curr) {
				continue
			}

			// Extract full number
			number := string(curr)
			startCol := col
			for col+1 <= maxCol {
				next, exists := grid.Get(row, col+1)
				if !exists || !unicode.IsDigit(next) {
					break
				}
				number += string(next)
				col++
			}
			endCol := col
			numberInt := conv.MustAtoi(number)

			// Check for symbol neighbors
			hasSymbolNeighbor := false
			for i := startCol; i <= endCol; i++ {
				coord := gridutil.Coordinate{Row: row, Col: i}
				neighbors := grid.GetNeighbours8C(coord)

				for _, neighbor := range neighbors {
					if !unicode.IsDigit(neighbor) && neighbor != '.' {
						hasSymbolNeighbor = true

						// Check for gear symbols
						for _, dir := range gridutil.Get8Directions() {
							testRow, testCol := row+dir.Row, i+dir.Col
							if testRow < minRow || testRow > maxRow || testCol < minCol || testCol > maxCol {
								continue
							}

							testRune, exists := grid.Get(testRow, testCol)
							if !exists {
								continue
							}

							if testRune == '*' {
								key := position{testRow, testCol}
								if _, ok := starNumbers[key]; !ok {
									starNumbers[key] = []int{}
								}
								starNumbers[key] = append(starNumbers[key], numberInt)
								break
							}
						}
						break
					}
				}
				if hasSymbolNeighbor {
					break
				}
			}

			if hasSymbolNeighbor {
				sum += numberInt
			}
		}
	}

	sumGears := 0
	for _, numbers := range starNumbers {
		if len(numbers) == 2 {
			sumGears += numbers[0] * numbers[1]
		}
	}

	fmt.Println(sum)
	fmt.Println(sumGears)
}
