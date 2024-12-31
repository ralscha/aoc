package main

import (
	"aoc/2019/intcomputer"
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/gridutil"
	"fmt"
	"log"
	"strconv"
)

func main() {
	input, err := download.ReadInput(2019, 17)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

type Camera struct {
	x, y int
	grid gridutil.Grid2D[rune]
	raw  string
}

func (c *Camera) scan(computer *intcomputer.IntcodeComputer) error {
	for {
		result, err := computer.Run()
		if err != nil {
			return err
		}

		switch result.Signal {
		case intcomputer.SignalOutput:
			char := rune(result.Value)
			c.raw += string(char)
			if char == '\n' {
				c.y++
				c.x = 0
			} else {
				c.grid.Set(c.y, c.x, char)
				c.x++
			}
		case intcomputer.SignalEnd:
			return nil
		case intcomputer.SignalInput:
			return fmt.Errorf("unexpected input request")
		}
	}
}

func (c *Camera) isIntersection(coord gridutil.Coordinate) bool {
	if val, ok := c.grid.GetC(coord); !ok {
		return false
	} else if val != '#' && val != '<' && val != '>' && val != '^' && val != 'v' {
		return false
	}

	away := 0
	for _, n := range c.grid.GetNeighbours4C(coord) {
		if n == '#' || n == '<' || n == '>' || n == '^' || n == 'v' {
			away++
		}
	}
	return away > 2
}

func part1(input string) {
	program := conv.ToIntSliceComma(input)
	computer := intcomputer.NewIntcodeComputer(program)
	camera := Camera{
		grid: gridutil.NewGrid2D[rune](false),
	}

	if err := camera.scan(computer); err != nil {
		log.Fatalf("scanning with camera failed: %v", err)
	}

	var intersects []gridutil.Coordinate
	minRow, maxRow := camera.grid.GetMinMaxRow()
	minCol, maxCol := camera.grid.GetMinMaxCol()
	for row := minRow; row <= maxRow; row++ {
		for col := minCol; col <= maxCol; col++ {
			coord := gridutil.Coordinate{Row: row, Col: col}
			if camera.isIntersection(coord) {
				intersects = append(intersects, coord)
			}
		}
	}

	alignmentSum := 0
	for _, i := range intersects {
		alignmentSum += i.Col * i.Row
	}
	fmt.Println("Part 1", alignmentSum)
}

type robot struct {
	grid   gridutil.Grid2D[rune]
	loc    gridutil.Coordinate
	curDir gridutil.Direction
}

func newRobot(grid gridutil.Grid2D[rune]) *robot {
	var start gridutil.Coordinate
	var facing rune
	for row := range grid.Height() {
		for col := range grid.Width() {
			coord := gridutil.Coordinate{Row: row, Col: col}
			if val, ok := grid.GetC(coord); ok {
				if val == '<' || val == '^' || val == '>' || val == 'v' {
					start = coord
					facing = val
					break
				}
			}
		}
	}

	curDir := map[rune]gridutil.Direction{
		'^': gridutil.DirectionN,
		'>': gridutil.DirectionE,
		'v': gridutil.DirectionS,
		'<': gridutil.DirectionW,
	}[facing]

	return &robot{
		grid:   grid,
		loc:    start,
		curDir: curDir,
	}
}

func (r *robot) findPath() []any {
	var steps []any
	move := r.nextMove()
	for move != nil {
		steps = append(steps, move)
		move = r.nextMove()
	}
	return steps
}

func (r *robot) isScaffold(loc gridutil.Coordinate) bool {
	val, ok := r.grid.GetC(loc)
	return ok && (val == '#' || val == '<' || val == '>' || val == '^' || val == 'v')
}

func (r *robot) nextMove() any {
	steps := 0
	nextLoc := gridutil.Coordinate{Row: r.loc.Row + r.curDir.Row, Col: r.loc.Col + r.curDir.Col}
	for r.isScaffold(nextLoc) {
		steps++
		r.loc = nextLoc
		nextLoc = gridutil.Coordinate{Row: r.loc.Row + r.curDir.Row, Col: r.loc.Col + r.curDir.Col}
	}
	if steps > 0 {
		return steps
	}

	r.curDir = gridutil.TurnLeft(r.curDir)
	nextLoc = gridutil.Coordinate{Row: r.loc.Row + r.curDir.Row, Col: r.loc.Col + r.curDir.Col}
	if r.isScaffold(nextLoc) {
		return "L"
	}

	r.curDir = gridutil.TurnRight(gridutil.TurnRight(r.curDir))
	nextLoc = gridutil.Coordinate{Row: r.loc.Row + r.curDir.Row, Col: r.loc.Col + r.curDir.Col}
	if r.isScaffold(nextLoc) {
		return "R"
	}
	return nil
}

func countRepetitions(pattern []any, full []any, start int) int {
	patLen := len(pattern)
	c := 0
	for idx := start; idx <= len(full)-patLen; idx++ {
		match := true
		for i := 0; i < patLen; i++ {
			if full[idx+i] != pattern[i] {
				match = false
				break
			}
		}
		if match {
			c++
		}
	}
	return c
}

func subsequenceFrequencies(path []any) map[string]int {
	freqs := map[string]int{}
	maxLen := (len(path) / 2) + 1
	for length := 1; length < maxLen; length++ {
		for start := 0; start <= len(path)-length; start++ {
			pattern := path[start : start+length]
			patternStr := fmt.Sprint(pattern)
			if _, ok := freqs[patternStr]; !ok {
				c := countRepetitions(pattern, path, start)
				freqs[patternStr] = c
			}
		}
	}
	return freqs
}

type progBuilder struct {
	path []any
	subs [][]any
	prog string
}

func newProgBuilder(path []any) *progBuilder {
	freqs := subsequenceFrequencies(path)
	var subs [][]any
	for k, v := range freqs {
		var sub []any
		inBracket := false
		currentNum := ""
		for _, char := range k {
			if char == '[' {
				inBracket = true
			} else if char == ']' {
				inBracket = false
				if currentNum != "" {
					if num, err := strconv.Atoi(currentNum); err == nil {
						sub = append(sub, num)
					}
					currentNum = ""
				}
			} else if inBracket {
				if char >= '0' && char <= '9' {
					currentNum += string(char)
				} else if char == 'L' || char == 'R' {
					if currentNum != "" {
						if num, err := strconv.Atoi(currentNum); err == nil {
							sub = append(sub, num)
						}
						currentNum = ""
					}
					sub = append(sub, string(char))
				} else if char == ',' {
					if currentNum != "" {
						if num, err := strconv.Atoi(currentNum); err == nil {
							sub = append(sub, num)
						}
						currentNum = ""
					}
				}
			}
		}

		subStr := ""
		for _, val := range sub {
			subStr += fmt.Sprintf("%v,", val)
		}
		if len(subStr) > 0 {
			subStr = subStr[:len(subStr)-1]
		}

		if len(subStr) <= 20 && v > 1 {
			subs = append(subs, sub)
		}
	}

	for i := range len(subs) {
		for j := i + 1; j < len(subs); j++ {
			if len(subs[i]) < len(subs[j]) {
				subs[i], subs[j] = subs[j], subs[i]
			}
		}
	}

	builder := &progBuilder{
		path: path,
		subs: subs,
	}

	main, parts := builder.buildProgram(nil, nil)

	mainStr := ""
	for _, p := range main {
		mainStr += string(rune(int('A')+p)) + ","
	}
	if len(mainStr) > 0 {
		mainStr = mainStr[:len(mainStr)-1]
	}

	prog := mainStr + "\n"
	for _, p := range parts {
		partStr := ""
		for _, v := range p {
			partStr += fmt.Sprintf("%v,", v)
		}
		if len(partStr) > 0 {
			partStr = partStr[:len(partStr)-1]
		}
		prog += partStr + "\n"
	}
	prog += "n\n"

	builder.prog = prog
	return builder
}

func (pb *progBuilder) buildProgram(main []int, parts [][]any) ([]int, [][]any) {
	expanded := pb.expand(main, parts)
	if fmt.Sprint(expanded) == fmt.Sprint(pb.path) {
		return main, parts
	}

	start := len(expanded)
	opts := pb.options(start)
	for _, o := range opts {
		nextParts := make([][]any, len(parts))
		copy(nextParts, parts)
		found := false
		idx := 0
		for i, p := range nextParts {
			if fmt.Sprint(p) == fmt.Sprint(o) {
				found = true
				idx = i
				break
			}
		}
		if !found {
			nextParts = append(nextParts, o)
		}
		if len(nextParts) > 3 {
			continue
		}
		if !found {
			idx = len(nextParts) - 1
		}

		nextMain := append(main, idx)

		if resMain, resParts := pb.buildProgram(nextMain, nextParts); resMain != nil {
			return resMain, resParts
		}
	}

	return nil, nil
}

func (pb *progBuilder) expand(main []int, parts [][]any) []any {
	var prog []any
	for _, idx := range main {
		prog = append(prog, parts[idx]...)
	}
	return prog
}

func (pb *progBuilder) options(start int) [][]any {
	var opts [][]any
	for _, i := range pb.subs {
		if len(i) < len(pb.path)-start+1 {
			match := true
			for j := 0; j < len(i); j++ {
				if pb.path[start+j] != i[j] {
					match = false
					break
				}
			}
			if match {
				opts = append(opts, i)
			}
		}
	}
	return opts
}

func (pb *progBuilder) getProg() string {
	return pb.prog
}

func part2(input string) {
	program := conv.ToIntSliceComma(input)
	computer := intcomputer.NewIntcodeComputer(program)
	camera := Camera{
		grid: gridutil.NewGrid2D[rune](false),
	}

	if err := camera.scan(computer); err != nil {
		log.Fatalf("scanning with camera failed: %v", err)
	}

	robot := newRobot(camera.grid)
	path := robot.findPath()
	progBuilder := newProgBuilder(path)
	programStr := progBuilder.getProg()

	program[0] = 2
	computer = intcomputer.NewIntcodeComputer(program)

	if err := computer.AddInput(programStr); err != nil {
		log.Fatalf("adding input failed: %v", err)
	}

	var lastOutput int
	for {
		result, err := computer.Run()
		if err != nil {
			log.Fatalf("running program failed: %v", err)
		}

		switch result.Signal {
		case intcomputer.SignalOutput:
			lastOutput = result.Value
		case intcomputer.SignalEnd:
			fmt.Println("Part 2", lastOutput)
			return
		case intcomputer.SignalInput:
			log.Fatal("unexpected input request")
		}
	}
}
