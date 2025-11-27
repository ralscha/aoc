package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"strings"
)

func main() {
	input, err := download.ReadInput(2020, 18)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	sum := 0
	for _, line := range lines {
		res := evaluatePart1(line)
		sum += res
	}
	fmt.Println("Part 1", sum)
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	sum := 0
	for _, line := range lines {
		res := evaluatePart2(line)
		sum += res
	}
	fmt.Println("Part 2", sum)
}

func evaluatePart1(expression string) int {
	return evaluate(expression, map[string]int{"+": 1, "*": 1})
}

func evaluatePart2(expression string) int {
	return evaluate(expression, map[string]int{"+": 2, "*": 1})
}

func evaluate(expression string, precedence map[string]int) int {
	values := make([]int, 0)
	operators := make([]string, 0)

	tokens := strings.Split(strings.ReplaceAll(expression, " ", ""), "")

	applyOp := func() {
		operator := operators[len(operators)-1]
		operators = operators[:len(operators)-1]
		b := values[len(values)-1]
		values = values[:len(values)-1]
		a := values[len(values)-1]
		values = values[:len(values)-1]

		var res int
		switch operator {
		case "+":
			res = a + b
		case "*":
			res = a * b
		}
		values = append(values, res)
	}

	for _, token := range tokens {
		if token == "(" {
			operators = append(operators, token)
		} else if token == ")" {
			for operators[len(operators)-1] != "(" {
				applyOp()
			}
			operators = operators[:len(operators)-1]
		} else if op, ok := precedence[token]; ok {
			for len(operators) > 0 && operators[len(operators)-1] != "(" && precedence[operators[len(operators)-1]] >= op {
				applyOp()
			}
			operators = append(operators, token)
		} else {
			num := conv.MustAtoi(token)
			values = append(values, num)
		}
	}

	for len(operators) > 0 {
		applyOp()
	}

	return values[0]
}
