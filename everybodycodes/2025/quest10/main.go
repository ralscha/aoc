package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"aoc/internal/gridutil"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func main() {
	partI()
	partII()
	partIII()
}

func partI() {
	input := `SSS..SSSS.SS.SSSSSSSS
.SSS.S.S.S...SSSSSSSS
SSS.SS..S.SS.SSSS.SSS
SSSS...S..SS.SSSSS.SS
SSS.SSSSS.SSSS.SSSS.S
SSSSSSSSSSSSSSSSSSS.S
.SS.SSSS.S.S.S..S.SS.
SSS.S..SSS.S.SSSSSSS.
.SSS..SS.SSSSS..SSSSS
.SS.S..SSS.S..SSSSSSS
SSS.SS.SSSDS.SSSS..SS
S.SSSS.S.SSSSSSSSSS.S
SS..S.S.SSSS.S..S.SSS
.S.SSSS..SSSSSSSSSS..
SSSSSS.SSSSS.SSSSSSSS
SSSSSS.S..SSSSSSS.SS.
SS.SS.SSSSSSS.S.SSSSS
SSSS.SSSSSSS.SSSSS..S
S.SSS..SSS..SSSS.SSS.
SSSSSS.S.SSS.SSSS.SSS
.SS..SS.S.S..S.SSSS.S`

	lines := conv.SplitNewline(input)
	grid := gridutil.NewCharGrid2D(lines)
	var dragonPos gridutil.Coordinate

	for row, line := range lines {
		for col, ch := range line {
			if ch == 'D' {
				dragonPos = gridutil.Coordinate{Row: row, Col: col}
			}
		}
	}

	reachable := container.NewSet[gridutil.Coordinate]()
	queue := []struct {
		pos   gridutil.Coordinate
		moves int
	}{{dragonPos, 0}}
	visited := make(map[gridutil.Coordinate]int)
	visited[dragonPos] = 0

	getMoves := func(p gridutil.Coordinate) []gridutil.Coordinate {
		var moves []gridutil.Coordinate
		moves = append(moves, gridutil.Coordinate{Row: p.Row - 2, Col: p.Col - 1})
		moves = append(moves, gridutil.Coordinate{Row: p.Row - 2, Col: p.Col + 1})
		moves = append(moves, gridutil.Coordinate{Row: p.Row + 2, Col: p.Col - 1})
		moves = append(moves, gridutil.Coordinate{Row: p.Row + 2, Col: p.Col + 1})
		moves = append(moves, gridutil.Coordinate{Row: p.Row - 1, Col: p.Col - 2})
		moves = append(moves, gridutil.Coordinate{Row: p.Row + 1, Col: p.Col - 2})
		moves = append(moves, gridutil.Coordinate{Row: p.Row - 1, Col: p.Col + 2})
		moves = append(moves, gridutil.Coordinate{Row: p.Row + 1, Col: p.Col + 2})

		var valid []gridutil.Coordinate
		for _, m := range moves {
			if _, exists := grid.GetC(m); exists {
				valid = append(valid, m)
			}
		}
		return valid
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.moves <= 4 {
			reachable.Add(current.pos)
		}

		if current.moves < 4 {
			for _, next := range getMoves(current.pos) {
				if prevMoves, seen := visited[next]; !seen || prevMoves > current.moves+1 {
					visited[next] = current.moves + 1
					queue = append(queue, struct {
						pos   gridutil.Coordinate
						moves int
					}{next, current.moves + 1})
				}
			}
		}
	}

	count := 0
	for _, pos := range reachable.Values() {
		if val, exists := grid.GetC(pos); exists && val == 'S' {
			count++
		}
	}

	fmt.Println(count)
}

func partII() {
	data, err := os.ReadFile("input2")
	if err != nil {
		data, err = os.ReadFile(filepath.Join("everybodycodes", "2025", "quest10", "input2"))
		if err != nil {
			log.Fatalf("reading input failed: %v", err)
		}
	}

	input := string(data)
	lines := conv.SplitNewline(input)

	grid := gridutil.NewCharGrid2D(lines)
	minRow, maxRow := grid.GetMinMaxRow()
	minCol, maxCol := grid.GetMinMaxCol()

	hideouts := container.NewSet[gridutil.Coordinate]()
	sheep := container.NewSet[gridutil.Coordinate]()
	var dragon gridutil.Coordinate

	for r := minRow; r <= maxRow; r++ {
		for c := minCol; c <= maxCol; c++ {
			coord := gridutil.Coordinate{Row: r, Col: c}
			if val, ok := grid.GetC(coord); ok {
				switch val {
				case '#':
					hideouts.Add(coord)
				case 'S':
					sheep.Add(coord)
				case 'D':
					dragon = coord
				}
			}
		}
	}

	knightMoves := []gridutil.Coordinate{
		{Row: -2, Col: -1},
		{Row: -2, Col: 1},
		{Row: -1, Col: -2},
		{Row: -1, Col: 2},
		{Row: 1, Col: -2},
		{Row: 1, Col: 2},
		{Row: 2, Col: -1},
		{Row: 2, Col: 1},
	}

	dragonPositions := container.NewSet[gridutil.Coordinate]()
	dragonPositions.Add(dragon)

	totalEaten := 0

	for range 20 {
		nextDragon := container.NewSet[gridutil.Coordinate]()
		for _, pos := range dragonPositions.Values() {
			for _, move := range knightMoves {
				nextPos := gridutil.Coordinate{Row: pos.Row + move.Row, Col: pos.Col + move.Col}
				if _, ok := grid.GetC(nextPos); ok {
					nextDragon.Add(nextPos)
				}
			}
		}

		for _, pos := range nextDragon.Values() {
			if sheep.Contains(pos) && !hideouts.Contains(pos) {
				sheep.Remove(pos)
				totalEaten++
			}
		}

		newSheep := container.NewSet[gridutil.Coordinate]()
		for _, pos := range sheep.Values() {
			nextPos := gridutil.Coordinate{Row: pos.Row + 1, Col: pos.Col}

			if _, ok := grid.GetC(nextPos); !ok {
				continue
			}

			if hideouts.Contains(nextPos) {
				newSheep.Add(nextPos)
				continue
			}

			if nextDragon.Contains(nextPos) {
				totalEaten++
				continue
			}

			newSheep.Add(nextPos)
		}

		sheep = newSheep
		dragonPositions = nextDragon
	}

	fmt.Println(totalEaten)
}

func partIII() {
	input := `.SSSSS.
.......
..#...#
#..###.
####.##
###D###`

	lines := conv.SplitNewline(strings.TrimSpace(input))
	result := countWinningSequences(lines)
	fmt.Println(result)
}

const eatenSheep = int(-1)

type sheepInfo struct {
	row int
	col int
}

type gameData struct {
	rows, cols  int
	hideouts    *container.Set[gridutil.Coordinate]
	sheepCols   []int
	colToIndex  map[int]int
	knightMoves []gridutil.Coordinate
}

func countWinningSequences(lines []string) int {
	if len(lines) == 0 {
		return 0
	}
	rows := len(lines)
	cols := len(lines[0])
	hideouts := container.NewSet[gridutil.Coordinate]()
	var dragon gridutil.Coordinate
	var sheepInfos []sheepInfo

	for r, line := range lines {
		for c, ch := range line {
			coord := gridutil.Coordinate{Row: r, Col: c}
			switch ch {
			case '#':
				hideouts.Add(coord)
			case 'S':
				sheepInfos = append(sheepInfos, sheepInfo{row: r, col: c})
			case 'D':
				dragon = coord
			}
		}
	}

	sort.Slice(sheepInfos, func(i, j int) bool {
		return sheepInfos[i].col < sheepInfos[j].col
	})

	sheepCols := make([]int, len(sheepInfos))
	sheepRows := make([]int, len(sheepInfos))
	colToIndex := make(map[int]int, len(sheepInfos))
	for idx, info := range sheepInfos {
		sheepCols[idx] = info.col
		sheepRows[idx] = int(info.row)
		colToIndex[info.col] = idx
	}

	knightMoves := []gridutil.Coordinate{
		{Row: -2, Col: -1},
		{Row: -2, Col: 1},
		{Row: -1, Col: -2},
		{Row: -1, Col: 2},
		{Row: 1, Col: -2},
		{Row: 1, Col: 2},
		{Row: 2, Col: -1},
		{Row: 2, Col: 1},
	}

	data := &gameData{
		rows:        rows,
		cols:        cols,
		hideouts:    hideouts,
		sheepCols:   sheepCols,
		colToIndex:  colToIndex,
		knightMoves: knightMoves,
	}

	memo := make(map[memoKey]int)
	return countSequences(dragon, sheepRows, true, data, memo)
}

type memoKey struct {
	dragonRow int
	dragonCol int
	sheepTurn bool
	sheep     string
}

func countSequences(dragon gridutil.Coordinate, sheepRows []int, sheepTurn bool, data *gameData, memo map[memoKey]int) int {
	if allSheepEaten(sheepRows) {
		return 1
	}

	key := memoKey{
		dragonRow: int(dragon.Row),
		dragonCol: int(dragon.Col),
		sheepTurn: sheepTurn,
		sheep:     encodeSheepRows(sheepRows),
	}
	if val, ok := memo[key]; ok {
		return val
	}

	var total int
	if sheepTurn {
		movesAvailable := false
		for idx, row := range sheepRows {
			if row == eatenSheep {
				continue
			}

			nextRow := int(row) + 1
			if nextRow >= data.rows {
				movesAvailable = true
				continue
			}

			target := gridutil.Coordinate{Row: nextRow, Col: data.sheepCols[idx]}
			if target == dragon && !data.hideouts.Contains(target) {
				continue
			}

			movesAvailable = true
			nextSheep := cloneSheepRows(sheepRows)
			nextSheep[idx] = int(nextRow)
			total += countSequences(dragon, nextSheep, false, data, memo)
		}

		if !movesAvailable {
			total = countSequences(dragon, sheepRows, false, data, memo)
		}
	} else {
		moveMade := false
		for _, move := range data.knightMoves {
			next := gridutil.Coordinate{Row: dragon.Row + move.Row, Col: dragon.Col + move.Col}
			if next.Row < 0 || next.Row >= data.rows || next.Col < 0 || next.Col >= data.cols {
				continue
			}

			moveMade = true
			nextSheep := sheepRows
			if idx, ok := data.colToIndex[next.Col]; ok {
				if nextSheep[idx] == int(next.Row) && !data.hideouts.Contains(next) {
					nextSheep = cloneSheepRows(nextSheep)
					nextSheep[idx] = eatenSheep
				}
			}

			total += countSequences(next, nextSheep, true, data, memo)
		}

		if !moveMade {
			total = 0
		}
	}

	memo[key] = total
	return total
}

func encodeSheepRows(rows []int) string {
	buf := make([]byte, 0, len(rows)*2)
	for _, row := range rows {
		val := uint(row + 2)
		buf = append(buf, byte(val>>8), byte(val))
	}
	return string(buf)
}

func cloneSheepRows(rows []int) []int {
	copyRows := make([]int, len(rows))
	copy(copyRows, rows)
	return copyRows
}

func allSheepEaten(rows []int) bool {
	for _, row := range rows {
		if row != eatenSheep {
			return false
		}
	}
	return true
}
