package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/mathx"
	"fmt"
	"log"
)

func main() {
	input, err := download.ReadInput(2019, 7)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

type amplifier1 struct {
	code       []int
	currentPos int
	inputs     []int
}

func newAmplifier(programCode []int, inputs []int) *amplifier1 {
	programCodeCopy := make([]int, len(programCode))
	copy(programCodeCopy, programCode)
	return &amplifier1{code: programCodeCopy, inputs: inputs}
}

func (c *amplifier1) run() (int, bool) {
	inputIdx := 0
	for {
		opcode := c.code[c.currentPos] % 100
		modes := []int{
			(c.code[c.currentPos] / 100) % 10,
			(c.code[c.currentPos] / 1000) % 10,
		}

		switch opcode {
		case 1: // add
			p1 := c.getParam(1, modes[0])
			p2 := c.getParam(2, modes[1])
			c.setParam(3, p1+p2)
			c.currentPos += 4
		case 2: // multiply
			p1 := c.getParam(1, modes[0])
			p2 := c.getParam(2, modes[1])
			c.setParam(3, p1*p2)
			c.currentPos += 4
		case 3: // input
			if inputIdx >= len(c.inputs) {
				return 0, false
			}
			c.setParam(1, c.inputs[inputIdx])
			inputIdx++
			c.currentPos += 2
		case 4: // output
			output := c.getParam(1, modes[0])
			c.currentPos += 2
			return output, true
		case 5: // jump-if-true
			p1 := c.getParam(1, modes[0])
			p2 := c.getParam(2, modes[1])
			if p1 != 0 {
				c.currentPos = p2
			} else {
				c.currentPos += 3
			}
		case 6: // jump-if-false
			p1 := c.getParam(1, modes[0])
			p2 := c.getParam(2, modes[1])
			if p1 == 0 {
				c.currentPos = p2
			} else {
				c.currentPos += 3
			}
		case 7: // less than
			p1 := c.getParam(1, modes[0])
			p2 := c.getParam(2, modes[1])
			if p1 < p2 {
				c.setParam(3, 1)
			} else {
				c.setParam(3, 0)
			}
			c.currentPos += 4
		case 8: // equals
			p1 := c.getParam(1, modes[0])
			p2 := c.getParam(2, modes[1])
			if p1 == p2 {
				c.setParam(3, 1)
			} else {
				c.setParam(3, 0)
			}
			c.currentPos += 4
		case 99: // halt
			return 0, false
		}
	}
}

func (c *amplifier1) getParam(offset, mode int) int {
	if mode == 0 {
		return c.code[c.code[c.currentPos+offset]]
	}
	return c.code[c.currentPos+offset]
}

func (c *amplifier1) setParam(offset, value int) {
	c.code[c.code[c.currentPos+offset]] = value
}

func runAmplifiers(program []int, phases []int) int {
	signal := 0
	for _, phase := range phases {
		comp := newAmplifier(program, []int{phase, signal})
		var ok bool
		signal, ok = comp.run()
		if !ok {
			log.Fatal("unexpected end of program")
		}
	}
	return signal
}

func part1(input string) {
	numbers := conv.ToIntSliceComma(input)
	phases := []int{0, 1, 2, 3, 4}
	maxSignal := 0

	for _, perm := range mathx.Permutations(phases) {
		signal := runAmplifiers(numbers, perm)
		if signal > maxSignal {
			maxSignal = signal
		}
	}
	fmt.Println("Result 1:", maxSignal)
}

type amplifier2 struct {
	code       []int
	currentPos int
	inputs     []int
	halted     bool
}

func newAmplifier2(program []int) *amplifier2 {
	copyProgram := make([]int, len(program))
	copy(copyProgram, program)
	return &amplifier2{code: copyProgram}
}

func (c *amplifier2) run() (int, bool) {
	for {
		opcode := c.code[c.currentPos] % 100
		modes := []int{
			(c.code[c.currentPos] / 100) % 10,
			(c.code[c.currentPos] / 1000) % 10,
		}

		switch opcode {
		case 1: // add
			p1 := c.getParam(1, modes[0])
			p2 := c.getParam(2, modes[1])
			c.setParam(3, p1+p2)
			c.currentPos += 4
		case 2: // multiply
			p1 := c.getParam(1, modes[0])
			p2 := c.getParam(2, modes[1])
			c.setParam(3, p1*p2)
			c.currentPos += 4
		case 3: // input
			if len(c.inputs) == 0 {
				return 0, false
			}
			c.setParam(1, c.inputs[0])
			c.inputs = c.inputs[1:]
			c.currentPos += 2
		case 4: // output
			output := c.getParam(1, modes[0])
			c.currentPos += 2
			return output, true
		case 5: // jump-if-true
			p1 := c.getParam(1, modes[0])
			p2 := c.getParam(2, modes[1])
			if p1 != 0 {
				c.currentPos = p2
			} else {
				c.currentPos += 3
			}
		case 6: // jump-if-false
			p1 := c.getParam(1, modes[0])
			p2 := c.getParam(2, modes[1])
			if p1 == 0 {
				c.currentPos = p2
			} else {
				c.currentPos += 3
			}
		case 7: // less than
			p1 := c.getParam(1, modes[0])
			p2 := c.getParam(2, modes[1])
			if p1 < p2 {
				c.setParam(3, 1)
			} else {
				c.setParam(3, 0)
			}
			c.currentPos += 4
		case 8: // equals
			p1 := c.getParam(1, modes[0])
			p2 := c.getParam(2, modes[1])
			if p1 == p2 {
				c.setParam(3, 1)
			} else {
				c.setParam(3, 0)
			}
			c.currentPos += 4
		case 99: // halt
			c.halted = true
			return 0, false
		}
	}
}

func (c *amplifier2) getParam(offset, mode int) int {
	if mode == 0 {
		return c.code[c.code[c.currentPos+offset]]
	}
	return c.code[c.currentPos+offset]
}

func (c *amplifier2) setParam(offset, value int) {
	c.code[c.code[c.currentPos+offset]] = value
}

func runAmplifiersFeedback(program []int, phases []int) int {
	amps := make([]*amplifier2, 5)
	for i := range amps {
		amps[i] = newAmplifier2(program)
		amps[i].inputs = []int{phases[i]}
	}

	signal := 0
	amps[0].inputs = append(amps[0].inputs, 0)

	for {
		allHalted := true
		for i := range amps {
			if !amps[i].halted {
				allHalted = false
				output, hasOutput := amps[i].run()
				if hasOutput {
					nextAmp := (i + 1) % 5
					amps[nextAmp].inputs = append(amps[nextAmp].inputs, output)
					if nextAmp == 0 {
						signal = output
					}
				}
			}
		}
		if allHalted {
			break
		}
	}
	return signal
}

func part2(input string) {
	numbers := conv.ToIntSliceComma(input)
	phases := []int{5, 6, 7, 8, 9}
	maxSignal := 0

	for _, perm := range mathx.Permutations(phases) {
		signal := runAmplifiersFeedback(numbers, perm)
		if signal > maxSignal {
			maxSignal = signal
		}
	}
	fmt.Println("Result 2:", maxSignal)
}
