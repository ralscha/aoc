package main

import (
	"aoc/internal/download"
	"fmt"
	"log"
)

func main() {
	inputFile := "./2022/17/input.txt"
	input, err := download.ReadInput(inputFile, 2022, 17)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
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

func part1(input string) {
	input = input[:len(input)-1]
	gas := &jet{
		jets: input,
		next: 0,
	}

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

	var chamber []row

	noOfRocks := 2022

	for i := 0; i < noOfRocks; i++ {
		rockX := 2
		rockY := 0
		nextRock := rocks[i%5]
		rockHeight := len(nextRock.grid)

		highestRock := 0
		for j := len(chamber) - 1; j >= 0; j-- {
			row := chamber[j]
			if !row[0] && !row[1] && !row[2] && !row[3] && !row[4] && !row[5] && !row[6] {
				highestRock = j + 1
				break
			}
		}

		emptyRows := len(chamber) - (len(chamber) - highestRock)
		missing := rockHeight + 3 - emptyRows
		if missing < 0 {
			chamber = chamber[missing*-1:]
		} else {
			for j := 0; j < missing; j++ {
				chamber = append([]row{{}}, chamber...)
			}
		}

		rockX += nextGas(gas)
		if rockX < 0 {
			rockX = 0
		} else if rockX+len(nextRock.grid[0])-1 > 6 {
			rockX -= 1
		}
		drawRock(chamber, nextRock, rockX, rockY)
		fall(chamber, nextRock, rockX, rockY, gas)
	}
	height := 0
	for i := len(chamber) - 1; i >= 0; i-- {
		row := chamber[i]
		if !row[0] && !row[1] && !row[2] && !row[3] && !row[4] && !row[5] && !row[6] {
			break
		}
		height++
	}
	fmt.Println(height)
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

func eraseRock(chamber []row, currentRock rock, rockX int, rockY int) {
	rockHeight := len(currentRock.grid)
	for y := 0; y < rockHeight; y++ {
		for x := 0; x < len(currentRock.grid[y]); x++ {
			if currentRock.grid[y][x] {
				chamber[rockY+y][rockX+x] = false
			}
		}
	}
}

func fall(chamber []row, rock rock, rockX int, rockY int, gas *jet) {
	currentRockX, currentRockY := rockX, rockY
	nextY := rockY + 1

	for nextY < len(chamber) {
		isCollisionY1 := collisionY(chamber, rock, currentRockX, nextY)
		if isCollisionY1 {
			return
		}
		eraseRock(chamber, rock, currentRockX, currentRockY)
		currentRockY = nextY
		// gas
		nextX := currentRockX + nextGas(gas)
		if !collisionX(chamber, rock, nextX, currentRockY) {
			currentRockX = nextX
		}
		drawRock(chamber, rock, currentRockX, currentRockY)
		nextY++
	}
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

func printChamber(chamber []row) {
	for y := 0; y < len(chamber); y++ {
		for x := 0; x < len(chamber[y]); x++ {
			if chamber[y][x] {
				fmt.Print("@")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
