package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
)

func main() {
	input, err := download.ReadInput(2019, 5)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	numbers := conv.ToIntSliceComma(input)

	currentPos := 0
	for {
		opcode := numbers[currentPos] % 100
		if opcode == 99 {
			break
		}
		parameterModeC := (numbers[currentPos] / 100) % 10
		parameterModeB := (numbers[currentPos] / 1000) % 10
		parameterModeA := (numbers[currentPos] / 10000) % 10

		switch opcode {
		case 1:
			var a, b int
			if parameterModeC == 0 {
				a = numbers[numbers[currentPos+1]]
			} else {
				a = numbers[currentPos+1]
			}
			if parameterModeB == 0 {
				b = numbers[numbers[currentPos+2]]
			} else {
				b = numbers[currentPos+2]
			}
			result := a + b
			if parameterModeA == 0 {
				numbers[numbers[currentPos+3]] = result
			} else {
				numbers[currentPos+3] = result
			}
			currentPos += 4
		case 2:
			var a, b int
			if parameterModeC == 0 {
				a = numbers[numbers[currentPos+1]]
			} else {
				a = numbers[currentPos+1]
			}
			if parameterModeB == 0 {
				b = numbers[numbers[currentPos+2]]
			} else {
				b = numbers[currentPos+2]
			}
			result := a * b
			if parameterModeA == 0 {
				numbers[numbers[currentPos+3]] = result
			} else {
				numbers[currentPos+3] = result
			}
			currentPos += 4
		case 3:
			numbers[numbers[currentPos+1]] = 1
			currentPos += 2
		case 4:
			if parameterModeC == 0 {
				fmt.Println(numbers[numbers[currentPos+1]])
			} else {
				fmt.Println(numbers[currentPos+1])
			}
			currentPos += 2
		default:
			log.Fatalf("unknown opcode: %d", opcode)
		}
	}

}

func part2(input string) {
	numbers := conv.ToIntSliceComma(input)

	currentPos := 0
	for {
		opcode := numbers[currentPos] % 100
		if opcode == 99 {
			break
		}
		parameterModeC := (numbers[currentPos] / 100) % 10
		parameterModeB := (numbers[currentPos] / 1000) % 10
		parameterModeA := (numbers[currentPos] / 10000) % 10

		switch opcode {
		case 1:
			var a, b int
			if parameterModeC == 0 {
				a = numbers[numbers[currentPos+1]]
			} else {
				a = numbers[currentPos+1]
			}
			if parameterModeB == 0 {
				b = numbers[numbers[currentPos+2]]
			} else {
				b = numbers[currentPos+2]
			}
			result := a + b
			if parameterModeA == 0 {
				numbers[numbers[currentPos+3]] = result
			} else {
				numbers[currentPos+3] = result
			}
			currentPos += 4
		case 2:
			var a, b int
			if parameterModeC == 0 {
				a = numbers[numbers[currentPos+1]]
			} else {
				a = numbers[currentPos+1]
			}
			if parameterModeB == 0 {
				b = numbers[numbers[currentPos+2]]
			} else {
				b = numbers[currentPos+2]
			}
			result := a * b
			if parameterModeA == 0 {
				numbers[numbers[currentPos+3]] = result
			} else {
				numbers[currentPos+3] = result
			}
			currentPos += 4
		case 3:
			numbers[numbers[currentPos+1]] = 5
			currentPos += 2
		case 4:
			if parameterModeC == 0 {
				fmt.Println(numbers[numbers[currentPos+1]])
			} else {
				fmt.Println(numbers[currentPos+1])
			}
			currentPos += 2
		case 5:
			var a, b int
			if parameterModeC == 0 {
				a = numbers[numbers[currentPos+1]]
			} else {
				a = numbers[currentPos+1]
			}
			if parameterModeB == 0 {
				b = numbers[numbers[currentPos+2]]
			} else {
				b = numbers[currentPos+2]
			}
			if a != 0 {
				currentPos = b
			} else {
				currentPos += 3
			}
		case 6:
			var a, b int
			if parameterModeC == 0 {
				a = numbers[numbers[currentPos+1]]
			} else {
				a = numbers[currentPos+1]
			}
			if parameterModeB == 0 {
				b = numbers[numbers[currentPos+2]]
			} else {
				b = numbers[currentPos+2]
			}
			if a == 0 {
				currentPos = b
			} else {
				currentPos += 3
			}
		case 7:
			var a, b int
			if parameterModeC == 0 {
				a = numbers[numbers[currentPos+1]]
			} else {
				a = numbers[currentPos+1]
			}
			if parameterModeB == 0 {
				b = numbers[numbers[currentPos+2]]
			} else {
				b = numbers[currentPos+2]
			}
			if a < b {
				numbers[numbers[currentPos+3]] = 1
			} else {
				numbers[numbers[currentPos+3]] = 0
			}
			currentPos += 4
		case 8:
			var a, b int
			if parameterModeC == 0 {
				a = numbers[numbers[currentPos+1]]
			} else {
				a = numbers[currentPos+1]
			}
			if parameterModeB == 0 {
				b = numbers[numbers[currentPos+2]]
			} else {
				b = numbers[currentPos+2]
			}
			if a == b {
				numbers[numbers[currentPos+3]] = 1
			} else {
				numbers[numbers[currentPos+3]] = 0
			}
			currentPos += 4

		default:
			log.Fatalf("unknown opcode: %d", opcode)
		}
	}

}
