package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
)

func main() {
	inputFile := "./2022/22/input.txt"
	input, err := download.ReadInput(inputFile, 2022, 22)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)

}

type position struct {
	row, col int
}

type grid []row

type row []rune

func part1(input string) {
	lines := conv.SplitNewline(input)
	g := createGrid(lines)
	instructions := lines[len(lines)-1]

	currentPosition := position{row: 0, col: findFirstEmpty(g[0])}
	// Facing is 0 for right (>), 1 for down (v), 2 for left (<), and 3 for up (^)
	facing := 0

	instrutionIndex := 0

	for instrutionIndex < len(instructions) {
		instruction := instructions[instrutionIndex]
		if instruction == 'R' || instruction == 'L' {
			facing = nextFacing(facing, rune(instruction))
			instrutionIndex++
		} else {
			numberStr := ""
			for instrutionIndex < len(instructions) {
				instruction = instructions[instrutionIndex]
				if instruction == 'R' || instruction == 'L' {
					break
				}
				numberStr += string(instruction)
				instrutionIndex++
			}
			number := conv.MustAtoi(numberStr)
			currentPosition = move(g, currentPosition, facing, number)
		}
	}

	fmt.Println((currentPosition.row+1)*1000 + (currentPosition.col+1)*4 + facing)
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	g := createGrid(lines)
	instructions := lines[len(lines)-1]

	currentPosition := position{row: 0, col: findFirstEmpty(g[0])}
	// Facing is 0 for right (>), 1 for down (v), 2 for left (<), and 3 for up (^)
	facing := 0

	instrutionIndex := 0
	for instrutionIndex < len(instructions) {
		instruction := instructions[instrutionIndex]
		if instruction == 'R' || instruction == 'L' {
			facing = nextFacing(facing, rune(instruction))
			instrutionIndex++
		} else {
			numberStr := ""
			for instrutionIndex < len(instructions) {
				instruction = instructions[instrutionIndex]
				if instruction == 'R' || instruction == 'L' {
					break
				}
				numberStr += string(instruction)
				instrutionIndex++
			}
			number := conv.MustAtoi(numberStr)
			currentPosition = move3d(g, currentPosition, facing, number)
		}
	}

	fmt.Println((currentPosition.row+1)*1000 + (currentPosition.col+1)*4 + facing)
}

func move3d(g grid, currentPosition position, facing, number int) position {
	for i := 0; i < number; i++ {
		switch facing {
		case 0: // right
			nextCol := (currentPosition.col + 1) % len(g[currentPosition.row])
			if g[currentPosition.row][nextCol] == '.' {
				currentPosition.col = nextCol
			} else if g[currentPosition.row][nextCol] == '#' {
				return currentPosition
			}

		case 1: // down
			nextRow := (currentPosition.row + 1) % len(g)
			if g[nextRow][currentPosition.col] == '.' {
				currentPosition.row = nextRow
			} else if g[nextRow][currentPosition.col] == '#' {
				return currentPosition
			}
		case 2: // left
			nextCol := currentPosition.col - 1
			for g[currentPosition.row][nextCol] == ' ' || g[currentPosition.row][nextCol] == 0 {
				nextCol -= 1
				if nextCol < 0 {
					nextCol = len(g[currentPosition.row]) - 1
				}
			}
			if g[currentPosition.row][nextCol] == '.' {
				currentPosition.col = nextCol
			} else if g[currentPosition.row][nextCol] == '#' {
				return currentPosition
			}

		case 3: // up
			nextRow := currentPosition.row - 1
			for g[nextRow][currentPosition.col] == ' ' || g[nextRow][currentPosition.col] == 0 {
				nextRow -= 1
				if nextRow < 0 {
					nextRow = len(g) - 1
				}
			}
			if g[nextRow][currentPosition.col] == '.' {
				currentPosition.row = nextRow
			} else if g[nextRow][currentPosition.col] == '#' {
				return currentPosition
			}
		}
	}
	return currentPosition
}

func createGrid(lines []string) grid {
	longestLine := 0
	for _, line := range lines {
		if line == "" {
			break
		}
		if len(line) > longestLine {
			longestLine = len(line)
		}
	}
	var g grid
	for _, line := range lines {
		if line == "" {
			break
		}
		row := make(row, longestLine)
		for i, r := range line {
			row[i] = r
		}
		g = append(g, row)
	}
	return g
}

func move(g grid, currentPosition position, facing, number int) position {
	for i := 0; i < number; i++ {
		switch facing {
		case 0: // right
			nextCol := (currentPosition.col + 1) % len(g[currentPosition.row])
			for g[currentPosition.row][nextCol] == ' ' || g[currentPosition.row][nextCol] == 0 {
				nextCol = (nextCol + 1) % len(g[currentPosition.row])
			}
			if g[currentPosition.row][nextCol] == '.' {
				currentPosition.col = nextCol
			} else if g[currentPosition.row][nextCol] == '#' {
				return currentPosition
			}

		case 1: // down
			nextRow := (currentPosition.row + 1) % len(g)
			for g[nextRow][currentPosition.col] == ' ' || g[nextRow][currentPosition.col] == 0 {
				nextRow = (nextRow + 1) % len(g)
			}
			if g[nextRow][currentPosition.col] == '.' {
				currentPosition.row = nextRow
			} else if g[nextRow][currentPosition.col] == '#' {
				return currentPosition
			}
		case 2: // left
			nextCol := currentPosition.col - 1
			if nextCol < 0 {
				nextCol = len(g[currentPosition.row]) - 1
			}
			for g[currentPosition.row][nextCol] == ' ' || g[currentPosition.row][nextCol] == 0 {
				nextCol -= 1
				if nextCol < 0 {
					nextCol = len(g[currentPosition.row]) - 1
				}
			}
			if g[currentPosition.row][nextCol] == '.' {
				currentPosition.col = nextCol
			} else if g[currentPosition.row][nextCol] == '#' {
				return currentPosition
			}

		case 3: // up
			nextRow := currentPosition.row - 1
			if nextRow < 0 {
				nextRow = len(g) - 1
			}
			for g[nextRow][currentPosition.col] == ' ' || g[nextRow][currentPosition.col] == 0 {
				nextRow -= 1
				if nextRow < 0 {
					nextRow = len(g) - 1
				}
			}
			if g[nextRow][currentPosition.col] == '.' {
				currentPosition.row = nextRow
			} else if g[nextRow][currentPosition.col] == '#' {
				return currentPosition
			}
		}
	}
	return currentPosition
}

func nextFacing(facing int, next rune) int {
	if next == 'R' {
		facing++
	} else {
		facing--
	}
	if facing < 0 {
		facing = 3
	}

	if facing > 3 {
		facing = 0
	}

	return facing
}

func findFirstEmpty(row []rune) int {
	for i, r := range row {
		if r == '.' {
			return i
		}
	}
	return -1
}
