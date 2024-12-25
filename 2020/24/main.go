package main

import (
	"aoc/internal/container"
	"aoc/internal/download"
	"aoc/internal/gridutil"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	input, err := download.ReadInput(2020, 24)
	if err != nil {
		panic(err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	lines := strings.Split(input, "\n")
	flipped := container.NewSet[gridutil.Coordinate]()

	re := regexp.MustCompile(`e|se|sw|w|nw|ne`)
	for _, line := range lines {
		matches := re.FindAllString(line, -1)
		c := gridutil.Coordinate{Col: 0, Row: 0}
		for _, move := range matches {
			switch move {
			case "e":
				c.Col++
			case "se":
				c.Row++
			case "sw":
				c.Col--
				c.Row++
			case "w":
				c.Col--
			case "nw":
				c.Row--
			case "ne":
				c.Col++
				c.Row--
			}
		}
		if flipped.Contains(c) {
			flipped.Remove(c)
		} else {
			flipped.Add(c)
		}
	}

	blackCount := flipped.Len()
	fmt.Println("Part 1", blackCount)
}

func part2(input string) {
	lines := strings.Split(input, "\n")
	flipped := container.NewSet[gridutil.Coordinate]()

	re := regexp.MustCompile(`e|se|sw|w|nw|ne`)
	for _, line := range lines {
		matches := re.FindAllString(line, -1)
		c := gridutil.Coordinate{Col: 0, Row: 0}
		for _, move := range matches {
			switch move {
			case "e":
				c.Col++
			case "se":
				c.Row++
			case "sw":
				c.Col--
				c.Row++
			case "w":
				c.Col--
			case "nw":
				c.Row--
			case "ne":
				c.Col++
				c.Row--
			}
		}
		if flipped.Contains(c) {
			flipped.Remove(c)
		} else {
			flipped.Add(c)
		}
	}

	blackTiles := container.NewSet[gridutil.Coordinate]()
	for _, c := range flipped.Values() {
		blackTiles.Add(c)
	}

	for range 100 {
		newBlackTiles := container.NewSet[gridutil.Coordinate]()
		checkTiles := container.NewSet[gridutil.Coordinate]()
		for _, c := range blackTiles.Values() {
			checkTiles.Add(c)
			for _, neighbor := range getNeighbors(c) {
				checkTiles.Add(neighbor)
			}
		}

		for _, c := range checkTiles.Values() {
			blackNeighborCount := 0
			for _, neighbor := range getNeighbors(c) {
				if blackTiles.Contains(neighbor) {
					blackNeighborCount++
				}
			}

			isCurrentlyBlack := blackTiles.Contains(c)
			if isCurrentlyBlack && (blackNeighborCount == 1 || blackNeighborCount == 2) {
				newBlackTiles.Add(c)
			}
			if !isCurrentlyBlack && blackNeighborCount == 2 {
				newBlackTiles.Add(c)
			}
		}
		blackTiles = newBlackTiles
	}

	fmt.Println("Part 2", blackTiles.Len())
}

func getNeighbors(c gridutil.Coordinate) []gridutil.Coordinate {
	return []gridutil.Coordinate{
		{Col: c.Col + 1, Row: c.Row},
		{Col: c.Col, Row: c.Row + 1},
		{Col: c.Col - 1, Row: c.Row + 1},
		{Col: c.Col - 1, Row: c.Row},
		{Col: c.Col, Row: c.Row - 1},
		{Col: c.Col + 1, Row: c.Row - 1},
	}
}
