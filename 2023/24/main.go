package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"strings"
)

func main() {
	input, err := download.ReadInput(2023, 24)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
}

type vector struct {
	x, y, z float64
}

type hailstone struct {
	position, velocity vector
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	var hailstones []hailstone

	for _, line := range lines {

		parts := strings.Split(line, " @ ")
		posStr := strings.Split(strings.TrimSpace(parts[0]), ",")
		velStr := strings.Split(strings.TrimSpace(parts[1]), ",")

		posX := conv.MustAtoi(strings.TrimSpace(posStr[0]))
		posY := conv.MustAtoi(strings.TrimSpace(posStr[1]))
		posZ := conv.MustAtoi(strings.TrimSpace(posStr[2]))
		velX := conv.MustAtoi(strings.TrimSpace(velStr[0]))
		velY := conv.MustAtoi(strings.TrimSpace(velStr[1]))
		velZ := conv.MustAtoi(strings.TrimSpace(velStr[2]))

		h := hailstone{
			position: vector{x: float64(posX), y: float64(posY), z: float64(posZ)},
			velocity: vector{x: float64(velX), y: float64(velY), z: float64(velZ)},
		}

		hailstones = append(hailstones, h)
	}

	var lower float64 = 200000000000000
	var upper float64 = 400000000000000

	intersections := 0

	for i := 0; i < len(hailstones)-1; i++ {
		for j := i + 1; j < len(hailstones); j++ {

			h1 := hailstones[i]
			h2 := hailstones[j]

			apx := h1.position.x
			apy := h1.position.y
			avx := h1.velocity.x
			avy := h1.velocity.y
			y := (h2.position.x - (h2.velocity.x / h2.velocity.y * h2.position.y) + (h1.velocity.x / h1.velocity.y * h1.position.y) - h1.position.x) / (h1.velocity.x/h1.velocity.y - h2.velocity.x/h2.velocity.y)
			x := ((y-h1.position.y)/h1.velocity.y)*h1.velocity.x + h1.position.x
			if lower <= x && x <= upper && lower <= y && y <= upper {
				if (x-apx)/avx > 0 && (y-apy)/avy > 0 && (x-h2.position.x)/h2.velocity.x > 0 && (y-h2.position.y)/h2.velocity.y > 0 {
					intersections++
				}
			}
		}
	}

	fmt.Println(intersections)
}
