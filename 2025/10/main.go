package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	input, err := download.ReadInput(2025, 10)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	total := 0

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		target, buttons := parseMachineLine(line)

		minPresses, solvable := minPressesMachine(target, buttons)
		if !solvable {
			continue
		}

		total += minPresses
	}

	fmt.Println("Part 1", total)
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	total := 0

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		requirements, buttons := parseJoltageMachineLine(line)

		minPresses := minPressesJoltage(requirements, buttons)
		total += minPresses
	}

	fmt.Println("Part 2", total)

}

func parseMachineLine(line string) ([]int, [][]int) {
	target, buttons, _ := parseStrictLine(line)
	return target, buttons
}

func parseJoltageMachineLine(line string) ([]int, [][]int) {
	_, buttons, requirements := parseStrictLine(line)
	return requirements, buttons
}

func parseLightDiagram(raw string) []int {
	target := make([]int, len(raw))
	for i, r := range raw {
		switch r {
		case '.':
			target[i] = 0
		case '#':
			target[i] = 1
		}
	}

	return target
}

func parseJoltageRequirements(raw string) []int {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return []int{}
	}

	parts := strings.Split(raw, ",")
	values := make([]int, len(parts))
	for i := range parts {
		values[i] = conv.MustAtoi(strings.TrimSpace(parts[i]))
	}

	return values
}

func parseButtons(raw string, size int) [][]int {
	if strings.TrimSpace(raw) == "" {
		return [][]int{}
	}

	buttonParts := strings.Split(raw, ") (")
	buttons := make([][]int, 0, len(buttonParts))

	for i := range buttonParts {
		segment := buttonParts[i]
		if i == 0 {
			segment = strings.TrimPrefix(segment, "(")
		}
		if i == len(buttonParts)-1 {
			segment = strings.TrimSuffix(segment, ")")
		}

		button := make([]int, size)
		if strings.TrimSpace(segment) != "" {
			parts := strings.SplitSeq(segment, ",")
			for p := range parts {
				idx := conv.MustAtoi(strings.TrimSpace(p))
				button[idx] = 1
			}
		}

		buttons = append(buttons, button)
	}

	return buttons
}

func parseStrictLine(line string) ([]int, [][]int, []int) {
	closeBracket := strings.Index(line, "]")
	closeBrace := strings.LastIndex(line, "}")
	openBrace := strings.LastIndex(line[:closeBrace], "{")

	lightRaw := line[1:closeBracket]
	buttonsRaw := strings.TrimSpace(line[closeBracket+1 : openBrace])
	requirementsRaw := line[openBrace+1 : closeBrace]

	target := parseLightDiagram(lightRaw)
	buttons := parseButtons(buttonsRaw, len(target))
	requirements := parseJoltageRequirements(requirementsRaw)

	return target, buttons, requirements
}

func minPressesMachine(target []int, buttons [][]int) (int, bool) {
	m := len(target)
	n := len(buttons)

	table := make([]int, m)
	best := -1
	limit := 1 << n

	for subset := range limit {
		for i := range m {
			table[i] = 0
		}

		presses := 0
		for button := range n {
			if ((subset >> button) & 1) == 0 {
				continue
			}
			presses++
			for i := range m {
				table[i] += buttons[button][i]
			}
		}

		matches := true
		for i := range m {
			if (table[i] & 1) != target[i] {
				matches = false
				break
			}
		}

		if !matches {
			continue
		}

		if best == -1 || presses < best {
			best = presses
		}
	}

	if best == -1 {
		return 0, false
	}

	return best, true
}

const inf = int(^uint(0) >> 2)

func minPressesJoltage(requirements []int, buttons [][]int) int {
	m := len(requirements)

	buttonIdxs := make([][]int, 0, len(buttons))
	for j := range buttons {
		idxs := make([]int, 0, m)
		for i := range m {
			if buttons[j][i] == 1 {
				idxs = append(idxs, i)
			}
		}
		if len(idxs) > 0 {
			buttonIdxs = append(buttonIdxs, idxs)
		}
	}

	n := len(buttonIdxs)

	patternFrom := func(values []int) string {
		var builder strings.Builder
		builder.Grow(len(values))
		for _, v := range values {
			if v&1 == 0 {
				builder.WriteByte('0')
			} else {
				builder.WriteByte('1')
			}
		}
		return builder.String()
	}

	type combo struct {
		presses int
		joltage []int
	}

	combosByPattern := make(map[string][]combo)
	totalCombos := 1 << n
	for subset := range totalCombos {
		presses := 0
		joltage := make([]int, m)

		for button := range n {
			if ((subset >> button) & 1) == 0 {
				continue
			}
			presses++
			for _, idx := range buttonIdxs[button] {
				joltage[idx]++
			}
		}

		pattern := patternFrom(joltage)
		combosByPattern[pattern] = append(combosByPattern[pattern], combo{presses: presses, joltage: joltage})
	}

	cache := make(map[string]int)

	stateKey := func(values []int) string {
		var builder strings.Builder
		builder.Grow(len(values) * 4)
		for i, v := range values {
			if i > 0 {
				builder.WriteByte(',')
			}
			builder.WriteString(strconv.Itoa(v))
		}
		return builder.String()
	}

	var countPresses func([]int) int
	countPresses = func(target []int) int {
		key := stateKey(target)
		if val, ok := cache[key]; ok {
			return val
		}

		onlyZeros := true
		for _, jolt := range target {
			if jolt < 0 {
				cache[key] = inf
				return inf
			}
			if jolt > 0 {
				onlyZeros = false
			}
		}
		if onlyZeros {
			cache[key] = 0
			return 0
		}

		pattern := patternFrom(target)
		combos, ok := combosByPattern[pattern]
		if !ok {
			cache[key] = inf
			return inf
		}

		best := inf
		for _, candidate := range combos {
			halfTarget := make([]int, m)
			valid := true
			for i := range m {
				delta := target[i] - candidate.joltage[i]
				if delta < 0 || (delta&1) != 0 {
					valid = false
					break
				}
				halfTarget[i] = delta / 2
			}
			if !valid {
				continue
			}

			sub := countPresses(halfTarget)
			if sub == inf {
				continue
			}

			total := candidate.presses + 2*sub
			if total < best {
				best = total
			}
		}

		cache[key] = best
		return best
	}

	answer := countPresses(requirements)
	return answer
}
