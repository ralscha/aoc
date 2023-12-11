package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"bufio"
	"fmt"
	"log"
	"strings"
)

func main() {
	input, err := download.ReadInput(2021, 2)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	horizontalPosition := 0
	depth := 0

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Fields(line)
		num := conv.MustAtoi(split[1])
		switch split[0] {
		case "down":
			depth += num
		case "up":
			depth -= num
		case "forward":
			horizontalPosition += num
		}
	}

	fmt.Println("Depth: ", depth)
	fmt.Println("Horizontal Position: ", horizontalPosition)
	fmt.Println(depth * horizontalPosition)
	fmt.Println()
}

func part2(input string) {
	horizontalPosition := 0
	depth := 0
	aim := 0

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Fields(line)
		num := conv.MustAtoi(split[1])
		switch split[0] {
		case "down":
			aim += num
		case "up":
			aim -= num
		case "forward":
			horizontalPosition += num
			depth += aim * num
		}
	}

	fmt.Println("Depth: ", depth)
	fmt.Println("Horizontal Position: ", horizontalPosition)
	fmt.Println("Aim: ", aim)
	fmt.Println(depth * horizontalPosition)
}
