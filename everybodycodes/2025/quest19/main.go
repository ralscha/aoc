package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strings"

	"aoc/internal/container"
	"aoc/internal/conv"
)

func main() {
	partI()
	partII()
	partIII()
}

type state struct {
	x      int
	height int
}

type opening struct {
	start int
	end   int
}

func solve(filename string) int {
	data, _ := os.ReadFile(filename)
	content := strings.ReplaceAll(string(data), "\r\n", "\n")
	lines := conv.SplitNewline(strings.TrimSpace(content))

	wallMap := make(map[int][]opening)
	maxX := 0
	for _, line := range lines {
		parts := conv.ToIntSliceComma(line)
		x, startHeight, size := parts[0], parts[1], parts[2]
		wallMap[x] = append(wallMap[x], opening{start: startHeight + 1, end: startHeight + size})
		if x > maxX {
			maxX = x
		}
	}

	pq := container.NewPriorityQueue[state]()
	pq.Push(state{x: 0, height: 0}, 0)

	visited := make(map[state]int)
	visited[state{x: 0, height: 0}] = 0

	for !pq.IsEmpty() {
		curr := pq.Pop()

		if curr.x == maxX {
			return visited[curr]
		}

		currCost := visited[curr]

		for _, dh := range []int{1, -1} {
			newX := curr.x + 1
			newHeight := curr.height + dh
			newCost := currCost
			if dh == 1 {
				newCost++
			}

			if newHeight < 0 {
				continue
			}

			if openings, exists := wallMap[newX]; exists {
				valid := false
				for _, o := range openings {
					if newHeight >= o.start && newHeight <= o.end {
						valid = true
						break
					}
				}
				if !valid {
					continue
				}
			}

			newState := state{x: newX, height: newHeight}
			if prevCost, exists := visited[newState]; !exists || newCost < prevCost {
				visited[newState] = newCost
				pq.Push(newState, newCost)
			}
		}
	}

	return -1
}

func minFlaps(fromHeight, toHeight, distance int) int {
	heightDiff := toHeight - fromHeight
	if distance < 0 || (heightDiff > distance) || (heightDiff < -distance) {
		return -1
	}
	if (distance-heightDiff)%2 != 0 {
		return -1
	}
	flaps := (distance + heightDiff) / 2
	return flaps
}

func solveFastOptimized(filename string) int {
	data, _ := os.ReadFile(filename)
	content := strings.ReplaceAll(string(data), "\r\n", "\n")
	lines := conv.SplitNewline(strings.TrimSpace(content))

	wallMap := make(map[int][]opening)
	var wallXs []int
	xSet := make(map[int]bool)
	maxX := 0
	for _, line := range lines {
		parts := conv.ToIntSliceComma(line)
		x, startHeight, size := parts[0], parts[1], parts[2]
		wallMap[x] = append(wallMap[x], opening{start: startHeight + 1, end: startHeight + size})
		if !xSet[x] {
			xSet[x] = true
			wallXs = append(wallXs, x)
		}
		if x > maxX {
			maxX = x
		}
	}
	sort.Ints(wallXs)

	bestAtWall := make([]map[int]int, len(wallXs))
	for i := range bestAtWall {
		bestAtWall[i] = make(map[int]int)
	}

	firstWallX := wallXs[0]
	firstOpenings := wallMap[firstWallX]
	for _, o := range firstOpenings {
		for h := o.start; h <= o.end; h++ {
			flaps := minFlaps(0, h, firstWallX)
			if flaps >= 0 {
				if prev, exists := bestAtWall[0][h]; !exists || flaps < prev {
					bestAtWall[0][h] = flaps
				}
			}
		}
	}

	for wallIdx := 0; wallIdx < len(wallXs)-1; wallIdx++ {
		currWallX := wallXs[wallIdx]
		nextWallIdx := wallIdx + 1
		nextWallX := wallXs[nextWallIdx]
		distance := nextWallX - currWallX

		nextOpenings := wallMap[nextWallX]

		for currHeight, currCost := range bestAtWall[wallIdx] {
			for _, o := range nextOpenings {
				for h := o.start; h <= o.end; h++ {
					flaps := minFlaps(currHeight, h, distance)
					if flaps >= 0 {
						newCost := currCost + flaps
						if prev, exists := bestAtWall[nextWallIdx][h]; !exists || newCost < prev {
							bestAtWall[nextWallIdx][h] = newCost
						}
					}
				}
			}
		}
	}

	minCost := math.MaxInt
	for _, cost := range bestAtWall[len(wallXs)-1] {
		if cost < minCost {
			minCost = cost
		}
	}

	return minCost
}

func partI() {
	fmt.Println(solve("input1"))
}

func partII() {
	fmt.Println(solve("input2"))
}

func partIII() {
	fmt.Println(solveFastOptimized("input3"))
}
