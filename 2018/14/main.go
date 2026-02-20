package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input, err := download.ReadInput(2018, 14)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	numRecipes := conv.MustAtoi(input)
	scoreboard := []int{3, 7}
	elf1Pos := 0
	elf2Pos := 1

	for len(scoreboard) < numRecipes+10 {
		sum := scoreboard[elf1Pos] + scoreboard[elf2Pos]
		if sum >= 10 {
			scoreboard = append(scoreboard, sum/10, sum%10)
		} else {
			scoreboard = append(scoreboard, sum)
		}

		elf1Pos = (elf1Pos + 1 + scoreboard[elf1Pos]) % len(scoreboard)
		elf2Pos = (elf2Pos + 1 + scoreboard[elf2Pos]) % len(scoreboard)
	}

	var result strings.Builder
	for i := numRecipes; i < numRecipes+10; i++ {
		result.WriteString(strconv.Itoa(scoreboard[i]))
	}
	fmt.Println("Part 1", result.String())
}

func part2(input string) {
	target := make([]int, len(input))
	for i, r := range input {
		target[i] = int(r - '0')
	}

	scoreboard := []int{3, 7}
	elf1Pos := 0
	elf2Pos := 1

	for {
		sum := scoreboard[elf1Pos] + scoreboard[elf2Pos]
		var digits []int
		if sum >= 10 {
			digits = append(digits, sum/10, sum%10)
		} else {
			digits = append(digits, sum)
		}

		for _, digit := range digits {
			scoreboard = append(scoreboard, digit)
			if len(scoreboard) >= len(target) {
				match := true
				for i := range len(target) {
					if scoreboard[len(scoreboard)-len(target)+i] != target[i] {
						match = false
						break
					}
				}
				if match {
					fmt.Println("Part 2", len(scoreboard)-len(target))
					return
				}
			}
		}

		elf1Pos = (elf1Pos + 1 + scoreboard[elf1Pos]) % len(scoreboard)
		elf2Pos = (elf2Pos + 1 + scoreboard[elf2Pos]) % len(scoreboard)
	}
}
