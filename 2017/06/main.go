package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"strings"
)

func main() {
	input, err := download.ReadInput(2017, 6)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}
	input = strings.TrimSuffix(input, "\n")
	banksSlice := conv.ToIntSlice(strings.Split(input, "\t"))
	var banks [16]int
	copy(banks[:], banksSlice)
	part1(banks)
	part2(banks)
}

func part1(banks [16]int) {
	seen := make(map[[16]int]bool)
	seen[banks] = true
	cycle := 0
	for {
		banks = redistribute(banks)
		cycle++
		if seen[banks] {
			break
		}
		seen[banks] = true
	}
	fmt.Println(cycle)
}

func part2(banks [16]int) {
	seen := make(map[[16]int]int)
	seen[banks] = 0
	cycle := 0
	for {
		banks = redistribute(banks)
		cycle++
		if seen[banks] > 0 {
			break
		}
		seen[banks] = cycle
	}
	fmt.Println(cycle - seen[banks])
}

func redistribute(banks [16]int) [16]int {
	mostBlockIndex := 0
	for i := range banks {
		if banks[i] > banks[mostBlockIndex] {
			mostBlockIndex = i
		}
	}
	redist := banks[mostBlockIndex]
	banks[mostBlockIndex] = 0
	for i := 1; i <= redist; i++ {
		banks[(mostBlockIndex+i)%len(banks)]++
	}
	return banks
}
