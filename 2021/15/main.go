package main

import (
	"aoc/internal/download"
	"container/heap"
	"container/list"
	"fmt"
	"log"
	"strings"
)

type point struct {
	x, y int
}

type qi struct {
	pos       point
	riskLevel int
	index     int
}

var dx = [4]int{0, 0, -1, 1}
var dy = [4]int{-1, 1, 0, 0}

func main() {
	input, err := download.ReadInput(2021, 15)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	lines := strings.Split(input, "\n")
	grid := make(map[point]int)
	for x, row := range lines {
		if row == "" {
			continue
		}
		for y, val := range row {
			grid[point{x: x, y: y}] = int(val - '0')
		}
	}

	var maxx, maxy = len(lines[0]) - 1, len(lines) - 2
	start := point{0, 0}
	target := point{maxx, maxy}

	queue := list.New()

	shortestAt := make(map[point]int)
	queue.PushBack(qi{pos: start, riskLevel: 0})
	for queue.Len() > 0 {
		front := queue.Front()
		head := front.Value.(qi)
		queue.Remove(front)

		for i := 0; i < 4; i++ {
			next := point{head.pos.x + dx[i], head.pos.y + dy[i]}
			if next.x > target.x || next.x < 0 || next.y > target.y || next.y < 0 {
				continue
			}
			nextRisk := head.riskLevel + grid[next]
			if sAt, ok := shortestAt[next]; ok && sAt <= nextRisk {
				continue
			}
			shortestAt[next] = nextRisk
			queue.PushBack(qi{pos: next, riskLevel: nextRisk})
		}
	}
	fmt.Println("Result: ", shortestAt[target])
}

func part2(input string) {
	lines := strings.Split(input, "\n")
	grid := make(map[point]int)

	for x, row := range lines {
		for y, val := range row {
			grid[point{x: x, y: y}] = int(val - '0')
		}
	}

	var maxx, maxy = len(lines[0]) - 1, len(lines) - 2
	start := point{0, 0}
	target := point{(maxx+1)*5 - 1, (maxy+1)*5 - 1}

	risk := func(pos point) int {
		og := point{pos.x % (maxx + 1), pos.y % (maxy + 1)}
		risk := grid[og] +
			(pos.x)/(maxx+1) + (pos.y)/(maxy+1)
		if risk > 9 {
			return risk - 9
		}
		return risk
	}
	shortestAt := make(map[point]int)
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, qi{pos: start, riskLevel: 0})
	for pq.Len() > 0 {
		head := heap.Pop(&pq).(qi)
		for i := 0; i < 4; i++ {
			next := point{head.pos.x + dx[i], head.pos.y + dy[i]}
			if next.x > target.x || next.x < 0 || next.y > target.y || next.y < 0 {
				continue
			}
			nextRisk := head.riskLevel + risk(next)
			if sAt, ok := shortestAt[next]; ok && sAt <= nextRisk {
				continue
			}
			shortestAt[next] = nextRisk
			heap.Push(&pq, qi{pos: next, riskLevel: nextRisk})
		}
	}
	fmt.Println("Result: ", shortestAt[target])
}

type PriorityQueue []qi

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].riskLevel < pq[j].riskLevel
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(qi)
	item.index = n
	*pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}
