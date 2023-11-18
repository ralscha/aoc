package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
)

func main() {
	inputFile := "./2016/07/input.txt"
	input, err := download.ReadInput(inputFile, 2016, 7)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}
	lines := conv.SplitNewline(input)
	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	supportTLS := 0
	for _, line := range lines {
		if supportsTLS(line) {
			supportTLS++
		}
	}
	fmt.Println(supportTLS)
}

func supportsTLS(line string) bool {
	insideBrackets := false
	hasOutside := false
	for i := 0; i < len(line)-3; i++ {
		if line[i] == '[' {
			insideBrackets = true
		} else if line[i] == ']' {
			insideBrackets = false
		}
		if line[i] == line[i+3] && line[i+1] == line[i+2] && line[i] != line[i+1] {
			if insideBrackets {
				return false
			}
			hasOutside = true
		}
	}
	return hasOutside
}

func part2(lines []string) {
	supportSSL := 0
	for _, line := range lines {
		if supportsSSL(line) {
			supportSSL++
		}
	}
	fmt.Println(supportSSL)
}

func supportsSSL(line string) bool {
	insideBrackets := false
	abas := make([]string, 0)
	babs := make([]string, 0)
	for i := 0; i < len(line)-2; i++ {
		if line[i] == '[' {
			insideBrackets = true
		} else if line[i] == ']' {
			insideBrackets = false
		}
		if line[i] == line[i+2] && line[i] != line[i+1] {
			if insideBrackets {
				babs = append(babs, line[i:i+3])
			} else {
				abas = append(abas, line[i:i+3])
			}
		}
	}
	for _, aba := range abas {
		for _, bab := range babs {
			if aba[0] == bab[1] && aba[1] == bab[0] {
				return true
			}
		}
	}
	return false
}
