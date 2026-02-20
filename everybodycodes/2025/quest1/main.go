package main

import (
	"aoc/internal/conv"
	"fmt"
	"strings"
)

type instruction struct {
	turn  string
	steps int
}

func main() {
	partI()
	partII()
	partIII()
}

func partI() {
	inputNames := "Sarnselor,Lirisis,Harnfeth,Havxelor,Araktaril,Gorathzeth,Jorathgaz,Shaelilor,Mardrith,Urakxal"
	inputInstructions := "L4,R4,L7,R8,L2,R8,L1,R9,L3,R2,L5"

	names := strings.Split(inputNames, ",")
	instructions := parseInstructions(inputInstructions)

	currentIndex := 0
	for _, instr := range instructions {
		switch instr.turn {
		case "L":
			currentIndex = max(currentIndex-instr.steps, 0)
		case "R":
			currentIndex = currentIndex + instr.steps
			if currentIndex >= len(names) {
				currentIndex = len(names) - 1
			}
		}
	}
	fmt.Println(names[currentIndex])
}

func partII() {
	inputNames := "Zalindor,Selthyris,Felmarthar,Jaersaral,Shaelzor,Zyrixluth,Thazloris,Ryssvyr,Felmariral,Ignmyr,Pylarxal,Cynvardaros,Hyralyr,Shaemidris,Lirnarel,Iskardra,Thyroslar,Felmarmir,Zraalfarin,Malithvor"
	inputInstructions := "L11,R18,L15,R16,L19,R5,L13,R8,L9,R16,L5,R5,L5,R9,L5,R17,L5,R10,L5,R19,L15,R18,L18,R7,L12,R18,L11,R7,L8"

	names := strings.Split(inputNames, ",")
	instructions := parseInstructions(inputInstructions)

	currentIndex := 0
	for _, instr := range instructions {
		switch instr.turn {
		case "L":
			currentIndex = (currentIndex - instr.steps) % len(names)
			if currentIndex < 0 {
				currentIndex += len(names)
			}
		case "R":
			currentIndex = (currentIndex + instr.steps) % len(names)
		}
	}
	fmt.Println(names[currentIndex])
}

func partIII() {
	inputNames := "Drelxelor,Ulmargnar,Iskarmirix,Darvor,Vanox,Zyrxal,Calath,Arkluth,Fersyron,Naldvalir,Iskarrax,Hyraxaril,Tharnnyn,Oronloris,Palthsin,Drelryn,Valimar,Helardith,Lorfeth,Hazor,Ulsyx,Lithural,Pylarketh,Paldthys,Tarlfelix,Faljor,Zarfal,Talindor,Ascaljoris,Elarryn"
	inputInstructions := "L17,R37,L48,R25,L49,R30,L37,R5,L19,R15,L18,R31,L18,R36,L47,R30,L20,R16,L13,R49,L5,R20,L5,R38,L5,R10,L5,R22,L5,R40,L5,R35,L5,R18,L5,R46,L5,R15,L5,R29,L28,R46,L27,R26,L45,R49,L26,R25,L14,R21,L5,R13,L32,R29,L45,R25,L5,R22,L17"
	names := strings.Split(inputNames, ",")
	instructions := parseInstructions(inputInstructions)

	currentIndex := 0
	for _, instr := range instructions {
		switch instr.turn {
		case "L":
			swapIndex := (currentIndex - instr.steps) % len(names)
			if swapIndex < 0 {
				swapIndex += len(names)
			}
			names[currentIndex], names[swapIndex] = names[swapIndex], names[currentIndex]
		case "R":
			swapIndex := (currentIndex + instr.steps) % len(names)
			names[currentIndex], names[swapIndex] = names[swapIndex], names[currentIndex]
		}
	}
	fmt.Println(names[0])
}

func parseInstructions(input string) []instruction {
	parts := strings.Split(input, ",")
	instructions := make([]instruction, len(parts))

	for i, part := range parts {
		turn := string(part[0])
		steps := conv.MustAtoi(part[1:])
		instructions[i] = instruction{turn: turn, steps: steps}
	}
	return instructions
}
