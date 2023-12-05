package main

import (
	"aoc/internal/container"
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
		winningNumbers := strings.Fields(parts[0])
		myNumbers := strings.Fields(parts[1])
		winningSet := container.NewSet[string]()

		for _, w := range winningNumbers {
			winningSet.Add(w)
		}

		points, _ := calcPoints(myNumbers, winningSet)
		totalPoints += points
	}

	println(totalPoints)
}

func calcPoints(myNumbers []string, winningSet container.Set[string]) (int, int) {
	points := 0
	numberOfWinningNumbers := 0
	for _, y := range myNumbers {
		if winningSet.Contains(y) {
			if points == 0 {
				points = 1
			} else {
				points *= 2
			}
			numberOfWinningNumbers++
		}
	}
	return points, numberOfWinningNumbers
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

			winningNumbers := strings.Fields(parts[0])
			myNumbers := strings.Fields(parts[1])
			winningSet := container.NewSet[string]()

			for _, w := range winningNumbers {
				winningSet.Add(w)
			}

			_, numberOfNextScratchcards := calcPoints(myNumbers, winningSet)
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
