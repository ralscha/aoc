package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"aoc/internal/gridutil"
	"fmt"
	"os"
	"strings"
)

func main() {
	partI()
	partII()
	partIII()
}

func partI() {
	input, _ := os.ReadFile("input1")
	lines := conv.SplitNewline(string(input))
	grid := gridutil.NewCharGrid2D(lines)

	var volcano gridutil.Coordinate
	minRow, maxRow := grid.GetMinMaxRow()
	minCol, maxCol := grid.GetMinMaxCol()
	for row := minRow; row <= maxRow; row++ {
		for col := minCol; col <= maxCol; col++ {
			if c, ok := grid.Get(row, col); ok && c == '@' {
				volcano = gridutil.Coordinate{Row: row, Col: col}
				break
			}
		}
	}

	radius := 10
	sum := 0
	for row := minRow; row <= maxRow; row++ {
		for col := minCol; col <= maxCol; col++ {
			c, ok := grid.Get(row, col)
			if !ok || c == '@' {
				continue
			}
			dx := volcano.Col - col
			dy := volcano.Row - row
			if dx*dx+dy*dy <= radius*radius {
				sum += int(c - '0')
			}
		}
	}

	fmt.Println(sum)
}

func partII() {
	input, _ := os.ReadFile("input2")
	lines := conv.SplitNewline(string(input))
	grid := gridutil.NewCharGrid2D(lines)

	var volcano gridutil.Coordinate
	minRow, maxRow := grid.GetMinMaxRow()
	minCol, maxCol := grid.GetMinMaxCol()
	for row := minRow; row <= maxRow; row++ {
		for col := minCol; col <= maxCol; col++ {
			if c, ok := grid.Get(row, col); ok && c == '@' {
				volcano = gridutil.Coordinate{Row: row, Col: col}
				break
			}
		}
	}

	distToTop := volcano.Row - minRow
	distToBottom := maxRow - volcano.Row
	distToLeft := volcano.Col - minCol
	distToRight := maxCol - volcano.Col
	maxRadius := min(distToTop, distToBottom, distToLeft, distToRight)

	destroyed := container.NewSet[gridutil.Coordinate]()
	maxDestruction := 0
	maxRadiusUsed := 0

	for radius := 1; radius <= maxRadius; radius++ {
		stepDestruction := 0
		r2 := radius * radius

		for row := minRow; row <= maxRow; row++ {
			for col := minCol; col <= maxCol; col++ {
				coord := gridutil.Coordinate{Row: row, Col: col}
				if destroyed.Contains(coord) {
					continue
				}
				c, ok := grid.Get(row, col)
				if !ok || c == '@' {
					continue
				}
				dx := volcano.Col - col
				dy := volcano.Row - row
				if dx*dx+dy*dy <= r2 {
					stepDestruction += int(c - '0')
					destroyed.Add(coord)
				}
			}
		}

		if stepDestruction > maxDestruction {
			maxDestruction = stepDestruction
			maxRadiusUsed = radius
		}
	}

	fmt.Println(maxRadiusUsed * maxDestruction)
}

type State struct {
	pos   gridutil.Coordinate
	state int
	cost  int
}

func partIII() {
	input, _ := os.ReadFile("input3")
	inputStr := strings.ReplaceAll(string(input), "\r", "")
	lines := conv.SplitNewline(inputStr)
	grid := gridutil.NewCharGrid2D(lines)

	var start, volcano gridutil.Coordinate
	minRow, maxRow := grid.GetMinMaxRow()
	minCol, maxCol := grid.GetMinMaxCol()
	for row := minRow; row <= maxRow; row++ {
		for col := minCol; col <= maxCol; col++ {
			if c, ok := grid.Get(row, col); ok {
				switch c {
				case 'S':
					start = gridutil.Coordinate{Row: row, Col: col}
				case '@':
					volcano = gridutil.Coordinate{Row: row, Col: col}
				}
			}
		}
	}

	const (
		StateStart = 0
		StateLeft  = 1
		StateRight = 2
	)

	directions := gridutil.Get4Directions()

	for radius := 0; ; radius++ {
		maxDist := (radius+1)*30 - 1
		r2 := radius * radius

		type StateKey struct {
			pos   gridutil.Coordinate
			state int
		}
		visited := make(map[StateKey]int)

		pq := container.NewPriorityQueue[State]()
		pq.Push(State{pos: start, state: StateStart, cost: 0}, 0)
		visited[StateKey{pos: start, state: StateStart}] = 0

		for !pq.IsEmpty() {
			curr := pq.Pop()

			key := StateKey{pos: curr.pos, state: curr.state}
			if prev, ok := visited[key]; ok && prev < curr.cost {
				continue
			}

			newState := curr.state
			if curr.state == StateStart {
				if curr.pos.Row == volcano.Row && curr.pos.Col < volcano.Col {
					newState = StateLeft
				} else if curr.pos.Row == volcano.Row && curr.pos.Col > volcano.Col {
					newState = StateRight
				} else if curr.pos.Row > volcano.Row {
					continue
				}
			} else {
				if curr.pos.Row < volcano.Row {
					continue
				}
				if curr.pos.Row > volcano.Row && curr.pos.Col == volcano.Col {
					continue
				}
			}

			if curr.cost > maxDist {
				continue
			}

			for _, dir := range directions {
				newPos := gridutil.Coordinate{
					Row: curr.pos.Row + dir.Row,
					Col: curr.pos.Col + dir.Col,
				}

				c, ok := grid.Get(newPos.Row, newPos.Col)
				if !ok {
					continue
				}

				dr := newPos.Row - volcano.Row
				dc := newPos.Col - volcano.Col
				if dr*dr+dc*dc <= r2 {
					continue
				}

				var moveCost int
				if c == 'S' || c == '@' {
					continue
				} else {
					moveCost = int(c - '0')
				}
				if moveCost == 0 {
					continue
				}

				newCost := curr.cost + moveCost

				newKey := StateKey{pos: newPos, state: newState}
				if prev, ok := visited[newKey]; ok && prev <= newCost {
					continue
				}
				visited[newKey] = newCost

				pq.Push(State{pos: newPos, state: newState, cost: newCost}, newCost)
			}
		}

		for row := volcano.Row + 1; row <= maxRow; row++ {
			meetPos := gridutil.Coordinate{Row: row, Col: volcano.Col}
			c, ok := grid.Get(meetPos.Row, meetPos.Col)
			if !ok || c == '@' || c == 'S' {
				continue
			}
			cellCost := int(c - '0')
			if cellCost == 0 {
				continue
			}

			leftKey := StateKey{pos: meetPos, state: StateLeft}
			rightKey := StateKey{pos: meetPos, state: StateRight}

			leftDist, leftOk := visited[leftKey]
			rightDist, rightOk := visited[rightKey]

			if leftOk && rightOk {
				totalDist := leftDist + rightDist - cellCost
				if totalDist <= maxDist {
					fmt.Println(totalDist * radius)
					return
				}
			}
		}
	}
}
