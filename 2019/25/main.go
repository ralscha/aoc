package main

import (
	"aoc/2019/intcomputer"
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"math/rand"
	"slices"
	"strings"
)

func main() {
	input, err := download.ReadInput(2019, 25)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}
	part1(input)
}

func part1(input string) {
	program := conv.ToIntSliceComma(input)
	computer := intcomputer.NewIntcodeComputer(program)
	if err := runAutoGame(computer); err != nil {
		log.Fatalf("running game failed: %v", err)
	}
}

type gameState struct {
	inventory []string
	floor     []string
}

func parseOutput(output string) ([]string, []string) {
	var items []string
	var doors []string

	lines := conv.SplitNewline(output)
	inItems := false
	inDoors := false

	for _, line := range lines {
		if strings.HasPrefix(line, "Items here:") {
			inItems = true
			inDoors = false
			continue
		}
		if strings.HasPrefix(line, "Doors here lead:") {
			inItems = false
			inDoors = true
			continue
		}
		if strings.HasPrefix(line, "- ") {
			item := strings.TrimPrefix(line, "- ")
			if inItems {
				items = append(items, item)
			} else if inDoors {
				doors = append(doors, item)
			}
		} else {
			inItems = false
			inDoors = false
		}
	}

	return items, doors
}

func isDangerousItem(item string) bool {
	dangerous := []string{
		"giant electromagnet",
		"infinite loop",
		"photons",
		"escape pod",
		"molten lava",
	}
	return slices.Contains(dangerous, item)
}

func sendCommand(computer *intcomputer.IntcodeComputer, command string) (string, error) {
	if err := computer.AddInput(command + "\n"); err != nil {
		return "", fmt.Errorf("adding input failed: %v", err)
	}

	result, err := computer.ReadString()
	if err != nil {
		return "", fmt.Errorf("reading output failed: %v", err)
	}

	return result.Str, nil
}

func runAutoGame(computer *intcomputer.IntcodeComputer) error {
	state := gameState{
		inventory: []string{},
		floor:     []string{},
	}

	result, err := computer.ReadString()
	if err != nil {
		return fmt.Errorf("reading initial output failed: %v", err)
	}
	output := result.Str

	for {
		items, doors := parseOutput(output)
		if strings.Contains(output, "== Pressure-Sensitive Floor ==") && len(state.inventory) == 8 {
			break
		}

		for _, item := range items {
			if !isDangerousItem(item) {
				output, err = sendCommand(computer, "take "+item)
				if err != nil {
					return err
				}
				state.inventory = append(state.inventory, item)
			}
		}

		if len(doors) > 0 {
			direction := doors[rand.Intn(len(doors))]
			output, err = sendCommand(computer, direction)
			if err != nil {
				return err
			}
		}
	}

	attempts := 0
	for {
		attempts++
		output, err = sendCommand(computer, "north")
		if err != nil {
			return err
		}

		if strings.Contains(output, "heavier") {
			if len(state.floor) > 0 {
				item := state.floor[rand.Intn(len(state.floor))]
				output, err = sendCommand(computer, "take "+item)
				if err != nil {
					return err
				}
				for i, f := range state.floor {
					if f == item {
						state.floor = append(state.floor[:i], state.floor[i+1:]...)
						break
					}
				}
				state.inventory = append(state.inventory, item)
			}
		} else if strings.Contains(output, "lighter") {
			if len(state.inventory) > 0 {
				item := state.inventory[rand.Intn(len(state.inventory))]
				output, err = sendCommand(computer, "drop "+item)
				if err != nil {
					return err
				}

				for i, inv := range state.inventory {
					if inv == item {
						state.inventory = append(state.inventory[:i], state.inventory[i+1:]...)
						break
					}
				}
				state.floor = append(state.floor, item)
			}
		} else {
			fmt.Println("Correct items in inventory:")
			for _, item := range state.inventory {
				fmt.Printf("- %s\n", item)
			}

			fmt.Println("Output:", output)
			return nil
		}
	}
}
