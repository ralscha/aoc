package main

import (
	"aoc/internal/download"
	"encoding/json"
	"log"
)

func main() {
	inputFile := "./2015/12/input.txt"
	input, err := download.ReadInput(inputFile, 2015, 12)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1and2(input)
}

func part1and2(input string) {
	var v []any
	err := json.Unmarshal([]byte(input), &v)
	if err != nil {
		log.Fatalf("unmarshalling failed: %v", err)
	}

	total := walk(v)
	log.Printf("Total: %d", total)

	total = walkIgnoreRed(v)
	log.Printf("Total: %d", total)

}

func walk(v any) int {
	total := 0
	switch v := v.(type) {
	case map[string]any:
		for _, value := range v {
			total += walk(value)
		}
	case []any:
		for _, value := range v {
			total += walk(value)
		}
	case float64:
		total += int(v)
	}
	return total
}

func walkIgnoreRed(v any) int {
	total := 0
	switch v := v.(type) {
	case map[string]any:
		for _, value := range v {
			if value == "red" {
				return 0
			}
		}
		for _, value := range v {
			total += walkIgnoreRed(value)
		}
	case []any:
		for _, value := range v {
			total += walkIgnoreRed(value)
		}
	case float64:
		total += int(v)
	}
	return total
}
