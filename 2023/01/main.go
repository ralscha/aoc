package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	input, err := download.ReadInput(2023, 1)
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
		firstDigit := ""
		lastDigit := ""
		for _, c := range line {
			if unicode.IsDigit(c) {
				if firstDigit == "" {
					firstDigit = string(c)
				}
				lastDigit = string(c)
			}
		}
		sum = sum + conv.MustAtoi(firstDigit+lastDigit)
	}
	fmt.Println(sum)
}

func part2(input string) {
	numbers := []string{
		"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
	}

	lines := conv.SplitNewline(input)
	sum := 0
	for _, line := range lines {
		firstDigit := ""
		lastDigit := ""

		indexOfFirstDigit := -1
		indexOfLastDigit := len(line)
		for i, c := range line {
			if unicode.IsDigit(c) {
				firstDigit = string(c)
				indexOfFirstDigit = i
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			if unicode.IsDigit(rune(line[i])) {
				lastDigit = string(line[i])
				indexOfLastDigit = i
				break
			}
		}

		if indexOfFirstDigit > 0 {
			for i := 0; i < indexOfFirstDigit; i++ {
				for j, number := range numbers {
					if strings.HasPrefix(line[i:], number) {
						firstDigit = strconv.Itoa(j)
						indexOfFirstDigit = i
						break
					}
				}
			}
		}

		if indexOfLastDigit < len(line)-1 {
			for i := indexOfLastDigit + 1; i < len(line); i++ {
				for j, number := range numbers {
					if strings.HasPrefix(line[i:], number) {
						lastDigit = strconv.Itoa(j)
						indexOfLastDigit = i + len(number) - 1
						break
					}
				}
			}
		}
		sum = sum + conv.MustAtoi(firstDigit+lastDigit)
	}
	fmt.Println(sum)
}
