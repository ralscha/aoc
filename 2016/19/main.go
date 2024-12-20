package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
)

func main() {
	input, err := download.ReadInput(2016, 19)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

type elf struct {
	position int
	presents int
}

func findWinningElf(numElves int) int {
	elves := container.NewQueue[*elf]()

	for i := 1; i <= numElves; i++ {
		elves.Push(&elf{
			position: i,
			presents: 1,
		})
	}

	for elves.Len() > 1 {
		currentElf := elves.Pop()
		targetElf := elves.Pop()
		currentElf.presents += targetElf.presents
		targetElf.presents = 0

		if currentElf.presents > 0 {
			elves.Push(currentElf)
		}
	}

	winner := elves.Pop()
	return winner.position
}

func findWinningElfPart2(numElves int) int {
	left := container.NewQueue[int]()
	right := container.NewQueue[int]()

	for i := 1; i <= numElves/2; i++ {
		left.Push(i)
	}
	for i := numElves/2 + 1; i <= numElves; i++ {
		right.Push(i)
	}

	for left.Len()+right.Len() > 1 {
		right.Pop()

		if left.Len()+right.Len() == 0 {
			break
		}

		val := left.Pop()
		right.Push(val)

		if (left.Len()+right.Len())%2 == 0 {
			val := right.Pop()
			left.Push(val)
		}
	}

	if right.Len() == 1 {
		return right.Pop()
	}
	return left.Pop()
}

func part1(input string) {
	numberOfElves := conv.MustAtoi(input)
	winnerPosition := findWinningElf(numberOfElves)
	fmt.Println("Part 1", winnerPosition)
}

func part2(input string) {
	numberOfElves := conv.MustAtoi(input)
	winnerPosition := findWinningElfPart2(numberOfElves)
	fmt.Println("Part 2", winnerPosition)
}
