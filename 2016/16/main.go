package main

import (
	"aoc/internal/download"
	"aoc/internal/stringutil"
	"fmt"
	"log"
)

func dragonCurve(a []byte) []byte {
	b := make([]byte, len(a))
	reversed := []byte(stringutil.Reverse(string(a)))
	for i := range len(reversed) {
		if reversed[i] == '0' {
			b[i] = '1'
		} else {
			b[i] = '0'
		}
	}
	res := make([]byte, len(a)+1+len(b))
	copy(res, a)
	res[len(a)] = '0'
	copy(res[len(a)+1:], b)
	return res
}

func checksum(data []byte) []byte {
	check := make([]byte, len(data)/2)
	for i := 0; i < len(data)-1; i += 2 {
		if data[i] == data[i+1] {
			check[i/2] = '1'
		} else {
			check[i/2] = '0'
		}
	}
	for len(check)%2 == 0 {
		nextCheck := make([]byte, len(check)/2)
		for i := 0; i < len(check)-1; i += 2 {
			if check[i] == check[i+1] {
				nextCheck[i/2] = '1'
			} else {
				nextCheck[i/2] = '0'
			}
		}
		check = nextCheck
	}
	return check
}

func main() {
	input, err := download.ReadInput(2016, 16)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func fillDisk(input string, diskSize int) string {
	data := []byte(input)
	for len(data) < diskSize {
		data = dragonCurve(data)
	}
	data = data[:diskSize]
	return string(checksum(data))
}

func part1(input string) {
	fmt.Println("Part 1:", fillDisk(input, 272))
}

func part2(input string) {
	fmt.Println("Part 2:", fillDisk(input, 35651584))
}
