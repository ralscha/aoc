package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"strings"
)

func main() {
	input, err := download.ReadInput(2021, 21)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	splitted1 := strings.Fields(lines[0])
	splitted2 := strings.Fields(lines[1])
	var playerPos [2]int
	playerPos[0] = conv.MustAtoi(splitted1[4])
	playerPos[1] = conv.MustAtoi(splitted2[4])
	var score [2]int

	rolls := 0
	currentPlayer := 0
	for score[0] < 1000 && score[1] < 1000 {
		round := 3*rolls + 6
		rolls += 3
		newPos := playerPos[currentPlayer] + round
		newPos = newPos % 10
		if newPos == 0 {
			newPos = 10
		}
		playerPos[currentPlayer] = newPos
		score[currentPlayer] += playerPos[currentPlayer]
		currentPlayer = (currentPlayer + 1) % 2
	}
	fmt.Println(min(score[0], score[1]) * rolls)
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	splitted1 := strings.Fields(lines[0])
	splitted2 := strings.Fields(lines[1])
	var playerPos [2]int
	playerPos[0] = conv.MustAtoi(splitted1[4])
	playerPos[1] = conv.MustAtoi(splitted2[4])

	score1, score2 := play2(playerPos[0], playerPos[1], 0, 0)

	fmt.Println(max(score1, score2))
}

type sum struct {
	move int
	n    int
}

func play2(pos1, pos2, score1, score2 int) (int, int) {
	if score2 >= 21 {
		return 0, 1
	}

	wins1, wins2 := 0, 0
	for _, p := range []sum{{3, 1}, {4, 3}, {5, 6}, {6, 7}, {7, 6}, {8, 3}, {9, 1}} {
		move, n := p.move, p.n
		pos1_ := (pos1 + move) % 10
		if pos1_ == 0 {
			pos1_ = 10
		}
		w2, w1 := play2(pos2, pos1_, score2, score1+pos1_)
		wins1, wins2 = wins1+n*w1, wins2+n*w2
	}
	return wins1, wins2
}
