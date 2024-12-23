package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/gridutil"
	"fmt"
	"log"
	"regexp"
)

func main() {
	input, err := download.ReadInput(2017, 20)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}
	part1(input)
	part2(input)
}

func part1(input string) {
	particles := parseParticles(input)
	fmt.Println("Part 1", findClosestParticle(particles))
}

func part2(input string) {
	particles := parseParticles(input)
	fmt.Println("Part 2", simulateCollisions(particles))
}

type particle struct {
	id           int
	position     gridutil.Coordinate3D
	velocity     gridutil.Coordinate3D
	acceleration gridutil.Coordinate3D
}

func (p *particle) update() {
	p.velocity = p.velocity.Add(p.acceleration)
	p.position = p.position.Add(p.velocity)
}

func (p *particle) distanceFromOrigin() (int, int, int) {
	origin := gridutil.Coordinate3D{}
	return p.acceleration.ManhattanDistance(origin),
		p.velocity.ManhattanDistance(origin),
		p.position.ManhattanDistance(origin)
}

var particleRegex = regexp.MustCompile(`p=<(-?\d+),(-?\d+),(-?\d+)>, v=<(-?\d+),(-?\d+),(-?\d+)>, a=<(-?\d+),(-?\d+),(-?\d+)>`)

func parseParticles(input string) []particle {
	lines := conv.SplitNewline(input)
	particles := make([]particle, len(lines))

	for i, line := range lines {
		matches := particleRegex.FindStringSubmatch(line)
		particles[i] = particle{
			id: i,
			position: gridutil.Coordinate3D{
				X: conv.MustAtoi(matches[1]),
				Y: conv.MustAtoi(matches[2]),
				Z: conv.MustAtoi(matches[3]),
			},
			velocity: gridutil.Coordinate3D{
				X: conv.MustAtoi(matches[4]),
				Y: conv.MustAtoi(matches[5]),
				Z: conv.MustAtoi(matches[6]),
			},
			acceleration: gridutil.Coordinate3D{
				X: conv.MustAtoi(matches[7]),
				Y: conv.MustAtoi(matches[8]),
				Z: conv.MustAtoi(matches[9]),
			},
		}
	}
	return particles
}

func findClosestParticle(particles []particle) int {
	if len(particles) == 0 {
		return -1
	}

	closestID := 0
	minAcc, minVel, minPos := particles[0].distanceFromOrigin()

	for i := 1; i < len(particles); i++ {
		acc, vel, pos := particles[i].distanceFromOrigin()
		if acc < minAcc ||
			(acc == minAcc && vel < minVel) ||
			(acc == minAcc && vel == minVel && pos < minPos) {
			closestID = i
			minAcc, minVel, minPos = acc, vel, pos
		}
	}

	return closestID
}

func simulateCollisions(initialParticles []particle) int {

	particles := make(map[int]particle, len(initialParticles))
	for _, p := range initialParticles {
		particles[p.id] = p
	}

	const iterations = 1000
	for range iterations {
		positions := make(map[gridutil.Coordinate3D]*container.Set[int])

		for id, p := range particles {
			p.update()
			particles[id] = p

			if positions[p.position] == nil {
				positions[p.position] = container.NewSet[int]()
			}
			positions[p.position].Add(id)
		}

		for _, ids := range positions {
			if ids.Len() > 1 {
				for _, id := range ids.Values() {
					delete(particles, id)
				}
			}
		}
	}

	return len(particles)
}
