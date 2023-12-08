package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"container/ring"
	"fmt"
	"log"
	"slices"
	"strings"
)

func main() {
	input, err := download.ReadInput(2018, 9)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	splitted := strings.Fields(input)
	noOfPlayers := conv.MustAtoi(splitted[0])
	lastMarble := conv.MustAtoi(splitted[6])

	part1and2(noOfPlayers, lastMarble)
	part1and2(noOfPlayers, lastMarble*100)
}

func part1and2(noOfPlayers, lastMarble int) {
	scores := make([]int, noOfPlayers)
	circle := ring.New(1)
	circle.Value = 0

	for marble := 1; marble <= lastMarble; marble++ {
		if marble%23 == 0 {
			player := (marble - 1) % noOfPlayers
			scores[player] += marble
			circle = circle.Move(-8)
			removedMarble := circle.Unlink(1).Value.(int)
			scores[player] += removedMarble
			circle = circle.Move(1)
		} else {
			circle = circle.Next()
			newMarble := ring.New(1)
			newMarble.Value = marble
			circle.Link(newMarble)
			circle = newMarble
		}
	}

	winningScore := slices.Max(scores)
	fmt.Println(winningScore)
}
