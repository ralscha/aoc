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
	input, err := download.ReadInput(2019, 11)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	program := conv.ToIntSliceComma(input)
	computer := intcomputer.NewIntcodeComputer(program)
	painted := runRobot(computer, 0)
	fmt.Println("Part 1", len(painted))
}

func part2(input string) {
	program := conv.ToIntSliceComma(input)
	computer := intcomputer.NewIntcodeComputer(program)
	painted := runRobot(computer, 1)
	fmt.Println("Part 2")
	printRegistration(painted)
}

func runRobot(computer *intcomputer.IntcodeComputer, startColor int) map[gridutil.Coordinate]int {
	painted := make(map[gridutil.Coordinate]int)
	pos := gridutil.Coordinate{}
	dir := gridutil.DirectionN

	painted[pos] = startColor

	for {
		// Provide current color as input when requested
		result, err := computer.Run()
		if err != nil {
			log.Fatalf("running program failed: %v", err)
		}

		switch result.Signal {
		case intcomputer.SignalInput:
			color := painted[pos]
			if err := computer.AddInput(color); err != nil {
				log.Fatalf("adding input failed: %v", err)
			}
		case intcomputer.SignalOutput:
			// First output is the color to paint
			newColor := result.Value
			painted[pos] = newColor

			// Get second output for turn direction
			result, err = computer.Run()
			if err != nil {
				log.Fatalf("running program failed: %v", err)
			}
			if result.Signal != intcomputer.SignalOutput {
				log.Fatal("expected turn direction output")
			}
			turn := result.Value

			// Update direction and move
			if turn == 0 {
				dir = gridutil.TurnLeft(dir)
			} else {
				dir = gridutil.TurnRight(dir)
			}

			switch dir {
			case gridutil.DirectionN:
				pos.Row--
			case gridutil.DirectionE:
				pos.Col++
			case gridutil.DirectionS:
				pos.Row++
			case gridutil.DirectionW:
				pos.Col--
			}
		case intcomputer.SignalEnd:
			return painted
		}
	}
}

func printRegistration(painted map[gridutil.Coordinate]int) {
	minCol, maxCol, minRow, maxRow := 0, 0, 0, 0
	for p := range painted {
		if p.Col < minCol {
			minCol = p.Col
		}
		if p.Col > maxCol {
			maxCol = p.Col
		}
		if p.Row < minRow {
			minRow = p.Row
		}
		if p.Row > maxRow {
			maxRow = p.Row
		}
	}

	for r := minRow; r <= maxRow; r++ {
		for c := minCol; c <= maxCol; c++ {
			if painted[gridutil.Coordinate{Col: c, Row: r}] == 1 {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
