package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"regexp"
	"strings"
)

func main() {
	input, err := download.ReadInput(2017, 25)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}
	part1(input)
}

func part1(input string) {
	tm, steps, err := newTuringMachine(input)
	if err != nil {
		log.Fatalf("creating Turing machine failed: %v", err)
	}

	for range steps {
		tm.step()
	}

	fmt.Println("Part 1", tm.checksum())
}

type rule struct {
	writeValue int
	move       int
	nextState  string
}

type state struct {
	name  string
	rules map[int]rule
}

type turingMachine struct {
	states       map[string]*state
	tape         map[int]int
	cursor       int
	currentState string
}

func newTuringMachine(input string) (*turingMachine, int, error) {
	patterns := map[string]*regexp.Regexp{
		"initial": regexp.MustCompile(`Begin in state (\w)\.`),
		"steps":   regexp.MustCompile(`Perform a diagnostic checksum after (\d+) steps\.`),
		"state":   regexp.MustCompile(`In state (\w):`),
		"ifZero":  regexp.MustCompile(`If the current value is 0:`),
		"ifOne":   regexp.MustCompile(`If the current value is 1:`),
		"write":   regexp.MustCompile(`- Write the value (\d)\.`),
		"move":    regexp.MustCompile(`- Move one slot to the (left|right)\.`),
		"next":    regexp.MustCompile(`- Continue with state (\w)\.`),
	}

	lines := conv.SplitNewline(input)
	initialState := patterns["initial"].FindStringSubmatch(lines[0])[1]
	steps := conv.MustAtoi(patterns["steps"].FindStringSubmatch(lines[1])[1])

	tm := &turingMachine{
		states:       make(map[string]*state),
		tape:         make(map[int]int),
		cursor:       0,
		currentState: initialState,
	}

	var currentState *state
	for i := 3; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if line == "" {
			continue
		}

		if matches := patterns["state"].FindStringSubmatch(line); matches != nil {
			stateName := matches[1]
			currentState = &state{
				name:  stateName,
				rules: make(map[int]rule),
			}
			tm.states[stateName] = currentState
			continue
		}

		if currentState == nil {
			continue
		}

		if patterns["ifZero"].MatchString(line) || patterns["ifOne"].MatchString(line) {
			value := 0
			if patterns["ifOne"].MatchString(line) {
				value = 1
			}

			rule := rule{
				writeValue: conv.MustAtoi(patterns["write"].FindStringSubmatch(lines[i+1])[1]),
				nextState:  patterns["next"].FindStringSubmatch(lines[i+3])[1],
			}

			if move := patterns["move"].FindStringSubmatch(lines[i+2])[1]; move == "left" {
				rule.move = -1
			} else {
				rule.move = 1
			}

			currentState.rules[value] = rule
			i += 3
		}
	}

	return tm, steps, nil
}

func (tm *turingMachine) step() {
	currentValue := tm.tape[tm.cursor]
	rule := tm.states[tm.currentState].rules[currentValue]

	tm.tape[tm.cursor] = rule.writeValue
	tm.cursor += rule.move
	tm.currentState = rule.nextState
}

func (tm *turingMachine) checksum() int {
	sum := 0
	for _, v := range tm.tape {
		if v == 1 {
			sum++
		}
	}
	return sum
}
