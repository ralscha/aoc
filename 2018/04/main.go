package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"sort"
	"strings"
)

func main() {
	input, err := download.ReadInput(2018, 4)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1and2(input)
}

type guard struct {
	id      int
	asleep  int
	minutes *container.Bag[int]
}

type guardLog struct {
	date   string
	action string
}

func part1and2(input string) {
	lines := conv.SplitNewline(input)
	guardLogs := make([]guardLog, len(lines))

	for i, line := range lines {
		secondBracket := strings.Index(line, "]") + 1
		date := line[0:secondBracket]
		action := line[secondBracket+1:]
		guardLogs[i] = guardLog{date, action}
	}

	sort.Slice(guardLogs, func(i, j int) bool {
		return guardLogs[i].date < guardLogs[j].date
	})

	guards := make(map[int]*guard)
	var currentGuard *guard
	var asleep int
	for _, gl := range guardLogs {
		switch gl.action {
		case "falls asleep":
			asleep = conv.MustAtoi(gl.date[15:17])
		case "wakes up":
			if currentGuard == nil {
				log.Fatalf("no guard found")
			}
			minute := conv.MustAtoi(gl.date[15:17])
			currentGuard.asleep += minute - asleep + 1
			for i := asleep; i < minute; i++ {
				currentGuard.minutes.Add(i)
			}
		default:
			id := conv.MustAtoi(gl.action[7:strings.Index(gl.action, " begins")])
			if v, ok := guards[id]; ok {
				currentGuard = v
			} else {
				currentGuard = &guard{
					id:      id,
					minutes: container.NewBag[int](),
				}
				guards[id] = currentGuard
			}
		}
	}

	var maxGuard *guard
	for _, g := range guards {
		if maxGuard == nil || g.asleep > maxGuard.asleep {
			maxGuard = g
		}
	}

	if maxGuard == nil {
		log.Fatalf("no max guard found")
	}

	var maxMinuteCount int
	var maxMinuteIndex int

	for minute := 0; minute < 60; minute++ {
		count := maxGuard.minutes.Count(minute)
		if count > maxMinuteCount {
			maxMinuteCount = count
			maxMinuteIndex = minute
		}
	}

	fmt.Println("Part 1", maxGuard.id*maxMinuteIndex)

	maxGuard = nil
	maxMinuteCount = 0
	maxMinuteIndex = 0

	for _, g := range guards {
		for minute := 0; minute < 60; minute++ {
			count := g.minutes.Count(minute)
			if count > maxMinuteCount {
				maxMinuteCount = count
				maxMinuteIndex = minute
				maxGuard = g
			}
		}
	}
	if maxGuard == nil {
		log.Fatalf("no max guard found")
	}
	fmt.Println("Part 2", maxGuard.id*maxMinuteIndex)
}
