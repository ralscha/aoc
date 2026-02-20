package main

import (
	"aoc/internal/download"
	"crypto/md5"
	"fmt"
	"log"
	"strings"
)

func main() {
	input, err := download.ReadInput(2015, 4)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1and2(input, 5)
	part1and2(input, 6)

}

func part1and2(input string, numberOfZeros int) {
	var zeros strings.Builder
	for range numberOfZeros {
		zeros.WriteString("0")
	}
	input = input[:len(input)-1]

	count := 0
	for {
		c := input + fmt.Sprintf("%d", count)
		hash := fmt.Sprintf("%x", md5.Sum([]byte(c)))
		if hash[0:numberOfZeros] == zeros.String() {
			fmt.Println(count)
			break
		}
		count++
	}
}
