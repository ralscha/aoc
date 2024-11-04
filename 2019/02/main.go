package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
)

func main() {
	input, err := download.ReadInput(2019, 2)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	numbers := conv.ToIntSliceComma(input)
	numbers[1] = 12
	numbers[2] = 2

	for i := 0; i < len(numbers); i += 4 {
		if numbers[i] == 99 {
			break
		}
		if numbers[i] == 1 {
			numbers[numbers[i+3]] = numbers[numbers[i+1]] + numbers[numbers[i+2]]
		}
		if numbers[i] == 2 {
			numbers[numbers[i+3]] = numbers[numbers[i+1]] * numbers[numbers[i+2]]
		}
	}

	fmt.Println("Result 1:", numbers[0])
}

func part2(input string) {
	originalNumbers := conv.ToIntSliceComma(input)

	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			numbers := make([]int, len(originalNumbers))
			copy(numbers, originalNumbers)
			numbers[1] = noun
			numbers[2] = verb

			for i := 0; i < len(numbers); i += 4 {
				if numbers[i] == 99 {
					break
				}
				if numbers[i] == 1 {
					numbers[numbers[i+3]] = numbers[numbers[i+1]] + numbers[numbers[i+2]]
				}
				if numbers[i] == 2 {
					numbers[numbers[i+3]] = numbers[numbers[i+1]] * numbers[numbers[i+2]]
				}
			}

			if numbers[0] == 19690720 {
				fmt.Println("Result 2:", 100*noun+verb)
				return
			}
		}
	}
}
