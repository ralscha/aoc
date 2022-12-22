package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
)

func main() {
	inputFile := "./2016/02/input.txt"
	input, err := download.ReadInput(inputFile, 2016, 2)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	lines := conv.SplitNewline(input)

	num := 5
	for _, line := range lines {
		for _, c := range line {
			switch c {
			case 'U':
				num = up(num)
			case 'D':
				num = down(num)
			case 'L':
				num = left(num)
			case 'R':
				num = right(num)
			}
		}
		fmt.Print(num)
	}
	fmt.Println()
}

func up(num int) int {
	if num <= 3 {
		return num
	}
	return num - 3
}

func down(num int) int {
	if num == 7 || num == 8 || num == 9 {
		return num
	}
	return num + 3
}

func left(num int) int {
	if num == 1 || num == 4 || num == 7 {
		return num
	}
	return num - 1
}

func right(num int) int {
	if num == 3 || num == 6 || num == 9 {
		return num
	}
	return num + 1
}

func part2(input string) {
	lines := conv.SplitNewline(input)

	num := 5
	for _, line := range lines {
		for _, c := range line {
			switch c {
			case 'U':
				num = up2(num)
			case 'D':
				num = down2(num)
			case 'L':
				num = left2(num)
			case 'R':
				num = right2(num)
			}
		}
		if num > 9 {
			fmt.Printf("%c", num-10+'A')
		} else {
			fmt.Print(num)
		}
	}
	fmt.Println()
}

func up2(num int) int {
	switch num {
	case 3:
		return 1
	case 6:
		return 2
	case 7:
		return 3
	case 8:
		return 4
	case 10:
		return 6
	case 11:
		return 7
	case 12:
		return 8
	case 13:
		return 11
	}
	return num
}

func down2(num int) int {
	switch num {
	case 1:
		return 3
	case 2:
		return 6
	case 3:
		return 7
	case 4:
		return 8
	case 6:
		return 10
	case 7:
		return 11
	case 8:
		return 12
	case 11:
		return 13
	}
	return num
}

func left2(num int) int {
	switch num {
	case 3:
		return 2
	case 4:
		return 3
	case 6:
		return 5
	case 7:
		return 6
	case 8:
		return 7
	case 9:
		return 8
	case 11:
		return 10
	case 12:
		return 11
	}
	return num
}

func right2(num int) int {
	switch num {
	case 2:
		return 3
	case 3:
		return 4
	case 5:
		return 6
	case 6:
		return 7
	case 7:
		return 8
	case 8:
		return 9
	case 10:
		return 11
	case 11:
		return 12
	}
	return num
}
