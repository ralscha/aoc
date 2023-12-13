package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
)

func main() {
	input, err := download.ReadInput(2023, 13)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	mirrors := part1(input)
	part2(input, mirrors)
}

func part1(input string) []int {
	var mirrors []int
	lines := conv.SplitNewline(input)
	summary := 0
	var pattern []string
	for _, line := range lines {
		if line == "" {
			s := findMirror(pattern)
			summary += s
			pattern = []string{}
			mirrors = append(mirrors, s)
		} else {
			pattern = append(pattern, line)
		}
	}
	if len(pattern) > 0 {
		s := findMirror(pattern)
		summary += s
		mirrors = append(mirrors, s)
	}

	fmt.Println(summary)
	return mirrors
}

func part2(input string, mirrors []int) {
	lines := conv.SplitNewline(input)
	summary := 0

	patternNo := 0
	var pattern []string
	for _, line := range lines {
		if line == "" {

		outer:
			for r := range pattern {
				for c := range pattern[r] {
					origStr := pattern[r]
					if pattern[r][c] == '#' {
						pattern[r] = pattern[r][:c] + "." + pattern[r][c+1:]
					} else {
						pattern[r] = pattern[r][:c] + "#" + pattern[r][c+1:]
					}
					s := findMirrorIgnore(pattern, mirrors[patternNo])
					pattern[r] = origStr

					if s != 0 {
						summary += s
						break outer
					}
				}
			}

			pattern = []string{}
			patternNo++
		} else {
			pattern = append(pattern, line)
		}
	}
	patternNo = len(mirrors) - 1
	if len(pattern) > 0 {
	outer2:
		for r := range pattern {
			for c := range pattern[r] {
				origStr := pattern[r]
				if pattern[r][c] == '#' {
					pattern[r] = pattern[r][:c] + "." + pattern[r][c+1:]
				} else {
					pattern[r] = pattern[r][:c] + "#" + pattern[r][c+1:]
				}
				s := findMirrorIgnore(pattern, mirrors[patternNo])
				pattern[r] = origStr
				if s != 0 {
					summary += s
					break outer2
				}
			}
		}
	}

	fmt.Println(summary)
}

func findMirror(pattern []string) int {
	mirror := findMirrorColumnVertically(pattern, -1)
	if mirror == -1 {
		mirror = findMirrorRowHorizontally(pattern, -1)
		if mirror != -1 {
			return mirror * 100
		}
	} else {
		return mirror
	}
	return 0
}

func findMirrorIgnore(pattern []string, ignore int) int {
	mirror := findMirrorColumnVertically(pattern, ignore)
	if mirror == -1 || mirror == ignore {
		mirror = findMirrorRowHorizontally(pattern, ignore/100)
		if mirror != -1 {
			return mirror * 100
		}
	} else {
		return mirror
	}
	return 0
}

func findMirrorRowHorizontally(pattern []string, ignore int) int {
	for row := 1; row < len(pattern); row++ {
		if isMirrorHorizontally(pattern, row) && row != ignore {
			return row
		}
	}
	return -1
}

func findMirrorColumnVertically(pattern []string, ignore int) int {
	for col := 1; col < len(pattern[0]); col++ {
		if isMirrorVertically(pattern, col) && col != ignore {
			return col
		}
	}
	return -1
}

func isMirrorHorizontally(pattern []string, middle int) bool {
	cols := len(pattern[0])
	for i := 0; i < middle; i++ {
		for col := 0; col < cols; col++ {
			upper := ""
			lower := ""
			for row := 0; row < middle; row++ {
				if middle-row-1 < 0 || middle+row >= len(pattern) {
					break
				}
				upper += string(pattern[middle-row-1][col])
				lower += string(pattern[middle+row][col])
			}

			if upper != "" && upper != lower {
				return false
			}
		}
	}
	return true
}

func isMirrorVertically(pattern []string, middle int) bool {
	for _, row := range pattern {
		for col := 0; col < middle; col++ {
			if middle-col-1 < 0 || middle+col >= len(row) {
				break
			}
			if row[middle-col-1] != row[middle+col] {
				return false
			}
		}
	}
	return true
}
