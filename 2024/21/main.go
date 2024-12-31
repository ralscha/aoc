package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/mathx"
	"fmt"
	"log"
	"strings"
)

var POSITIONS = map[string][]int{
	"7": {0, 0},
	"8": {0, 1},
	"9": {0, 2},
	"4": {1, 0},
	"5": {1, 1},
	"6": {1, 2},
	"1": {2, 0},
	"2": {2, 1},
	"3": {2, 2},
	"0": {3, 1},
	"A": {3, 2},
	"^": {0, 1},
	"a": {0, 2},
	"<": {1, 0},
	"v": {1, 1},
	">": {1, 2},
}

var DIRECTIONS = map[string][]int{
	"^": {-1, 0},
	"v": {1, 0},
	"<": {0, -1},
	">": {0, 1},
}

func subtract(a, b []int) []int {
	return []int{a[0] - b[0], a[1] - b[1]}
}

func add(a, b []int) []int {
	return []int{a[0] + b[0], a[1] + b[1]}
}

func equal(a, b []int) bool {
	return len(a) == len(b) && a[0] == b[0] && a[1] == b[1]
}

func seeToMoveSet(start, finish []int, avoid []int) []string {
	delta := subtract(finish, start)
	var sb strings.Builder
	dX, dY := delta[0], delta[1]

	if dX < 0 {
		sb.WriteString(strings.Repeat("^", mathx.Abs(dX)))
	} else {
		sb.WriteString(strings.Repeat("v", dX))
	}
	if dY < 0 {
		sb.WriteString(strings.Repeat("<", mathx.Abs(dY)))
	} else {
		sb.WriteString(strings.Repeat(">", dY))
	}
	moveString := sb.String()

	var rv []string
	seen := container.NewSet[string]()
	moveRunes := []rune(moveString)
	perms := mathx.Permutations(moveRunes)

	for _, p := range perms {
		perm := string(p)
		if !seen.Contains(perm) {
			valid := true
			currentPos := start
			for _, moveChar := range perm {
				move := DIRECTIONS[string(moveChar)]
				nextPos := add(currentPos, move)
				if equal(nextPos, avoid) {
					valid = false
					break
				}
				currentPos = nextPos
			}
			if valid {
				rv = append(rv, perm+"a")
			}
			seen.Add(perm)
		}
	}

	if len(rv) == 0 {
		return []string{"a"}
	}
	return rv
}

type memoKey struct {
	s     string
	depth int
	lim   int
}

var memoization = make(map[memoKey]int)

func minLength(s string, lim int, depth int) int {
	key := memoKey{s, depth, lim}
	if val, ok := memoization[key]; ok {
		return val
	}

	avoid := []int{3, 0}
	if depth > 0 {
		avoid = []int{0, 0}
	}
	cur := POSITIONS["A"]
	if depth > 0 {
		cur = POSITIONS["a"]
	}
	length := 0

	for _, char := range s {
		nextCurrent := POSITIONS[string(char)]
		moveSet := seeToMoveSet(cur, nextCurrent, avoid)
		if depth == lim {
			length += len(moveSet[0])
		} else {
			minSubLength := -1
			for _, move := range moveSet {
				subLength := minLength(move, lim, depth+1)
				if minSubLength == -1 || subLength < minSubLength {
					minSubLength = subLength
				}
			}
			length += minSubLength
		}
		cur = nextCurrent
	}

	memoization[key] = length
	return length
}

func sumOfComplexities(codes []string, limit int) int {
	complexity := 0
	for _, code := range codes {
		length := minLength(code, limit, 0)
		numeric := conv.MustAtoi(code[:3])
		complexity += length * numeric
	}
	return complexity
}

func main() {
	input, err := download.ReadInput(2024, 21)
	if err != nil {
		log.Fatal(err)
	}
	part1(input)
	part2(input)

}

func part1(input string) {
	codes := conv.SplitNewline(input)
	fmt.Println("Part 1", sumOfComplexities(codes, 2))

}

func part2(input string) {
	codes := conv.SplitNewline(input)
	fmt.Println("Part 2", sumOfComplexities(codes, 25))
}
