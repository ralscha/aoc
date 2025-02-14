package main

import (
	"aoc/2019/intcomputer"
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
)

func main() {
	input, err := download.ReadInput(2019, 23)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

type packet struct {
	x, y int
}

type nat struct {
	lastPacket *packet
	lastY      *int
}

func (n *nat) Receive(p packet) {
	n.lastPacket = &p
}

func (n *nat) CheckIdle(queues [][]packet) bool {
	for _, q := range queues {
		if len(q) > 0 {
			return false
		}
	}
	return true
}

func (n *nat) SendToZero(queues [][]packet) bool {
	if n.lastPacket == nil {
		return false
	}

	queues[0] = append(queues[0], *n.lastPacket)

	if n.lastY != nil && *n.lastY == n.lastPacket.y {
		return true
	}
	y := n.lastPacket.y
	n.lastY = &y
	return false
}

func initializeComputers(program []int) ([]*intcomputer.IntcodeComputer, [][]packet) {
	computers := make([]*intcomputer.IntcodeComputer, 50)
	queues := make([][]packet, 50)
	for i := range computers {
		computers[i] = intcomputer.NewIntcodeComputer(program)
		computers[i].AddInput(i)
		queues[i] = make([]packet, 0)
	}
	return computers, queues
}

func runNetwork(computers []*intcomputer.IntcodeComputer, queues [][]packet, nat *nat, part int) {
	idleCount := 0

	for {
		allWaiting := true
		for i := range computers {
			computer := computers[i]

			if len(queues[i]) > 0 {
				p := queues[i][0]
				queues[i] = queues[i][1:]
				computer.AddInput(p.x, p.y)
			}

			result, err := computer.Run()
			if err != nil {
				panic(err)
			}

			if result.Signal == intcomputer.SignalInput {
				computer.AddInput(-1)
			} else if result.Signal == intcomputer.SignalOutput {
				addr := result.Value

				result, err = computer.Run()
				if err != nil {
					panic(err)
				}
				x := result.Value

				result, err = computer.Run()
				if err != nil {
					panic(err)
				}
				y := result.Value

				p := packet{x: x, y: y}

				if addr == 255 {
					if part == 1 {
						fmt.Println("Part 1", y)
						return
					}
					nat.Receive(p)
				} else if addr >= 0 && addr < 50 {
					queues[addr] = append(queues[addr], p)
				}

				allWaiting = false
			}
		}

		if allWaiting {
			idleCount++
			if idleCount > 100 {
				if nat.CheckIdle(queues) {
					if nat.SendToZero(queues) {
						fmt.Println("Part 2", *nat.lastY)
						return
					}
					idleCount = 0
				}
			}
		} else {
			idleCount = 0
		}
	}
}

func part1(input string) {
	program := conv.ToIntSliceComma(input)
	computers, queues := initializeComputers(program)
	runNetwork(computers, queues, nil, 1)
}

func part2(input string) {
	program := conv.ToIntSliceComma(input)
	computers, queues := initializeComputers(program)
	nat := &nat{}
	runNetwork(computers, queues, nat, 2)
}
