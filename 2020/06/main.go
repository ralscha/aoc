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
	input, err := download.ReadInput(2020, 6)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	groups := strings.Split(input, "\n\n")
	totalCount := 0
	for _, group := range groups {
		s := container.NewSet[rune]()
		for _, personAnswers := range conv.SplitNewline(group) {
			for _, answer := range personAnswers {
				s.Add(answer)
			}
		}
		totalCount += s.Len()
	}
	fmt.Println("Part 1", totalCount)
}

func part2(input string) {
	groups := strings.Split(input, "\n\n")
	totalCount := 0
	for _, group := range groups {
		persons := conv.SplitNewline(group)
		if len(persons) == 0 || persons[0] == "" {
			continue
		}
		groupAnswers := container.NewSet[rune]()
		for _, answer := range persons[0] {
			groupAnswers.Add(answer)
		}

		for i := 1; i < len(persons); i++ {
			personAnswers := container.NewSet[rune]()
			for _, answer := range persons[i] {
				personAnswers.Add(answer)
			}
			intersection := container.NewSet[rune]()
			for _, val := range groupAnswers.Values() {
				if personAnswers.Contains(val) {
					intersection.Add(val)
				}
			}
			groupAnswers = intersection
		}
		totalCount += groupAnswers.Len()
	}
	fmt.Println("Part 2", totalCount)
}
