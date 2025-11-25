package main

import (
	"fmt"
	"os"
	"strings"

	"aoc/internal/conv"
)

func main() {
	partI()
	partII()
	partIII()
}

func partI() {
	input, _ := os.ReadFile("./input1")
	spell := conv.ToIntSliceComma(strings.TrimSpace(string(input)))

	wallLength := 90
	columns := make([]int, wallLength+1)

	for _, n := range spell {
		for col := n; col <= wallLength; col += n {
			columns[col]++
		}
	}

	total := 0
	for _, height := range columns {
		total += height
	}

	fmt.Println(total)
}

func partII() {
	input, _ := os.ReadFile("./input2")
	wall := conv.ToIntSliceComma(strings.TrimSpace(string(input)))

	wallLength := len(wall)
	columns := make([]int, wallLength+1)
	for i, v := range wall {
		columns[i+1] = v
	}

	spell := []int{}

	for n := 1; n <= wallLength; n++ {
		explained := 0
		for _, s := range spell {
			if n%s == 0 {
				explained++
			}
		}

		if columns[n] > explained {
			spell = append(spell, n)
		}
	}

	product := 1
	for _, s := range spell {
		product *= s
	}

	fmt.Println(product)
}

func partIII() {
	input, _ := os.ReadFile("./input3")
	wall := conv.ToIntSliceComma(strings.TrimSpace(string(input)))

	wallLength := len(wall)
	columns := make([]int, wallLength+1)
	for i, v := range wall {
		columns[i+1] = v
	}

	spell := []int{}
	for n := 1; n <= wallLength; n++ {
		explained := 0
		for _, s := range spell {
			if n%s == 0 {
				explained++
			}
		}
		if columns[n] > explained {
			spell = append(spell, n)
		}
	}

	calcBlocks := func(length int64) int64 {
		var total int64 = 0
		for _, s := range spell {
			total += length / int64(s)
		}
		return total
	}

	targetBlocks := int64(202520252025000)

	low := int64(1)
	high := int64(202520252025000)

	for low < high {
		mid := (low + high + 1) / 2
		blocks := calcBlocks(mid)
		if blocks <= targetBlocks {
			low = mid
		} else {
			high = mid - 1
		}
	}

	fmt.Println(low)
}
