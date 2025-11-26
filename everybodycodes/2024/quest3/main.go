package main

import (
	"aoc/internal/conv"
	"aoc/internal/gridutil"
	"fmt"
	"log"
	"os"
)

func main() {
	partIandII("everybodycodes/quest3/partI.txt")
	partIandII("everybodycodes/quest3/partII.txt")
	partIII()
}

func partIandII(file string) {
	input, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	lines := conv.SplitNewline(string(input))

	grid := make([][]rune, len(lines))
	for i, l := range lines {
		grid[i] = []rune(l)
	}
	for i, l := range grid {
		for j, c := range l {
			if c == '#' {
				grid[i][j] = '1'
			}
		}
	}

	level := 1
	for {
		changed := false
		for i := 1; i < len(grid)-1; i++ {
			for j := 1; j < len(grid[i])-1; j++ {
				levelString := rune(level + '0')
				if grid[i][j] == levelString {
					if isSurrounded4(grid, i, j, levelString, rune(level+1+'0')) {
						grid[i][j] = rune(level + 1 + '0')
						changed = true
					}
				}
			}
		}
		if !changed {
			break
		}
		level++
	}

	total := 0
	for _, l := range grid {
		for _, c := range l {
			if c != '.' {
				total += int(c - '0')
			}
		}
	}

	fmt.Println(total)
}

func partIII() {
	input, err := os.ReadFile("everybodycodes/quest3/partIII.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := conv.SplitNewline(string(input))

	grid := make([][]rune, len(lines))
	for i, l := range lines {
		grid[i] = []rune(l)
	}
	for i, l := range grid {
		for j, c := range l {
			if c == '#' {
				grid[i][j] = '1'
			}
		}
	}

	level := 1
	for {
		changed := false
		for i := 1; i < len(grid)-1; i++ {
			for j := 1; j < len(grid[i])-1; j++ {
				levelString := rune(level + '0')
				if grid[i][j] == levelString {
					if isSurrounded8(grid, i, j, levelString, rune(level+1+'0')) {
						grid[i][j] = rune(level + 1 + '0')
						changed = true
					}
				}
			}
		}
		if !changed {
			break
		}
		level++
	}

	total := 0
	for _, l := range grid {
		for _, c := range l {
			if c != '.' {
				total += int(c - '0')
			}
		}
	}

	fmt.Println(total)
}

func isSurrounded4(grid [][]rune, i, j int, surrounded ...rune) bool {
	return (grid[i-1][j] == surrounded[0] || grid[i-1][j] == surrounded[1]) &&
		(grid[i+1][j] == surrounded[0] || grid[i+1][j] == surrounded[1]) &&
		(grid[i][j-1] == surrounded[0] || grid[i][j-1] == surrounded[1]) &&
		(grid[i][j+1] == surrounded[0] || grid[i][j+1] == surrounded[1])
}

func isSurrounded8(grid [][]rune, i, j int, surrounded ...rune) bool {
	rows := len(grid)
	cols := len(grid[0])

	directions := gridutil.Get8Directions()

	for _, d := range directions {
		ni, nj := i+d.Row, j+d.Col
		if ni < 0 || ni >= rows || nj < 0 || nj >= cols {
			return false
		}
		if grid[ni][nj] != surrounded[0] && grid[ni][nj] != surrounded[1] {
			return false
		}
	}

	return true
}
