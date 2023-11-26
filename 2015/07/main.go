package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input, err := download.ReadInput(2015, 7)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1and2(input, false)
	part1and2(input, true)

}

func part1and2(input string, overrideB bool) {
	lines := conv.SplitNewline(input)
	re := regexp.MustCompile(`([a-z0-9]+|NOT)\s*([A-Z]*)\s*([a-z0-9]*)`)

	instructions := make(map[string]string)
	for _, line := range lines {
		parts := strings.Split(line, "->")
		instructions[strings.TrimSpace(parts[1])] = strings.TrimSpace(parts[0])

		if overrideB {
			instructions["b"] = "3176"
		}
	}

	wires := make(map[string]uint16)
	for {
		updated := false
		for dest, src := range instructions {
			matches := re.FindStringSubmatch(src)
			src1, op, src2 := matches[1], matches[2], matches[3]

			if _, err := strconv.Atoi(src1); err != nil && strings.ToLower(src1) == src1 {
				if _, ok := wires[src1]; !ok {
					continue
				}
			}

			if src2 != "" {
				if _, err := strconv.Atoi(src2); err != nil && strings.ToLower(src2) == src2 {
					if _, ok := wires[src2]; !ok {
						continue
					}
				}
			}

			if src1 == "NOT" {
				wires[dest] = ^value(wires, src2)
			} else {
				switch op {
				case "AND":
					wires[dest] = value(wires, src1) & value(wires, src2)
				case "OR":
					wires[dest] = value(wires, src1) | value(wires, src2)
				case "LSHIFT":
					wires[dest] = value(wires, src1) << conv.MustAtoi(src2)
				case "RSHIFT":
					wires[dest] = value(wires, src1) >> conv.MustAtoi(src2)
				default:
					wires[dest] = value(wires, src1)
				}
			}
			updated = true
			delete(instructions, dest)
		}
		if !updated {
			break
		}
	}
	fmt.Println(wires["a"])
}

func value(wires map[string]uint16, wireName string) uint16 {
	if isWire(wireName) {
		return wires[wireName]
	}
	return uint16(conv.MustAtoi(wireName))
}

func isWire(s string) bool {
	if _, err := strconv.Atoi(s); err != nil && strings.ToLower(s) == s {
		return true
	}
	return false
}
