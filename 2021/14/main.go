package main

import (
	"aoc/internal/download"
	"bufio"
	"fmt"
	"log"
	"slices"
	"strings"
)

func main() {
	input, err := download.ReadInput(2021, 14)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Scan()
	template := scanner.Text()
	rules := make(map[string]string)

	for scanner.Scan() {
		line := scanner.Text()
		before, after, ok := strings.Cut(line, " -> ")
		if ok {
			rules[before] = after
		}
	}

	for range 10 {
		newTemplate := template[:1]
		ix := 0
		for ix < len(template)-1 {
			k := template[ix : ix+2]
			if v, ok := rules[k]; ok {
				newTemplate = newTemplate + v + k[1:]
			}
			ix++
		}

		template = newTemplate
	}

	counts := make(map[string]int)
	for _, c := range template {
		counts[string(c)]++
	}

	minCounts := ""
	maxCounts := ""
	for k, v := range counts {
		if minCounts == "" || v < counts[minCounts] {
			minCounts = k
		}
		if maxCounts == "" || v > counts[maxCounts] {
			maxCounts = k
		}
	}

	fmt.Println("Part 1", counts[maxCounts]-counts[minCounts])
}

func part2(input string) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Scan()
	template := scanner.Text()
	rules := make(map[string]string)

	for scanner.Scan() {
		line := scanner.Text()
		before, after, ok := strings.Cut(line, " -> ")
		if ok {
			rules[before] = after
		}
	}

	pairs := make(map[string]int)
	for i := range len(template) - 1 {
		pairs[template[i:i+2]] += 1
	}

	for range 40 {
		newpairs := make(map[string]int)
		for k, count := range pairs {
			if v, ok := rules[k]; ok {
				newpairs[k[:1]+v] += count
				newpairs[v+k[1:]] += count
			}
		}
		pairs = newpairs
	}

	chars := make(map[string]int)
	for k, v := range pairs {
		chars[k[:1]] += v
		chars[k[1:]] += v
	}

	var values []int
	for _, v := range chars {
		values = append(values, v)
	}
	maxValue := slices.Max(values)
	minValue := slices.Min(values)

	fmt.Println("Part 2", (maxValue-minValue)/2+1)
}
