package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"strings"
)

func main() {
	input, err := download.ReadInput(2023, 4)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	totalPoints := 0
	for _, line := range lines {
		splitted := strings.Split(line, ":")
		parts := strings.Split(splitted[1], "|")

		winningStr := strings.Fields(parts[0])
		winningNumbers := make([]int, len(winningStr))
		for i, numStr := range winningStr {
			winningNumbers[i] = conv.MustAtoi(numStr)
		}

		myStr := strings.Fields(parts[1])
		myNumbers := make([]int, len(myStr))
		for i, numStr := range myStr {
			myNumbers[i] = conv.MustAtoi(numStr)
		}

		points, _ := calcPoints(winningNumbers, myNumbers)
		totalPoints += points
	}

	fmt.Println(totalPoints)
}

func calcPoints(winningNumbers, myNumbers []int) (int, int) {
	numberOfWinningNumbers := 0
	for _, num := range myNumbers {
		for _, winNum := range winningNumbers {
			if num == winNum {
				numberOfWinningNumbers++
				break
			}
		}
	}
	if numberOfWinningNumbers == 0 {
		return 0, 0
	}
	return 1 << (numberOfWinningNumbers - 1), numberOfWinningNumbers
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	totalScratchcards := countScratchcards(lines)
	fmt.Println(totalScratchcards + len(lines))
}

var cardCache = make(map[string]int)

func countScratchcards(lines []string) int {
	totalScratchcards := 0

	for ix, line := range lines {
		if _, ok := cardCache[line]; !ok {
			splitted := strings.Split(line, ":")
			parts := strings.Split(splitted[1], "|")

			winningStr := strings.Fields(parts[0])
			winningNumbers := make([]int, len(winningStr))
			for i, numStr := range winningStr {
				winningNumbers[i] = conv.MustAtoi(numStr)
			}

			myStr := strings.Fields(parts[1])
			myNumbers := make([]int, len(myStr))
			for i, numStr := range myStr {
				myNumbers[i] = conv.MustAtoi(numStr)
			}

			_, numberOfNextScratchcards := calcPoints(winningNumbers, myNumbers)
			cardCache[line] = numberOfNextScratchcards
		}

		numberOfNextScratchcards, _ := cardCache[line]
		totalScratchcards += numberOfNextScratchcards
		if numberOfNextScratchcards > 0 {
			totalScratchcards += countScratchcards(lines[ix+1 : ix+1+numberOfNextScratchcards])
		}
	}
	return totalScratchcards
}
