package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Gate struct {
	a      string
	op     string
	b      string
	output string
}

func main() {
	content, err := os.ReadFile("2024/24/input.txt")
	if err != nil {
		panic(err)
	}

	parts := strings.Split(strings.TrimSpace(string(content)), "\n\n")
	wiresRaw := strings.Split(parts[0], "\n")
	gatesRaw := strings.Split(parts[1], "\n")

	wires := make(map[string]int)
	for _, line := range wiresRaw {
		parts := strings.Split(line, ": ")
		var value int
		conv.MustSscanf(parts[1], "%d", &value)
		wires[parts[0]] = value
	}

	inputBitCount := len(wiresRaw) / 2

	var gates []Gate
	for _, line := range gatesRaw {
		parts := strings.Split(line, " -> ")
		inputs := strings.Split(parts[0], " ")

		var gate Gate
		if len(inputs) == 3 {
			gate = Gate{
				a:      inputs[0],
				op:     inputs[1],
				b:      inputs[2],
				output: parts[1],
			}
		}

		gates = append(gates, gate)

		if _, exists := wires[gate.a]; !exists {
			wires[gate.a] = -1
		}
		if _, exists := wires[gate.b]; !exists {
			wires[gate.b] = -1
		}
		if _, exists := wires[gate.output]; !exists {
			wires[gate.output] = -1
		}
	}

	isDirect := func(gate Gate) bool {
		return strings.HasPrefix(gate.a, "x") || strings.HasPrefix(gate.b, "x")
	}

	isOutput := func(gate Gate) bool {
		return strings.HasPrefix(gate.output, "z")
	}

	isGate := func(opType string) func(Gate) bool {
		return func(gate Gate) bool {
			return gate.op == opType
		}
	}

	hasOutput := func(output string) func(Gate) bool {
		return func(gate Gate) bool {
			return gate.output == output
		}
	}

	hasInput := func(input string) func(Gate) bool {
		return func(gate Gate) bool {
			return gate.a == input || gate.b == input
		}
	}

	flags := container.NewSet[string]()

	var FAGate0s []Gate
	for _, gate := range gates {
		if isDirect(gate) && gate.op == "XOR" {
			FAGate0s = append(FAGate0s, gate)
		}
	}

	for _, gate := range FAGate0s {
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

	var FAGate3s []Gate
	for _, gate := range gates {
		if isGate("XOR")(gate) && !isDirect(gate) {
			FAGate3s = append(FAGate3s, gate)
			if !isOutput(gate) {
				flags.Add(gate.output)
			}
		}
	}

	for _, gate := range gates {
		if isOutput(gate) {
			lastOutput := fmt.Sprintf("z%03d", inputBitCount)
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

	var checkNext []Gate
	for _, gate := range FAGate0s {
		if flags.Contains(gate.output) {
			continue
		}
		if gate.output == "z00" {
			continue
		}

		var matches []Gate
		for _, g := range FAGate3s {
			if hasInput(gate.output)(g) {
				matches = append(matches, g)
			}
		}

		if len(matches) == 0 {
			checkNext = append(checkNext, gate)
			flags.Add(gate.output)
		}
	}

	for _, gate := range checkNext {
		intendedResult := "z" + gate.a[1:]
		var matches []Gate
		for _, g := range FAGate3s {
			if hasOutput(intendedResult)(g) {
				matches = append(matches, g)
			}
		}

		if len(matches) != 1 {
			panic("Critical Error! Is your input correct?")
		}

		match := matches[0]
		toCheck := []string{match.a, match.b}

		var orMatches []Gate
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

	var flagsSlice []string
	for _, flag := range flags.Values() {
		flagsSlice = append(flagsSlice, flag)
	}
	sort.Strings(flagsSlice)

	if len(flagsSlice) > 8 {
		flagsSlice = flagsSlice[:8]
	}

	fmt.Println(strings.Join(flagsSlice, ","))
}
