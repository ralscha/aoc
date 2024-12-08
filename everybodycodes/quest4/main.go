package main

import (
	"aoc/internal/conv"
	"aoc/internal/mathx"
	"log"
	"os"
	"slices"
	"strings"
)

func main() {
	partI()
	partII()
	partIII()
}

func partI() {
	input := "18\n14\n7\n13\n16\n18\n16\n17\n15\n17\n13"

	nails := conv.ToIntSlice(strings.Split(input, "\n"))
	minNail := slices.Min(nails)
	strikes := 0
	for _, nail := range nails {
		strikes += nail - minNail
	}
	println(strikes)
}

func partII() {
	input, err := os.ReadFile("everybodycodes/quest4/partII.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(input), "\n")
	for i, l := range lines {
		lines[i] = strings.TrimRight(l, "\r\n")
	}
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	nails := conv.ToIntSlice(lines)
	minNail := slices.Min(nails)
	strikes := 0
	for _, nail := range nails {
		strikes += nail - minNail
	}
	println(strikes)
}

func partIII() {
	input, err := os.ReadFile("everybodycodes/quest4/partIII.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(input), "\n")
	for i, l := range lines {
		lines[i] = strings.TrimRight(l, "\r\n")
	}
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	nails := conv.ToIntSlice(lines)
	sortedNails := make([]int, len(nails))
	copy(sortedNails, nails)
	slices.Sort(sortedNails)

	medianHeight := sortedNails[len(sortedNails)/2]

	strikes := 0
	for _, nail := range nails {
		strikes += mathx.Abs(nail - medianHeight)
	}
	println(strikes)

}
