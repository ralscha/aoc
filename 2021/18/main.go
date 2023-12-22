package main

import (
	"aoc/internal/download"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"strings"
)

// https://github.com/benediktwerner/AdventOfCode/blob/master/2021/day18/sol.py

func main() {
	input, err := download.ReadInput(2021, 18)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1and2(input)
}

func addLeft(x interface{}, n interface{}) interface{} {
	if n == nil {
		return x
	}
	switch v := x.(type) {
	case float64:
		return v + n.(float64)
	case []interface{}:
		return []interface{}{addLeft(v[0], n), v[1]}
	}
	return x
}

func addRight(x interface{}, n interface{}) interface{} {
	if n == nil {
		return x
	}
	switch v := x.(type) {
	case float64:
		return v + n.(float64)
	case []interface{}:
		return []interface{}{v[0], addRight(v[1], n)}
	}
	return x
}

func explode(x interface{}, n float64) (bool, interface{}, interface{}, interface{}) {
	switch v := x.(type) {
	case float64:
		return false, nil, v, nil
	case []interface{}:
		if n == 0 {
			return true, v[0], float64(0), v[1]
		}
		exp, left, a, right := explode(v[0], n-1)
		if exp {
			return true, left, []interface{}{a, addLeft(v[1], right)}, nil
		}
		exp, left, b, right := explode(v[1], n-1)
		if exp {
			return true, nil, []interface{}{addRight(v[0], left), b}, right
		}
	}
	return false, nil, x, nil
}

func split(x interface{}) (bool, interface{}) {
	switch v := x.(type) {
	case float64:
		if int(v) >= 10 {
			return true, []interface{}{float64(int(v / 2)), float64(int(math.Ceil(float64(v) / 2)))}
		}
		return false, v
	case []interface{}:
		change, a := split(v[0])
		if change {
			return true, []interface{}{a, v[1]}
		}
		change, b := split(v[1])
		return change, []interface{}{v[0], b}
	}
	return false, x
}

func add(a, b interface{}) interface{} {
	x := []interface{}{a, b}
	for {
		change, _, xNew, _ := explode(x, 4)
		x = xNew.([]interface{})
		if change {
			continue
		}
		change, xNew = split(x)
		x = xNew.([]interface{})
		if !change {
			break
		}
	}
	return x
}

func magnitude(x interface{}) float64 {
	switch v := x.(type) {
	case float64:
		return v
	case []interface{}:
		return 3*magnitude(v[0]) + 2*magnitude(v[1])
	}
	return 0
}

func reduce(fn func(a, b interface{}) interface{}, list []interface{}) interface{} {
	result := list[0]
	for i := 1; i < len(list); i++ {
		result = fn(result, list[i])
	}
	return result
}

func splitLines(s string) []string {
	return strings.Split(strings.TrimSpace(s), "\n")
}

func part1and2(input string) {
	input = strings.TrimSpace(input)
	lines := make([]interface{}, 0)
	for _, line := range splitLines(input) {
		var parsed interface{}
		if err := json.Unmarshal([]byte(line), &parsed); err != nil {
			panic(err)
		}
		lines = append(lines, parsed)
	}
	fmt.Println(magnitude(reduce(add, lines)))

	var maxMagnitude float64 = 0
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines); j++ {
			if i != j {
				mag := magnitude(add(lines[i], lines[j]))
				if mag > maxMagnitude {
					maxMagnitude = mag
				}
			}
		}
	}
	fmt.Println(maxMagnitude)
}
