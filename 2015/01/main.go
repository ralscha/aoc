package main

import (
	"aoc/internal/download"
	"fmt"
	"log"
)

func main() {
	input, err := download.ReadInput(2015, 1)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	floor := 0
	posBasement := 0
	for pos, c := range input {
		if c == '(' {
			floor++
		} else if c == ')' {
			floor--
		}

		if floor == -1 && posBasement == 0 {
			posBasement = pos + 1
		}
	}

	fmt.Println("Current floor:", floor)
	fmt.Println("Position of basement:", posBasement)
}
