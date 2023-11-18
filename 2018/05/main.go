package main

import (
	"aoc/internal/download"
	"fmt"
	"log"
)

func main() {
	inputFile := "./2018/05/input.txt"
	input, err := download.ReadInput(inputFile, 2018, 5)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	if input[len(input)-1] == '\n' {
		input = input[:len(input)-1]
	}
	for {
		changed := false
		for i := 0; i < len(input)-1; i++ {
			if input[i] == input[i+1]+32 || input[i] == input[i+1]-32 {
				input = input[:i] + input[i+2:]
				changed = true
			}
		}
		if !changed {
			break
		}
	}
	fmt.Println(len(input))
}

func part2(input string) {
	if input[len(input)-1] == '\n' {
		input = input[:len(input)-1]
	}
	minInputLen := len(input)
	var c uint8
	for c = 'a'; c <= 'z'; c++ {
		out := ""
		for i := 0; i < len(input); i++ {
			if input[i] != c && input[i] != c-32 {
				out += string(input[i])
			}
		}
		for {
			changed := false
			for i := 0; i < len(out)-1; i++ {
				if out[i] == out[i+1]+32 || out[i] == out[i+1]-32 {
					out = out[:i] + out[i+2:]
					changed = true
				}
			}
			if !changed {
				break
			}
		}
		if len(out) < minInputLen {
			minInputLen = len(out)
		}
	}
	fmt.Println(minInputLen)
}
