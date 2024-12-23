package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
)

func main() {
	input, err := download.ReadInput(2020, 15)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	starting := conv.ToIntSliceComma(input)
	spoken := make(map[int]int)
	var lastNum int

	for i, num := range starting {
		spoken[num] = i + 1
		lastNum = num
	}

	for turn := len(starting) + 1; turn <= 2020; turn++ {
		var nextNum int
		if lastTurn, ok := spoken[lastNum]; ok {
			nextNum = turn - 1 - lastTurn
		} else {
			nextNum = 0
		}
		spoken[lastNum] = turn - 1
		lastNum = nextNum
	}

	fmt.Println("Part 1", lastNum)
}

func part2(input string) {
	starting := conv.ToIntSliceComma(input)
	spoken := make(map[int]int)
	var lastNum int

	for i, num := range starting {
		spoken[num] = i + 1
		lastNum = num
	}

	for turn := len(starting) + 1; turn <= 30000000; turn++ {
		var nextNum int
		if lastTurn, ok := spoken[lastNum]; ok {
			nextNum = turn - 1 - lastTurn
		} else {
			nextNum = 0
		}
		spoken[lastNum] = turn - 1
		lastNum = nextNum
	}

	fmt.Println("Part 2", lastNum)
}
