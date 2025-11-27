package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/gridutil"
	"fmt"
	"log"

	"aoc/2019/intcomputer"
)

func main() {
	input, err := download.ReadInput(2019, 13)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	program := conv.ToIntSliceComma(input)
	computer := intcomputer.NewIntcodeComputer(program)
	tiles := make(map[gridutil.Coordinate]int)

	var outputs []int
	for {
		result, err := computer.Run()
		if err != nil {
			log.Fatalf("running program failed: %v", err)
		}

		switch result.Signal {
		case intcomputer.SignalOutput:
			outputs = append(outputs, result.Value)
			if len(outputs) == 3 {
				col, row, tileID := outputs[0], outputs[1], outputs[2]
				tiles[gridutil.Coordinate{Row: row, Col: col}] = tileID
				outputs = outputs[:0]
			}
		case intcomputer.SignalEnd:
			blockCount := 0
			for _, tile := range tiles {
				if tile == 2 {
					blockCount++
				}
			}
			fmt.Println("Part 1", blockCount)
			return
		case intcomputer.SignalInput:
			log.Fatal("unexpected input request")
		}
	}
}

func part2(input string) {
	program := conv.ToIntSliceComma(input)
	program[0] = 2
	computer := intcomputer.NewIntcodeComputer(program)
	tiles := make(map[gridutil.Coordinate]int)
	var score int
	var ballCol, paddleCol int
	var outputs []int

	for {
		result, err := computer.Run()
		if err != nil {
			log.Fatalf("running program failed: %v", err)
		}

		switch result.Signal {
		case intcomputer.SignalOutput:
			outputs = append(outputs, result.Value)
			if len(outputs) == 3 {
				col, row, tileID := outputs[0], outputs[1], outputs[2]
				if col == -1 && row == 0 {
					score = tileID
				} else {
					tiles[gridutil.Coordinate{Row: row, Col: col}] = tileID
					switch tileID {
					case 3:
						paddleCol = col
					case 4:
						ballCol = col
					}
				}
				outputs = outputs[:0]
			}
		case intcomputer.SignalInput:
			var joystick int
			if ballCol < paddleCol {
				joystick = -1
			} else if ballCol > paddleCol {
				joystick = 1
			}
			if err := computer.AddInput(joystick); err != nil {
				log.Fatalf("adding input failed: %v", err)
			}
		case intcomputer.SignalEnd:
			fmt.Println("Part 2", score)
			return
		}
	}
}
