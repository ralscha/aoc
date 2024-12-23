package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"strings"
)

func main() {
	input, err := download.ReadInput(2017, 13)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func buildFirewall(input string) (map[int]int, int) {
	firewall := make(map[int]int)
	maxDepth := 0
	for _, line := range conv.SplitNewline(input) {
		parts := strings.Split(line, ": ")
		depth := conv.MustAtoi(parts[0])
		rangeVal := conv.MustAtoi(parts[1])
		firewall[depth] = rangeVal
		if depth > maxDepth {
			maxDepth = depth
		}
	}
	return firewall, maxDepth
}

func part1(input string) {
	firewall, maxDepth := buildFirewall(input)

	severity := 0
	for depth := 0; depth <= maxDepth; depth++ {
		rangeVal, ok := firewall[depth]
		if ok {
			if rangeVal == 1 {
				severity += depth * rangeVal
			} else {
				cycle := 2 * (rangeVal - 1)
				if depth%cycle == 0 {
					severity += depth * rangeVal
				}
			}
		}
	}

	fmt.Println("Part 1", severity)
}

func part2(input string) {
	firewall, maxDepth := buildFirewall(input)

	delay := 0
	for {
		caught := false
		for depth := 0; depth <= maxDepth; depth++ {
			rangeVal, ok := firewall[depth]
			if ok {
				if rangeVal == 1 {
					caught = true
					break
				} else {
					cycle := 2 * (rangeVal - 1)
					if (delay+depth)%cycle == 0 {
						caught = true
						break
					}
				}
			}
		}
		if !caught {
			fmt.Println("Part 2", delay)
			return
		}
		delay++
	}
}
