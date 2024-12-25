package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/gridutil"
	"fmt"
	"log"

	"aoc/2019/intcomputer"
)

func main() {
	input, err := download.ReadInput(2019, 13)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	program := conv.ToIntSliceComma(input)
	computer := intcomputer.NewIntcodeComputer(program)
	tiles := make(map[gridutil.Coordinate]int)

	for !computer.Halted {
		col := computer.Run()
		row := computer.Run()
		tileID := computer.Run()
		tiles[gridutil.Coordinate{Row: row, Col: col}] = tileID
	}

	blockCount := 0
	for _, tile := range tiles {
		if tile == 2 {
			blockCount++
		}
	}
	fmt.Println("Part 1", blockCount)
}

func part2(input string) {
	program := conv.ToIntSliceComma(input)
	program[0] = 2
	computer := intcomputer.NewIntcodeComputer(program)
	tiles := make(map[gridutil.Coordinate]int)
	var score int
	var ballCol, paddleCol int

	for !computer.Halted {
		col := computer.Run()
		row := computer.Run()
		tileID := computer.Run()

		if col == -1 && row == 0 {
			score = tileID
		} else {
			tiles[gridutil.Coordinate{Row: row, Col: col}] = tileID
			if tileID == 3 {
				paddleCol = col
			} else if tileID == 4 {
				ballCol = col
			}
		}

		if ballCol < paddleCol {
			computer.Input = -1
		} else if ballCol > paddleCol {
			computer.Input = 1
		} else {
			computer.Input = 0
		}
	}

	fmt.Println("Part 2", score)
}
