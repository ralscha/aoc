package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"log"
	"strings"
)

func main() {
	inputFile := "./2015/14/input.txt"
	input, err := download.ReadInput(inputFile, 2015, 14)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

type reindeer struct {
	speed    int
	duration int
	rest     int
}

func part1(input string) {
	reindeers := makeReindeers(input)

	maxDistance := 0
	for _, r := range reindeers {
		distance := r.distance(2503)
		if distance > maxDistance {
			maxDistance = distance
		}
	}

	println(maxDistance)
}

func part2(input string) {
	reindeers := makeReindeers(input)

	points := make(map[string]int)
	for i := 0; i < 2503; i++ {
		maxDistance := 0
		for _, r := range reindeers {
			distance := r.distance(i + 1)
			if distance > maxDistance {
				maxDistance = distance
			}
		}

		for name, r := range reindeers {
			distance := r.distance(i + 1)
			if distance == maxDistance {
				points[name]++
			}
		}
	}

	maxPoints := 0
	for _, p := range points {
		if p > maxPoints {
			maxPoints = p
		}
	}

	println(maxPoints)
}

func makeReindeers(input string) map[string]*reindeer {
	lines := conv.SplitNewline(input)

	reindeers := make(map[string]*reindeer)
	for _, line := range lines {
		splitted := strings.Fields(line)
		name := splitted[0]
		speed := conv.MustAtoi(splitted[3])
		duration := conv.MustAtoi(splitted[6])
		rest := conv.MustAtoi(splitted[13])

		reindeers[name] = &reindeer{
			speed:    speed,
			duration: duration,
			rest:     rest,
		}
	}
	return reindeers
}

func (r *reindeer) distance(raceDurationInSeconds int) int {
	period := r.duration + r.rest
	periods := raceDurationInSeconds / period
	distance := periods * r.duration * r.speed
	remaining := raceDurationInSeconds % period
	if remaining > r.duration {
		remaining = r.duration
	}
	distance += remaining * r.speed
	return distance
}
