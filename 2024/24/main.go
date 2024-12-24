package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type gate struct {
	a      string
	op     string
	b      string
	output string
}

func main() {
	input, err := download.ReadInput(2024, 24)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func parseInput(input string) (map[string]int, map[string]gate) {
	lines := conv.SplitNewline(input)
	wireValues := make(map[string]int)
	gates := make(map[string]gate)

	for _, line := range lines {
		if strings.Contains(line, ":") {
			parts := strings.Split(line, ": ")
			value := conv.MustAtoi(parts[1])
			wireValues[parts[0]] = value
		} else if strings.Contains(line, "->") {
			parts := strings.Split(line, " ")
			output := parts[len(parts)-1]
			gate := gate{a: parts[0], op: parts[1], b: parts[2], output: output}
			gates[output] = gate
		}
	}
	return wireValues, gates
}

func evaluateCircuit(initialWireValues map[string]int, gates map[string]gate) map[string]int {
	wireValues := make(map[string]int)
	for k, v := range initialWireValues {
		wireValues[k] = v
	}

	evaluated := container.NewSet[string]()

	canEvaluate := func(gate gate) bool {
		_, ok1 := wireValues[gate.a]
		_, ok2 := wireValues[gate.b]
		return ok1 && ok2
	}

	allEvaluated := false
	for !allEvaluated {
		numEvaluated := 0
		for _, gate := range gates {
			if !evaluated.Contains(gate.output) {
				switch gate.op {
				case "AND":
					if canEvaluate(gate) {
						wireValues[gate.output] = wireValues[gate.a] & wireValues[gate.b]
						evaluated.Add(gate.output)
						numEvaluated++
					}
				case "OR":
					if canEvaluate(gate) {
						wireValues[gate.output] = wireValues[gate.a] | wireValues[gate.b]
						evaluated.Add(gate.output)
						numEvaluated++
					}
				case "XOR":
					if canEvaluate(gate) {
						wireValues[gate.output] = wireValues[gate.a] ^ wireValues[gate.b]
						evaluated.Add(gate.output)
						numEvaluated++
					}
				}
			}
		}
		if numEvaluated == 0 && evaluated.Len() < len(gates) {
			return wireValues
		}
		if evaluated.Len() == len(gates) {
			allEvaluated = true
		}
	}
	return wireValues
}

func getBinaryValue(wireValues map[string]int, prefix string) string {
	binaryResult := ""
	for i := 0; ; i++ {
		wire := fmt.Sprintf("%s%02d", prefix, i)
		if value, ok := wireValues[wire]; ok {
			binaryResult = strconv.Itoa(value) + binaryResult
		} else {
			break
		}
	}
	return binaryResult
}

func part1(input string) {
	initialWireValues, gatesMap := parseInput(input)
	wireValues := evaluateCircuit(initialWireValues, gatesMap)
	binaryResult := getBinaryValue(wireValues, "z")
	decimalResult, _ := strconv.ParseInt(binaryResult, 2, 64)
	fmt.Println("Part 1", decimalResult)
}

func part2(input string) {

}
