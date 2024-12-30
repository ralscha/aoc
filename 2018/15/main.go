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

type field [][]cell

func parseInput(input string) (field, []*unit) {
	var field field
	var units []*unit
	y := 0

	lines := conv.SplitNewline(input)
	for _, line := range lines {
		fieldRow := make([]cell, len(line))

		for x, ch := range line {
			if ch == 'G' || ch == 'E' {
				unit := &unit{200, y, x, ch}
				units = append(units, unit)
				fieldRow[x] = cell{isWall: false, unit: unit}
			} else {
				fieldRow[x] = cell{isWall: ch == '#', unit: nil}
			}
		}

		field = append(field, fieldRow)
		y++
	}

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

func findMovementTarget(field field, unit *unit) *gridutil.Coordinate {
	height, width := len(field), len(field[0])
	start := gridutil.Coordinate{Row: unit.row, Col: unit.col}
	queue := []gridutil.Coordinate{start}
	seen := make(map[gridutil.Coordinate]gridutil.Coordinate)
	seen[start] = gridutil.Coordinate{Row: -1, Col: -1}

	dirs := [4]gridutil.Direction{gridutil.DirectionN, gridutil.DirectionW, gridutil.DirectionE, gridutil.DirectionS}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, dir := range dirs {
			newRow, newCol := current.Row+dir.Row, current.Col+dir.Col

			if newRow >= 0 && newRow < height && newCol >= 0 && newCol < width {
				newPos := gridutil.Coordinate{Row: newRow, Col: newCol}

				if _, exists := seen[newPos]; !exists {
					cell := field[newRow][newCol]

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

func findAttackTarget(field field, u *unit) *unit {
	height, width := len(field), len(field[0])
	dirs := [4]gridutil.Direction{gridutil.DirectionN, gridutil.DirectionW, gridutil.DirectionE, gridutil.DirectionS}

	var targets []*unit

	for _, dir := range dirs {
		newRow, newCol := u.row+dir.Row, u.col+dir.Col

		if newRow >= 0 && newRow < height && newCol >= 0 && newCol < width {
			if targetUnit := field[newRow][newCol].unit; targetUnit != nil {
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

func simulateBattle(field field, units []*unit, elfPower int) (map[rune]int, int) {
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
				field[unit.row][unit.col] = cell{isWall: false, unit: nil}
				unit.row, unit.col = target.Row, target.Col
				field[unit.row][unit.col] = cell{isWall: false, unit: unit}
			}

			if target := findAttackTarget(field, unit); target != nil {
				if unit.utype == 'E' {
					target.health -= elfPower
				} else {
					target.health -= 3
				}

				if target.health <= 0 {
					field[target.row][target.col] = cell{isWall: false, unit: nil}
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

func copyState(f field) (field, []*unit) {
	newField := make(field, len(f))
	var newUnits []*unit

	for i := range f {
		newField[i] = make([]cell, len(f[i]))
		for j := range f[i] {
			if u := f[i][j].unit; u != nil {
				newUnit := &unit{u.health, u.row, u.col, u.utype}
				newUnits = append(newUnits, newUnit)
				newField[i][j] = cell{isWall: false, unit: newUnit}
			} else {
				newField[i][j] = f[i][j]
			}
		}
	}

	return newField, newUnits
}
