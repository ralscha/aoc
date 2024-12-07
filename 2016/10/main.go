package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"aoc/internal/download"
	"log"
	"slices"
	"strings"
)

func main() {
	input, err := download.ReadInput(2016, 10)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}
	lines := conv.SplitNewline(input)
	part1and2(lines)
}

func part1and2(lines []string) {
	bots := make(map[int]*container.Queue[int])
	outputs := make(map[int]int)

	// Initialize queues for each bot
	for _, line := range lines {
		if strings.HasPrefix(line, "value") {
			splitted := strings.Split(line, " ")
			value := conv.MustAtoi(splitted[1])
			bot := conv.MustAtoi(splitted[5])
			if _, ok := bots[bot]; !ok {
				bots[bot] = container.NewQueue[int]()
			}
			bots[bot].Push(value)
		}
	}

	part1Solved := false
	part2Solved := false

	for {
		for _, line := range lines {
			if strings.HasPrefix(line, "bot") {
				splitted := strings.Split(line, " ")
				bot := conv.MustAtoi(splitted[1])
				if _, ok := bots[bot]; !ok || bots[bot].Len() != 2 {
					continue
				}

				// Get both values from the queue
				values := []int{bots[bot].Pop(), bots[bot].Pop()}
				if (values[0] == 17 && values[1] == 61) || (values[0] == 61 && values[1] == 17) {
					log.Printf("Part 1: %v", bot)
					part1Solved = true
					if part2Solved {
						return
					}
				}

				low := conv.MustAtoi(splitted[6])
				high := conv.MustAtoi(splitted[11])
				minVal, maxVal := slices.Min(values), slices.Max(values)

				if splitted[5] == "output" {
					outputs[low] = minVal
				} else {
					if _, ok := bots[low]; !ok {
						bots[low] = container.NewQueue[int]()
					}
					bots[low].Push(minVal)
				}

				if splitted[10] == "output" {
					outputs[high] = maxVal
				} else {
					if _, ok := bots[high]; !ok {
						bots[high] = container.NewQueue[int]()
					}
					bots[high].Push(maxVal)
				}

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
