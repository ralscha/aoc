package main

import (
	"aoc/internal/download"
	"aoc/internal/mathx"
	"fmt"
	"log"
	"math"
)

// https://github.com/jonathanpaulson/AdventOfCode/blob/master/2021/23.py

type state struct {
	bot map[string][]string
	top []string
}

func (s state) String() string {
	return fmt.Sprintf("%v %v", s.bot, s.top)
}

var cost = map[string]int{
	"A": 1,
	"B": 10,
	"C": 100,
	"D": 1000,
}

func done(state state) bool {
	for k, v := range state.bot {
		for _, vv := range v {
			if vv != k {
				return false
			}
		}
	}
	return true
}

func canMoveFrom(k string, col []string) bool {
	for _, c := range col {
		if c != k && c != "E" {
			return true
		}
	}
	return false
}

func canMoveTo(k string, col []string) bool {
	for _, c := range col {
		if c != k && c != "E" {
			return false
		}
	}
	return true
}

func botIdx(bot string) int {
	return map[string]int{"A": 2, "B": 4, "C": 6, "D": 8}[bot]
}

func topIdx(col []string) int {
	for i, c := range col {
		if c != "E" {
			return i
		}
	}
	return -1
}

func destIdx(col []string) int {
	for i := len(col) - 1; i >= 0; i-- {
		if col[i] == "E" {
			return i
		}
	}
	return -1
}

func between(a int, bot string, top int) bool {
	return (botIdx(bot) < a && a < top) || (top < a && a < botIdx(bot))
}

func clearPath(bot string, topIdx int, top []string) bool {
	for ti := range top {
		if between(ti, bot, topIdx) && top[ti] != "E" {
			return false
		}
	}
	return true
}

var memo = make(map[string]int)

func f(s state) int {
	if done(s) {
		return 0
	}
	if cost, ok := memo[s.String()]; ok {
		return cost
	}

	ans := math.MaxInt32
	for i, c := range s.top {
		if c != "E" && canMoveTo(c, s.bot[c]) && clearPath(c, i, s.top) {
			di := destIdx(s.bot[c])
			dist := di + 1 + mathx.Abs(botIdx(c)-i)
			cost := cost[c] * dist
			newTop := make([]string, len(s.top))
			copy(newTop, s.top)
			newTop[i] = "E"
			newBot := deepcopy(s.bot)
			newBot[c][di] = c
			ans = min(ans, cost+f(state{bot: newBot, top: newTop}))
		}
	}

	for k, col := range s.bot {
		if !canMoveFrom(k, col) {
			continue
		}
		ki := topIdx(col)
		if ki == -1 {
			continue
		}
		c := col[ki]
		for to := range s.top {
			if to == 2 || to == 4 || to == 6 || to == 8 || s.top[to] != "E" {
				continue
			}
			if clearPath(k, to, s.top) {
				dist := ki + 1 + mathx.Abs(to-botIdx(k))
				newTop := make([]string, len(s.top))
				copy(newTop, s.top)
				newTop[to] = c
				newBot := deepcopy(s.bot)
				newBot[k][ki] = "E"
				ans = min(ans, cost[c]*dist+f(state{bot: newBot, top: newTop}))
			}
		}
	}

	memo[s.String()] = ans
	return ans
}

func deepcopy(m map[string][]string) map[string][]string {
	newMap := make(map[string][]string)
	for k, v := range m {
		newSlice := make([]string, len(v))
		copy(newSlice, v)
		newMap[k] = newSlice
	}
	return newMap
}

func main() {
	_, err := download.ReadInput(2021, 23)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1and2()
}

func part1and2() {

	A1 := []string{"D", "B"}
	B1 := []string{"D", "A"}
	C1 := []string{"C", "B"}
	D1 := []string{"C", "A"}

	A := []string{"D", "D", "D", "B"}
	B := []string{"D", "C", "B", "A"}
	C := []string{"C", "B", "A", "B"}
	D := []string{"C", "A", "C", "A"}

	top := make([]string, 11)
	for i := range top {
		top[i] = "E"
	}

	part1 := state{bot: map[string][]string{"A": A1, "B": B1, "C": C1, "D": D1}, top: top}
	part2 := state{bot: map[string][]string{"A": A, "B": B, "C": C, "D": D}, top: top}

	fmt.Println("Part 1", f(part1))
	fmt.Println("Part 2", f(part2))

}
