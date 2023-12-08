package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/mathx"
	"fmt"
	"log"
	"strings"
)

func main() {
	input, err := download.ReadInput(2023, 8)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

type next struct {
	left  string
	right string
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	instructions := lines[0]
	network := buildNetwork(lines)

	steps := 0
	current := "AAA"
	currentInstruction := 0
	for {
		instruction := instructions[currentInstruction]
		currentInstruction = (currentInstruction + 1) % len(instructions)
		steps += 1

		if instruction == 'L' {
			current = network[current].left
		} else {
			current = network[current].right
		}
		if current == "ZZZ" {
			break
		}
	}

	fmt.Println(steps)
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	instructions := lines[0]
	network := buildNetwork(lines)

	currentKeys := make([]string, 0)
	for key := range network {
		if key[len(key)-1] == 'A' {
			currentKeys = append(currentKeys, key)
		}
	}

	var keySteps []int
	for _, key := range currentKeys {
		steps := 0
		current := key
		currentInstruction := 0
		for {
			instruction := instructions[currentInstruction]
			currentInstruction = (currentInstruction + 1) % len(instructions)
			steps += 1

			if instruction == 'L' {
				current = network[current].left
			} else {
				current = network[current].right
			}
			if current[len(current)-1] == 'Z' {
				break
			}
		}
		keySteps = append(keySteps, steps)
	}

	fmt.Println(mathx.Lcm(keySteps))
}

func buildNetwork(lines []string) map[string]next {
	network := make(map[string]next)
	for i := 2; i < len(lines); i++ {
		line := lines[i]
		split := strings.Split(line, "=")
		key := strings.TrimSpace(split[0])
		splitNext := strings.Split(split[1], ",")
		nextLeft := strings.TrimSpace(splitNext[0])
		nextRight := strings.TrimSpace(splitNext[1])
		network[key] = next{left: nextLeft[1:], right: nextRight[:len(nextRight)-1]}
	}
	return network
}
