package main

import (
	"aoc/internal/container"
	"aoc/internal/download"
	"fmt"
	"log"
	"slices"
)

type state struct {
	elevator   int
	generators []int
	microchips []int
	steps      int
}

func (s state) isValid() bool {
	for i, mFloor := range s.microchips {
		if mFloor == s.generators[i] {
			continue
		}
		if slices.Contains(s.generators, mFloor) {
			return false
		}
	}
	return true
}

func (s state) isFinal() bool {
	topFloor := 4
	if s.elevator != topFloor {
		return false
	}
	for _, floor := range s.generators {
		if floor != topFloor {
			return false
		}
	}
	for _, floor := range s.microchips {
		if floor != topFloor {
			return false
		}
	}
	return true
}

func (s state) hash() string {
	return fmt.Sprintf("%d:%v:%v", s.elevator, s.generators, s.microchips)
}

func getNextStates(current state) []state {
	var nextStates []state
	var possibleFloors []int

	if current.elevator < 4 {
		possibleFloors = append(possibleFloors, current.elevator+1)
	}
	if current.elevator > 1 {
		possibleFloors = append(possibleFloors, current.elevator-1)
	}

	var items []struct {
		isGenerator bool
		index       int
	}

	for i, floor := range current.generators {
		if floor == current.elevator {
			items = append(items, struct {
				isGenerator bool
				index       int
			}{true, i})
		}
	}

	for i, floor := range current.microchips {
		if floor == current.elevator {
			items = append(items, struct {
				isGenerator bool
				index       int
			}{false, i})
		}
	}

	for _, newFloor := range possibleFloors {
		for i := range items {
			newState := state{
				elevator:   newFloor,
				generators: make([]int, len(current.generators)),
				microchips: make([]int, len(current.microchips)),
			}
			copy(newState.generators, current.generators)
			copy(newState.microchips, current.microchips)

			if items[i].isGenerator {
				newState.generators[items[i].index] = newFloor
			} else {
				newState.microchips[items[i].index] = newFloor
			}

			if newState.isValid() {
				nextStates = append(nextStates, newState)
			}
		}

		for i := range items {
			for j := i + 1; j < len(items); j++ {
				newState := state{
					elevator:   newFloor,
					generators: make([]int, len(current.generators)),
					microchips: make([]int, len(current.microchips)),
				}
				copy(newState.generators, current.generators)
				copy(newState.microchips, current.microchips)

				if items[i].isGenerator {
					newState.generators[items[i].index] = newFloor
				} else {
					newState.microchips[items[i].index] = newFloor
				}

				if items[j].isGenerator {
					newState.generators[items[j].index] = newFloor
				} else {
					newState.microchips[items[j].index] = newFloor
				}

				if newState.isValid() {
					nextStates = append(nextStates, newState)
				}
			}
		}
	}

	return nextStates
}

func findMinSteps(initial state) int {
	queue := container.NewPriorityQueue[state]()
	seen := container.NewSet[string]()

	initial.steps = 0
	queue.Push(initial, 0)
	seen.Add(initial.hash())

	for !queue.IsEmpty() {
		current := queue.Pop()

		if current.isFinal() {
			return current.steps
		}

		for _, next := range getNextStates(current) {
			next.steps = current.steps + 1
			hash := next.hash()
			if !seen.Contains(hash) {
				seen.Add(hash)
				queue.Push(next, next.steps)
			}
		}
	}

	return -1
}

func main() {
	_, err := download.ReadInput(2016, 11)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}
	part1()
	part2()
}

func part1() {
	initialState := state{
		elevator:   1,
		generators: []int{1, 1, 2, 2, 2}, // SG, PG, TG, RG, CG
		microchips: []int{1, 1, 3, 2, 2}, // SM, PM, TM, RM, CM
		steps:      0,
	}

	steps := findMinSteps(initialState)
	fmt.Println("Part 1", steps)
}

func part2() {
	initialState := state{
		elevator:   1,
		generators: []int{1, 1, 2, 2, 2, 1, 1}, // SG, PG, TG, RG, CG, EG, DG
		microchips: []int{1, 1, 3, 2, 2, 1, 1}, // SM, PM, TM, RM, CM, EM, DM
		steps:      0,
	}

	steps := findMinSteps(initialState)
	fmt.Println("Part 2", steps)
}
