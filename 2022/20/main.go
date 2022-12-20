package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
)

func main() {
	inputFile := "./2022/20/input.txt"
	input, err := download.ReadInput(inputFile, 2022, 20)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)

}

type number struct {
	value int
	index int
}

func part1(input string) {
	numberStr := conv.SplitNewline(input)
	var numbers = make([]number, len(numberStr))
	for ix, n := range numberStr {
		numbers[ix] = number{value: conv.MustAtoi(n), index: ix}
	}
	numbers = mix(numbers)
	printResult(numbers)
}

func part2(input string) {
	decriptionKey := 811589153
	numberStr := conv.SplitNewline(input)
	var numbers = make([]number, len(numberStr))
	for ix, n := range numberStr {
		numbers[ix] = number{value: conv.MustAtoi(n) * decriptionKey, index: ix}
	}

	for r := 0; r < 10; r++ {
		numbers = mix(numbers)
	}
	printResult(numbers)
}

func mix(numbers []number) []number {
	for ix := 0; ix < len(numbers); ix++ {
		currentIx := findIndex(numbers, ix)
		n := numbers[currentIx]

		numbers = append(numbers[:currentIx], numbers[currentIx+1:]...)

		insertIndex := (currentIx + n.value) % len(numbers)
		if n.value < 0 && insertIndex == 0 {
			insertIndex = len(numbers)
		}
		if insertIndex < 0 {
			insertIndex += len(numbers)
		}
		numbers = append(numbers, number{})
		copy(numbers[insertIndex+1:], numbers[insertIndex:])
		numbers[insertIndex] = n
	}
	return numbers
}

func printResult(numbers []number) {
	target := []int{1000, 2000, 3000}
	zeroIndex := 0
	for ix, n := range numbers {
		if n.value == 0 {
			zeroIndex = ix
			break
		}
	}
	sum := 0
	for _, t := range target {
		ix := (zeroIndex + t) % len(numbers)
		sum += numbers[ix].value
	}
	fmt.Println(sum)
}

func findIndex(numbers []number, ix int) int {
	for i, n := range numbers {
		if n.index == ix {
			return i
		}
	}
	return -1
}
