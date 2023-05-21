package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"strings"
)

func main() {
	inputFile := "./2015/25/input.txt"
	input, err := download.ReadInput(inputFile, 2015, 25)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
}

func part1(input string) {
	splitted := strings.Fields(input)
	fmt.Println(splitted)
	row := conv.MustAtoi(splitted[15][:len(splitted[15])-1])
	col := conv.MustAtoi(splitted[17][:len(splitted[17])-1])
	code := 20151125
	total := row + col - 2
	total = (total * (total + 1)) / 2
	total += col
	for i := 1; i < total; i++ {
		code = (code * 252533) % 33554393
	}
	fmt.Println(code)
}
