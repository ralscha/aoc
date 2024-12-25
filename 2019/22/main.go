package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"strings"
)

func main() {
	input, err := download.ReadInput(2019, 22)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	deckSize := 10007
	cardToFind := 2019
	deck := make([]int, deckSize)
	for i := range deckSize {
		deck[i] = i
	}

	lines := conv.SplitNewline(input)
	for _, line := range lines {
		if line == "deal into new stack" {
			for i, j := 0, deckSize-1; i < j; i, j = i+1, j-1 {
				deck[i], deck[j] = deck[j], deck[i]
			}
		} else if strings.HasPrefix(line, "cut ") {
			nStr := line[4:]
			n := conv.MustAtoi(nStr)
			if n > 0 {
				top := deck[:n]
				rest := deck[n:]
				deck = append(rest, top...)
			} else {
				n = -n
				bottom := deck[deckSize-n:]
				rest := deck[:deckSize-n]
				deck = append(bottom, rest...)
			}
		} else if strings.HasPrefix(line, "deal with increment ") {
			nStr := line[20:]
			increment := conv.MustAtoi(nStr)
			newDeck := make([]int, deckSize)
			for i := range deckSize {
				newDeck[i] = -1
			}
			currentPos := 0
			for _, card := range deck {
				newDeck[currentPos] = card
				currentPos = (currentPos + increment) % deckSize
			}
			deck = newDeck
		}
	}

	position := -1
	for i, card := range deck {
		if card == cardToFind {
			position = i
			break
		}
	}

	fmt.Println("Part 1", position)
}

func part2(input string) {
	fmt.Println("Part 2 TODO")
}
