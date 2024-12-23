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
	input, err := download.ReadInput(2017, 16)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	moves := strings.Split(input, ",")
	programs := []rune("abcdefghijklmnop")

	for _, move := range moves {
		switch move[0] {
		case 's':
			x := conv.MustAtoi(move[1:])
			spin(programs, x)
		case 'x':
			parts := strings.Split(move[1:], "/")
			a := conv.MustAtoi(parts[0])
			b := conv.MustAtoi(parts[1])
			exchange(programs, a, b)
		case 'p':
			parts := strings.Split(move[1:], "/")
			partner(programs, rune(parts[0][0]), rune(parts[1][0]))
		}
	}

	fmt.Println("Part 1", string(programs))
}

func part2(input string) {
	moves := strings.Split(input, ",")
	programs := []rune("abcdefghijklmnop")

	seen := make(map[string]int)
	var finalOrder string

	for i := range 1000000000 {
		currentOrder := string(programs)
		if prev, ok := seen[currentOrder]; ok {
			remainingCycles := 1000000000 - i
			cycleLength := i - prev
			remaining := remainingCycles % cycleLength
			for k := 0; k < remaining; k++ {
				for _, move := range moves {
					applyMove(programs, move)
				}
			}
			finalOrder = string(programs)
			break
		}
		seen[currentOrder] = i
		for _, move := range moves {
			applyMove(programs, move)
		}
		if i == 999999999 {
			finalOrder = string(programs)
		}
	}

	fmt.Println("Part 2", finalOrder)
}

func spin(programs []rune, x int) {
	n := len(programs)
	x = x % n
	temp := slices.Clone(programs[n-x:])
	programs = slices.Replace(programs, x, n, programs[:n-x]...)
	programs = slices.Replace(programs, 0, x, temp...)
}

func exchange(programs []rune, a, b int) {
	programs[a], programs[b] = programs[b], programs[a]
}

func partner(programs []rune, a, b rune) {
	var indexA, indexB int
	for i, p := range programs {
		if p == a {
			indexA = i
		}
		if p == b {
			indexB = i
		}
	}
	programs[indexA], programs[indexB] = programs[indexB], programs[indexA]
}

func applyMove(programs []rune, move string) {
	switch move[0] {
	case 's':
		x := conv.MustAtoi(move[1:])
		spin(programs, x)
	case 'x':
		parts := strings.Split(move[1:], "/")
		a := conv.MustAtoi(parts[0])
		b := conv.MustAtoi(parts[1])
		exchange(programs, a, b)
	case 'p':
		parts := strings.Split(move[1:], "/")
		partner(programs, rune(parts[0][0]), rune(parts[1][0]))
	}
}
