package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"strings"
)

func main() {
	input, err := download.ReadInput(2017, 24)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	bridge := newBridge(input)

	strongest := bridge.findBestBridge(0, make([]bool, len(bridge.components)), func(curr, best bridgeResult) bool {
		return curr.strength > best.strength
	})
	fmt.Println("Part 1", strongest.strength)
}

func part2(input string) {
	bridge := newBridge(input)
	longest := bridge.findBestBridge(0, make([]bool, len(bridge.components)), func(curr, best bridgeResult) bool {
		return curr.length > best.length ||
			(curr.length == best.length && curr.strength > best.strength)
	})
	fmt.Println("Part 2", longest.strength)
}

type component struct {
	a, b int
}

func (c component) strength() int {
	return c.a + c.b
}

func (c component) canConnect(port int) (bool, int) {
	switch port {
	case c.a:
		return true, c.b
	case c.b:
		return true, c.a
	default:
		return false, 0
	}
}

type bridge struct {
	components []component
}

func newBridge(input string) *bridge {
	lines := conv.SplitNewline(input)
	components := make([]component, len(lines))

	for i, line := range lines {
		parts := strings.Split(line, "/")
		components[i] = component{
			a: conv.MustAtoi(parts[0]),
			b: conv.MustAtoi(parts[1]),
		}
	}

	return &bridge{components: components}
}

type bridgeResult struct {
	length   int
	strength int
}

func (b *bridge) findBestBridge(port int, used []bool, scoreFn func(curr, other bridgeResult) bool) bridgeResult {
	var best bridgeResult

	for i, comp := range b.components {
		if used[i] {
			continue
		}

		if canConnect, nextPort := comp.canConnect(port); canConnect {
			used[i] = true
			result := b.findBestBridge(nextPort, used, scoreFn)
			used[i] = false

			current := bridgeResult{
				length:   result.length + 1,
				strength: result.strength + comp.strength(),
			}

			if scoreFn(current, best) {
				best = current
			}
		}
	}

	return best
}
