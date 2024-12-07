package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/mathx"
	"fmt"
	"log"
)

func main() {
	input, err := download.ReadInput(2015, 20)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	target := conv.MustAtoi(input)
	house := 1
	for {
		presents := 0
		for _, factor := range mathx.Factors(house) {
			presents += factor * 10
		}
		if presents >= target {
			break
		}
		house++
	}

	fmt.Println("Part 1", house)
}

func part2(input string) {
	target := conv.MustAtoi(input)
	house := 2
	stopAfter := 50
	for {
		presents := 0
		for _, factor := range mathx.Factors(house) {
			if house/factor <= stopAfter {
				presents += factor * 11
			}
		}
		if presents >= target {
			break
		}
		house++
	}

	fmt.Println("Part 2", house)
}
