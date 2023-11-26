package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"strings"
)

func main() {
	input, err := download.ReadInput(2016, 9)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}
	input = strings.Trim(input, "\n")
	part1(input)
	part2(input)
}

func part1(input string) {
	decompressed := ""
	for i := 0; i < len(input); i++ {
		if input[i] == '(' {
			end := strings.Index(input[i:], ")")
			marker := input[i+1 : i+end]
			parts := strings.Split(marker, "x")
			length := conv.MustAtoi(parts[0])
			repeat := conv.MustAtoi(parts[1])
			for j := 0; j < repeat; j++ {
				decompressed += input[i+end+1 : i+end+1+length]
			}
			i += end + length
		} else {
			decompressed += string(input[i])
		}
	}
	fmt.Println(len(decompressed))
}

func part2(input string) {
	fmt.Println(decompress(input))
}

func decompress(input string) int {
	decompressed := 0
	for i := 0; i < len(input); i++ {
		if input[i] == '(' {
			end := strings.Index(input[i:], ")")
			marker := input[i+1 : i+end]
			parts := strings.Split(marker, "x")
			length := conv.MustAtoi(parts[0])
			repeat := conv.MustAtoi(parts[1])
			decompressed += decompress(input[i+end+1:i+end+1+length]) * repeat
			i += end + length
		} else {
			decompressed++
		}
	}
	return decompressed
}
