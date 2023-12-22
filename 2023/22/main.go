package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, err := download.ReadInput(2023, 22)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1and2(input)
}

func part1and2(input string) {
	lines := conv.SplitNewline(input)
	bricks := make([]brick, 0, len(lines))
	for _, line := range lines {
		ints := convToInts(line)
		bricks = append(bricks, brick{ints[0], ints[1], ints[2], ints[3], ints[4], ints[5]})
	}
	sort.Slice(bricks, func(i, j int) bool {
		return bricks[i].z1 < bricks[j].z1
	})
	_, fallen := drop(bricks)
	p1 := 0
	p2 := 0
	for i := range fallen {
		removed := append(fallen[:i:i], fallen[i+1:]...)
		falls, _ := drop(removed)
		if falls == 0 {
			p1++
		} else {
			p2 += falls
		}
	}
	fmt.Println(p1)
	fmt.Println(p2)
}

type brick struct {
	x1, y1, z1, x2, y2, z2 int
}

func droppedBrick(tallest map[[2]int]int, b brick) brick {
	peak := 0
	for x := b.x1; x <= b.x2; x++ {
		for y := b.y1; y <= b.y2; y++ {
			if h, ok := tallest[[2]int{x, y}]; ok && h > peak {
				peak = h
			}
		}
	}
	dz := max(b.z1-peak-1, 0)
	return brick{b.x1, b.y1, b.z1 - dz, b.x2, b.y2, b.z2 - dz}
}

func drop(tower []brick) (int, []brick) {
	tallest := make(map[[2]int]int)
	newTower := make([]brick, 0, len(tower))
	falls := 0
	for _, brick := range tower {
		newBrick := droppedBrick(tallest, brick)
		if newBrick.z1 != brick.z1 {
			falls++
		}
		newTower = append(newTower, newBrick)
		for x := brick.x1; x <= brick.x2; x++ {
			for y := brick.y1; y <= brick.y2; y++ {
				tallest[[2]int{x, y}] = newBrick.z2
			}
		}
	}
	return falls, newTower
}

func convToInts(s string) []int {
	splitted := strings.Split(s, "~")
	ints := make([]int, 6)
	i := 0
	for s := range splitted {
		fields := strings.Split(splitted[s], ",")
		for _, field := range fields {
			n, err := strconv.Atoi(field)
			if err != nil {
				panic(err)
			}
			ints[i] = n
			i++
		}
	}
	return ints
}
