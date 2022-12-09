package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"strings"
)

func main() {
	inputFile := "./2022/09/input.txt"
	input, err := download.ReadInput(inputFile, 2022, 9)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	lines := conv.SplitNewline(input)

	headX, headY := 0, 0
	tailX, tailY := 0, 0
	visited := make(map[string]bool)
	visited["0,0"] = true

	for _, line := range lines {
		splitted := strings.Fields(line)
		direction := splitted[0]
		steps := conv.MustAtoi(splitted[1])

		for i := 0; i < steps; i++ {
			if direction == "U" {
				headY -= 1
			} else if direction == "D" {
				headY += 1
			} else if direction == "L" {
				headX -= 1
			} else if direction == "R" {
				headX += 1
			}

			if (headX != tailX || headY != tailY) && !isAdjacent(headX, headY, tailX, tailY) {
				if headX == tailX {
					if headY > tailY {
						tailY++
					} else {
						tailY--
					}
				} else if headY == tailY {
					if headX > tailX {
						tailX++
					} else {
						tailX--
					}
				} else {
					if headX > tailX && headY > tailY {
						tailX++
						tailY++
					} else if headX > tailX && headY < tailY {
						tailX++
						tailY--
					} else if headX < tailX && headY > tailY {
						tailX--
						tailY++
					} else if headX < tailX && headY < tailY {
						tailX--
						tailY--
					}
				}
				key := fmt.Sprintf("%d,%d", tailX, tailY)
				visited[key] = true
			}
		}
	}
	fmt.Println("Visited:", len(visited))
}

func part2(input string) {
	lines := conv.SplitNewline(input)

	headX, headY := 0, 0
	tails := make([][2]int, 9)
	for i := 0; i < 9; i++ {
		tails[i] = [2]int{0, 0}
	}

	visited := make(map[string]bool)
	visited["0,0"] = true

	for _, line := range lines {
		splitted := strings.Fields(line)
		direction := splitted[0]
		steps := conv.MustAtoi(splitted[1])

		for i := 0; i < steps; i++ {
			if direction == "U" {
				headY -= 1
			} else if direction == "D" {
				headY += 1
			} else if direction == "L" {
				headX -= 1
			} else if direction == "R" {
				headX += 1
			}

			prevX, prevY := headX, headY
			for i := 0; i < len(tails); i++ {
				if (prevX != tails[i][0] || prevY != tails[i][1]) && !isAdjacent(prevX, prevY, tails[i][0], tails[i][1]) {
					if prevX == tails[i][0] {
						if prevY > tails[i][1] {
							tails[i][1]++
						} else {
							tails[i][1]--
						}
					} else if prevY == tails[i][1] {
						if prevX > tails[i][0] {
							tails[i][0]++
						} else {
							tails[i][0]--
						}
					} else {
						if prevX > tails[i][0] && prevY > tails[i][1] {
							tails[i][0]++
							tails[i][1]++
						} else if prevX > tails[i][0] && prevY < tails[i][1] {
							tails[i][0]++
							tails[i][1]--
						} else if prevX < tails[i][0] && prevY > tails[i][1] {
							tails[i][0]--
							tails[i][1]++
						} else if prevX < tails[i][0] && prevY < tails[i][1] {
							tails[i][0]--
							tails[i][1]--
						}
					}
					if i == len(tails)-1 {
						key := fmt.Sprintf("%d,%d", tails[i][0], tails[i][1])
						visited[key] = true
					}
				}
				prevX, prevY = tails[i][0], tails[i][1]
			}
		}
	}
	fmt.Println("Visited:", len(visited))
}

func isAdjacent(x1, y1, x2, y2 int) bool {
	if x1 == x2 && (y1 == y2+1 || y1 == y2-1) {
		return true
	}
	if (x1 == x2+1 || x1 == x2-1) && y1 == y2 {
		return true
	}
	if x1 == x2+1 && (y1 == y2+1 || y1 == y2-1) {
		return true
	}
	if (x1 == x2-1) && (y1 == y2+1 || y1 == y2-1) {
		return true
	}
	return false
}
