package main

import (
	"aoc/internal/conv"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type SpineSegment struct {
	value int
	left  *int
	right *int
}

type Sword struct {
	id      int
	quality int
	levels  []int
}

func main() {
	partI()
	partII()
	partIII()
}

func partI() {
	input := `12:5,3,3,6,2,8,7,6,6,5,5,4,4,7,1,9,3,8,7,6,7,8,4,4,9,5,2,7,2,4`

	numbers := conv.ToIntSliceComma(strings.Split(input, ":")[1])
	spine := buildFishbone(numbers)
	quality := calculateQuality(spine)

	fmt.Println(quality)
}

func buildFishbone(numbers []int) []SpineSegment {
	spine := []SpineSegment{}

	if len(numbers) > 0 {
		spine = append(spine, SpineSegment{value: numbers[0]})
	}

	for i := 1; i < len(numbers); i++ {
		num := numbers[i]
		placed := false

		for j := range len(spine) {
			if num < spine[j].value && spine[j].left == nil {
				spine[j].left = &num
				placed = true
				break
			} else if num > spine[j].value && spine[j].right == nil {
				spine[j].right = &num
				placed = true
				break
			}
		}

		if !placed {
			spine = append(spine, SpineSegment{value: num})
		}
	}

	return spine
}

func partII() {
	input := `1:8,6,8,6,3,5,8,6,6,5,6,6,9,2,5,4,9,6,1,4,7,3,3,2,2,5,5,4,8,5
2:5,2,8,3,1,8,2,3,4,5,4,9,8,4,1,2,6,6,7,2,1,1,7,3,2,3,9,5,2,7`

	lines := strings.Split(strings.TrimSpace(input), "\n")

	var qualities []int
	for _, line := range lines {
		numbers := conv.ToIntSliceComma(strings.Split(line, ":")[1])
		spine := buildFishbone(numbers)
		qualities = append(qualities, calculateQuality(spine))
	}

	difference := slices.Max(qualities) - slices.Min(qualities)
	fmt.Println(difference)
}

func partIII() {
	input := `1:9,9,4,9,5,6,1,4,7,8,4,1,6,4,6,5,5,2,6,6,8,2,4,1,5,7,8,7,6,6
2:1,9,4,5,9,5,5,7,4,6,9,9,1,1,7,9,8,3,4,3,3,6,5,1,2,3,7,4,4,1
3:7,6,7,8,6,8,2,7,7,1,1,5,4,3,6,1,7,4,4,3,4,7,5,4,7,6,3,2,1,4`

	lines := strings.Split(strings.TrimSpace(input), "\n")

	var swords []Sword
	for _, line := range lines {
		parts := strings.Split(line, ":")
		id := conv.MustAtoi(parts[0])
		numbers := conv.ToIntSliceComma(parts[1])

		spine := buildFishbone(numbers)
		quality := calculateQuality(spine)
		levels := calculateLevels(spine)
		swords = append(swords, Sword{id: id, quality: quality, levels: levels})
	}

	slices.SortFunc(swords, func(a, b Sword) int {
		if a.quality != b.quality {
			return b.quality - a.quality
		}

		maxLen := max(len(b.levels), len(a.levels))
		for i := range maxLen {
			aLevel := 0
			bLevel := 0
			if i < len(a.levels) {
				aLevel = a.levels[i]
			}
			if i < len(b.levels) {
				bLevel = b.levels[i]
			}

			if aLevel != bLevel {
				return bLevel - aLevel
			}
		}

		return b.id - a.id
	})

	checksum := 0
	for pos, sword := range swords {
		checksum += (pos + 1) * sword.id
	}

	fmt.Println(checksum)
}

func calculateQuality(spine []SpineSegment) int {
	var quality strings.Builder
	for _, seg := range spine {
		quality.WriteString(strconv.Itoa(seg.value))
	}
	return conv.MustAtoi(quality.String())
}

func calculateLevels(spine []SpineSegment) []int {
	var levels []int

	for i := range spine {
		var levelStr strings.Builder

		if spine[i].left != nil {
			levelStr.WriteString(strconv.Itoa(*spine[i].left))
		}

		levelStr.WriteString(strconv.Itoa(spine[i].value))

		if spine[i].right != nil {
			levelStr.WriteString(strconv.Itoa(*spine[i].right))
		}

		levels = append(levels, conv.MustAtoi(levelStr.String()))
	}

	return levels
}
