package intcomputer

import (
	"fmt"
	"log"
)

type IntcodeComputer struct {
	Memory  []int
	Input   int
	Output  int
	Ip      int
	RelBase int
	Halted  bool
}

func NewIntcodeComputer(program []int) *IntcodeComputer {
	memory := make([]int, len(program)*10)
	copy(memory, program)
	return &IntcodeComputer{
		Memory:  memory,
		Ip:      0,
		RelBase: 0,
		Halted:  false,
	}
}

func (c *IntcodeComputer) Reset() {
	c.Ip = 0
	c.RelBase = 0
	c.Halted = false
	// Keep the original memory
}

func (c *IntcodeComputer) getParam(mode int, offset int) int {
	switch mode {
	case 0: // Position mode
		return c.Memory[c.Memory[c.Ip+offset]]
	case 1: // Immediate mode
		return c.Memory[c.Ip+offset]
	case 2: // Relative mode
		return c.Memory[c.RelBase+c.Memory[c.Ip+offset]]
	default:
		panic(fmt.Sprintf("invalid parameter mode: %d", mode))
	}
}

func (c *IntcodeComputer) setParam(mode int, offset int, value int) {
	switch mode {
	case 0: // Position mode
		c.Memory[c.Memory[c.Ip+offset]] = value
	case 2: // Relative mode
		c.Memory[c.RelBase+c.Memory[c.Ip+offset]] = value
	default:
		panic(fmt.Sprintf("invalid parameter mode for write: %d", mode))
	}
}

func (c *IntcodeComputer) Run() int {
	for {
		opcode := c.Memory[c.Ip] % 100
		mode1 := (c.Memory[c.Ip] / 100) % 10
		mode2 := (c.Memory[c.Ip] / 1000) % 10
		mode3 := (c.Memory[c.Ip] / 10000) % 10

		switch opcode {
		case 1: // Add
			param1 := c.getParam(mode1, 1)
			param2 := c.getParam(mode2, 2)
			c.setParam(mode3, 3, param1+param2)
			c.Ip += 4
		case 2: // Multiply
			param1 := c.getParam(mode1, 1)
			param2 := c.getParam(mode2, 2)
			c.setParam(mode3, 3, param1*param2)
			c.Ip += 4
		case 3: // Input
			c.setParam(mode1, 1, c.Input)
			c.Ip += 2
		case 4: // Output
			c.Output = c.getParam(mode1, 1)
			c.Ip += 2
			return c.Output
		case 5: // Jump-if-true
			if c.getParam(mode1, 1) != 0 {
				c.Ip = c.getParam(mode2, 2)
			} else {
				c.Ip += 3
			}
		case 6: // Jump-if-false
			if c.getParam(mode1, 1) == 0 {
				c.Ip = c.getParam(mode2, 2)
			} else {
				c.Ip += 3
			}
		case 7: // Less than
			param1 := c.getParam(mode1, 1)
			param2 := c.getParam(mode2, 2)
			if param1 < param2 {
				c.setParam(mode3, 3, 1)
			} else {
				c.setParam(mode3, 3, 0)
			}
			c.Ip += 4
		case 8: // Equals
			param1 := c.getParam(mode1, 1)
			param2 := c.getParam(mode2, 2)
			if param1 == param2 {
				c.setParam(mode3, 3, 1)
			} else {
				c.setParam(mode3, 3, 0)
			}
			c.Ip += 4
		case 9: // Adjust relative base
			c.RelBase += c.getParam(mode1, 1)
			c.Ip += 2
		case 99: // Halt
			c.Halted = true
			return -1
		default:
			log.Fatalf("unknown opcode: %d at position %d", opcode, c.Ip)
		}
	}
}
