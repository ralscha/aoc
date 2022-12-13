package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/mathx"
	"fmt"
	"log"
	"strings"
)

func main() {
	inputFile := "./2015/13/input.txt"
	input, err := download.ReadInput(inputFile, 2015, 13)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

type person struct {
	neighbours map[string]int
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	persons := createPersons(lines)
	maxHappiness(persons)
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	persons := createPersons(lines)
	persons["me"] = &person{neighbours: make(map[string]int)}
	for name := range persons {
		if name != "me" {
			persons[name].neighbours["me"] = 0
			persons["me"].neighbours[name] = 0
		}
	}
	maxHappiness(persons)
}

func createPersons(lines []string) map[string]*person {
	persons := make(map[string]*person)
	for _, line := range lines {
		splitted := strings.Fields(line)
		name := splitted[0]
		gainLoseValue := conv.MustAtoi(splitted[3])
		if splitted[2] == "lose" {
			gainLoseValue = -gainLoseValue
		}
		neighbour := splitted[10]
		neighbour = neighbour[:len(neighbour)-1]
		if p, ok := persons[name]; ok {
			p.neighbours[neighbour] = gainLoseValue
		} else {
			persons[name] = &person{
				neighbours: map[string]int{neighbour: gainLoseValue},
			}
		}
	}
	return persons
}

func maxHappiness(persons map[string]*person) {
	names := make([]string, 0, len(persons))
	for name := range persons {
		names = append(names, name)
	}

	perms := mathx.Permutations(names)
	maxHappiness := 0
	for _, perm := range perms {
		happiness := 0
		for i, name := range perm {
			neighbour1 := perm[(i-1+len(perm))%len(perm)]
			neighbour2 := perm[(i+1)%len(perm)]
			happiness += persons[name].neighbours[neighbour1]
			happiness += persons[name].neighbours[neighbour2]
		}
		if happiness > maxHappiness {
			maxHappiness = happiness
		}
	}

	fmt.Println(maxHappiness)
}
