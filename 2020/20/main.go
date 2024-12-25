package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/gridutil"
	"aoc/internal/stringutil"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

func main() {
	input, err := download.ReadInput(2020, 20)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

type tile struct {
	id    int
	grid  gridutil.Grid2D[rune]
	edges edges
}

type edges struct {
	top    string
	bottom string
	left   string
	right  string
}

func getEdges(tile tile) edges {
	top, bottom, left, right := "", "", "", ""
	for c := range tile.grid.Width() {
		val, _ := tile.grid.Get(0, c)
		top += string(val)
		val, _ = tile.grid.Get(tile.grid.Height()-1, c)
		bottom += string(val)
	}
	for r := range tile.grid.Height() {
		val, _ := tile.grid.Get(r, 0)
		left += string(val)
		val, _ = tile.grid.Get(r, tile.grid.Width()-1)
		right += string(val)
	}
	return edges{top: top, bottom: bottom, left: left, right: right}
}

func parseTiles(input string) map[int]tile {
	tiles := make(map[int]tile)
	blocks := strings.Split(input, "\n\n")
	for _, block := range blocks {
		lines := conv.SplitNewline(block)
		if len(lines) == 0 {
			continue
		}
		idStr := strings.TrimSuffix(strings.TrimPrefix(lines[0], "Tile "), ":")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			panic(err)
		}
		grid := gridutil.NewCharGrid2D(lines[1:])
		tiles[id] = tile{id: id, grid: grid, edges: getEdges(tile{grid: grid})}
	}
	return tiles
}

func part1(input string) {
	tiles := parseTiles(input)
	edgeCounts := make(map[string]int)
	for _, tile := range tiles {
		edges := getEdges(tile)
		edgeCounts[edges.top]++
		edgeCounts[stringutil.Reverse(edges.top)]++
		edgeCounts[edges.bottom]++
		edgeCounts[stringutil.Reverse(edges.bottom)]++
		edgeCounts[edges.left]++
		edgeCounts[stringutil.Reverse(edges.left)]++
		edgeCounts[edges.right]++
		edgeCounts[stringutil.Reverse(edges.right)]++
	}

	product := 1
	for id, tile := range tiles {
		count := 0
		edges := getEdges(tile)
		if edgeCounts[edges.top] == 1 {
			count++
		}
		if edgeCounts[stringutil.Reverse(edges.top)] == 1 {
			count++
		}
		if edgeCounts[edges.bottom] == 1 {
			count++
		}
		if edgeCounts[stringutil.Reverse(edges.bottom)] == 1 {
			count++
		}
		if edgeCounts[edges.left] == 1 {
			count++
		}
		if edgeCounts[stringutil.Reverse(edges.left)] == 1 {
			count++
		}
		if edgeCounts[edges.right] == 1 {
			count++
		}
		if edgeCounts[stringutil.Reverse(edges.right)] == 1 {
			count++
		}

		if count == 4 {
			product *= id
		}
	}
	fmt.Println("Part 1", product)
}

func part2(input string) {
	tiles := parseTiles(input)
	answer := countNonMonsterTiles(generateImage(tiles))
	fmt.Println("Part 2", answer)
}

func generateGridSymmetries(grid gridutil.Grid2D[rune]) []gridutil.Grid2D[rune] {
	candidates := make([]gridutil.Grid2D[rune], 0, 8)

	for range 4 {
		last := grid
		newGrid := gridutil.NewGrid2D[rune](false)
		for r := range last.Height() {
			for c := range last.Width() {
				val, _ := last.Get(r, c)
				newGrid.Set(c, last.Height()-1-r, val)
			}
		}

		candidates = append(candidates, newGrid)

		mirror := gridutil.NewGrid2D[rune](false)
		for r := range newGrid.Height() {
			for c := range newGrid.Width() {
				val, _ := newGrid.Get(r, c)
				mirror.Set(newGrid.Height()-1-r, c, val)
			}
		}
		candidates = append(candidates, mirror)

		grid = newGrid
	}
	return candidates
}

func chooseTile(tiles map[int]tile, borderToTiles map[string]*container.Set[int],
	tiling map[gridutil.Coordinate]tile, col, row int, unusedTiles *container.Set[int]) (int, tile) {

	for _, num := range unusedTiles.Values() {
		for _, orientation := range generateGridSymmetries(tiles[num].grid) {
			borders := getEdges(tile{grid: orientation})
			top, _, _, left := borders.top, borders.right, borders.bottom, borders.left

			if col > 0 {
				leftNeighbor := getEdges(tiling[gridutil.Coordinate{Row: row, Col: col - 1}]).right
				if leftNeighbor != left {
					continue
				}
			} else {
				if borderToTiles[canonicalBorder(left)].Len() > 1 {
					continue
				}
			}

			if row > 0 {
				topNeighbor := getEdges(tiling[gridutil.Coordinate{Row: row - 1, Col: col}]).bottom
				if topNeighbor != top {
					continue
				}
			} else {
				if borderToTiles[canonicalBorder(top)].Len() > 1 {
					continue
				}
			}
			return num, tile{id: num, grid: orientation, edges: getEdges(tile{grid: orientation})}
		}
	}
	panic("No matching tile found")
}

func generateTiling(tiles map[int]tile) (int, map[gridutil.Coordinate]tile) {
	borderToTiles := getBorderToTilesMapping(tiles)
	unusedTiles := container.NewSet[int]()
	for num := range tiles {
		unusedTiles.Add(num)
	}

	dim := int(math.Sqrt(float64(len(tiles))))
	tiling := make(map[gridutil.Coordinate]tile)

	for x := range dim {
		for y := range dim {
			num, tile := chooseTile(tiles, borderToTiles, tiling, x, y, unusedTiles)
			unusedTiles.Remove(num)
			tiling[gridutil.Coordinate{Row: y, Col: x}] = tile
		}
	}
	return dim, tiling
}

func canonicalBorder(border string) string {
	reversed := stringutil.Reverse(border)
	if border < reversed {
		return border
	}
	return reversed
}

func getBorderToTilesMapping(tiles map[int]tile) map[string]*container.Set[int] {
	borderToTiles := make(map[string]*container.Set[int])
	for id, tile := range tiles {
		edges := getEdges(tile)
		borders := []string{edges.top, edges.bottom, edges.left, edges.right}
		for _, border := range borders {
			canonical := canonicalBorder(border)
			if borderToTiles[canonical] == nil {
				borderToTiles[canonical] = container.NewSet[int]()
			}
			borderToTiles[canonical].Add(id)
		}
	}
	return borderToTiles
}

func generateImage(tiles map[int]tile) gridutil.Grid2D[rune] {
	dim, tiling := generateTiling(tiles)
	image := gridutil.NewGrid2D[rune](false)

	for y := range dim {
		for x := range dim {
			grid := tiling[gridutil.Coordinate{Row: y, Col: x}].grid
			for r := 1; r < grid.Height()-1; r++ {
				for c := 1; c < grid.Width()-1; c++ {
					val, _ := grid.Get(r, c)
					image.Set(y*(grid.Height()-2)+r-1, x*(grid.Width()-2)+c-1, val)
				}
			}
		}
	}
	return image
}

func getMonsterTiles(image gridutil.Grid2D[rune]) map[gridutil.Coordinate]bool {
	monsterPattern := []string{
		"                  # ",
		"#    ##    ##    ###",
		" #  #  #  #  #  #   ",
	}

	var monsterCoords []gridutil.Coordinate
	for y, line := range monsterPattern {
		for x, c := range line {
			if c == '#' {
				monsterCoords = append(monsterCoords, gridutil.Coordinate{Row: y, Col: x})
			}
		}
	}

	maxX := 0
	maxY := 0
	for _, p := range monsterCoords {
		if p.Col > maxX {
			maxX = p.Col
		}
		if p.Row > maxY {
			maxY = p.Row
		}
	}

	monsters := make(map[gridutil.Coordinate]bool)
	for y := range image.Height() - maxY {
		for x := range image.Width() - maxX {
			isMonster := true
			for _, coord := range monsterCoords {
				val, _ := image.Get(y+coord.Row, x+coord.Col)
				if val != '#' {
					isMonster = false
					break
				}
			}
			if isMonster {
				for _, coord := range monsterCoords {
					monsters[gridutil.Coordinate{Row: y + coord.Row, Col: x + coord.Col}] = true
				}
			}
		}
	}
	return monsters
}

func countNonMonsterTiles(image gridutil.Grid2D[rune]) int {
	for _, img := range generateGridSymmetries(image) {
		monsterTiles := getMonsterTiles(img)
		if len(monsterTiles) == 0 {
			continue
		}

		allPounds := make(map[gridutil.Coordinate]bool)
		for y := range img.Height() {
			for x := range img.Width() {
				val, _ := img.Get(y, x)
				if val == '#' {
					allPounds[gridutil.Coordinate{Row: y, Col: x}] = true
				}
			}
		}

		count := 0
		for p := range allPounds {
			if !monsterTiles[p] {
				count++
			}
		}
		return count
	}
	panic("No monsters found in any orientation")
}
