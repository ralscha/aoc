package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/gridutil"
	"fmt"
	"log"
)

func main() {
	input, err := download.ReadInput(2024, 12)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	garden := gridutil.NewCharGrid2D(lines)
	totalPrice := totalFencePrice(&garden)
	fmt.Println("Part 1", totalPrice)
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	garden := gridutil.NewCharGrid2D(lines)
	totalPrice := totalFencePriceWithSides(&garden)
	fmt.Println("Part 2", totalPrice)
}

func totalFencePrice(garden *gridutil.Grid2D[rune]) int {
	visited := container.NewSet[gridutil.Coordinate]()
	totalPrice := 0

	for row := 0; row <= garden.Height(); row++ {
		for col := 0; col <= garden.Width(); col++ {
			coord := gridutil.Coordinate{Row: row, Col: col}
			if !visited.Contains(coord) {
				if plant, exists := garden.GetC(coord); exists {
					area, permiter, _ := exploreRegion(coord, plant, garden, visited)
					price := area * permiter
					totalPrice += price
				}
			}
		}
	}

	return totalPrice
}

func totalFencePriceWithSides(garden *gridutil.Grid2D[rune]) int {
	visited := container.NewSet[gridutil.Coordinate]()
	totalPrice := 0

	for row := 0; row <= garden.Height(); row++ {
		for col := 0; col <= garden.Width(); col++ {
			coord := gridutil.Coordinate{Row: row, Col: col}
			if !visited.Contains(coord) {
				if plant, exists := garden.GetC(coord); exists {
					area, _, region := exploreRegion(coord, plant, garden, visited)
					sides := countCorners(region)
					price := area * sides
					totalPrice += price
				}
			}
		}
	}

	return totalPrice
}

func countCorners(coords []gridutil.Coordinate) int {
	corners := 0

	for _, coord := range coords {
		unattachedCount := 0
		for _, dir := range gridutil.Get4Directions() {
			neighbour := gridutil.Coordinate{
				Row: coord.Row + dir.Row,
				Col: coord.Col + dir.Col,
			}
			if !contains(coords, neighbour) {
				unattachedCount++
			}
		}

		top := gridutil.Coordinate{Row: coord.Row - 1, Col: coord.Col}
		bottom := gridutil.Coordinate{Row: coord.Row + 1, Col: coord.Col}
		left := gridutil.Coordinate{Row: coord.Row, Col: coord.Col - 1}
		right := gridutil.Coordinate{Row: coord.Row, Col: coord.Col + 1}

		switch unattachedCount {
		case 4:
			corners += 4
		case 3:
			corners += 2
		case 2:
			if (contains(coords, top) && contains(coords, bottom)) ||
				(contains(coords, left) && contains(coords, right)) {
				corners += 0
			} else {
				corners += 1
			}
		}

		// Inner corners
		topRight := gridutil.Coordinate{Row: coord.Row - 1, Col: coord.Col + 1}
		if !contains(coords, topRight) &&
			contains(coords, top) &&
			contains(coords, right) {
			corners += 1
		}

		bottomRight := gridutil.Coordinate{Row: coord.Row + 1, Col: coord.Col + 1}
		if !contains(coords, bottomRight) &&
			contains(coords, bottom) &&
			contains(coords, right) {
			corners += 1
		}

		topLeft := gridutil.Coordinate{Row: coord.Row - 1, Col: coord.Col - 1}
		if !contains(coords, topLeft) &&
			contains(coords, top) &&
			contains(coords, left) {
			corners += 1
		}

		bottomLeft := gridutil.Coordinate{Row: coord.Row + 1, Col: coord.Col - 1}
		if !contains(coords, bottomLeft) &&
			contains(coords, bottom) &&
			contains(coords, left) {
			corners += 1
		}
	}

	return corners
}

func contains(coords []gridutil.Coordinate, coord gridutil.Coordinate) bool {
	for _, c := range coords {
		if c == coord {
			return true
		}
	}
	return false
}

func exploreRegion(start gridutil.Coordinate, regionPlant rune, garden *gridutil.Grid2D[rune], visited *container.Set[gridutil.Coordinate]) (int, int, []gridutil.Coordinate) {
	queue := container.NewQueue[gridutil.Coordinate]()
	queue.Push(start)
	visited.Add(start)

	region := []gridutil.Coordinate{start}

	area := 0
	perimeter := 0

	for !queue.IsEmpty() {
		current := queue.Pop()
		area++

		for _, dir := range gridutil.Get4Directions() {
			nextCoord := gridutil.Coordinate{
				Row: current.Row + dir.Row,
				Col: current.Col + dir.Col,
			}

			nextPlant, exists := garden.GetC(nextCoord)
			if !exists || nextPlant != regionPlant {
				perimeter++
			} else {
				if !visited.Contains(nextCoord) {
					region = append(region, nextCoord)
					visited.Add(nextCoord)
					queue.Push(nextCoord)
				}
			}
		}
	}

	return area, perimeter, region
}
