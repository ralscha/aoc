package main

import (
	"aoc/internal/container"
	"aoc/internal/download"
	"crypto/md5"
	"fmt"
	"log"
)

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

func part1(input string) {
	salt := input
	queue := container.NewQueue[struct {
		x    int
		y    int
		path string
	}]()
	queue.Push(struct {
		x, y int
		path string
	}{0, 0, ""})

	for !queue.IsEmpty() {
		curr := queue.Pop()
		x := curr.x
		y := curr.y
		path := curr.path

		if x == 3 && y == 3 {
			fmt.Println("Part 1:", path)
			return
		}

		open := getOpenDoors(path, salt)
		if open[0] && y > 0 {
			queue.Push(struct {
				x, y int
				path string
			}{x, y - 1, path + "U"})
		}
		if open[1] && y < 3 {
			queue.Push(struct {
				x, y int
				path string
			}{x, y + 1, path + "D"})
		}
		if open[2] && x > 0 {
			queue.Push(struct {
				x, y int
				path string
			}{x - 1, y, path + "L"})
		}
		if open[3] && x < 3 {
			queue.Push(struct {
				x, y int
				path string
			}{x + 1, y, path + "R"})
		}
	}
}

func part2(input string) {
	salt := input
	queue := container.NewQueue[struct {
		x    int
		y    int
		path string
	}]()
	queue.Push(struct {
		x, y int
		path string
	}{0, 0, ""})
	longest := 0

	for !queue.IsEmpty() {
		curr := queue.Pop()
		x := curr.x
		y := curr.y
		path := curr.path

		if x == 3 && y == 3 {
			if len(path) > longest {
				longest = len(path)
			}
			continue
		}

		open := getOpenDoors(path, salt)
		if open[0] && y > 0 {
			queue.Push(struct {
				x, y int
				path string
			}{x, y - 1, path + "U"})
		}
		if open[1] && y < 3 {
			queue.Push(struct {
				x, y int
				path string
			}{x, y + 1, path + "D"})
		}
		if open[2] && x > 0 {
			queue.Push(struct {
				x, y int
				path string
			}{x - 1, y, path + "L"})
		}
		if open[3] && x < 3 {
			queue.Push(struct {
				x, y int
				path string
			}{x + 1, y, path + "R"})
		}
	}
	fmt.Println("Part 2:", longest)
}
