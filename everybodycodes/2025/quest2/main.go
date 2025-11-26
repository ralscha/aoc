package main

import (
	"aoc/internal/conv"
	"fmt"
)

func main() {
	partI()
	partII()
	partIII()
}

func partI() {
	input := "A=[158,55]"

	var ax, ay int
	conv.MustSscanf(input, "A=[%d,%d]", &ax, &ay)

	rx, ry := 0, 0

	for range 3 {
		rxNew := rx*rx - ry*ry
		ryNew := rx*ry + ry*rx
		rx, ry = rxNew, ryNew

		rx /= 10
		ry /= 10

		rx += ax
		ry += ay
	}

	fmt.Printf("[%d,%d]\n", rx, ry)
}

func partII() {
	input := "A=[-79067,14068]"
	var ax, ay int
	conv.MustSscanf(input, "A=[%d,%d]", &ax, &ay)

	engravedCount := solve(ax, ay, 101, 10)
	fmt.Println(engravedCount)
}

func partIII() {
	input := "A=[-79067,14068]"
	var ax, ay int
	conv.MustSscanf(input, "A=[%d,%d]", &ax, &ay)

	engravedCount := solve(ax, ay, 1001, 1)
	fmt.Println(engravedCount)
}

func solve(ax, ay, gridSize, step int) int {
	engravedCount := 0

	for i := range gridSize {
		for j := range gridSize {
			px := ax + i*step
			py := ay + j*step

			rx, ry := 0, 0
			engrave := true
			for range 100 {
				rxNew := rx*rx - ry*ry
				ryNew := rx*ry + ry*rx
				rx, ry = rxNew, ryNew

				rx /= 100000
				ry /= 100000

				rx += px
				ry += py

				if rx > 1000000 || rx < -1000000 || ry > 1000000 || ry < -1000000 {
					engrave = false
					break
				}
			}
			if engrave {
				engravedCount++
			}
		}
	}

	return engravedCount
}
