package main

import (
	"aoc/internal/conv"
	"aoc/internal/mathx"
	"log"
	"os"
	"slices"
)

func main() {
	partI()
	partII()
	partIII()
}

func partI() {
	input := "18\n14\n7\n13\n16\n18\n16\n17\n15\n17\n13"

	nails := conv.ToIntSlice(conv.SplitNewline(input))
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
	lines := conv.SplitNewline(string(input))

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
	lines := conv.SplitNewline(string(input))

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
