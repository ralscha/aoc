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
	inputFile := "./2015/07/input.txt"
	input, err := download.ReadInput(inputFile, 2015, 7)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	//	part2(input)

}

func part1(input string) {
	lines := conv.SplitNewline(input)
	re := regexp.MustCompile(`([a-z0-9]+|NOT)\s*([A-Z]*)\s*([a-z0-9]*)`)

	instructions := make(map[string]string)
	for _, line := range lines {
		parts := strings.Split(line, "->")
		instructions[strings.TrimSpace(parts[1])] = strings.TrimSpace(parts[0])
	}

	wires := make(map[string]uint16)
	for {
		updated := false
		for dest, src := range instructions {
			matches := re.FindStringSubmatch(src)
			src1, op, src2 := matches[1], matches[2], matches[3]
			// fmt.Println(src1, op, src2)
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
				fmt.Println("NOT", src2, "->", dest)
				wires[dest] = ^wires[src2]
			} else {
				switch op {
				case "AND":
					fmt.Println(src1, "AND", src2, "->", dest)
					wires[dest] = wires[src1] & wires[src2]
				case "OR":
					fmt.Println(src1, "OR", src2, "->", dest)
					wires[dest] = wires[src1] | wires[src2]
				case "LSHIFT":
					fmt.Println(src1, "LSHIFT", src2, "->", dest)
					wires[dest] = wires[src1] << conv.MustAtoi(src2)
				case "RSHIFT":
					fmt.Println(src1, "RSHIFT", src2, "->", dest)
					wires[dest] = wires[src1] >> conv.MustAtoi(src2)
				default:
					fmt.Println(src1, "->", dest)
					if val, err := strconv.Atoi(src1); err == nil {
						wires[dest] = uint16(val)
					} else {
						fmt.Println("src1", src1)
						fmt.Println("value", wires[src1])
						wires[dest] = wires[src1]
					}
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

	for k, v := range wires {
		fmt.Println(k, v)
	}

}
