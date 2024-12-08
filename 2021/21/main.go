package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"strings"
)

type player struct {
	position int
	score    int
}

type gameState struct {
	players [2]player
	turn    int
}

type diceRoll struct {
	sum   int
	count int
}

// Pre-calculated possible sums of three dice rolls in quantum game
var quantumRolls = []diceRoll{
	{3, 1}, // 1,1,1
	{4, 3}, // 1,1,2 1,2,1 2,1,1
	{5, 6}, // 1,1,3 1,2,2 1,3,1 2,1,2 2,2,1 3,1,1
	{6, 7}, // 1,2,3 1,3,2 2,1,3 2,2,2 2,3,1 3,1,2 3,2,1
	{7, 6}, // 1,3,3 2,2,3 2,3,2 3,1,3 3,2,2 3,3,1
	{8, 3}, // 2,3,3 3,2,3 3,3,2
	{9, 1}, // 3,3,3
}

func main() {
	input, err := download.ReadInput(2021, 21)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	p1, p2 := parseStartPositions(input)
	game := gameState{
		players: [2]player{{position: p1}, {position: p2}},
	}

	rolls := 0
	die := 1

	for game.players[0].score < 1000 && game.players[1].score < 1000 {
		// Roll three times and sum
		sum := 0
		for i := 0; i < 3; i++ {
			sum += die
			die = die%100 + 1
		}
		rolls += 3

		// Move current player
		currentPlayer := &game.players[game.turn]
		currentPlayer.position = ((currentPlayer.position + sum - 1) % 10) + 1
		currentPlayer.score += currentPlayer.position

		// Switch turns
		game.turn = 1 - game.turn
	}

	losingScore := min(game.players[0].score, game.players[1].score)
	fmt.Println("Part 1", losingScore*rolls)
}

func part2(input string) {
	p1, p2 := parseStartPositions(input)
	wins := playQuantumGame(player{position: p1}, player{position: p2}, true)
	fmt.Println("Part 2", max(wins[0], wins[1]))
}

func parseStartPositions(input string) (int, int) {
	lines := conv.SplitNewline(input)
	p1 := conv.MustAtoi(strings.Fields(lines[0])[4])
	p2 := conv.MustAtoi(strings.Fields(lines[1])[4])
	return p1, p2
}

func playQuantumGame(p1, p2 player, p1Turn bool) [2]int64 {
	// Base case: if either player has won
	if p2.score >= 21 {
		return [2]int64{0, 1}
	}
	if p1.score >= 21 {
		return [2]int64{1, 0}
	}

	// Track total wins for each player across all universes
	totalWins := [2]int64{0, 0}

	// Try each possible roll sum
	for _, roll := range quantumRolls {
		var nextP1, nextP2 player
		if p1Turn {
			// Move p1
			nextP1.position = ((p1.position + roll.sum - 1) % 10) + 1
			nextP1.score = p1.score + nextP1.position
			nextP2 = p2
		} else {
			// Move p2
			nextP1 = p1
			nextP2.position = ((p2.position + roll.sum - 1) % 10) + 1
			nextP2.score = p2.score + nextP2.position
		}

		// Recursively play out this universe
		wins := playQuantumGame(nextP1, nextP2, !p1Turn)

		// Add wins from this universe (multiplied by number of ways to get this roll)
		totalWins[0] += wins[0] * int64(roll.count)
		totalWins[1] += wins[1] * int64(roll.count)
	}

	return totalWins
}
