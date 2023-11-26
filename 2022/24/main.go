package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"container/list"
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
	startPos := position{row: 0, col: 1}
	minutes := valley.findFastestPath(0, startPos, endPos)
	fmt.Println(minutes)
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	valley, endPos := makeValley(lines)
	startPos := position{row: 0, col: 1}
	minutes := valley.findFastestPath(0, startPos, endPos)
	minutes = valley.findFastestPath(minutes+1, endPos, startPos)
	minutes = valley.findFastestPath(minutes+1, startPos, endPos)
	fmt.Println(minutes)
}

func makeValley(lines []string) (*valley, position) {
	var blizzards []blizzard
	width, height := len(lines[0]), len(lines)
	for row, line := range lines {
		for col, c := range line {
			pos := position{row, col}
			if c == '>' {
				blizzards = append(blizzards, blizzard{startPos: pos, dir: right})
			} else if c == '<' {
				blizzards = append(blizzards, blizzard{startPos: pos, dir: left})
			} else if c == '^' {
				blizzards = append(blizzards, blizzard{startPos: pos, dir: up})
			} else if c == 'v' {
				blizzards = append(blizzards, blizzard{startPos: pos, dir: down})
			}
		}
	}

	valley := &valley{
		width:     width,
		height:    height,
		blizzards: blizzards,
	}
	return valley, position{row: height - 1, col: width - 2}
}

const (
	up = iota
	down
	left
	right
)

type position struct {
	row, col int
}

type path struct {
	pos    position
	minute int
}

type valley struct {
	width     int
	height    int
	blizzards []blizzard
}

type blizzard struct {
	startPos position
	dir      int
}

func (b *blizzard) posAtMinute(minute int, v valley) position {
	switch b.dir {
	case up:
		move := minute % (v.height - 2)
		newRow := b.startPos.row - move
		if newRow <= 0 {
			newRow = v.height - 2 + newRow
		}
		return position{row: newRow, col: b.startPos.col}
	case down:
		move := minute % (v.height - 2)
		newRow := b.startPos.row + move
		if newRow >= v.height-1 {
			newRow = newRow - (v.height - 2)
		}
		return position{row: newRow, col: b.startPos.col}
	case left:
		move := minute % (v.width - 2)
		newCol := b.startPos.col - move
		if newCol <= 0 {
			newCol = v.width - 2 + newCol
		}
		return position{row: b.startPos.row, col: newCol}
	case right:
		move := minute % (v.width - 2)
		newCol := b.startPos.col + move
		if newCol >= v.width-1 {
			newCol = newCol - (v.width - 2)
		}
		return position{row: b.startPos.row, col: newCol}
	}
	panic("invalid direction")
}

func (v *valley) moveExpedition(minute int, pos position, dir int, toPos position) (position, bool) {
	nextRow, nextCol := pos.row, pos.col
	switch dir {
	case up:
		nextRow--
	case down:
		nextRow++
	case left:
		nextCol--
	case right:
		nextCol++
	}

	nextPos := position{row: nextRow, col: nextCol}

	if nextPos == toPos {
		return toPos, true
	}

	if nextRow < 1 || nextRow >= v.height-1 || nextCol < 1 || nextCol >= v.width-1 {
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

func (v *valley) findFastestPath(minute int, fromPos, toPos position) int {
	paths := list.New()
	paths.PushBack(path{pos: fromPos, minute: minute})
	existingPaths := make(map[path]bool)

	for paths.Len() > 0 {
		e := paths.Front()
		p := e.Value.(path)
		paths.Remove(e)

		p.minute++

		for _, dir := range []int{up, down, left, right} {
			nextPos, ok := v.moveExpedition(p.minute, p.pos, dir, toPos)
			if !ok {
				continue
			}
			if nextPos == toPos {
				return p.minute
			}
			nextPath := path{pos: nextPos, minute: p.minute}
			if !existingPaths[nextPath] {
				existingPaths[nextPath] = true
				paths.PushBack(nextPath)
			}
		}

		nextPos, ok := v.moveExpedition(p.minute, p.pos, -1, toPos)
		if ok {
			nextPath := path{pos: nextPos, minute: p.minute}
			if !existingPaths[nextPath] {
				existingPaths[nextPath] = true
				paths.PushBack(nextPath)
			}
		}

	}
	return -1
}
