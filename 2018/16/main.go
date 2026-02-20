package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"strings"
)

func main() {
	input, err := download.ReadInput(2018, 16)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func addr(registers []int, a, b, c int) []int {
	registers[c] = registers[a] + registers[b]
	return registers
}

func addi(registers []int, a, b, c int) []int {
	registers[c] = registers[a] + b
	return registers
}

func mulr(registers []int, a, b, c int) []int {
	registers[c] = registers[a] * registers[b]
	return registers
}

func muli(registers []int, a, b, c int) []int {
	registers[c] = registers[a] * b
	return registers
}

func banr(registers []int, a, b, c int) []int {
	registers[c] = registers[a] & registers[b]
	return registers
}

func bani(registers []int, a, b, c int) []int {
	registers[c] = registers[a] & b
	return registers
}

func borr(registers []int, a, b, c int) []int {
	registers[c] = registers[a] | registers[b]
	return registers
}

func bori(registers []int, a, b, c int) []int {
	registers[c] = registers[a] | b
	return registers
}

func setr(registers []int, a, _, c int) []int {
	registers[c] = registers[a]
	return registers
}

func seti(registers []int, a, _, c int) []int {
	registers[c] = a
	return registers
}

func gtir(registers []int, a, b, c int) []int {
	if a > registers[b] {
		registers[c] = 1
	} else {
		registers[c] = 0
	}
	return registers
}

func gtri(registers []int, a, b, c int) []int {
	if registers[a] > b {
		registers[c] = 1
	} else {
		registers[c] = 0
	}
	return registers
}

func gtrr(registers []int, a, b, c int) []int {
	if registers[a] > registers[b] {
		registers[c] = 1
	} else {
		registers[c] = 0
	}
	return registers
}

func eqir(registers []int, a, b, c int) []int {
	if a == registers[b] {
		registers[c] = 1
	} else {
		registers[c] = 0
	}
	return registers
}

func eqri(registers []int, a, b, c int) []int {
	if registers[a] == b {
		registers[c] = 1
	} else {
		registers[c] = 0
	}
	return registers
}

func eqrr(registers []int, a, b, c int) []int {
	if registers[a] == registers[b] {
		registers[c] = 1
	} else {
		registers[c] = 0
	}
	return registers
}

func part1(input string) {
	samples := strings.Split(input, "\n\n")
	count := 0
	opcodes := []func([]int, int, int, int) []int{
		addr, addi, mulr, muli, banr, bani, borr, bori,
		setr, seti, gtir, gtri, gtrr, eqir, eqri, eqrr,
	}

	for _, sample := range samples {
		if !strings.HasPrefix(sample, "Before:") {
			break
		}
		lines := conv.SplitNewline(sample)
		var before [4]int
		var instruction [4]int
		var after [4]int

		conv.MustSscanf(lines[0], "Before: [%d, %d, %d, %d]", &before[0], &before[1], &before[2], &before[3])
		conv.MustSscanf(lines[1], "%d %d %d %d", &instruction[0], &instruction[1], &instruction[2], &instruction[3])
		conv.MustSscanf(lines[2], "After:  [%d, %d, %d, %d]", &after[0], &after[1], &after[2], &after[3])

		matchingOpcodes := 0
		for _, opcodeFunc := range opcodes {
			registers := [4]int{before[0], before[1], before[2], before[3]}
			result := opcodeFunc(registers[:], instruction[1], instruction[2], instruction[3])
			if result[0] == after[0] && result[1] == after[1] && result[2] == after[2] && result[3] == after[3] {
				matchingOpcodes++
			}
		}
		if matchingOpcodes >= 3 {
			count++
		}
	}
	fmt.Println("Part 1", count)
}

func part2(input string) {
	parts := strings.Split(input, "\n\n\n\n")
	samplesStr := parts[0]
	programStr := parts[1]

	samples := strings.Split(samplesStr, "\n\n")
	opcodesMap := make(map[int]*container.Set[string])
	opNames := []string{"addr", "addi", "mulr", "muli", "banr", "bani", "borr", "bori", "setr", "seti", "gtir", "gtri", "gtrr", "eqir", "eqri", "eqrr"}
	opFuncs := map[string]func([]int, int, int, int) []int{
		"addr": addr, "addi": addi, "mulr": mulr, "muli": muli, "banr": banr, "bani": bani, "borr": borr, "bori": bori,
		"setr": setr, "seti": seti, "gtir": gtir, "gtri": gtri, "gtrr": gtrr, "eqir": eqir, "eqri": eqri, "eqrr": eqrr,
	}

	for i := range 16 {
		opcodesMap[i] = container.NewSet[string]()
		for _, name := range opNames {
			opcodesMap[i].Add(name)
		}
	}

	for _, sample := range samples {
		lines := strings.Split(sample, "\n")
		var before [4]int
		var instruction [4]int
		var after [4]int

		conv.MustSscanf(lines[0], "Before: [%d, %d, %d, %d]", &before[0], &before[1], &before[2], &before[3])
		conv.MustSscanf(lines[1], "%d %d %d %d", &instruction[0], &instruction[1], &instruction[2], &instruction[3])
		conv.MustSscanf(lines[2], "After:  [%d, %d, %d, %d]", &after[0], &after[1], &after[2], &after[3])

		opcodeNum := instruction[0]
		for name, opcodeFunc := range opFuncs {
			registers := [4]int{before[0], before[1], before[2], before[3]}
			result := opcodeFunc(registers[:], instruction[1], instruction[2], instruction[3])
			match := result[0] == after[0] && result[1] == after[1] && result[2] == after[2] && result[3] == after[3]
			if !match {
				opcodesMap[opcodeNum].Remove(name)
			}
		}
	}

	finalOpcodes := make(map[int]string)
	for len(finalOpcodes) < 16 {
		for opcodeNum, possibleNames := range opcodesMap {
			if possibleNames.Len() == 1 {
				for _, name := range possibleNames.Values() {
					finalOpcodes[opcodeNum] = name
					for i := range 16 {
						if i != opcodeNum {
							opcodesMap[i].Remove(name)
						}
					}
				}
			}
		}
	}

	programLines := conv.SplitNewline(programStr)
	registers := []int{0, 0, 0, 0}
	for _, line := range programLines {
		var instruction [4]int
		conv.MustSscanf(line, "%d %d %d %d", &instruction[0], &instruction[1], &instruction[2], &instruction[3])
		opName := finalOpcodes[instruction[0]]
		registers = opFuncs[opName](registers, instruction[1], instruction[2], instruction[3])
	}

	fmt.Println("Part 2", registers[0])
}
