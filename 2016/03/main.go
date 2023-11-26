package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"strings"
)

func main() {
	input, err := download.ReadInput(2016, 3)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	lines := conv.SplitNewline(input)

	valid := 0
	for _, line := range lines {
		splitted := strings.Fields(line)
		n1 := conv.MustAtoi(splitted[0])
		n2 := conv.MustAtoi(splitted[1])
		n3 := conv.MustAtoi(splitted[2])

		if n1+n2 > n3 && n1+n3 > n2 && n2+n3 > n1 {
			valid++
		}
	}
	fmt.Println(valid)
}

func part2(input string) {
	lines := conv.SplitNewline(input)

	valid := 0
	for i := 0; i < len(lines); i += 3 {
		line1 := strings.Fields(lines[i])
		line2 := strings.Fields(lines[i+1])
		line3 := strings.Fields(lines[i+2])
		for c := 0; c < 3; c++ {
			n1 := conv.MustAtoi(line1[c])
			n2 := conv.MustAtoi(line2[c])
			n3 := conv.MustAtoi(line3[c])

			if n1+n2 > n3 && n1+n3 > n2 && n2+n3 > n1 {
				valid++
			}
		}
	}

	fmt.Println(valid)
}
