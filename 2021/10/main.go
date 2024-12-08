package main

import (
	"aoc/internal/download"
	"bufio"
	"fmt"
	"log"
	"slices"
	"strings"
)

func main() {
	input, err := download.ReadInput(2021, 10)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	openP := "([{<"
	closeP := ")]}>"
	var stack []string

	scanner := bufio.NewScanner(strings.NewReader(input))
	score := 0
	for scanner.Scan() {
		line := scanner.Text()
		for _, c := range line {
			s := string(c)
			pos := strings.Index(openP, s)
			if pos != -1 {
				cl := string(closeP[pos])
				stack = append(stack, cl)
			} else {
				n := len(stack) - 1
				cl := stack[n]
				if s != cl {
					// log.Printf("Expected %s, but found %s instead.", cl, s)
					switch s {
					case ")":
						score += 3
					case "]":
						score += 57
					case "}":
						score += 1197
					case ">":
						score += 25137
					}
					break
				}
				stack = stack[:n]
			}
		}
	}

	fmt.Println("Part 1", score)

}

func part2(input string) {
	openP := "([{<"
	closeP := ")]}>"

	var scores []int
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		var stack []string
		invalid := false
		for _, c := range line {
			s := string(c)
			pos := strings.Index(openP, s)
			if pos != -1 {
				cl := string(closeP[pos])
				stack = append(stack, cl)
			} else {
				n := len(stack) - 1
				cl := stack[n]
				if s != cl {
					invalid = true
					break
				}
				stack = stack[:n]
			}
		}
		if !invalid && len(stack) > 0 {
			score := 0
			for len(stack) > 0 {
				score = score * 5
				n := len(stack) - 1
				cl := stack[n]
				switch cl {
				case ")":
					score += 1
				case "]":
					score += 2
				case "}":
					score += 3
				case ">":
					score += 4
				}
				stack = stack[:n]
			}

			scores = append(scores, score)
		}
	}

	slices.Sort(scores)
	fmt.Println("Part 2", scores[len(scores)/2])
}
