package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"slices"
	"strings"
)

func main() {
	inputFile := "./2016/04/input.txt"
	input, err := download.ReadInput(inputFile, 2016, 4)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

type char struct {
	c          rune
	occurrence int
}

func part1(input string) {
	lines := conv.SplitNewline(input)

	sum := 0
	for _, line := range lines {
		bracket := strings.Index(line, "[")
		check := line[bracket+1 : len(line)-1]

		lastDash := strings.LastIndex(line, "-")
		occurrences := make(map[rune]*char)
		for _, r := range line[:lastDash] {
			if r == '-' {
				continue
			}
			if v, ok := occurrences[r]; ok {
				v.occurrence++
			} else {
				occurrences[r] = &char{r, 1}
			}
		}

		var chars []char
		for _, v := range occurrences {
			chars = append(chars, *v)
		}
		slices.SortFunc(chars, func(i, j char) int {
			if i.occurrence == j.occurrence {
				return int(i.c - j.c)
			}
			return j.occurrence - i.occurrence
		})
		computedCheck := ""
		for i := 0; i < 5; i++ {
			computedCheck += string(chars[i].c)
		}

		if computedCheck == check {
			sectorID := conv.MustAtoi(line[lastDash+1 : bracket])
			sum += sectorID
		}
	}
	fmt.Println(sum)
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
			fmt.Println(decrypted, sectorID)
		}
	}
}
