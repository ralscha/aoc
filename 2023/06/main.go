package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	input, err := download.ReadInput(2023, 6)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

type race struct {
	time           int
	recordDistance int
}

func part1(input string) {
	races := parseRaces(input)

	totalWays := 1
	for _, race := range races {
		ways := calcWays(race)
		totalWays *= ways
	}

	fmt.Println(totalWays)
}

func calcWays(race race) int {
	ways := 0
	for chargeTime := 1; chargeTime < race.time; chargeTime++ {
		distance := (race.time - chargeTime) * chargeTime
		if distance > race.recordDistance {
			ways++
		}
	}
	return ways
}

func part2(input string) {
	races := parseRaces(input)
	raceTimeStr := ""
	raceDistanceStr := ""

	for _, race := range races {
		raceTimeStr = raceTimeStr + strconv.Itoa(race.time)
		raceDistanceStr = raceDistanceStr + strconv.Itoa(race.recordDistance)
	}
	combinedRace := race{
		time:           conv.MustAtoi(raceTimeStr),
		recordDistance: conv.MustAtoi(raceDistanceStr),
	}

	ways := calcWays(combinedRace)
	fmt.Println(ways)
}

func parseRaces(input string) []race {
	lines := conv.SplitNewline(input)

	splitted0 := strings.Split(lines[0], ":")
	splitted1 := strings.Split(lines[1], ":")
	races0 := strings.Fields(splitted0[1])
	races1 := strings.Fields(splitted1[1])

	var races []race
	for i := range races0 {
		race := race{
			time:           conv.MustAtoi(races0[i]),
			recordDistance: conv.MustAtoi(races1[i]),
		}
		races = append(races, race)
	}
	return races
}
