package main

import (
	"aoc/2019/intcomputer"
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
)

func main() {
	input, err := download.ReadInput(2019, 21)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func runSpringScript(program []int, script string) int {
	computer := intcomputer.NewIntcodeComputer(program)
	if err := computer.AddInput(script, "\n"); err != nil {
		log.Fatalf("adding input failed: %v", err)
	}

	var damage int
	for {
		result, err := computer.Run()
		if err != nil {
			log.Fatalf("running program failed: %v", err)
		}

		switch result.Signal {
		case intcomputer.SignalOutput:
			if result.Value > 127 {
				damage = result.Value
			}
		case intcomputer.SignalEnd:
			return damage
		case intcomputer.SignalInput:
			log.Fatal("unexpected input request")
		}
	}
}

func part1(input string) {
	program := conv.ToIntSliceComma(input)

	script := `NOT A J
NOT B T
OR T J
NOT C T
OR T J
AND D J
WALK
`
	damage := runSpringScript(program, script)
	fmt.Println("Part 1", damage)
}

func part2(input string) {
	program := conv.ToIntSliceComma(input)

	script := `NOT A J
NOT B T
OR T J
NOT C T
OR T J
AND D J
NOT E T
NOT T T
OR H T
AND T J
RUN
`
	damage := runSpringScript(program, script)
	fmt.Println("Part 2", damage)
}
