package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"strings"
)

func main() {
	inputFile := "./2015/05/input.txt"
	input, err := download.ReadInput(inputFile, 2015, 5)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)

}

func part1(input string) {
	lines := conv.SplitNewline(input)
	count := 0
	for _, line := range lines {
		if isNice(line) {
			count++
		}
	}
	fmt.Println(count)
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	count := 0
	for _, line := range lines {
		if isNice2(line) {
			count++
		}
	}
	fmt.Println(count)
}

func isNice(line string) bool {
	// It contains at least three vowels
	vowels := 0
	for _, c := range line {
		if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
			vowels++
		}
	}
	if vowels < 3 {
		return false
	}

	// It contains at least one letter that appears twice in a row
	doubleChar := false
	for i := 0; i < len(line)-1; i++ {
		if line[i] == line[i+1] {
			doubleChar = true
			break
		}
	}
	if !doubleChar {
		return false
	}

	// It does not contain the strings: ab, cd, pq, or xy
	if strings.Contains(line, "ab") || strings.Contains(line, "cd") || strings.Contains(line, "pq") || strings.Contains(line, "xy") {
		return false
	}

	return true
}

func isNice2(line string) bool {
	// It contains a pair of any two letters that appears at least twice in the string without overlapping,
	// like xyxy (xy) or aabcdefgaa (aa), but not like aaa (aa, but it overlaps)
	pair := false
	for i := 0; i < len(line)-1; i++ {
		if strings.Count(line, line[i:i+2]) > 1 {
			pair = true
			break
		}
	}
	if !pair {
		return false
	}

	// It contains at least one letter which repeats with exactly one letter
	// between them, like xyx, abcdefeghi (efe), or even aaa.
	repeat := false
	for i := 0; i < len(line)-2; i++ {
		if line[i] == line[i+2] {
			repeat = true
			break
		}
	}
	if !repeat {
		return false
	}

	return true
}
