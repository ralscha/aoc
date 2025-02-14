package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/gridutil"
	"fmt"
	"slices"
)

func main() {
	input, err := download.ReadInput(2018, 15)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	initialField, _ := parseInput(input)
	field, units := copyState(initialField)
	_, score := simulateBattle(field, units, 3)
	fmt.Println("Part 1", score)
}

func part2(input string) {
	initialField, initialUnits := parseInput(input)
	initialElves := countElves(initialUnits)
	for elfPower := 4; ; elfPower++ {
		field, units := copyState(initialField)
		result, score := simulateBattle(field, units, elfPower)
		if result['E'] == initialElves {
			fmt.Println("Part 2", score)
			return
		}
	}
}

type unit struct {
	health   int
	row, col int
	utype    rune
}

type cell struct {
	isWall bool
	unit   *unit
}

func parseInput(input string) (gridutil.Grid2D[cell], []*unit) {
	field := gridutil.NewGrid2D[cell](false)
	var units []*unit

	lines := conv.SplitNewline(input)
	for row, line := range lines {
		for col, ch := range line {
			if ch == 'G' || ch == 'E' {
				unit := &unit{200, row, col, ch}
				units = append(units, unit)
				field.Set(row, col, cell{isWall: false, unit: unit})
			} else {
				field.Set(row, col, cell{isWall: ch == '#', unit: nil})
			}
		}
	}

	field.SetMaxRowCol(len(lines)-1, len(lines[0])-1)
	return field, units
}

func countElves(units []*unit) int {
	count := 0
	for _, unit := range units {
		if unit.utype == 'E' {
			count++
		}
	}
	return count
}

func findMovementTarget(field gridutil.Grid2D[cell], unit *unit) *gridutil.Coordinate {
	start := gridutil.Coordinate{Row: unit.row, Col: unit.col}
	queue := []gridutil.Coordinate{start}
	seen := make(map[gridutil.Coordinate]gridutil.Coordinate)
	seen[start] = gridutil.Coordinate{Row: -1, Col: -1}

	dirs := [4]gridutil.Direction{gridutil.DirectionN, gridutil.DirectionW, gridutil.DirectionE, gridutil.DirectionS}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, dir := range dirs {
			newPos := gridutil.Coordinate{Row: current.Row + dir.Row, Col: current.Col + dir.Col}

			if _, exists := seen[newPos]; !exists {
				if cell, ok := field.GetC(newPos); ok {
					if targetUnit := cell.unit; targetUnit != nil {
						if targetUnit.utype != unit.utype {
							if current == start {
								return nil
							}

							pos := current
							for seen[pos] != start {
								pos = seen[pos]
							}
							return &pos
						}
					} else if !cell.isWall {
						queue = append(queue, newPos)
						seen[newPos] = current
					}
				}
			}
		}
	}

	return nil
}

func findAttackTarget(field gridutil.Grid2D[cell], u *unit) *unit {
	var targets []*unit
	dirs := [4]gridutil.Direction{gridutil.DirectionN, gridutil.DirectionW, gridutil.DirectionE, gridutil.DirectionS}

	for _, dir := range dirs {
		newPos := gridutil.Coordinate{Row: u.row + dir.Row, Col: u.col + dir.Col}
		if cell, ok := field.GetC(newPos); ok {
			if targetUnit := cell.unit; targetUnit != nil {
				if targetUnit.utype != u.utype {
					targets = append(targets, targetUnit)
				}
			}
		}
	}

	if len(targets) == 0 {
		return nil
	}

	slices.SortFunc(targets, func(a, b *unit) int {
		if a.health == b.health {
			if a.row == b.row {
				return a.col - b.col
			}
			return a.row - b.row
		}
		return a.health - b.health
	})

	return targets[0]
}

func simulateBattle(field gridutil.Grid2D[cell], units []*unit, elfPower int) (map[rune]int, int) {
	round := 0

	for {
		slices.SortFunc(units, func(a, b *unit) int {
			if a.row == b.row {
				return a.col - b.col
			}
			return a.row - b.row
		})

		for _, unit := range units {
			if unit.health <= 0 {
				continue
			}

			counts := make(map[rune]int)
			for _, u := range units {
				if u.health > 0 {
					counts[u.utype]++
				}
			}

			if counts['G'] == 0 || counts['E'] == 0 {
				totalHealth := 0
				for _, u := range units {
					if u.health > 0 {
						totalHealth += u.health
					}
				}
				return counts, round * totalHealth
			}

			if target := findMovementTarget(field, unit); target != nil {
				field.Set(unit.row, unit.col, cell{isWall: false, unit: nil})
				unit.row, unit.col = target.Row, target.Col
				field.Set(unit.row, unit.col, cell{isWall: false, unit: unit})
			}

			if target := findAttackTarget(field, unit); target != nil {
				if unit.utype == 'E' {
					target.health -= elfPower
				} else {
					target.health -= 3
				}

				if target.health <= 0 {
					field.Set(target.row, target.col, cell{isWall: false, unit: nil})
				}
			}
		}

		var aliveUnits []*unit
		for _, unit := range units {
			if unit.health > 0 {
				aliveUnits = append(aliveUnits, unit)
			}
		}
		units = aliveUnits
		round++
	}
}

func copyState(f gridutil.Grid2D[cell]) (gridutil.Grid2D[cell], []*unit) {
	newField := f.Copy()
	var newUnits []*unit

	minRow, maxRow := f.GetMinMaxRow()
	minCol, maxCol := f.GetMinMaxCol()

	for row := minRow; row <= maxRow; row++ {
		for col := minCol; col <= maxCol; col++ {
			if c, ok := f.Get(row, col); ok {
				if u := c.unit; u != nil {
					newUnit := &unit{u.health, u.row, u.col, u.utype}
					newUnits = append(newUnits, newUnit)
					newField.Set(row, col, cell{isWall: false, unit: newUnit})
				}
			}
		}
	}

	return newField, newUnits
}
