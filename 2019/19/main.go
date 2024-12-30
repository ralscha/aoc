package main

import (
	"aoc/2019/intcomputer"
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
)

func main() {
	input, err := download.ReadInput(2019, 19)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func checkPoint(program []int, x, y int) bool {
	computer := intcomputer.NewIntcodeComputer(program)
	if err := computer.AddInput(x, y); err != nil {
		log.Fatalf("adding input failed: %v", err)
	}

	result, err := computer.Run()
	if err != nil {
		log.Fatalf("running program failed: %v", err)
	}

	if result.Signal != intcomputer.SignalOutput {
		log.Fatal("expected output signal")
	}

	return result.Value == 1
}

func part1(input string) {
	program := conv.ToIntSliceComma(input)
	count := 0

	for y := 0; y < 50; y++ {
		for x := 0; x < 50; x++ {
			if checkPoint(program, x, y) {
				count++
			}
		}
	}

	fmt.Println("Part 1", count)
}

func findSquare(program []int, size int) (int, int) {
	y := size
	leftX := 0

	for {
		for !checkPoint(program, leftX, y) {
			leftX++
		}

		rightX := leftX + size - 1
		if !checkPoint(program, rightX, y) {
			y++
			continue
		}

		topY := y - size + 1
		if !checkPoint(program, leftX, topY) || !checkPoint(program, rightX, topY) {
			y++
			continue
		}
		
		return leftX, topY
	}
}

func part2(input string) {
	program := conv.ToIntSliceComma(input)
	x, y := findSquare(program, 100)
	result := x*10000 + y
	fmt.Println("Part 2", result)
}
