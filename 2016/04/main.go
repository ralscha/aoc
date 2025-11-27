package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"slices"
	"strings"
)

func main() {
	input, err := download.ReadInput(2016, 4)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

type char struct {
	c     rune
	count int
}

func part1(input string) {
	lines := conv.SplitNewline(input)

	sum := 0
	for _, line := range lines {
		bracket := strings.Index(line, "[")
		check := line[bracket+1 : len(line)-1]

		lastDash := strings.LastIndex(line, "-")
		bag := container.NewBag[rune]()
		for _, r := range line[:lastDash] {
			if r == '-' {
				continue
			}
			bag.Add(r)
		}

		var chars []char
		for r, count := range bag.Values() {
			chars = append(chars, char{r, count})
		}
		slices.SortFunc(chars, func(i, j char) int {
			if i.count == j.count {
				return int(i.c - j.c)
			}
			return j.count - i.count
		})

		computedCheck := ""
		for i := range 5 {
			computedCheck += string(chars[i].c)
		}

		if computedCheck == check {
			sectorID := conv.MustAtoi(line[lastDash+1 : bracket])
			sum += sectorID
		}
	}
	fmt.Println("Part 1", sum)
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	for _, line := range lines {
		bracket := strings.Index(line, "[")
		lastDash := strings.LastIndex(line, "-")
		sectorID := conv.MustAtoi(line[lastDash+1 : bracket])

		decrypted := ""
		for _, r := range line[:lastDash] {
			if r == '-' {
				decrypted += " "
				continue
			}
			decrypted += string(rune((int(r)-'a'+sectorID)%26 + 'a'))
		}
		if strings.Contains(decrypted, "north") {
			fmt.Println("Part 2", decrypted, sectorID)
		}
	}
}
