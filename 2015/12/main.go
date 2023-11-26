package main

import (
	"aoc/internal/download"
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	input, err := download.ReadInput(2015, 12)
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
	fmt.Printf("Total: %d\n", total)

	total = walkIgnoreRed(v)
	fmt.Printf("Total: %d", total)

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
