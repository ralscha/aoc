package intcomputer

import (
	"fmt"
)

type Signal string

const (
	SignalInput  Signal = "input"
	SignalOutput Signal = "output"
	SignalEnd    Signal = "end"
)

type Result struct {
	Signal Signal
	Value  int
	Str    string
}

type State struct {
	Memory     []int
	Ip         int
	RelBase    int
	InputQueue []int
}

type IntcodeComputer struct {
	Memory     []int
	InputQueue []int
	Ip         int
	RelBase    int
}

func NewIntcodeComputer(program []int) *IntcodeComputer {
	memory := make([]int, len(program)*10)
	copy(memory, program)
	return &IntcodeComputer{
		Memory:     memory,
		InputQueue: make([]int, 0),
		Ip:         0,
		RelBase:    0,
	}
}

func (c *IntcodeComputer) Clone() *IntcodeComputer {
	memory := make([]int, len(c.Memory))
	copy(memory, c.Memory)
	inputQueue := make([]int, len(c.InputQueue))
	copy(inputQueue, c.InputQueue)
	return &IntcodeComputer{
		Memory:     memory,
		InputQueue: inputQueue,
		Ip:         c.Ip,
		RelBase:    c.RelBase,
	}
}

func (c *IntcodeComputer) GetState() State {
	memory := make([]int, len(c.Memory))
	copy(memory, c.Memory)
	inputQueue := make([]int, len(c.InputQueue))
	copy(inputQueue, c.InputQueue)
	return State{
		Memory:     memory,
		Ip:         c.Ip,
		RelBase:    c.RelBase,
		InputQueue: inputQueue,
	}
}

func (c *IntcodeComputer) AddInput(inputs ...interface{}) error {
	for _, input := range inputs {
		switch v := input.(type) {
		case int:
			c.InputQueue = append(c.InputQueue, v)
		case string:
			for _, ch := range v {
				c.InputQueue = append(c.InputQueue, int(ch))
			}
		default:
			return fmt.Errorf("unsupported input type: %T", input)
		}
	}
	return nil
}

func (c *IntcodeComputer) ReadString() (*Result, error) {
	result, err := c.Run()
	if err != nil {
		return nil, err
	}

	var str string
	for result.Signal == SignalOutput && result.Value > 0 && result.Value <= 255 {
		str += string(rune(result.Value))
		result, err = c.Run()
		if err != nil {
			return nil, err
		}
	}

	result.Str = str
	return result, nil
}

func (c *IntcodeComputer) Reset() {
	c.Ip = 0
	c.RelBase = 0
	c.InputQueue = c.InputQueue[:0]
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

func (c *IntcodeComputer) Run() (*Result, error) {
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
			if len(c.InputQueue) == 0 {
				return &Result{Signal: SignalInput}, nil
			}
			c.setParam(mode1, 1, c.InputQueue[0])
			c.InputQueue = c.InputQueue[1:]
			c.Ip += 2
		case 4: // Output
			output := c.getParam(mode1, 1)
			c.Ip += 2
			return &Result{Signal: SignalOutput, Value: output}, nil
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
			return &Result{Signal: SignalEnd}, nil
		default:
			return nil, fmt.Errorf("unknown opcode: %d at position %d", opcode, c.Ip)
		}
	}
}
