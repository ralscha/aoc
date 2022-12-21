package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"strings"
)

func main() {
	inputFile := "./2022/21/input.txt"
	input, err := download.ReadInput(inputFile, 2022, 21)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)

}

type monkey struct {
	name string

	leftMonkey  string
	operation   string
	rightMonkey string

	yell int
}

var monkeys = make(map[string]*monkey)

func (m *monkey) Yell() {
	if m.leftMonkey != "" {
		monkeys[m.leftMonkey].Yell()
	}
	if m.rightMonkey != "" {
		monkeys[m.rightMonkey].Yell()
	}

	if m.operation != "" {
		switch m.operation {
		case "+":
			m.yell = monkeys[m.leftMonkey].yell + monkeys[m.rightMonkey].yell
		case "-":
			m.yell = monkeys[m.leftMonkey].yell - monkeys[m.rightMonkey].yell
		case "/":
			m.yell = monkeys[m.leftMonkey].yell / monkeys[m.rightMonkey].yell
		case "*":
			m.yell = monkeys[m.leftMonkey].yell * monkeys[m.rightMonkey].yell
		}
	}
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	createMonkeys(lines)
	root := monkeys["root"]
	root.Yell()
	fmt.Println(root.yell)
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	createMonkeys(lines)

	human := monkeys["humn"]
	root := monkeys["root"]
	leftMonkey := monkeys[root.leftMonkey]
	rightMonkey := monkeys[root.rightMonkey]

	human.yell = 155
	root.Yell()
	y1l := leftMonkey.yell
	y1r := rightMonkey.yell

	human.yell = 153
	root.Yell()
	y2l := leftMonkey.yell
	// y2r := rightMonkey.yell

	var targetMonkey *monkey
	target := 0
	if y1l == y2l {
		target = y1l
		targetMonkey = monkeys[root.rightMonkey]
	} else {
		target = y1r
		targetMonkey = monkeys[root.leftMonkey]
	}

	low := 1
	high := 9000000000000
	for low < high {
		middle := (low + high) / 2
		human.yell = middle
		targetMonkey.Yell()
		if targetMonkey.yell == target {
			fmt.Println(middle - 1)
			break
		}
		if targetMonkey.yell < target {
			high = middle
		} else {
			low = middle + 1
		}
	}
}

func createMonkeys(lines []string) {
	for _, line := range lines {
		colonIx := strings.Index(line, ":")
		name := line[:colonIx]
		job := line[colonIx+2:]
		yelled := 0
		operation := ""
		leftMonkey := ""
		rightMonkey := ""
		if job[0] >= '0' && job[0] <= '9' {
			yelled = conv.MustAtoi(job)
		} else {
			splitted := strings.Fields(job)
			leftMonkey = splitted[0]
			operation = splitted[1]
			rightMonkey = splitted[2]
		}

		monkeys[name] = &monkey{name: name, yell: yelled,
			operation: operation, leftMonkey: leftMonkey, rightMonkey: rightMonkey}
	}
}
