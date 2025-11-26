package main

import (
	"aoc/internal/conv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Dance struct {
	cols [][]int
}

func NewDance(input []string) *Dance {
	firstRow := strings.Fields(input[0])
	cols := make([][]int, len(firstRow))
	for i, n := range firstRow {
		cols[i] = []int{conv.MustAtoi(n)}
	}
	for _, l := range input[1:] {
		nums := strings.Fields(l)
		for i, n := range nums {
			cols[i] = append(cols[i], conv.MustAtoi(n))
		}
	}

	return &Dance{
		cols: cols,
	}
}

func (d *Dance) simulateClap(round int) int {
	clapperCol := (round - 1) % len(d.cols)
	clapperNum := d.cols[clapperCol][0]
	d.moveUp(clapperCol)

	targetCol := (clapperCol + 1) % len(d.cols)
	numberOfDancers := len(d.cols[targetCol])

	div := clapperNum / numberOfDancers
	if div%2 == 0 {
		targetPos := clapperNum%numberOfDancers - 1
		if targetPos == -1 {
			targetPos = 0
		}
		d.insertBefore(targetPos, targetCol, clapperNum)
	} else {
		targetPos := numberOfDancers - clapperNum%numberOfDancers
		d.insertAfter(targetPos, targetCol, clapperNum)
	}

	var result string
	for _, c := range d.cols {
		result += strconv.Itoa(c[0])
	}
	return conv.MustAtoi(result)
}

func (d *Dance) moveUp(col int) {
	currentCol := d.cols[col]
	for i := range len(currentCol) - 1 {
		currentCol[i] = currentCol[i+1]
	}
	d.cols[col] = currentCol[:len(currentCol)-1]
}

func (d *Dance) insertBefore(pos, col, num int) {
	currentCol := d.cols[col]
	newCol := make([]int, 0, len(currentCol)+1)
	newCol = append(newCol, currentCol[:pos]...)
	newCol = append(newCol, num)
	newCol = append(newCol, currentCol[pos:]...)
	d.cols[col] = newCol
}

func (d *Dance) insertAfter(pos, col, num int) {
	currentCol := d.cols[col]
	newCol := make([]int, 0, len(currentCol)+1)
	if pos == len(currentCol) {
		newCol = append(newCol, currentCol...)
		newCol = append(newCol, num)
	} else {
		newCol = append(newCol, currentCol[:pos+1]...)
		newCol = append(newCol, num)
		newCol = append(newCol, currentCol[pos+1:]...)
	}
	d.cols[col] = newCol
}

func main() {
	partI()
	partII()
	partIII()
}

func partI() {
	input, err := os.ReadFile("everybodycodes/quest5/partI.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := conv.SplitNewline(string(input))

	dance := NewDance(lines)
	var result int
	for round := 1; round <= 10; round++ {
		result = dance.simulateClap(round)
	}
	fmt.Println("Part I", result)
}

func partII() {
	input, err := os.ReadFile("everybodycodes/quest5/partII.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := conv.SplitNewline(string(input))

	dance := NewDance(lines)
	counts := make(map[int]int)
	for round := 1; ; round++ {
		result := dance.simulateClap(round)
		counts[result]++
		if counts[result] == 2024 {
			fmt.Println("Part II", result*round)
			break
		}
	}
}

func partIII() {
	input, err := os.ReadFile("everybodycodes/quest5/partIII.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := conv.SplitNewline(string(input))

	dance := NewDance(lines)

	highest := 0
	highestRound := 0
	for round := 1; ; round++ {
		result := dance.simulateClap(round)
		if result > highest {
			highest = result
			highestRound = round
		}
		if round-highestRound > 100 {
			break
		}
	}
	fmt.Println("Part III", highest)
}
