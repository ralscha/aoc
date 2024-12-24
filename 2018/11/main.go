package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
)

func main() {
	input, err := download.ReadInput(2018, 11)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func getPowerLevel(x, y, serialNumber int) int {
	rackID := x + 10
	powerLevel := rackID * y
	powerLevel += serialNumber
	powerLevel *= rackID
	hundredsDigit := 0
	if powerLevel >= 100 {
		hundredsDigit = (powerLevel / 100) % 10
	} else if powerLevel <= -100 {
		hundredsDigit = ((-powerLevel) / 100) % 10
	}
	return hundredsDigit - 5
}

func part1(input string) {
	serialNumber := conv.MustAtoi(input)

	maxPower := -1000000
	bestX, bestY := -1, -1

	for y := 1; y <= 298; y++ {
		for x := 1; x <= 298; x++ {
			totalPower := 0
			for dy := 0; dy < 3; dy++ {
				for dx := 0; dx < 3; dx++ {
					totalPower += getPowerLevel(x+dx, y+dy, serialNumber)
				}
			}

			if totalPower > maxPower {
				maxPower = totalPower
				bestX, bestY = x, y
			}
		}
	}
	output := fmt.Sprintf("%d,%d", bestX, bestY)
	fmt.Println("Part 1", output)
}

func part2(input string) {
	serialNumber := conv.MustAtoi(input)

	maxPower := -1000000
	bestX, bestY, bestSize := -1, -1, -1

	powerGrid := make([][]int, 300)
	for i := range powerGrid {
		powerGrid[i] = make([]int, 300)
		for j := range powerGrid[i] {
			powerGrid[i][j] = getPowerLevel(j+1, i+1, serialNumber)
		}
	}

	summedAreaTable := make([][]int, 301)
	for i := range summedAreaTable {
		summedAreaTable[i] = make([]int, 301)
	}

	for y := 1; y <= 300; y++ {
		for x := 1; x <= 300; x++ {
			summedAreaTable[y][x] = powerGrid[y-1][x-1] +
				summedAreaTable[y-1][x] +
				summedAreaTable[y][x-1] -
				summedAreaTable[y-1][x-1]
		}
	}

	for size := 1; size <= 300; size++ {
		for y := size; y <= 300; y++ {
			for x := size; x <= 300; x++ {
				totalPower := summedAreaTable[y][x] -
					summedAreaTable[y-size][x] -
					summedAreaTable[y][x-size] +
					summedAreaTable[y-size][x-size]

				if totalPower > maxPower {
					maxPower = totalPower
					bestX, bestY = x-size+1, y-size+1
					bestSize = size
				}
			}
		}
	}

	output := fmt.Sprintf("%d,%d,%d", bestX, bestY, bestSize)
	fmt.Println("Part 2", output)
}
