package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/geomutil"
	"aoc/internal/gridutil"
	"aoc/internal/mathx"
	"fmt"
	"log"
	"strings"
)

func main() {
	input, err := download.ReadInput(2022, 9)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	lines := conv.SplitNewline(input)

	head := gridutil.Coordinate{Row: 0, Col: 0}
	tail := gridutil.Coordinate{Row: 0, Col: 0}
	visited := container.NewSet[gridutil.Coordinate]()
	visited.Add(tail)

	for _, line := range lines {
		splitted := strings.Fields(line)
		direction := splitted[0]
		steps := conv.MustAtoi(splitted[1])

		for range steps {
			switch direction {
			case "U":
				head.Row--
			case "D":
				head.Row++
			case "L":
				head.Col--
			case "R":
				head.Col++
			}

			if (head != tail) && !isAdjacent(head, tail) {
				if head.Col == tail.Col {
					if head.Row > tail.Row {
						tail.Row++
					} else {
						tail.Row--
					}
				} else if head.Row == tail.Row {
					if head.Col > tail.Col {
						tail.Col++
					} else {
						tail.Col--
					}
				} else {
					if head.Col > tail.Col && head.Row > tail.Row {
						tail.Col++
						tail.Row++
					} else if head.Col > tail.Col && head.Row < tail.Row {
						tail.Col++
						tail.Row--
					} else if head.Col < tail.Col && head.Row > tail.Row {
						tail.Col--
						tail.Row++
					} else if head.Col < tail.Col && head.Row < tail.Row {
						tail.Col--
						tail.Row--
					}
				}
				visited.Add(tail)
			}
		}
	}
	fmt.Println("Part 1", visited.Len())
}

func part2(input string) {
	lines := conv.SplitNewline(input)

	head := gridutil.Coordinate{Row: 0, Col: 0}
	tails := make([]gridutil.Coordinate, 9)
	for i := range tails {
		tails[i] = gridutil.Coordinate{Row: 0, Col: 0}
	}

	visited := container.NewSet[gridutil.Coordinate]()
	visited.Add(tails[len(tails)-1])

	for _, line := range lines {
		splitted := strings.Fields(line)
		direction := splitted[0]
		steps := conv.MustAtoi(splitted[1])

		for range steps {
			switch direction {
			case "U":
				head.Row--
			case "D":
				head.Row++
			case "L":
				head.Col--
			case "R":
				head.Col++
			}

			prev := head
			for i := range tails {
				if (prev != tails[i]) && !isAdjacent(prev, tails[i]) {
					if prev.Col == tails[i].Col {
						if prev.Row > tails[i].Row {
							tails[i].Row++
						} else {
							tails[i].Row--
						}
					} else if prev.Row == tails[i].Row {
						if prev.Col > tails[i].Col {
							tails[i].Col++
						} else {
							tails[i].Col--
						}
					} else {
						if prev.Col > tails[i].Col && prev.Row > tails[i].Row {
							tails[i].Col++
							tails[i].Row++
						} else if prev.Col > tails[i].Col && prev.Row < tails[i].Row {
							tails[i].Col++
							tails[i].Row--
						} else if prev.Col < tails[i].Col && prev.Row > tails[i].Row {
							tails[i].Col--
							tails[i].Row++
						} else if prev.Col < tails[i].Col && prev.Row < tails[i].Row {
							tails[i].Col--
							tails[i].Row--
						}
					}
					if i == len(tails)-1 {
						visited.Add(tails[i])
					}
				}
				prev = tails[i]
			}
		}
	}
	fmt.Println("Part 2", visited.Len())
}

func isAdjacent(p1, p2 gridutil.Coordinate) bool {
	// Points are adjacent if they are at most 1 step away in any direction
	// This means their Manhattan distance must be 1, or they must be diagonally adjacent
	manhattanDist := geomutil.ManhattanDistance(p1, p2)
	return manhattanDist <= 1 || (manhattanDist == 2 && mathx.Abs(p1.Row-p2.Row) == 1 && mathx.Abs(p1.Col-p2.Col) == 1)
}
