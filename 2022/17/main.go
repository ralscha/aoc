package main

import (
	"aoc/internal/download"
	"fmt"
	"log"
	"time"
)

func main() {
	input, err := download.ReadInput(2022, 17)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1and2(input, 2022)
	part1and2(input, 1000000000000)
}

type rock struct {
	grid   [][]bool
	left   []int
	right  []int
	bottom []int
}

type row [7]bool

type jet struct {
	jets string
	next int
}

func part1and2(input string, noOfRocks int) {
	input = input[:len(input)-1]
	gas := &jet{
		jets: input,
		next: 0,
	}

	rocks := createRocks()

	var chamber []row

	highestRock := 0

	cycleDetection := make(map[string]int)
	var heights []int
	cycle := 0

	start := time.Now()
	for i := 0; i < noOfRocks; i++ {
		rockX := 2
		rockY := 0
		nextRock := rocks[i%5]
		rockHeight := len(nextRock.grid)

		highestRock = getHighestRock(chamber)

		emptyRows := len(chamber) - (len(chamber) - highestRock)
		missing := rockHeight + 3 - emptyRows
		if missing < 0 {
			chamber = chamber[missing*-1:]
		} else {
			newRows := make([]row, missing)
			chamber = append(newRows, chamber...)
		}

		heights = append(heights, len(chamber)-highestRock)
		if i > 110 && cycle == 0 {
			nextRockIx := i % 5
			nextWindIx := gas.next
			var rowNumbers [100]int
			for i := 0; i < 100; i++ {
				rowNumbers[i] = convertRowToInt(chamber[highestRock+i])
			}
			key := fmt.Sprintf("%v-%v-%v", rowNumbers, nextRockIx, nextWindIx)
			if startCycle, ok := cycleDetection[key]; ok {
				cycle = i - startCycle
				increase := heights[i] - heights[startCycle]
				noOfRocksLeft := noOfRocks - startCycle
				totalIncrease := noOfRocksLeft / cycle * increase
				remainingHeight := heights[startCycle+(noOfRocksLeft%cycle)-1]
				fmt.Println("Answer", totalIncrease+remainingHeight+2)
				break
			} else {
				cycleDetection[key] = i
			}
		}

		rockX += nextGas(gas)
		if rockX < 0 {
			rockX = 0
		} else if rockX+len(nextRock.grid[0])-1 > 6 {
			rockX -= 1
		}
		fall(chamber, nextRock, rockX, rockY, gas)
	}
	end := time.Now()
	fmt.Println("Time taken:", end.Sub(start))

	fmt.Println(len(chamber) - highestRock)
}

func convertRowToInt(row row) int {
	var result int
	for i, v := range row {
		if v {
			result += 1 << i
		}
	}
	return result
}

func getHighestRock(chamber []row) int {
	for i := 0; i < len(chamber); i++ {
		row := chamber[i]
		for x := 0; x < len(row); x++ {
			if row[x] {
				return i
			}
		}
	}
	return 0
}

func createRocks() [5]rock {
	var rocks [5]rock

	rocks[0] = rock{
		grid: [][]bool{
			{true, true, true, true},
		},
		left:   []int{0},
		right:  []int{3},
		bottom: []int{0, 0, 0, 0},
	}

	rocks[1] = rock{
		grid: [][]bool{
			{false, true, false},
			{true, true, true},
			{false, true, false},
		},
		left:   []int{1, 0, 1},
		right:  []int{1, 2, 1},
		bottom: []int{1, 2, 1},
	}

	rocks[2] = rock{
		grid: [][]bool{
			{false, false, true},
			{false, false, true},
			{true, true, true},
		},
		left:   []int{2, 2, 0},
		right:  []int{2, 2, 2},
		bottom: []int{2, 2, 2},
	}

	rocks[3] = rock{
		grid: [][]bool{
			{true},
			{true},
			{true},
			{true},
		},
		left:   []int{0, 0, 0, 0},
		right:  []int{0, 0, 0, 0},
		bottom: []int{3},
	}

	rocks[4] = rock{
		grid: [][]bool{
			{true, true},
			{true, true},
		},
		left:   []int{0, 0},
		right:  []int{1, 1},
		bottom: []int{1, 1},
	}
	return rocks
}

func drawRock(chamber []row, currentRock rock, rockX int, rockY int) {
	rockHeight := len(currentRock.grid)
	for y := 0; y < rockHeight; y++ {
		for x := 0; x < len(currentRock.grid[y]); x++ {
			if currentRock.grid[y][x] {
				chamber[rockY+y][rockX+x] = currentRock.grid[y][x]
			}
		}
	}
}

func fall(chamber []row, rock rock, rockX int, rockY int, gas *jet) {
	currentRockX, currentRockY := rockX, rockY
	nextY := rockY + 1

	for nextY < len(chamber) {
		if collisionY(chamber, rock, currentRockX, nextY) {
			break
		}
		currentRockY = nextY
		// gas
		nextX := currentRockX + nextGas(gas)
		if !collisionX(chamber, rock, nextX, currentRockY) {
			currentRockX = nextX
		}
		nextY++
	}
	drawRock(chamber, rock, currentRockX, currentRockY)
}

func nextGas(jetpattern *jet) int {
	direction := jetpattern.jets[jetpattern.next]
	jetpattern.next++
	if jetpattern.next == len(jetpattern.jets) {
		jetpattern.next = 0
	}
	if direction == '<' {
		return -1
	} else if direction == '>' {
		return 1
	} else {
		return 0
	}
}

func collisionX(chamber []row, rock rock, rockX int, rockY int) bool {
	if rockX < 0 || rockX+len(rock.grid[0])-1 > 6 {
		return true
	}
	rockHeight := len(rock.grid)
	for y := 0; y < rockHeight; y++ {
		if rockX+rock.left[y] < 0 || rockX+rock.right[y] > 6 {
			return true
		}
	}

	for y := 0; y < rockHeight; y++ {
		if chamber[rockY+y][rockX+rock.left[y]] || chamber[rockY+y][rockX+rock.right[y]] {
			return true
		}
	}

	return false
}

func collisionY(chamber []row, rock rock, rockX int, rockY int) bool {
	rockHeight := len(rock.grid)
	if rockY+rockHeight-1 >= len(chamber) {
		return true
	}

	for x := 0; x < len(rock.grid[0]); x++ {
		if chamber[rockY+rock.bottom[x]][rockX+x] {
			return true
		}
	}

	return false
}
