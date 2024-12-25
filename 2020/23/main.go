package main

import (
	"aoc/internal/download"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	input, err := download.ReadInput(2020, 23)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	cups := parseInput(input)
	result := playGame(cups, 100)
	fmt.Println("Part 1", result)
}

func part2(input string) {
	cups := parseInput(input)
	cups = append(cups, make([]int, 1000000-len(cups))...)
	for i := len(input); i < 1000000; i++ {
		cups[i] = i + 1
	}
	result := playGame(cups, 10000000)
	fmt.Println("Part 2", result)
}

func parseInput(input string) []int {
	cups := make([]int, len(input))
	for i, r := range input {
		cups[i] = int(r - '0')
	}
	return cups
}

func playGame(cups []int, moves int) string {
	numCups := len(cups)
	next := make([]int, numCups+1)
	for i := 0; i < numCups-1; i++ {
		next[cups[i]] = cups[i+1]
	}
	next[cups[numCups-1]] = cups[0]

	currentCup := cups[0]
	for i := 0; i < moves; i++ {
		pickup1 := next[currentCup]
		pickup2 := next[pickup1]
		pickup3 := next[pickup2]

		next[currentCup] = next[pickup3]

		destinationCup := currentCup - 1
		if destinationCup < 1 {
			destinationCup = numCups
		}
		for destinationCup == pickup1 || destinationCup == pickup2 || destinationCup == pickup3 {
			destinationCup--
			if destinationCup < 1 {
				destinationCup = numCups
			}
		}

		next[pickup3] = next[destinationCup]
		next[destinationCup] = pickup1

		currentCup = next[currentCup]
	}

	if moves == 100 {
		result := strings.Builder{}
		cup := next[1]
		for i := 0; i < numCups-1; i++ {
			result.WriteString(strconv.Itoa(cup))
			cup = next[cup]
		}
		return result.String()
	}

	cup1 := next[1]
	cup2 := next[cup1]
	return fmt.Sprintf("%d", uint64(cup1)*uint64(cup2))
}
