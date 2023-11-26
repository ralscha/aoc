package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"math"
)

func main() {
	input, err := download.ReadInput(2015, 20)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	target := conv.MustAtoi(input[:len(input)-1])
	house := 1
	for {
		presents := 0
		for _, factor := range factors(house) {
			presents += factor * 10
		}
		if presents >= target {
			break
		}
		house++
	}

	fmt.Println(house)
}

func part2(input string) {
	target := conv.MustAtoi(input[:len(input)-1])
	house := 2
	stopAfter := 50
	for {
		presents := 0
		for _, factor := range factors(house) {
			if house/factor <= stopAfter {
				presents += factor * 11
			}
		}
		if presents >= target {
			break
		}
		house++
	}

	fmt.Println(house)
}

func factors(n int) []int {
	if n == 1 {
		return []int{1}
	}
	factors := []int{1, n}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			factors = append(factors, i)
			if i != n/i {
				factors = append(factors, n/i)
			}
		}
	}
	return factors
}
