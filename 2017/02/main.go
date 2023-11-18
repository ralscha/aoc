package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"golang.org/x/exp/slices"
	"log"
	"strings"
)

func main() {
	inputFile := "./2017/02/input.txt"
	input, err := download.ReadInput(inputFile, 2017, 2)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	sum := 0
	for _, line := range lines {
		numbers := strings.Fields(line)

		minNumbers := slices.MinFunc(numbers, func(i, j string) int {
			return conv.MustAtoi(i) - conv.MustAtoi(j)
		})
		maxNumbers := slices.MaxFunc(numbers, func(i, j string) int {
			return conv.MustAtoi(i) - conv.MustAtoi(j)
		})

		sum += conv.MustAtoi(maxNumbers) - conv.MustAtoi(minNumbers)
	}
	fmt.Println(sum)
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	sum := 0
	for _, line := range lines {
		numbers := strings.Fields(line)
		for i, number := range numbers {
			n := conv.MustAtoi(number)
			for j, number2 := range numbers {
				if i == j {
					continue
				}
				n2 := conv.MustAtoi(number2)
				if n%n2 == 0 {
					sum += n / n2
				}
			}
		}
	}
	fmt.Println(sum)
}
