package main

import (
	"aoc/internal/download"
	"bytes"
	"fmt"
	"log"
	"strconv"
)

func main() {
	inputFile := "./2015/10/input.txt"
	input, err := download.ReadInput(inputFile, 2015, 10)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1and2(input, 40)
	part1and2(input, 50)
}

func part1and2(input string, iterations int) {
	str := input[:len(input)-1]
	for i := 0; i < iterations; i++ {
		str = lookAndSay(str)
	}
	fmt.Println(len(str))
}

func lookAndSay(s string) string {
	var result bytes.Buffer
	i := 0
	for i < len(s) {
		count := 1
		for i+1 < len(s) && s[i] == s[i+1] {
			count++
			i++
		}
		result.WriteString(strconv.Itoa(count))
		result.WriteByte(s[i])
		i++
	}
	return result.String()
}
