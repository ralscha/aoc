package main

import (
	"aoc/internal/download"
	"fmt"
	"log"
	"strings"
)

type rule struct {
	opponent string
	win      string
	draw     string
	lose     string
}

var rules = []rule{
	{opponent: "A", win: "Y", draw: "X", lose: "Z"},
	{opponent: "B", win: "Z", draw: "Y", lose: "X"},
	{opponent: "C", win: "X", draw: "Z", lose: "Y"},
}

func main() {
	input, err := download.ReadInput(2022, 2)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	totalScore := 0
	for line := range strings.SplitSeq(input, "\n") {
		in := strings.Fields(line)
		if len(in) == 2 {
			totalScore += score(in[0], in[1])
		}
	}
	fmt.Println(totalScore)
}

func part2(input string) {
	totalScore := 0
	for line := range strings.SplitSeq(input, "\n") {
		in := strings.Fields(line)
		if len(in) == 2 {
			opponent := in[0]
			expected := in[1]
			myMove := getMyMove(opponent, expected)
			totalScore += score(opponent, myMove)
		}
	}
	fmt.Println(totalScore)
}

// X means you need to lose, Y means you need to end the round in a draw, and Z means you need to win
func getMyMove(o, e string) string {
	for _, r := range rules {
		if r.opponent == o {
			if e == "Y" {
				return r.draw
			} else if e == "X" {
				return r.lose
			} else if e == "Z" {
				return r.win
			}
		}
	}
	return ""
}

func score(o, m string) int {
	// 1 for Rock, 2 for Paper, and 3 for Scissors
	score := 0
	if m == "X" {
		score += 1
	} else if m == "Y" {
		score += 2
	} else if m == "Z" {
		score += 3
	}

	for _, rule := range rules {
		if rule.opponent == o {
			if m == rule.win {
				score += 6
			} else if m == rule.draw {
				score += 3
			}
			break
		}
	}
	return score
}
