package main

import (
	"aoc/internal/download"
	"fmt"
	"log"
	"strings"
)

func main() {
	input, err := download.ReadInput(2017, 9)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}
	part1and2(input)
}

func part1and2(input string) {
	input = strings.TrimSuffix(input, "\n")
	depth := 0
	score := 0
	garbage := false
	ignore := false
	nonCanceledCharacters := 0

	for _, char := range input {
		if ignore {
			ignore = false
			continue
		}
		if char == '!' {
			ignore = true
			continue
		}
		if garbage {
			if char == '>' {
				garbage = false
			} else {
				nonCanceledCharacters++
			}
			continue
		}
		if char == '<' {
			garbage = true
			continue
		}
		if char == '{' {
			depth++
			continue
		}
		if char == '}' {
			score += depth
			depth--
			continue
		}
	}
	fmt.Println(score)
	fmt.Println(nonCanceledCharacters)
}
