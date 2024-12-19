package main

import (
	"aoc/internal/container"
	"aoc/internal/download"
	"aoc/internal/gridutil"
	"crypto/md5"
	"fmt"
	"log"
)

type state struct {
	pos  gridutil.Coordinate
	path string
}

func getOpenDoors(path string, salt string) [4]bool {
	hash := fmt.Sprintf("%x", md5.Sum([]byte(salt+path)))
	open := [4]bool{}
	for i := 0; i < 4; i++ {
		if hash[i] >= 'b' && hash[i] <= 'f' {
			open[i] = true
		}
	}
	return open
}

func main() {
	input, err := download.ReadInput(2016, 17)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func findPaths(salt string, findLongest bool) string {
	queue := container.NewQueue[state]()
	queue.Push(state{pos: gridutil.Coordinate{Row: 0, Col: 0}, path: ""})
	target := gridutil.Coordinate{Row: 3, Col: 3}
	longest := 0
	shortestPath := ""

	dirs := []struct {
		dir  gridutil.Direction
		char string
	}{
		{gridutil.DirectionN, "U"},
		{gridutil.DirectionS, "D"},
		{gridutil.DirectionW, "L"},
		{gridutil.DirectionE, "R"},
	}

	for !queue.IsEmpty() {
		curr := queue.Pop()

		if curr.pos == target {
			if findLongest {
				if len(curr.path) > longest {
					longest = len(curr.path)
				}
				continue
			} else {
				return curr.path
			}
		}

		open := getOpenDoors(curr.path, salt)
		for i, d := range dirs {
			if !open[i] {
				continue
			}

			newPos := gridutil.Coordinate{Row: curr.pos.Row + d.dir.Row, Col: curr.pos.Col + d.dir.Col}
			if newPos.Row < 0 || newPos.Row > 3 || newPos.Col < 0 || newPos.Col > 3 {
				continue
			}

			queue.Push(state{
				pos:  newPos,
				path: curr.path + d.char,
			})
		}
	}

	if findLongest {
		return fmt.Sprintf("%d", longest)
	}
	return shortestPath
}

func part1(input string) {
	fmt.Println("Part 1:", findPaths(input, false))
}

func part2(input string) {
	fmt.Println("Part 2:", findPaths(input, true))
}
