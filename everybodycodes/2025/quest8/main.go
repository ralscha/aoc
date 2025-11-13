package main

import (
	"aoc/internal/conv"
	"fmt"
	"slices"
)

func main() {
	partI()
	partII()
	partIII()
}

func partI() {
	input := "1,32,16,31,15,29,12,28"
	nails := conv.ToIntSliceComma(input)

	const totalNails = 32
	const half = totalNails / 2

	count := 0
	for i := range len(nails) - 1 {
		nail1 := nails[i]
		nail2 := nails[i+1]

		diff := nail1 - nail2
		if diff < 0 {
			diff = -diff
		}
		if diff == half {
			count++
		}
	}

	fmt.Println(count)
}

type Segment struct {
	a, b int
}

func partII() {
	input := "1,140,9,148,26,147,29,146"
	nails := conv.ToIntSliceComma(input)

	var segments []Segment
	for i := range len(nails) - 1 {
		a, b := nails[i], nails[i+1]
		if a > b {
			a, b = b, a
		}
		segments = append(segments, Segment{a, b})
	}

	knots := 0
	for i := range len(segments) {
		for j := i + 1; j < len(segments); j++ {
			seg1, seg2 := segments[i], segments[j]
			if (seg1.a < seg2.a && seg2.a < seg1.b && seg1.b < seg2.b) ||
				(seg2.a < seg1.a && seg1.a < seg2.b && seg2.b < seg1.b) {
				knots++
			}
		}
	}

	fmt.Println(knots)
}

func partIII() {
	input := "1,85,199,63,187,64"
	nails := conv.ToIntSliceComma(input)

	const totalNails = 256

	links := make(map[int][]int)
	for i := range len(nails) - 1 {
		a, b := nails[i], nails[i+1]
		links[a] = append(links[a], b)
		links[b] = append(links[b], a)
	}

	bestCuts := 0
	for a := 1; a <= totalNails; a++ {
		cuts := 0
		for b := a + 2; b <= totalNails; b++ {
			for _, c := range links[b] {
				if a < c && c < b-1 {
					cuts--
				}
			}
			for _, c := range links[b-1] {
				if !(a <= c && c <= b) {
					cuts++
				}
			}
			directLink := 0
			if slices.Contains(links[a], b) {
				directLink = 1
			}
			if cuts+directLink > bestCuts {
				bestCuts = cuts + directLink
			}
		}
	}

	fmt.Println(bestCuts)
}
