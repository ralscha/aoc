package main

import (
	"aoc/internal/download"
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type point struct {
	x, y int
}

type fold struct {
	dir string
	val int
}

func main() {
	inputFile := "./2021/13/input.txt"
	input, err := download.ReadInput(inputFile, 2021, 13)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	var folds []fold

	grid := make(map[point]struct{})

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "fold along ") {
			blank := strings.LastIndex(line, " ")
			l := line[blank+1:]
			ls := strings.Split(l, "=")
			val, err := strconv.Atoi(ls[1])
			if err != nil {
				log.Fatalf("conversion failed: %s %v", ls[1], err)
			}
			folds = append(folds, fold{
				dir: ls[0],
				val: val,
			})
		} else if len(line) > 0 {
			splitted := strings.Split(line, ",")
			x, err := strconv.Atoi(splitted[0])
			if err != nil {
				log.Fatalf("conversion failed: %s %v", splitted[0], err)
			}
			y, err := strconv.Atoi(splitted[1])
			if err != nil {
				log.Fatalf("conversion failed: %s %v", splitted[1], err)
			}
			grid[point{
				x: x,
				y: y,
			}] = struct{}{}
		}
	}

	ff := folds[0]

	for k := range grid {
		var np point
		hasnp := false
		if k.x > ff.val && ff.dir == "x" {
			np = point{
				x: ff.val - (k.x - ff.val),
				y: k.y,
			}
			hasnp = true
		} else if k.y > ff.val && ff.dir == "y" {
			np = point{
				x: k.x,
				y: ff.val - (k.y - ff.val),
			}
			hasnp = true
		}
		if hasnp {
			if _, ok := grid[np]; !ok {
				grid[np] = struct{}{}
			}
			delete(grid, k)
		}

	}

	fmt.Println("Result: ", len(grid))
}

func part2(input string) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	var folds []fold

	grid := make(map[point]struct{})

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "fold along ") {
			blank := strings.LastIndex(line, " ")
			l := line[blank+1:]
			ls := strings.Split(l, "=")
			val, err := strconv.Atoi(ls[1])
			if err != nil {
				log.Fatalf("conversion failed: %s %v", ls[1], err)
			}
			folds = append(folds, fold{
				dir: ls[0],
				val: val,
			})
		} else if len(line) > 0 {
			splitted := strings.Split(line, ",")
			x, err := strconv.Atoi(splitted[0])
			if err != nil {
				log.Fatalf("conversion failed: %s %v", splitted[0], err)
			}
			y, err := strconv.Atoi(splitted[1])
			if err != nil {
				log.Fatalf("conversion failed: %s %v", splitted[1], err)
			}
			grid[point{
				x: x,
				y: y,
			}] = struct{}{}
		}
	}

	for _, ff := range folds {
		for k := range grid {
			var np point
			hasnp := false
			if k.x > ff.val && ff.dir == "x" {
				np = point{
					x: ff.val - (k.x - ff.val),
					y: k.y,
				}
				hasnp = true
			} else if k.y > ff.val && ff.dir == "y" {
				np = point{
					x: k.x,
					y: ff.val - (k.y - ff.val),
				}
				hasnp = true
			}
			if hasnp {
				if _, ok := grid[np]; !ok {
					grid[np] = struct{}{}
				}
				delete(grid, k)
			}
		}
	}

	fmt.Println("Result: ", len(grid))

	minX := 0
	minY := 0
	for k := range grid {
		if k.x > minX {
			minX = k.x
		}
		if k.y > minY {
			minY = k.y
		}
	}

	for y := 0; y <= minY; y++ {
		for x := 0; x <= minX; x++ {
			if _, ok := grid[point{
				x: x,
				y: y,
			}]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
