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
	input, err := download.ReadInput(2021, 6)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Scan()
	line := scanner.Text()
	splitted := strings.Split(line, ",")

	var eco []*byte
	for _, s := range splitted {
		n := conv.MustAtoi(s)
		b := byte(n)
		eco = append(eco, &b)
	}

	for range 80 {
		for _, f := range eco {
			if *f > 0 {
				*f = *f - 1
			} else {
				*f = 6
				b := byte(8)
				eco = append(eco, &b)
			}
		}
	}

	fmt.Println("Part 1", len(eco))
}

func part2(input string) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Scan()
	line := scanner.Text()
	splitted := strings.Split(line, ",")

	var eco [9]int
	for _, s := range splitted {
		n := conv.MustAtoi(s)

		eco[n] = eco[n] + 1
	}

	for range 256 {
		var neweco [9]int
		for j := 1; j < len(eco); j++ {
			neweco[j-1] = eco[j]
		}
		neweco[6] += eco[0]
		neweco[8] += eco[0]
		eco = neweco
	}

	total := 0
	for _, f := range eco {
		total += f
	}
	fmt.Println("Part 2", total)
}
