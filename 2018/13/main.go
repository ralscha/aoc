package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/gridutil"
	"fmt"
	"log"
	"slices"
)

type cart struct {
	col     int
	row     int
	dir     gridutil.Direction
	turn    int
	crashed bool
}

func main() {
	input, err := download.ReadInput(2018, 13)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	grid, carts := parseInput(input)

	for {
		slices.SortFunc(carts, func(cartA, cartB *cart) int {
			if cartA.row < cartB.row {
				return -1
			}
			if cartA.row > cartB.row {
				return 1
			}
			if cartA.col < cartB.col {
				return -1
			}
			if cartA.col > cartB.col {
				return 1
			}
			return 0
		})

		for i := range len(carts) {
			if carts[i].crashed {
				continue
			}
			moveCart(grid, carts[i])
			if collisionOccurred(carts) {
				for _, c := range carts {
					if c.crashed {
						fmt.Println("Part 1", fmt.Sprintf("%d,%d", c.col, c.row))
						return
					}
				}
			}
		}
	}
}

func part2(input string) {
	grid, carts := parseInput(input)

	for len(carts) > 1 {
		slices.SortFunc(carts, func(cartA, cartB *cart) int {
			if cartA.row < cartB.row {
				return -1
			}
			if cartA.row > cartB.row {
				return 1
			}
			if cartA.col < cartB.col {
				return -1
			}
			if cartA.col > cartB.col {
				return 1
			}
			return 0
		})

		moved := container.NewSet[*cart]()
		for i := range carts {
			if carts[i].crashed || moved.Contains(carts[i]) {
				continue
			}
			moveCart(grid, carts[i])
			moved.Add(carts[i])
			handleCollisionsPart2(carts)
		}

		var newCarts []*cart
		for _, c := range carts {
			if !c.crashed {
				newCarts = append(newCarts, c)
			}
		}
		carts = newCarts
	}

	fmt.Println("Part 2", fmt.Sprintf("%d,%d", carts[0].col, carts[0].row))
}

func parseInput(input string) (gridutil.Grid2D[rune], []*cart) {
	lines := conv.SplitNewline(input)
	grid := gridutil.NewCharGrid2D(lines)
	var carts []*cart

	for r := range grid.Height() {
		for c := range grid.Width() {
			v, ok := grid.Get(r, c)
			if !ok {
				continue
			}
			var dir gridutil.Direction
			isCart := true
			switch v {
			case '^':
				dir = gridutil.DirectionN
			case 'v':
				dir = gridutil.DirectionS
			case '<':
				dir = gridutil.DirectionW
			case '>':
				dir = gridutil.DirectionE
			default:
				isCart = false
			}
			if isCart {
				carts = append(carts, &cart{col: c, row: r, dir: dir})
			}
		}
	}
	return grid, carts
}

func moveCart(grid gridutil.Grid2D[rune], c *cart) {
	nextX := c.col + c.dir.Col
	nextY := c.row + c.dir.Row
	c.col = nextX
	c.row = nextY

	track, _ := grid.GetC(gridutil.Coordinate{Row: c.row, Col: c.col})

	switch track {
	case '/':
		switch c.dir {
		case gridutil.DirectionE:
			c.dir = gridutil.DirectionN
		case gridutil.DirectionW:
			c.dir = gridutil.DirectionS
		case gridutil.DirectionS:
			c.dir = gridutil.DirectionW
		case gridutil.DirectionN:
			c.dir = gridutil.DirectionE
		}
	case '\\':
		switch c.dir {
		case gridutil.DirectionE:
			c.dir = gridutil.DirectionS
		case gridutil.DirectionW:
			c.dir = gridutil.DirectionN
		case gridutil.DirectionS:
			c.dir = gridutil.DirectionE
		case gridutil.DirectionN:
			c.dir = gridutil.DirectionW
		}
	case '+':
		switch c.turn % 3 {
		case 0:
			c.dir = gridutil.TurnLeft(c.dir)
		case 2:
			c.dir = gridutil.TurnRight(c.dir)
		}
		c.turn++
	}
}

func collisionOccurred(carts []*cart) bool {
	locations := make(map[string]int)
	for _, c := range carts {
		key := fmt.Sprintf("%d,%d", c.col, c.row)
		locations[key]++
	}
	for _, count := range locations {
		if count > 1 {
			for i := range carts {
				for j := i + 1; j < len(carts); j++ {
					if carts[i].col == carts[j].col && carts[i].row == carts[j].row {
						carts[i].crashed = true
						carts[j].crashed = true
						return true
					}
				}
			}
		}
	}
	return false
}

func handleCollisionsPart2(carts []*cart) {
	locations := make(map[string][]*cart)
	for _, c := range carts {
		if !c.crashed {
			key := fmt.Sprintf("%d,%d", c.col, c.row)
			locations[key] = append(locations[key], c)
		}
	}
	for _, cl := range locations {
		if len(cl) > 1 {
			for _, c := range cl {
				c.crashed = true
			}
		}
	}
}
