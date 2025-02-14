package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/mathx"
	"aoc/internal/stringutil"
	"fmt"
	"log"
	"strings"
)

func scramble(password string, instructions []string) string {
	for _, instruction := range instructions {
		parts := strings.Split(instruction, " ")
		switch parts[0] {
		case "swap":
			if parts[1] == "position" {
				x := conv.MustAtoi(parts[2])
				y := conv.MustAtoi(parts[5])
				temp := password[x]
				password = password[:x] + string(password[y]) + password[x+1:]
				password = password[:y] + string(temp) + password[y+1:]
			} else {
				x := parts[2][0]
				y := parts[5][0]
				password = strings.ReplaceAll(password, string(x), "#")
				password = strings.ReplaceAll(password, string(y), string(x))
				password = strings.ReplaceAll(password, "#", string(y))
			}
		case "rotate":
			if parts[1] == "based" {
				x := strings.Index(password, string(parts[6][0]))
				rot := 1 + x
				if x >= 4 {
					rot++
				}
				rot = rot % len(password)
				password = password[len(password)-rot:] + password[:len(password)-rot]

			} else {
				x := conv.MustAtoi(parts[2])
				if parts[1] == "left" {
					x = len(password) - x
				}
				password = password[len(password)-x:] + password[:len(password)-x]
			}
		case "reverse":
			x := conv.MustAtoi(parts[2])
			y := conv.MustAtoi(parts[4])
			password = password[:x] + stringutil.Reverse(password[x:y+1]) + password[y+1:]
		case "move":
			x := conv.MustAtoi(parts[2])
			y := conv.MustAtoi(parts[5])
			temp := password[x]
			password = password[:x] + password[x+1:]
			password = password[:y] + string(temp) + password[y:]
		}
	}
	return password
}

func main() {
	input, err := download.ReadInput(2016, 21)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	instructions := conv.SplitNewline(input)
	password := "abcdefgh"
	fmt.Println("Part 1", scramble(password, instructions))
}

func part2(input string) {
	instructions := conv.SplitNewline(input)
	target := "fbgdceah"
	for _, perm := range mathx.Permutations([]rune("abcdefgh")) {
		if scramble(string(perm), instructions) == target {
			fmt.Println("Part 2", string(perm))
			return
		}
	}
}
