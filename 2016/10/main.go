package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"golang.org/x/exp/slices"
	"log"
	"strings"
)

func main() {
	inputFile := "./2016/10/input.txt"
	input, err := download.ReadInputAuto(inputFile)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}
	lines := conv.SplitNewline(input)
	part1and2(lines)
}

func part1and2(lines []string) {

	bots := make(map[int][]int)
	outputs := make(map[int]int)

	for _, line := range lines {
		if strings.HasPrefix(line, "value") {
			splitted := strings.Split(line, " ")
			value := conv.MustAtoi(splitted[1])
			bot := conv.MustAtoi(splitted[5])
			if _, ok := bots[bot]; !ok {
				bots[bot] = []int{value}
			} else {
				bots[bot] = append(bots[bot], value)
			}
		}
	}

	part1Solved := false
	part2Solved := false

	for {
		for _, line := range lines {
			if strings.HasPrefix(line, "bot") {
				splitted := strings.Split(line, " ")
				bot := conv.MustAtoi(splitted[1])
				payload := bots[bot]
				if len(payload) != 2 {
					continue
				}
				low := conv.MustAtoi(splitted[6])
				high := conv.MustAtoi(splitted[11])
				if _, ok := bots[bot]; !ok {
					continue
				}

				if (payload[0] == 17 && payload[1] == 61) || (payload[0] == 61 && payload[1] == 17) {
					log.Printf("Part 1: %v", bot)
					part1Solved = true
					if part2Solved {
						return
					}
				}
				if splitted[5] == "output" {
					outputs[low] = slices.Min(payload)
				} else {
					bots[low] = append(bots[low], slices.Min(payload))
				}
				if splitted[10] == "output" {
					outputs[high] = slices.Max(payload)
				} else {
					bots[high] = append(bots[high], slices.Max(payload))
				}
				bots[bot] = []int{}

				if outputs[0] != 0 && outputs[1] != 0 && outputs[2] != 0 {
					log.Printf("Part 2: %v", outputs[0]*outputs[1]*outputs[2])
					part2Solved = true
					if part1Solved {
						return
					}
				}
			}
		}
	}
}
