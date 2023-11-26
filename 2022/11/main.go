package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"strings"
)

type monkey struct {
	items          []int
	operation      string
	operationValue int
	test           int
	ifTrue         int
	ifFalse        int
}

func main() {
	input, err := download.ReadInput(2022, 11)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	lines := conv.SplitNewline(input)

	monkeys := createMonkeys(lines)
	inspectCount := make(map[int]int)

	for i := 0; i < 20; i++ {
		for m := 0; m < len(monkeys); m++ {
			monkey := monkeys[m]
			for _, item := range monkey.items {
				inspectCount[m]++
				newWorryLevel := item
				if monkey.operation == "*" {
					if monkey.operationValue == -1 {
						newWorryLevel *= newWorryLevel
					} else {
						newWorryLevel *= monkey.operationValue
					}
				} else if monkey.operation == "+" {
					newWorryLevel += monkey.operationValue
				} else if monkey.operation == "-" {
					newWorryLevel -= monkey.operationValue
				} else if monkey.operation == "/" {
					newWorryLevel /= monkey.operationValue
				}

				newWorryLevel /= 3

				if newWorryLevel%monkey.test == 0 {
					trueMonkey := monkeys[monkey.ifTrue]
					trueMonkey.items = append(trueMonkey.items, newWorryLevel)
				} else {
					falseMonkey := monkeys[monkey.ifFalse]
					falseMonkey.items = append(falseMonkey.items, newWorryLevel)
				}
			}
			monkey.items = []int{}
		}
	}

	printTop2Multiplication(inspectCount)
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	monkeys := createMonkeys(lines)

	prod := 1
	for _, m := range monkeys {
		prod *= m.test
	}

	inspectCount := make(map[int]int)

	for i := 0; i < 10000; i++ {
		for m := 0; m < len(monkeys); m++ {
			monkey := monkeys[m]
			for _, item := range monkey.items {
				inspectCount[m]++
				newWorryLevel := item
				if monkey.operation == "*" {
					if monkey.operationValue == -1 {
						newWorryLevel *= newWorryLevel
					} else {
						newWorryLevel *= monkey.operationValue
					}
				} else if monkey.operation == "+" {
					newWorryLevel += monkey.operationValue
				} else if monkey.operation == "-" {
					newWorryLevel -= monkey.operationValue
				} else if monkey.operation == "/" {
					newWorryLevel /= monkey.operationValue
				}

				newWorryLevel = newWorryLevel % prod

				if newWorryLevel%monkey.test == 0 {
					trueMonkey := monkeys[monkey.ifTrue]
					trueMonkey.items = append(trueMonkey.items, newWorryLevel)
				} else {
					falseMonkey := monkeys[monkey.ifFalse]
					falseMonkey.items = append(falseMonkey.items, newWorryLevel)
				}
			}
			monkey.items = []int{}
		}
	}

	printTop2Multiplication(inspectCount)
}

func printTop2Multiplication(inspectCount map[int]int) {
	var top1, top2 int
	for _, count := range inspectCount {
		if count > top1 {
			top2 = top1
			top1 = count
		} else if count > top2 {
			top2 = count
		}
	}

	fmt.Println(top1 * top2)
}

func createMonkeys(lines []string) map[int]*monkey {
	monkeys := make(map[int]*monkey)
	for i := 0; i < len(lines); i += 7 {
		first := strings.Fields(lines[i])
		monkeyNo := conv.MustAtoi(first[1][:len(first[1])-1])
		items := strings.Fields(lines[i+1])[2:]
		itemsNo := make([]int, len(items))
		for j := range items {
			if items[j][len(items[j])-1] == ',' {
				itemsNo[j] = conv.MustAtoi(items[j][:len(items[j])-1])
			} else {
				itemsNo[j] = conv.MustAtoi(items[j][:len(items[j])])
			}
		}

		operationStr := strings.Fields(lines[i+2])[3:]
		operation := operationStr[1]
		operationValue := operationStr[2]
		operationValueInt := -1
		if operationValue != "old" {
			operationValueInt = conv.MustAtoi(operationValue)
		}

		test := conv.MustAtoi(strings.Fields(lines[i+3])[3])
		ifTrue := conv.MustAtoi(strings.Fields(lines[i+4])[5])
		ifFalse := conv.MustAtoi(strings.Fields(lines[i+5])[5])

		monkeys[monkeyNo] = &monkey{
			items:          itemsNo,
			operation:      operation,
			operationValue: operationValueInt,
			test:           test,
			ifTrue:         ifTrue,
			ifFalse:        ifFalse,
		}
	}
	return monkeys
}
