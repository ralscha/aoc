package main

import (
	"aoc/internal/download"
	"crypto/md5"
	"fmt"
	"log"
	"strings"
)

func main() {
	input, err := download.ReadInput(2016, 14)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func getHash(salt string, index int, stretch int, cache map[string]string) string {
	key := fmt.Sprintf("%s%d-%d", salt, index, stretch)
	if hash, ok := cache[key]; ok {
		return hash
	}
	hash := fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s%d", salt, index))))
	for range stretch {
		hash = fmt.Sprintf("%x", md5.Sum([]byte(hash)))
	}
	cache[key] = hash
	return hash
}

func findRepeatingChar(hash string) (byte, bool) {
	for i := range len(hash) - 2 {
		if hash[i] == hash[i+1] && hash[i] == hash[i+2] {
			return hash[i], true
		}
	}
	return 0, false
}

func part1(input string) {
	salt := input
	index := 0
	found := 0
	cache := make(map[string]string)
	for found < 64 {
		hash := getHash(salt, index, 0, cache)
		if char, ok := findRepeatingChar(hash); ok {
			for j := 1; j <= 1000; j++ {
				nextHash := getHash(salt, index+j, 0, cache)
				if strings.Contains(nextHash, string(char)+string(char)+string(char)+string(char)+string(char)) {
					found++
					break
				}
			}
		}
		if found == 64 {
			fmt.Println("Part 1:", index)
		}
		index++
	}
}

func part2(input string) {
	salt := input
	index := 0
	found := 0
	cache := make(map[string]string)
	for found < 64 {
		hash := getHash(salt, index, 2016, cache)
		if char, ok := findRepeatingChar(hash); ok {
			for j := 1; j <= 1000; j++ {
				nextHash := getHash(salt, index+j, 2016, cache)
				if strings.Contains(nextHash, string(char)+string(char)+string(char)+string(char)+string(char)) {
					found++
					break
				}
			}
		}
		if found == 64 {
			fmt.Println("Part 2:", index)
		}
		index++
	}
}
