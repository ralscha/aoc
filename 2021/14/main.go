package main

import (
	"aoc/internal/download"
	"bufio"
	"fmt"
	"log"
	"math"
	"strings"
)

func main() {
	inputFile := "./2021/14/input.txt"
	input, err := download.ReadInput(inputFile, 2021, 14)
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
		pos := strings.Index(line, " -> ")
		if pos != -1 {
			rules[line[:pos]] = line[pos+4:]
		}
	}

	for i := 0; i < 10; i++ {
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

	min := ""
	max := ""
	for k, v := range counts {
		if min == "" || v < counts[min] {
			min = k
		}
		if max == "" || v > counts[max] {
			max = k
		}
	}

	fmt.Println("Result: ", counts[max]-counts[min])
}

func part2(input string) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Scan()
	template := scanner.Text()
	rules := make(map[string]string)

	for scanner.Scan() {
		line := scanner.Text()
		pos := strings.Index(line, " -> ")
		if pos != -1 {
			rules[line[:pos]] = line[pos+4:]
		}
	}

	pairs := make(map[string]int)
	for i := 0; i < len(template)-1; i++ {
		pairs[template[i:i+2]] += 1
	}

	for i := 0; i < 40; i++ {
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

	max, min := math.MinInt, math.MaxInt

	for _, v := range chars {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	fmt.Println("Result: ", (max-min)/2+1)
}
