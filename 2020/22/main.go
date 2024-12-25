package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"strings"
)

func main() {
	input, err := download.ReadInput(2020, 22)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func parseDecks(input string) (*container.Queue[int], *container.Queue[int]) {
	parts := strings.Split(input, "\n\n")
	player1Lines := conv.SplitNewline(parts[0])[1:]
	player2Lines := conv.SplitNewline(parts[1])[1:]

	player1Deck := container.NewQueue[int]()
	for _, line := range player1Lines {
		player1Deck.Push(conv.MustAtoi(line))
	}

	player2Deck := container.NewQueue[int]()
	for _, line := range player2Lines {
		player2Deck.Push(conv.MustAtoi(line))
	}

	return player1Deck, player2Deck
}

func calculateScore(deck *container.Queue[int]) int {
	score := 0
	size := deck.Len()
	for i := range size {
		card := deck.Pop()
		score += card * (size - i)
	}
	return score
}

func part1(input string) {
	player1Deck, player2Deck := parseDecks(input)

	for player1Deck.Len() > 0 && player2Deck.Len() > 0 {
		card1 := player1Deck.Pop()
		card2 := player2Deck.Pop()

		if card1 > card2 {
			player1Deck.Push(card1)
			player1Deck.Push(card2)
		} else {
			player2Deck.Push(card2)
			player2Deck.Push(card1)
		}
	}

	var winnerScore int
	if player1Deck.Len() > 0 {
		winnerScore = calculateScore(player1Deck)
	} else {
		winnerScore = calculateScore(player2Deck)
	}

	fmt.Println("Part 1", winnerScore)
}

func part2(input string) {
	player1Deck, player2Deck := parseDecks(input)
	winner := playRecursiveCombat(player1Deck, player2Deck)

	var winnerScore int
	if winner == 1 {
		winnerScore = calculateScore(player1Deck)
	} else {
		winnerScore = calculateScore(player2Deck)
	}

	fmt.Println("Part 2", winnerScore)
}

func playRecursiveCombat(player1Deck *container.Queue[int], player2Deck *container.Queue[int]) int {
	seen := container.NewSet[string]()

	for player1Deck.Len() > 0 && player2Deck.Len() > 0 {
		state := fmt.Sprintf("1:%v 2:%v", player1Deck.Values(), player2Deck.Values())
		if seen.Contains(state) {
			return 1
		}
		seen.Add(state)

		card1 := player1Deck.Pop()
		card2 := player2Deck.Pop()

		player1WinsRound := false
		if player1Deck.Len() >= card1 && player2Deck.Len() >= card2 {
			subGamePlayer1Deck := player1Deck.Values()[:card1]
			subGamePlayer2Deck := player2Deck.Values()[:card2]

			subGameWinner := playRecursiveCombat(container.NewQueueFromSlice(subGamePlayer1Deck), container.NewQueueFromSlice(subGamePlayer2Deck))
			player1WinsRound = subGameWinner == 1
		} else {
			player1WinsRound = card1 > card2
		}

		if player1WinsRound {
			player1Deck.Push(card1)
			player1Deck.Push(card2)
		} else {
			player2Deck.Push(card2)
			player2Deck.Push(card1)
		}
	}

	if player1Deck.Len() > 0 {
		return 1
	}
	return 2
}
