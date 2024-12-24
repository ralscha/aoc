package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"maps"
	"slices"
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
	initialWireValues, gatesMap := parseInput(input)
	inputBitCount := len(initialWireValues) / 2
	gates := slices.Collect(maps.Values(gatesMap))

	isDirect := func(gate gate) bool {
		return strings.HasPrefix(gate.a, "x") || strings.HasPrefix(gate.b, "x")
	}

	isOutput := func(gate gate) bool {
		return strings.HasPrefix(gate.output, "z")
	}

	isGate := func(opType string) func(gate) bool {
		return func(gate gate) bool {
			return gate.op == opType
		}
	}

	hasOutput := func(output string) func(gate) bool {
		return func(gate gate) bool {
			return gate.output == output
		}
	}

	hasInput := func(input string) func(gate) bool {
		return func(gate gate) bool {
			return gate.a == input || gate.b == input
		}
	}

	flags := container.NewSet[string]()

	var faGate0s []gate
	for _, gate := range gates {
		if isDirect(gate) && gate.op == "XOR" {
			faGate0s = append(faGate0s, gate)
		}
	}

	for _, gate := range faGate0s {
		isFirst := gate.a == "x00" || gate.b == "x00"
		if isFirst {
			if gate.output != "z00" {
				flags.Add(gate.output)
			}
			continue
		} else if gate.output == "z00" {
			flags.Add(gate.output)
		}

		if isOutput(gate) {
			flags.Add(gate.output)
		}
	}

	var faGate3s []gate
	for _, gate := range gates {
		if isGate("XOR")(gate) && !isDirect(gate) {
			faGate3s = append(faGate3s, gate)
			if !isOutput(gate) {
				flags.Add(gate.output)
			}
		}
	}

	for _, gate := range gates {
		if isOutput(gate) {
			lastOutput := fmt.Sprintf("z%02d", inputBitCount)
			isLast := gate.output == lastOutput

			if isLast {
				if gate.op != "OR" {
					flags.Add(gate.output)
				}
				continue
			} else if gate.op != "XOR" {
				flags.Add(gate.output)
			}
		}
	}

	var checkNext []gate
	for _, ga := range faGate0s {
		if flags.Contains(ga.output) {
			continue
		}
		if ga.output == "z00" {
			continue
		}

		var matches []gate
		for _, g := range faGate3s {
			if hasInput(ga.output)(g) {
				matches = append(matches, g)
			}
		}

		if len(matches) == 0 {
			checkNext = append(checkNext, ga)
			flags.Add(ga.output)
		}
	}

	for _, ga := range checkNext {
		intendedResult := "z" + ga.a[1:]
		var matches []gate
		for _, g := range faGate3s {
			if hasOutput(intendedResult)(g) {
				matches = append(matches, g)
			}
		}

		match := matches[0]
		toCheck := []string{match.a, match.b}

		var orMatches []gate
		for _, g := range gates {
			if isGate("OR")(g) {
				for _, output := range toCheck {
					if g.output == output {
						orMatches = append(orMatches, g)
					}
				}
			}
		}

		orMatchOutput := orMatches[0].output
		var correctOutput string
		for _, output := range toCheck {
			if output != orMatchOutput {
				correctOutput = output
				break
			}
		}
		flags.Add(correctOutput)
	}

	flagsSlice := flags.Values()
	slices.Sort(flagsSlice)

	if len(flagsSlice) > 8 {
		flagsSlice = flagsSlice[:8]
	}

	fmt.Println("Part 2", strings.Join(flagsSlice, ","))
}
