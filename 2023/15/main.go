package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"strings"
)

type lens struct {
	label       string
	focalLength int
}

func main() {
	input, err := download.ReadInput(2023, 15)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	input = strings.TrimSpace(input)
	steps := strings.Split(input, ",")

	sum := 0
	for _, step := range steps {
		sum += hash(step)
	}

	fmt.Println(sum)
}

func part2(input string) {
	input = strings.TrimSpace(input)
	boxes := make([][]lens, 256)

	steps := strings.Split(input, ",")
	for _, step := range steps {
		var label string
		var operation byte
		var focalLength int
		parts := strings.Split(step, "=")
		if len(parts) == 2 {
			label = parts[0]
			operation = '='
			focalLength = conv.MustAtoi(parts[1])
		} else {
			label = step[:len(step)-1]
			operation = '-'
			focalLength = 0
		}

		boxIndex := hash(label)
		switch operation {
		case '=':
			replaced := false
			for i, l := range boxes[boxIndex] {
				if l.label == label {
					boxes[boxIndex][i].focalLength = focalLength
					replaced = true
					break
				}
			}
			if !replaced {
				boxes[boxIndex] = append(boxes[boxIndex], lens{label: label, focalLength: focalLength})
			}
		case '-':
			for i, l := range boxes[boxIndex] {
				if l.label == label {
					boxes[boxIndex] = append(boxes[boxIndex][:i], boxes[boxIndex][i+1:]...)
					break
				}
			}
		}
	}

	totalPower := 0
	for boxIndex, box := range boxes {
		for slotIndex, l := range box {
			totalPower += (boxIndex + 1) * (slotIndex + 1) * l.focalLength
		}
	}
	fmt.Println(totalPower)
}

func hash(input string) int {
	currentValue := 0
	for _, char := range input {
		currentValue = ((currentValue + int(char)) * 17) % 256
	}
	return currentValue
}
