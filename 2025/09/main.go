package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/gridutil"
	"aoc/internal/mathx"
	"fmt"
	"log"
	"sort"
	"strings"
)

func main() {
	input, err := download.ReadInput(2025, 9)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func parsePoints(input string) []gridutil.Coordinate {
	var points []gridutil.Coordinate
	for _, line := range conv.SplitNewline(strings.TrimSpace(input)) {
		nums := conv.ToIntSliceComma(line)
		points = append(points, gridutil.Coordinate{Col: nums[0], Row: nums[1]})
	}
	return points
}

func part1(input string) {
	points := parsePoints(input)
	best := 0
	for i := range points {
		for j := i + 1; j < len(points); j++ {
			dx := mathx.Abs(points[i].Col-points[j].Col) + 1
			dy := mathx.Abs(points[i].Row-points[j].Row) + 1
			area := dx * dy
			if area > best {
				best = area
			}
		}
	}
	fmt.Println("Part 1", best)
}

func part2(input string) {
	points := parsePoints(input)
	n := len(points)

	colSet := container.NewSet[int]()
	rowSet := container.NewSet[int]()
	for _, p := range points {
		colSet.Add(p.Col)
		colSet.Add(p.Col + 1)
		rowSet.Add(p.Row)
		rowSet.Add(p.Row + 1)
	}

	cols := colSet.Values()
	rows := rowSet.Values()
	sort.Ints(cols)
	sort.Ints(rows)

	colIdx := make(map[int]int, len(cols))
	rowIdx := make(map[int]int, len(rows))
	for i, c := range cols {
		colIdx[c] = i
	}
	for i, r := range rows {
		rowIdx[r] = i
	}

	cw := len(cols)
	ch := len(rows)

	onBoundary := make([][]bool, ch)
	for r := range ch {
		onBoundary[r] = make([]bool, cw)
	}

	for i := range n {
		a := points[i]
		b := points[(i+1)%n]
		if a.Row == b.Row {
			ri := rowIdx[a.Row]
			clo, chi := min(colIdx[a.Col], colIdx[b.Col]), max(colIdx[a.Col], colIdx[b.Col])
			for c := clo; c < chi; c++ {
				onBoundary[ri][c] = true
			}
		} else {
			ci := colIdx[a.Col]
			rlo, rhi := min(rowIdx[a.Row], rowIdx[b.Row]), max(rowIdx[a.Row], rowIdx[b.Row])
			for r := rlo; r < rhi; r++ {
				onBoundary[r][ci] = true
			}
		}
	}

	vertBound := make([][]bool, ch)
	for r := range ch {
		vertBound[r] = make([]bool, cw)
	}
	for i := range n {
		a := points[i]
		b := points[(i+1)%n]
		if a.Col == b.Col {
			ci := colIdx[a.Col]
			rlo, rhi := min(rowIdx[a.Row], rowIdx[b.Row]), max(rowIdx[a.Row], rowIdx[b.Row])
			for r := rlo; r < rhi; r++ {
				vertBound[r][ci] = true
			}
		}
	}

	inside := make([][]bool, ch)
	for r := range ch {
		inside[r] = make([]bool, cw)
		crossings := 0
		for c := range cw {
			if vertBound[r][c] {
				crossings++
			}
			if onBoundary[r][c] || crossings%2 == 1 {
				inside[r][c] = true
			}
		}
	}

	prefix := make([][]int64, ch+1)
	for r := 0; r <= ch; r++ {
		prefix[r] = make([]int64, cw+1)
	}
	for r := range ch - 1 {
		h := int64(rows[r+1] - rows[r])
		for c := range cw - 1 {
			w := int64(cols[c+1] - cols[c])
			v := int64(0)
			if inside[r][c] {
				v = w * h
			}
			prefix[r+1][c+1] = v + prefix[r][c+1] + prefix[r+1][c] - prefix[r][c]
		}
	}

	best := int64(0)
	for i := range points {
		for j := i + 1; j < len(points); j++ {
			pi, pj := points[i], points[j]
			if pi.Row == pj.Row || pi.Col == pj.Col {
				continue
			}
			r1, r2 := min(pi.Row, pj.Row), max(pi.Row, pj.Row)
			c1, c2 := min(pi.Col, pj.Col), max(pi.Col, pj.Col)
			area := int64(r2-r1+1) * int64(c2-c1+1)
			if area <= best {
				continue
			}
			cr1, cr2 := rowIdx[r1], rowIdx[r2]
			cc1, cc2 := colIdx[c1], colIdx[c2]
			if prefix[cr2+1][cc2+1]-prefix[cr1][cc2+1]-prefix[cr2+1][cc1]+prefix[cr1][cc1] == area {
				best = area
			}
		}
	}
	fmt.Println("Part 2", best)
}
