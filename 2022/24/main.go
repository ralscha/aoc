package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/gridutil"
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
	return valley, gridutil.Coordinate{Row: height - 1, Col: width - 2}
}

const (
	up = iota
	down
	left
	right
)

var directions = []gridutil.Direction{
	{Row: -1, Col: 0}, // up
	{Row: 1, Col: 0},  // down
	{Row: 0, Col: -1}, // left
	{Row: 0, Col: 1},  // right
}

type path struct {
	pos    gridutil.Coordinate
	minute int
}

type valley struct {
	width     int
	height    int
	blizzards []blizzard
}

type blizzard struct {
	startPos gridutil.Coordinate
	dir      int
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
	for _, b := range v.blizzards {
		bPos := b.posAtMinute(minute, *v)
		if bPos == nextPos {
			return pos, false
		}
	}
	return nextPos, true
}

func (v *valley) findFastestPath(minute int, fromPos, toPos gridutil.Coordinate) int {
	paths := container.NewQueue[path]()
	paths.Push(path{pos: fromPos, minute: minute})
	existingPaths := container.NewSet[path]()

	for !paths.IsEmpty() {
		p := paths.Pop()
		p.minute++

		// Try all directions
		for _, dir := range directions {
			nextPos, ok := v.moveExpedition(p.minute, p.pos, dir, toPos)
			if !ok {
				continue
			}
			if nextPos == toPos {
				return p.minute
			}
			nextPath := path{pos: nextPos, minute: p.minute}
			if !existingPaths.Contains(nextPath) {
				existingPaths.Add(nextPath)
				paths.Push(nextPath)
			}
		}

		// Try waiting in place
		nextPos, ok := v.moveExpedition(p.minute, p.pos, gridutil.Direction{}, toPos)
		if ok {
			nextPath := path{pos: nextPos, minute: p.minute}
			if !existingPaths.Contains(nextPath) {
				existingPaths.Add(nextPath)
				paths.Push(nextPath)
			}
		}
	}
	return -1
}
