package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
)

type robot struct {
	posX, posY int
	velX, velY int
}

func main() {
	input, err := download.ReadInput(2024, 14)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	robots := parseInput(input)
	finalPositions := simulateRobots(robots, 101, 103, 100)
	q1, q2, q3, q4 := countRobots(finalPositions, 101, 103)
	fmt.Println("Part 1", q1*q2*q3*q4)
}

func part2(input string) {
	robots := parseInput(input)

	for seconds := 0; ; seconds++ {
		currentPositions := simulateRobots(robots, 101, 103, seconds)
		if allUniquePositions(currentPositions) {
			fmt.Println("Part 2", seconds)
			return
		}
	}
}

func parseInput(input string) []robot {
	lines := conv.SplitNewline(input)
	robots := make([]robot, len(lines))

	for i, line := range lines {
		newRobot := robot{}
		conv.MustSscanf(line, "p=%d,%d v=%d,%d\n", &newRobot.posX, &newRobot.posY, &newRobot.velX, &newRobot.velY)
		robots[i] = newRobot
	}

	return robots
}

func simulateRobots(robots []robot, width, height, seconds int) []robot {
	result := make([]robot, len(robots))

	for i, r := range robots {
		finalX := (r.posX + r.velX*seconds) % width
		if finalX < 0 {
			finalX += width
		}

		finalY := (r.posY + r.velY*seconds) % height
		if finalY < 0 {
			finalY += height
		}

		result[i] = robot{
			posX: finalX,
			posY: finalY,
			velX: r.velX,
			velY: r.velY,
		}
	}

	return result
}

func countRobots(robots []robot, width, height int) (int, int, int, int) {
	midX := width / 2
	midY := height / 2

	var q1, q2, q3, q4 int

	for _, robot := range robots {
		if robot.posX == midX || robot.posY == midY {
			continue
		}

		if robot.posX < midX {
			if robot.posY < midY {
				q1++
			} else {
				q3++
			}
		} else {
			if robot.posY < midY {
				q2++
			} else {
				q4++
			}
		}
	}

	return q1, q2, q3, q4
}

type key struct {
	x, y int
}

func allUniquePositions(robots []robot) bool {
	positionSet := container.NewSet[key]()

	for _, r := range robots {
		k := key{r.posX, r.posY}
		if positionSet.Contains(k) {
			return false
		}
		positionSet.Add(k)
	}
	return positionSet.Len() == len(robots)
}
