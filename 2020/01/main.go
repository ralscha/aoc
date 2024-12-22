package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
)

func main() {
	input, err := download.ReadInput(2020, 1)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	nums := conv.ToIntSlice(lines)

	for i, num1 := range nums {
		for j, num2 := range nums {
			if i == j {
				continue
			}
			if num1+num2 == 2020 {
				fmt.Println("Part 1", num1*num2)
				return
			}
		}
	}
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	nums := conv.ToIntSlice(lines)

	for i, num1 := range nums {
		for j, num2 := range nums {
			for k, num3 := range nums {
				if i == j || i == k || j == k {
					continue
				}
				if num1+num2+num3 == 2020 {
					fmt.Println("Part 2", num1*num2*num3)
					return
				}
			}
		}
	}
}
