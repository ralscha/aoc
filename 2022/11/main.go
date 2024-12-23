package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/mathx"
	"fmt"
	"log"
	"strings"
)

type monkey struct {
	items          *container.Queue[int]
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

	for range 20 {
		for m := 0; m < len(monkeys); m++ {
			monkey := monkeys[m]
			for !monkey.items.IsEmpty() {
				inspectCount[m]++
				item := monkey.items.Pop()
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
					monkeys[monkey.ifTrue].items.Push(newWorryLevel)
				} else {
					monkeys[monkey.ifFalse].items.Push(newWorryLevel)
				}
			}
		}
	}

	printTop2Multiplication(inspectCount)
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	monkeys := createMonkeys(lines)

	// Calculate LCM of all test values to keep worry levels manageable
	testValues := make([]int, 0, len(monkeys))
	for _, m := range monkeys {
		testValues = append(testValues, m.test)
	}
	lcm := mathx.Lcm(testValues)

	inspectCount := make(map[int]int)

	for range 10000 {
		for m := 0; m < len(monkeys); m++ {
			monkey := monkeys[m]
			for !monkey.items.IsEmpty() {
				inspectCount[m]++
				item := monkey.items.Pop()
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

				newWorryLevel = newWorryLevel % lcm

				if newWorryLevel%monkey.test == 0 {
					monkeys[monkey.ifTrue].items.Push(newWorryLevel)
				} else {
					monkeys[monkey.ifFalse].items.Push(newWorryLevel)
				}
			}
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

		items := container.NewQueue[int]()
		itemsStr := strings.Fields(lines[i+1])[2:]
		for _, item := range itemsStr {
			if item[len(item)-1] == ',' {
				items.Push(conv.MustAtoi(item[:len(item)-1]))
			} else {
				items.Push(conv.MustAtoi(item))
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
			items:          items,
			operation:      operation,
			operationValue: operationValueInt,
			test:           test,
			ifTrue:         ifTrue,
			ifFalse:        ifFalse,
		}
	}
	return monkeys
}
