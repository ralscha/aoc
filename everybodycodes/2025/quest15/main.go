package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"aoc/internal/gridutil"
	"aoc/internal/mathx"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	partI()
	partII()
	partIII()
}

func solve(filename string) int {
	data, _ := os.ReadFile(filename)
	input := strings.TrimSpace(string(data))
	instructions := strings.Split(input, ",")
	grid := gridutil.NewGrid2D[rune](false)
	pos := gridutil.Coordinate{Row: 0, Col: 0}
	dir := gridutil.DirectionE

	grid.SetC(pos, 'S')

	for _, inst := range instructions {
		turn := inst[0]
		length := conv.MustAtoi(inst[1:])

		switch turn {
		case 'L':
			dir = gridutil.TurnLeft(dir)
		case 'R':
			dir = gridutil.TurnRight(dir)
		}

		for range length {
			pos = gridutil.Coordinate{Row: pos.Row + dir.Row, Col: pos.Col + dir.Col}
			grid.SetC(pos, '#')
		}
	}

	grid.SetC(pos, 'E')

	var start gridutil.Coordinate
	minRow, maxRow := grid.GetMinMaxRow()
	minCol, maxCol := grid.GetMinMaxCol()
	for r := minRow; r <= maxRow; r++ {
		for c := minCol; c <= maxCol; c++ {
			if val, ok := grid.Get(r, c); ok && val == 'S' {
				start = gridutil.Coordinate{Row: r, Col: c}
				break
			}
		}
	}

	type state struct {
		pos  gridutil.Coordinate
		dist int
	}

	queue := container.NewQueue[state]()
	visited := container.NewSet[gridutil.Coordinate]()

	queue.Push(state{pos: start, dist: 0})
	visited.Add(start)

	directions := gridutil.Get4Directions()

	for !queue.IsEmpty() {
		curr := queue.Pop()

		if val, ok := grid.GetC(curr.pos); ok && val == 'E' {
			return curr.dist
		}

		for _, d := range directions {
			next := gridutil.Coordinate{Row: curr.pos.Row + d.Row, Col: curr.pos.Col + d.Col}

			if visited.Contains(next) {
				continue
			}

			if next.Row < minRow || next.Row > maxRow || next.Col < minCol || next.Col > maxCol {
				continue
			}

			val, ok := grid.GetC(next)
			if ok && (val == '#' || val == 'E') {
				if val == 'E' {
					return curr.dist + 1
				}
				continue
			}

			visited.Add(next)
			queue.Push(state{pos: next, dist: curr.dist + 1})
		}
	}

	return -1
}

func solveCompressed(filename string) int {
	data, _ := os.ReadFile(filename)
	input := strings.TrimSpace(string(data))

	instructions := strings.Split(input, ",")

	xs := []int{}
	ys := []int{}

	pos := gridutil.Coordinate{Row: 0, Col: 0}
	dir := gridutil.DirectionE
	for _, inst := range instructions {
		turn := inst[0]
		length := conv.MustAtoi(inst[1:])

		switch turn {
		case 'L':
			dir = gridutil.TurnLeft(dir)
		case 'R':
			dir = gridutil.TurnRight(dir)
		}

		next := gridutil.Coordinate{
			Row: pos.Row + dir.Row*length,
			Col: pos.Col + dir.Col*length,
		}

		before := gridutil.Coordinate{Row: pos.Row - dir.Row, Col: pos.Col - dir.Col}
		after := gridutil.Coordinate{Row: next.Row + dir.Row, Col: next.Col + dir.Col}

		if dir.Col != 0 {
			xs = append(xs, before.Col, pos.Col, next.Col, after.Col)
			ys = append(ys, pos.Row-1, pos.Row, pos.Row+1)
		} else {
			xs = append(xs, pos.Col-1, pos.Col, pos.Col+1)
			ys = append(ys, before.Row, pos.Row, next.Row, after.Row)
		}

		pos = next
	}

	slices.Sort(xs)
	xs = slices.Compact(xs)
	slices.Sort(ys)
	ys = slices.Compact(ys)

	xIndex := make(map[int]int)
	for i, x := range xs {
		xIndex[x] = i
	}
	yIndex := make(map[int]int)
	for i, y := range ys {
		yIndex[y] = i
	}

	walls := container.NewSet[gridutil.Coordinate]()

	pos = gridutil.Coordinate{Row: 0, Col: 0}
	dir = gridutil.DirectionE

	for _, inst := range instructions {
		turn := inst[0]
		length := conv.MustAtoi(inst[1:])

		switch turn {
		case 'L':
			dir = gridutil.TurnLeft(dir)
		case 'R':
			dir = gridutil.TurnRight(dir)
		}

		next := gridutil.Coordinate{
			Row: pos.Row + dir.Row*length,
			Col: pos.Col + dir.Col*length,
		}

		if dir.Col != 0 {
			y := pos.Row
			xStart := xIndex[pos.Col]
			xEnd := xIndex[next.Col]
			if xStart > xEnd {
				xStart, xEnd = xEnd, xStart
			}
			for xi := xStart; xi <= xEnd; xi++ {
				walls.Add(gridutil.Coordinate{Row: y, Col: xs[xi]})
			}
		} else {
			x := pos.Col
			yStart := yIndex[pos.Row]
			yEnd := yIndex[next.Row]
			if yStart > yEnd {
				yStart, yEnd = yEnd, yStart
			}
			for yi := yStart; yi <= yEnd; yi++ {
				walls.Add(gridutil.Coordinate{Row: ys[yi], Col: x})
			}
		}

		pos = next
	}

	goal := pos
	walls.Remove(goal)

	type state struct {
		xi, yi int
		dist   int
	}

	start := gridutil.Coordinate{Row: 0, Col: 0}
	startXi := xIndex[start.Col]
	startYi := yIndex[start.Row]

	queue := container.NewQueue[state]()
	visited := container.NewSet[gridutil.Coordinate]()

	queue.Push(state{xi: startXi, yi: startYi, dist: 0})
	visited.Add(start)

	deltas := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for !queue.IsEmpty() {
		curr := queue.Pop()
		currPos := gridutil.Coordinate{Row: ys[curr.yi], Col: xs[curr.xi]}

		if currPos.Row == goal.Row && currPos.Col == goal.Col {
			return curr.dist
		}

		for _, d := range deltas {
			nxi := curr.xi + d[0]
			nyi := curr.yi + d[1]

			if nxi < 0 || nxi >= len(xs) || nyi < 0 || nyi >= len(ys) {
				continue
			}

			nextPos := gridutil.Coordinate{Row: ys[nyi], Col: xs[nxi]}

			if visited.Contains(nextPos) {
				continue
			}

			if walls.Contains(nextPos) {
				continue
			}

			visited.Add(nextPos)
			dist := mathx.Abs(nextPos.Row-currPos.Row) + mathx.Abs(nextPos.Col-currPos.Col)
			queue.Push(state{xi: nxi, yi: nyi, dist: curr.dist + dist})
		}
	}

	return -1
}

func partI() {
	fmt.Println(solve("input1"))
}

func partII() {
	fmt.Println(solve("input2"))
}

func partIII() {
	fmt.Println(solveCompressed("input3"))
}
