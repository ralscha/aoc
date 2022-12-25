package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
)

func main() {
	inputFile := "./2022/25/input.txt"
	input, err := download.ReadInput(inputFile, 2022, 25)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
}

func part1(input string) {
	lines := conv.SplitNewline(input)

	totalFuel := 0
	for _, line := range lines {
		decimalValue := snafuToDecimal(line)
		totalFuel += decimalValue
	}
	fmt.Println(decimalToSnafu(totalFuel))
}

func snafuToDecimal(snafu string) int {
	result := 0
	currentPlace := 1

	for i := len(snafu) - 1; i >= 0; i-- {
		digit := snafu[i]

		switch digit {
		case '1':
			result += currentPlace
		case '2':
			result += currentPlace * 2
		case '-':
			result -= currentPlace
		case '=':
			result -= currentPlace * 2
		}
		currentPlace *= 5
	}

	return result
}

func decimalToSnafu(decimal int) string {
	snafus := []string{"=", "-", "0", "1", "2"}
	d := decimal
	result := ""
	for d > 0 {
		pos := (d + 2) % 5
		d = (d + 2) / 5
		result = snafus[pos] + result
	}
	return result
}
