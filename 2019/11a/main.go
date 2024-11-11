package main

import (
	"aoc/internal/download"
	"fmt"
	"log"
	"strconv"
	"strings"
)

// Point represents a position on the grid
type Point struct {
	x, y int
}

// Direction represents the robot's current facing direction
type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

// Robot represents the painting robot
type Robot struct {
	pos       Point
	dir       Direction
	painted   map[Point]int
	program   []int64
	ip        int // instruction pointer
	relBase   int // relative base for relative mode
	memory    map[int]int64
	halted    bool
	nextInput int
}

// NewRobot creates a new robot with the given program
func NewRobot(program []int64) *Robot {
	r := &Robot{
		pos:     Point{0, 0},
		dir:     Up,
		painted: make(map[Point]int),
		program: program,
		memory:  make(map[int]int64),
	}

	// Copy program into memory
	for i, v := range program {
		r.memory[i] = v
	}

	return r
}

// getParam returns the parameter value based on the parameter mode
func (r *Robot) getParam(pos int, mode int) int64 {
	switch mode {
	case 0: // position mode
		return r.memory[int(r.memory[pos])]
	case 1: // immediate mode
		return r.memory[pos]
	case 2: // relative mode
		return r.memory[int(r.memory[pos])+r.relBase]
	default:
		panic(fmt.Sprintf("unknown parameter mode: %d", mode))
	}
}

// getWriteAddr returns the address to write to based on the parameter mode
func (r *Robot) getWriteAddr(pos int, mode int) int {
	if mode == 2 { // relative mode
		return int(r.memory[pos]) + r.relBase
	}
	return int(r.memory[pos])
}

// run executes the Intcode program and returns the next output value
func (r *Robot) runUntilOutput() (int64, bool) {
	for !r.halted {
		opcode := int(r.memory[r.ip])
		op := opcode % 100
		modes := []int{
			(opcode / 100) % 10,
			(opcode / 1000) % 10,
			(opcode / 10000) % 10,
		}

		switch op {
		case 1: // add
			p1 := r.getParam(r.ip+1, modes[0])
			p2 := r.getParam(r.ip+2, modes[1])
			addr := r.getWriteAddr(r.ip+3, modes[2])
			r.memory[addr] = p1 + p2
			r.ip += 4

		case 2: // multiply
			p1 := r.getParam(r.ip+1, modes[0])
			p2 := r.getParam(r.ip+2, modes[1])
			addr := r.getWriteAddr(r.ip+3, modes[2])
			r.memory[addr] = p1 * p2
			r.ip += 4

		case 3: // input
			addr := r.getWriteAddr(r.ip+1, modes[0])
			r.memory[addr] = int64(r.nextInput)
			r.ip += 2

		case 4: // output
			p1 := r.getParam(r.ip+1, modes[0])
			r.ip += 2
			return p1, true

		case 5: // jump-if-true
			p1 := r.getParam(r.ip+1, modes[0])
			p2 := r.getParam(r.ip+2, modes[1])
			if p1 != 0 {
				r.ip = int(p2)
			} else {
				r.ip += 3
			}

		case 6: // jump-if-false
			p1 := r.getParam(r.ip+1, modes[0])
			p2 := r.getParam(r.ip+2, modes[1])
			if p1 == 0 {
				r.ip = int(p2)
			} else {
				r.ip += 3
			}

		case 7: // less than
			p1 := r.getParam(r.ip+1, modes[0])
			p2 := r.getParam(r.ip+2, modes[1])
			addr := r.getWriteAddr(r.ip+3, modes[2])
			if p1 < p2 {
				r.memory[addr] = 1
			} else {
				r.memory[addr] = 0
			}
			r.ip += 4

		case 8: // equals
			p1 := r.getParam(r.ip+1, modes[0])
			p2 := r.getParam(r.ip+2, modes[1])
			addr := r.getWriteAddr(r.ip+3, modes[2])
			if p1 == p2 {
				r.memory[addr] = 1
			} else {
				r.memory[addr] = 0
			}
			r.ip += 4

		case 9: // adjust relative base
			p1 := r.getParam(r.ip+1, modes[0])
			r.relBase += int(p1)
			r.ip += 2

		case 99: // halt
			r.halted = true
			return 0, false

		default:
			panic(fmt.Sprintf("unknown opcode: %d", op))
		}
	}
	return 0, false
}

// turn updates the robot's direction based on the turn instruction
func (r *Robot) turn(turnRight bool) {
	if turnRight {
		r.dir = (r.dir + 1) % 4
	} else {
		r.dir = (r.dir + 3) % 4
	}
}

// move moves the robot forward one panel in its current direction
func (r *Robot) move() {
	switch r.dir {
	case Up:
		r.pos.y++
	case Right:
		r.pos.x++
	case Down:
		r.pos.y--
	case Left:
		r.pos.x--
	}
}

// paint runs the robot and returns the number of panels painted at least once
func (r *Robot) paint() int {
	ix := 0
	for !r.halted {
		// Get current panel color (0 for black, 1 for white)
		r.nextInput = r.painted[r.pos]

		// Get paint color
		paintColor, ok := r.runUntilOutput()
		if !ok {
			break
		}

		// Get turn direction
		turnInstruction, ok := r.runUntilOutput()
		if !ok {
			break
		}

		fmt.Println(ix, paintColor, turnInstruction)
		ix += 1

		// Paint the current panel
		r.painted[r.pos] = int(paintColor)

		// Turn and move
		r.turn(turnInstruction == 1)
		r.move()
	}

	return len(r.painted)
}

func parseInput(input string) []int64 {
	parts := strings.Split(strings.TrimSpace(input), ",")
	program := make([]int64, len(parts))
	for i, p := range parts {
		val, err := strconv.ParseInt(p, 10, 64)
		if err != nil {
			panic(err)
		}
		program[i] = val
	}
	return program
}

func main() {
	input, err := download.ReadInput(2019, 11)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	program := parseInput(input)

	robot := NewRobot(program)
	panelsPainted := robot.paint()

	fmt.Printf("Number of panels painted at least once: %d\n", panelsPainted)
}
