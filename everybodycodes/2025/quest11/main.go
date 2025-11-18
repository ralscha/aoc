package main

import (
	"aoc/internal/conv"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	partI()
	partII()
	partIII()
}

func moveBlocksLeft(columns []int) bool {
	moved := false
	for i := 0; i < len(columns)-1; i++ {
		if columns[i] > columns[i+1] {
			columns[i]--
			columns[i+1]++
			moved = true
		}
	}
	return moved
}

func moveBlocksRight(columns []int) bool {
	moved := false
	for i := 0; i < len(columns)-1; i++ {
		if columns[i] < columns[i+1] {
			columns[i+1]--
			columns[i]++
			moved = true
		}
	}
	return moved
}

func runRounds(columns []int, moveFunc func([]int) bool, maxRounds int, startRound int) int {
	roundCount := startRound
	for roundCount < maxRounds {
		if moveFunc(columns) {
			roundCount++
		} else {
			break
		}
	}
	return roundCount
}

func calculateChecksum(columns []int) int {
	checksum := 0
	for i, count := range columns {
		checksum += (i + 1) * count
	}
	return checksum
}

func isBalanced(cols []int) bool {
	if len(cols) == 0 {
		return true
	}
	first := cols[0]
	for _, v := range cols {
		if v != first {
			return false
		}
	}
	return true
}

func partI() {
	input := `2
2
9
19
18
10`

	lines := conv.SplitNewline(input)
	columns := conv.ToIntSlice(lines)
	roundCount := runRounds(columns, moveBlocksLeft, 10, 0)
	roundCount = runRounds(columns, moveBlocksRight, 10, roundCount)
	checksum := calculateChecksum(columns)
	fmt.Println(checksum)
}

func partII() {
	input, err := os.ReadFile("input2")
	if err != nil {
		log.Fatalf("reading input2 failed: %v", err)
	}

	lines := conv.SplitNewline(string(input))
	for line := range lines {
		lines[line] = strings.TrimSpace(lines[line])
	}
	columns := conv.ToIntSlice(lines)

	roundCount := 0

	for !isBalanced(columns) {
		if moveBlocksLeft(columns) {
			roundCount++
		} else {
			break
		}
	}

	for !isBalanced(columns) {
		if moveBlocksRight(columns) {
			roundCount++
		} else {
			break
		}
	}

	fmt.Println(roundCount)
}

func partIII() {
	input, err := os.ReadFile("input3")
	if err != nil {
		log.Fatalf("reading input3 failed: %v", err)
	}

	lines := conv.SplitNewline(string(input))
	for line := range lines {
		lines[line] = strings.TrimSpace(lines[line])
	}
	columns := conv.ToIntSlice(lines)

	total := 0
	for _, col := range columns {
		total += col
	}
	target := total / len(columns)

	moves := 0
	for _, col := range columns {
		if col < target {
			moves += target - col
		} else {
			break
		}
	}

	fmt.Println(moves)
}
