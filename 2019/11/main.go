package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	grid2 "aoc/internal/gridutil"
	"fmt"
	"log"
)

func main() {
	input, err := download.ReadInput(2019, 11)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

type intCodeComputer struct {
	memory       map[int64]int64
	ip           int64
	relativeBase int64
	input        int64
	halted       bool
}

func newIntcodeComputer(program []int64) *intCodeComputer {
	memory := make(map[int64]int64)
	for i, v := range program {
		memory[int64(i)] = v
	}
	return &intCodeComputer{memory: memory}
}

func (comp *intCodeComputer) getValue(addr int64) int64 {
	return comp.memory[addr]
}

func (comp *intCodeComputer) setValue(addr int64, value int64) {
	comp.memory[addr] = value
}

func decodeInstruction(instruction int64) (int64, [3]int64) {
	opcode := instruction % 100
	modes := [3]int64{
		(instruction / 100) % 10,
		(instruction / 1000) % 10,
		(instruction / 10000) % 10,
	}
	return opcode, modes
}

func (comp *intCodeComputer) getParameter(mode int64, offset int64) int64 {
	switch mode {
	case 0:
		return comp.getValue(comp.getValue(comp.ip + offset))
	case 1:
		return comp.getValue(comp.ip + offset)
	case 2:
		return comp.getValue(comp.relativeBase + comp.getValue(comp.ip+offset))
	default:
		log.Fatalf("Invalid parameter mode: %d", mode)
		return 0
	}
}

func (comp *intCodeComputer) getAddress(mode int64, offset int64) int64 {
	switch mode {
	case 0:
		return comp.getValue(comp.ip + offset)
	case 2:
		return comp.relativeBase + comp.getValue(comp.ip+offset)
	default:
		log.Fatalf("Invalid address mode: %d", mode)
		return 0
	}
}

func (comp *intCodeComputer) run() int64 {
	for {
		opcode, modes := decodeInstruction(comp.getValue(comp.ip))

		switch opcode {
		case 1: // Addition
			a := comp.getParameter(modes[0], 1)
			b := comp.getParameter(modes[1], 2)
			dst := comp.getAddress(modes[2], 3)
			comp.setValue(dst, a+b)
			comp.ip += 4
		case 2: // Multiplication
			a := comp.getParameter(modes[0], 1)
			b := comp.getParameter(modes[1], 2)
			dst := comp.getAddress(modes[2], 3)
			comp.setValue(dst, a*b)
			comp.ip += 4
		case 3: // Input
			dst := comp.getAddress(modes[0], 1)
			comp.setValue(dst, comp.input)
			comp.ip += 2
		case 4: // Output
			a := comp.getParameter(modes[0], 1)
			comp.ip += 2
			return a
		case 5: // Jump-if-true
			a := comp.getParameter(modes[0], 1)
			b := comp.getParameter(modes[1], 2)
			if a != 0 {
				comp.ip = b
			} else {
				comp.ip += 3
			}
		case 6: // Jump-if-false
			a := comp.getParameter(modes[0], 1)
			b := comp.getParameter(modes[1], 2)
			if a == 0 {
				comp.ip = b
			} else {
				comp.ip += 3
			}
		case 7: // Less than
			a := comp.getParameter(modes[0], 1)
			b := comp.getParameter(modes[1], 2)
			dst := comp.getAddress(modes[2], 3)
			if a < b {
				comp.setValue(dst, 1)
			} else {
				comp.setValue(dst, 0)
			}
			comp.ip += 4
		case 8: // Equals
			a := comp.getParameter(modes[0], 1)
			b := comp.getParameter(modes[1], 2)
			dst := comp.getAddress(modes[2], 3)
			if a == b {
				comp.setValue(dst, 1)
			} else {
				comp.setValue(dst, 0)
			}
			comp.ip += 4
		case 9: // Adjust relative base
			a := comp.getParameter(modes[0], 1)
			comp.relativeBase += a
			comp.ip += 2
		case 99: // Halt
			comp.halted = true
			return 0
		default:
			log.Fatalf("Unknown opcode: %d", opcode)
		}
	}
}

func part1(input string) {
	code := conv.ToInt64SliceComma(input)
	computer := newIntcodeComputer(code)

	direction := grid2.DirectionN
	grid := grid2.NewGrid2D[int](false)
	currentPos := grid2.Coordinate{Row: 0, Col: 0}

	for !computer.halted {
		color, _ := grid.GetWithCoordinate(currentPos)
		computer.input = int64(color)

		output := computer.run()
		if computer.halted {
			break
		}
		turnDirection := computer.run()
		if computer.halted {
			break
		}

		grid.SetWithCoordinate(currentPos, int(output))
		if turnDirection == 0 {
			direction = grid2.TurnLeft(direction)
		} else {
			direction = grid2.TurnRight(direction)
		}
		currentPos = grid2.Coordinate{
			Row: currentPos.Row + direction.Row,
			Col: currentPos.Col + direction.Col,
		}
	}

	fmt.Println("Result 1", grid.Count())
}

func part2(input string) {
	code := conv.ToInt64SliceComma(input)
	computer := newIntcodeComputer(code)

	direction := grid2.DirectionN
	grid := grid2.NewGrid2D[int](false)
	currentPos := grid2.Coordinate{Row: 0, Col: 0}

	for !computer.halted {
		color, ok := grid.GetWithCoordinate(currentPos)

		// always start with white
		if !ok {
			color = 1
		}
		computer.input = int64(color)

		output := computer.run()
		if computer.halted {
			break
		}
		turnDirection := computer.run()
		if computer.halted {
			break
		}

		grid.SetWithCoordinate(currentPos, int(output))
		if turnDirection == 0 {
			direction = grid2.TurnLeft(direction)
		} else {
			direction = grid2.TurnRight(direction)
		}
		currentPos = grid2.Coordinate{
			Row: currentPos.Row + direction.Row,
			Col: currentPos.Col + direction.Col,
		}
	}

	minCol, maxCol := grid.GetMinMaxCol()
	minRow, maxRow := grid.GetMinMaxRow()

	for row := minRow; row <= maxRow; row++ {
		for col := minCol; col <= maxCol; col++ {
			color, ok := grid.Get(row, col)
			if !ok || color == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print("O")
			}
		}
		fmt.Println()
	}

	fmt.Println("Result 2", grid.Count())
}
