package main

import (
	"aoc/internal/download"
	"crypto/md5"
	"fmt"
	"log"
	"strings"
)

func main() {
	input, err := download.ReadInput(2016, 5)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}
	input = strings.Trim(input, "\n")
	part1(input)
	part2(input)
}

func part1(start string) {
	password := ""
	for i := 0; ; i++ {
		hash := md5.Sum([]byte(fmt.Sprintf("%s%d", start, i)))
		if hash[0] == 0 && hash[1] == 0 && hash[2] < 16 {
			hashStr := fmt.Sprintf("%x", hash)
			password += hashStr[5:6]
			if len(password) == 8 {
				break
			}
		}
	}
	fmt.Println(password)
}

func part2(start string) {
	password := make([]string, 8)
	for i := 0; ; i++ {
		hash := md5.Sum([]byte(fmt.Sprintf("%s%d", start, i)))
		if hash[0] == 0 && hash[1] == 0 && hash[2] < 8 {
			hashStr := fmt.Sprintf("%x", hash)
			pos := hashStr[5]
			if pos >= '0' && pos <= '7' {
				posInt := pos - '0'
				if password[posInt] == "" {
					password[posInt] = hashStr[6:7]
				}
			}
			if len(strings.Join(password, "")) == 8 {
				break
			}
		}
	}
	fmt.Println(strings.Join(password, ""))
}
