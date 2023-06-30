package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
)

func main() {
	inputFile := "./2018/02/input.txt"
	input, err := download.ReadInput(inputFile, 2018, 2)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	lines := conv.SplitNewline(input)

	twos := 0
	threes := 0

	for _, line := range lines {
		bag := container.NewBag[rune]()
		for _, char := range line {
			bag.Add(char)
		}

		hasTwo := false
		hasThree := false
		for _, count := range bag.Values() {
			if count == 2 {
				hasTwo = true
			}
			if count == 3 {
				hasThree = true
			}
			if hasTwo && hasThree {
				break
			}
		}
		if hasTwo {
			twos++
		}
		if hasThree {
			threes++
		}
	}
	checksum := twos * threes
	fmt.Println(checksum)
}

func part2(input string) {
	lines := conv.SplitNewline(input)

	for i, line := range lines {
		for _, line2 := range lines[i+1:] {
			difference := 0
			for j := range line {
				if line[j] != line2[j] {
					difference++
				}
				if difference > 1 {
					break
				}
			}
			if difference == 1 {
				for j := range line {
					if line[j] == line2[j] {
						fmt.Printf("%c", line[j])
					}
				}
				fmt.Println()
				return
			}
		}
	}
}
