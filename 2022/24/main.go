package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/gridutil"
	"aoc/internal/mathx"
	"fmt"
	"log"
)

func main() {
	input, err := download.ReadInput(2022, 24)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	valley, endPos := makeValley(lines)
	startPos := gridutil.Coordinate{Row: 0, Col: 1}
	minutes := valley.findFastestPath(0, startPos, endPos)
	fmt.Println(minutes)
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	valley, endPos := makeValley(lines)
	startPos := gridutil.Coordinate{Row: 0, Col: 1}
	minutes := valley.findFastestPath(0, startPos, endPos)
	minutes = valley.findFastestPath(minutes+1, endPos, startPos)
	minutes = valley.findFastestPath(minutes+1, startPos, endPos)
	fmt.Println(minutes)
}

func makeValley(lines []string) (*valley, gridutil.Coordinate) {
	var blizzards []blizzard
	width, height := len(lines[0]), len(lines)
	for row, line := range lines {
		for col, c := range line {
			pos := gridutil.Coordinate{Row: row, Col: col}
			switch c {
			case '>':
				blizzards = append(blizzards, blizzard{startPos: pos, dir: right})
			case '<':
				blizzards = append(blizzards, blizzard{startPos: pos, dir: left})
			case '^':
				blizzards = append(blizzards, blizzard{startPos: pos, dir: up})
			case 'v':
				blizzards = append(blizzards, blizzard{startPos: pos, dir: down})
			}
		}
	}

	valley := &valley{
		width:     width,
		height:    height,
		blizzards: blizzards,
	}
	valley.initBlizzardCache()
	return valley, gridutil.Coordinate{Row: height - 1, Col: width - 2}
}

const (
	up = iota
	down
	left
	right
)

type valley struct {
	width         int
	height        int
	blizzards     []blizzard
	blizzardCache []*container.Set[gridutil.Coordinate]
	cycleLength   int
}

type blizzard struct {
	startPos gridutil.Coordinate
	dir      int
}

func (v *valley) initBlizzardCache() {
	// Calculate cycle length (LCM of width-2 and height-2)
	v.cycleLength = mathx.Lcm([]int{v.width - 2, v.height - 2})
	v.blizzardCache = make([]*container.Set[gridutil.Coordinate], v.cycleLength)

	// Pre-calculate all blizzard positions for the entire cycle
	for minute := 0; minute < v.cycleLength; minute++ {
		positions := container.NewSet[gridutil.Coordinate]()
		for _, b := range v.blizzards {
			pos := b.posAtMinute(minute, *v)
			positions.Add(pos)
		}
		v.blizzardCache[minute] = positions
	}
}

func (b *blizzard) posAtMinute(minute int, v valley) gridutil.Coordinate {
	switch b.dir {
	case up:
		move := minute % (v.height - 2)
		newRow := b.startPos.Row - move
		if newRow <= 0 {
			newRow = v.height - 2 + newRow
		}
		return gridutil.Coordinate{Row: newRow, Col: b.startPos.Col}
	case down:
		move := minute % (v.height - 2)
		newRow := b.startPos.Row + move
		if newRow >= v.height-1 {
			newRow = newRow - (v.height - 2)
		}
		return gridutil.Coordinate{Row: newRow, Col: b.startPos.Col}
	case left:
		move := minute % (v.width - 2)
		newCol := b.startPos.Col - move
		if newCol <= 0 {
			newCol = v.width - 2 + newCol
		}
		return gridutil.Coordinate{Row: b.startPos.Row, Col: newCol}
	case right:
		move := minute % (v.width - 2)
		newCol := b.startPos.Col + move
		if newCol >= v.width-1 {
			newCol = newCol - (v.width - 2)
		}
		return gridutil.Coordinate{Row: b.startPos.Row, Col: newCol}
	}
	panic("invalid direction")
}

func (v *valley) isBlizzardAt(pos gridutil.Coordinate, minute int) bool {
	return v.blizzardCache[minute%v.cycleLength].Contains(pos)
}

func (v *valley) moveExpedition(minute int, pos gridutil.Coordinate, dir gridutil.Direction, toPos gridutil.Coordinate) (gridutil.Coordinate, bool) {
	nextPos := gridutil.Coordinate{
		Row: pos.Row + dir.Row,
		Col: pos.Col + dir.Col,
	}

	if nextPos == toPos {
		return toPos, true
	}

	if nextPos.Row < 1 || nextPos.Row >= v.height-1 || nextPos.Col < 1 || nextPos.Col >= v.width-1 {
		return pos, false
	}

	if v.isBlizzardAt(nextPos, minute) {
		return pos, false
	}

	return nextPos, true
}

type state struct {
	pos    gridutil.Coordinate
	minute int
}

func (v *valley) findFastestPath(minute int, fromPos, toPos gridutil.Coordinate) int {
	queue := container.NewQueue[state]()
	queue.Push(state{pos: fromPos, minute: minute})
	seen := container.NewSet[state]()

	for !queue.IsEmpty() {
		current := queue.Pop()
		nextMinute := current.minute + 1

		// Try all directions
		for _, dir := range gridutil.Get4Directions() {
			nextPos, ok := v.moveExpedition(nextMinute, current.pos, dir, toPos)
			if !ok {
				continue
			}
			if nextPos == toPos {
				return nextMinute
			}
			nextState := state{pos: nextPos, minute: nextMinute % v.cycleLength}
			if !seen.Contains(nextState) {
				seen.Add(nextState)
				queue.Push(state{pos: nextPos, minute: nextMinute})
			}
		}

		// Try waiting in place
		nextPos, ok := v.moveExpedition(nextMinute, current.pos, gridutil.Direction{}, toPos)
		if ok {
			nextState := state{pos: nextPos, minute: nextMinute % v.cycleLength}
			if !seen.Contains(nextState) {
				seen.Add(nextState)
				queue.Push(state{pos: nextPos, minute: nextMinute})
			}
		}
	}
	return -1
}
