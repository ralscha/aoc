package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"golang.org/x/exp/maps"
	"log"
	"slices"
	"strings"
)

func main() {
	input, err := download.ReadInput(2017, 8)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}
	part1and2(input)
}

func part1and2(input string) {
	lines := conv.SplitNewline(input)
	registers := make(map[string]int)
	highestValue := 0
	for _, line := range lines {
		splitted := strings.Split(line, " ")
		register := splitted[0]
		instruction := splitted[1]
		value := conv.MustAtoi(splitted[2])
		conditionRegister := splitted[4]
		condition := splitted[5]
		conditionValue := conv.MustAtoi(splitted[6])

		if _, ok := registers[register]; !ok {
			registers[register] = 0
		}

		if _, ok := registers[conditionRegister]; !ok {
			registers[conditionRegister] = 0
		}

		if checkCondition(registers[conditionRegister], condition, conditionValue) {
			if instruction == "inc" {
				registers[register] += value
			} else {
				registers[register] -= value
			}
			if registers[register] > highestValue {
				highestValue = registers[register]
			}
		}
	}

	maxValue := slices.Max(maps.Values(registers))
	fmt.Println(maxValue)
	fmt.Println(highestValue)
}

func checkCondition(registerValue int, condition string, conditionValue int) bool {
	switch condition {
	case ">":
		return registerValue > conditionValue
	case "<":
		return registerValue < conditionValue
	case ">=":
		return registerValue >= conditionValue
	case "<=":
		return registerValue <= conditionValue
	case "==":
		return registerValue == conditionValue
	case "!=":
		return registerValue != conditionValue
	default:
		panic(fmt.Sprintf("unknown condition %s", condition))
	}
}
