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

type Range struct {
	start, end, count int
}

func partI() {
	input, err := os.ReadFile("input1")
	if err != nil {
		log.Fatalf("reading %s failed: %v", "input1", err)
	}

	lines := conv.SplitNewline(string(input))
	numbers := conv.ToIntSlice(lines)

	wheelSize := len(numbers) + 1
	wheel := make([]int, wheelSize)
	wheel[0] = 1

	left := wheelSize - 1
	right := 1

	for i, n := range numbers {
		if i%2 == 0 {
			wheel[right] = n
			right++
		} else {
			wheel[left] = n
			left--
		}
	}

	targetPos := 2025 % wheelSize
	fmt.Println(wheel[targetPos])
}

func partII() {
	input, err := os.ReadFile("input2")
	if err != nil {
		log.Fatalf("reading %s failed: %v", "input2", err)
	}

	lines := conv.SplitNewline(string(input))

	var rightSide []int
	var leftSide []int

	idx := 0
	for _, line := range lines {
		if line == "" {
			continue
		}

		parts := strings.Split(line, "-")
		start := conv.MustAtoi(strings.TrimSpace(parts[0]))
		end := conv.MustAtoi(strings.TrimSpace(parts[1]))

		var nums []int
		for i := start; i <= end; i++ {
			nums = append(nums, i)
		}

		if idx%2 == 0 {
			rightSide = append(rightSide, nums...)
		} else {
			leftSide = append(leftSide, nums...)
		}
		idx++
	}

	for i, j := 0, len(leftSide)-1; i < j; i, j = i+1, j-1 {
		leftSide[i], leftSide[j] = leftSide[j], leftSide[i]
	}

	wheel := []int{1}
	wheel = append(wheel, rightSide...)
	wheel = append(wheel, leftSide...)

	targetPos := 20252025 % len(wheel)
	fmt.Println(wheel[targetPos])
}

func partIII() {
	input, err := os.ReadFile("input3")
	if err != nil {
		log.Fatalf("reading %s failed: %v", "input3", err)
	}

	lines := conv.SplitNewline(string(input))

	var rightRanges []Range
	var leftRanges []Range

	idx := 0
	for _, line := range lines {
		if line == "" {
			continue
		}

		parts := strings.Split(line, "-")
		start := conv.MustAtoi(strings.TrimSpace(parts[0]))
		end := conv.MustAtoi(strings.TrimSpace(parts[1]))
		count := end - start + 1

		r := Range{start: start, end: end, count: count}

		if idx%2 == 0 {
			rightRanges = append(rightRanges, r)
		} else {
			leftRanges = append(leftRanges, r)
		}
		idx++
	}

	var rightLen, leftLen int
	for _, r := range rightRanges {
		rightLen += r.count
	}
	for _, r := range leftRanges {
		leftLen += r.count
	}

	totalLen := 1 + rightLen + leftLen
	targetPos := 202520252025 % totalLen

	targetPos--

	if targetPos < rightLen {
		for _, r := range rightRanges {
			if targetPos < r.count {
				fmt.Println(r.start + targetPos)
				return
			}
			targetPos -= r.count
		}
	} else {
		targetPos -= rightLen
		for i := len(leftRanges) - 1; i >= 0; i-- {
			r := leftRanges[i]
			if targetPos < r.count {
				fmt.Println(r.end - targetPos)
				return
			}
			targetPos -= r.count
		}
	}
}
