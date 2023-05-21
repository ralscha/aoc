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

	instructionIndex := 0

	for instructionIndex < len(instructions) {
		instruction := instructions[instructionIndex]
		if instruction == 'R' || instruction == 'L' {
			facing = nextFacing(facing, rune(instruction))
			instructionIndex++
		} else {
			numberStr := ""
			for instructionIndex < len(instructions) {
				instruction = instructions[instructionIndex]
				if instruction == 'R' || instruction == 'L' {
					break
				}
				numberStr += string(instruction)
				instructionIndex++
			}
			number := conv.MustAtoi(numberStr)
			currentPosition = move(g, currentPosition, facing, number)
		}
	}

	fmt.Println((currentPosition.row+1)*1000 + (currentPosition.col+1)*4 + facing)
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

func part2(input string) {
	lines := conv.SplitNewline(input)
	g := createGrid(lines)
	instructions := lines[len(lines)-1]

	currentPosition := position{row: 0, col: findFirstEmpty(g[0])}
	// Facing is 0 for right (>), 1 for down (v), 2 for left (<), and 3 for up (^)
	facing := 0

	instructionIndex := 0
	for instructionIndex < len(instructions) {
		instruction := instructions[instructionIndex]
		if instruction == 'R' || instruction == 'L' {
			facing = nextFacing(facing, rune(instruction))
			instructionIndex++
		} else {
			numberStr := ""
			for instructionIndex < len(instructions) {
				instruction = instructions[instructionIndex]
				if instruction == 'R' || instruction == 'L' {
					break
				}
				numberStr += string(instruction)
				instructionIndex++
			}
			number := conv.MustAtoi(numberStr)
			currentPosition, facing = move3d(g, currentPosition, facing, number)
		}
	}
	fmt.Println((currentPosition.row+1)*1000 + (currentPosition.col+1)*4 + facing)
}

func move3d(g grid, currentPosition position, facing, number int) (position, int) {

	directions := make(map[int]position)
	directions[0] = position{row: 0, col: 1}
	directions[1] = position{row: 1, col: 0}
	directions[2] = position{row: 0, col: -1}
	directions[3] = position{row: -1, col: 0}

	for i := 0; i < number; i++ {
		nextPosition := position{
			row: currentPosition.row + directions[facing].row,
			col: currentPosition.col + directions[facing].col,
		}
		nextFacing := facing
		if nextPosition.row < 0 || nextPosition.col < 0 ||
			nextPosition.row == len(g) || nextPosition.col == len(g[nextPosition.row]) ||
			g[nextPosition.row][nextPosition.col] == ' ' ||
			g[nextPosition.row][nextPosition.col] == 0 {

			r := region(currentPosition.row, currentPosition.col)

			switch r {
			case 1:
				switch facing {
				case 2:
					nextPosition, nextFacing = position{col: 0, row: 149 - currentPosition.row}, 0
				case 3:
					nextPosition, nextFacing = position{col: 0, row: currentPosition.col + 100}, 0
				}
			case 2:
				switch facing {
				case 0:
					nextPosition, nextFacing = position{col: 99, row: 149 - currentPosition.row}, 2
				case 1:
					nextPosition, nextFacing = position{col: 99, row: currentPosition.col - 50}, 2
				case 3:
					nextPosition, nextFacing = position{col: currentPosition.col - 100, row: 199}, 3
				}
			case 3:
				switch facing {
				case 0:
					nextPosition, nextFacing = position{col: currentPosition.row + 50, row: 49}, 3
				case 2:
					nextPosition, nextFacing = position{col: currentPosition.row - 50, row: 100}, 1
				}
			case 4:
				switch facing {
				case 2:
					nextPosition, nextFacing = position{col: 50, row: 149 - currentPosition.row}, 0
				case 3:
					nextPosition, nextFacing = position{col: 50, row: currentPosition.col + 50}, 0
				}
			case 5:
				switch facing {
				case 0:
					nextPosition, nextFacing = position{col: 149, row: 149 - currentPosition.row}, 2
				case 1:
					nextPosition, nextFacing = position{col: 49, row: currentPosition.col + 100}, 2
				}
			case 6:
				switch facing {
				case 0:
					nextPosition, nextFacing = position{col: currentPosition.row - 100, row: 149}, 3
				case 1:
					nextPosition, nextFacing = position{col: currentPosition.col + 100, row: 0}, 1
				case 2:
					nextPosition, nextFacing = position{col: currentPosition.row - 100, row: 0}, 1
				}
			}
		}
		if g[nextPosition.row][nextPosition.col] == '#' {
			return currentPosition, facing
		}

		currentPosition = nextPosition
		facing = nextFacing
	}

	return currentPosition, facing
}

func region(row, col int) int {
	if col >= 50 && col <= 99 && row >= 0 && row <= 49 {
		return 1
	}
	if col >= 100 && col <= 149 && row >= 0 && row <= 49 {
		return 2
	}
	if col >= 50 && col <= 99 && row >= 50 && row <= 99 {
		return 3
	}
	if col >= 50 && col <= 99 && row >= 100 && row <= 149 {
		return 5
	}
	if col >= 0 && col <= 49 && row >= 100 && row <= 149 {
		return 4
	}
	if col >= 0 && col <= 49 && row >= 150 && row <= 199 {
		return 6
	}
	panic("invalid region")
	return 0
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
