package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"slices"
	"strings"
)

// Port of
// https://github.com/tymscar/Advent-Of-Code/tree/master/2022/typescript/day16

func main() {
	input, err := download.ReadInput(2022, 16)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

type valve struct {
	name     string
	flowRate int
	tunnelTo []string
}

type path struct {
	currentValve  string
	toVisit       []string
	timeLeft      int
	finished      bool
	steps         []string
	finalPressure int
}

func part1(input string) {
	valves := createValves(input)

	var startingValves []string
	startingValves = append(startingValves, "AA")

	var destinationValves []string
	for _, v := range valves {
		if v.flowRate != 0 {
			startingValves = append(startingValves, v.name)
			destinationValves = append(destinationValves, v.name)
		}
	}

	valvesPrices := make(map[string]map[string]int)
	for _, startingValve := range startingValves {
		valvesPrices[startingValve] = dijkstra(valves, startingValve, destinationValves)
	}

	paths := []*path{{
		currentValve:  "AA",
		toVisit:       destinationValves,
		timeLeft:      30,
		finished:      false,
		finalPressure: 0,
	}}

	for n := 0; n < len(paths); n++ {
		p := paths[n]
		if p.timeLeft <= 0 || p.finished {
			p.finished = true
			continue
		}

		currPrices := valvesPrices[p.currentValve]
		madeNewPath := false
		for _, valveName := range p.toVisit {
			if p.timeLeft-currPrices[valveName] > 1 {
				madeNewPath = true
				var newToVisit []string
				for _, v := range p.toVisit {
					if v != valveName {
						newToVisit = append(newToVisit, v)
					}
				}
				paths = append(paths, &path{
					currentValve:  valveName,
					toVisit:       newToVisit,
					timeLeft:      p.timeLeft - currPrices[valveName] - 1,
					finished:      false,
					finalPressure: p.finalPressure + (p.timeLeft-currPrices[valveName]-1)*valves[valveName].flowRate,
				})
			}
		}
		if !madeNewPath {
			p.finished = true
		}
	}

	maxPressure := 0
	for _, p := range paths {
		if p.finished {
			if p.finalPressure > maxPressure {
				maxPressure = p.finalPressure
			}
		}
	}
	fmt.Println(maxPressure)
}

func part2(input string) {
	valves := createValves(input)

	var startingValves []string
	startingValves = append(startingValves, "AA")

	var destinationValves []string
	for _, v := range valves {
		if v.flowRate != 0 {
			startingValves = append(startingValves, v.name)
			destinationValves = append(destinationValves, v.name)
		}
	}

	valvesPrices := make(map[string]map[string]int)
	for _, startingValve := range startingValves {
		valvesPrices[startingValve] = dijkstra(valves, startingValve, destinationValves)
	}

	paths := []*path{{
		currentValve:  "AA",
		toVisit:       destinationValves,
		timeLeft:      26,
		finished:      false,
		steps:         []string{},
		finalPressure: 0,
	}}

	for n := 0; n < len(paths); n++ {
		p := paths[n]
		if p.timeLeft <= 0 || p.finished {
			p.finished = true
			continue
		}

		currPrices := valvesPrices[p.currentValve]
		madeNewPath := false
		for _, valveName := range p.toVisit {
			if p.timeLeft-currPrices[valveName] > 1 {
				madeNewPath = true
				var newToVisit []string
				for _, v := range p.toVisit {
					if v != valveName {
						newToVisit = append(newToVisit, v)
					}
				}
				var newSteps []string
				for _, v := range p.steps {
					newSteps = append(newSteps, v)
				}
				newSteps = append(newSteps, valveName)
				paths = append(paths, &path{
					currentValve:  valveName,
					toVisit:       newToVisit,
					timeLeft:      p.timeLeft - currPrices[valveName] - 1,
					finished:      false,
					steps:         newSteps,
					finalPressure: p.finalPressure + (p.timeLeft-currPrices[valveName]-1)*valves[valveName].flowRate,
				})
				paths = append(paths, &path{
					currentValve:  valveName,
					toVisit:       []string{},
					timeLeft:      p.timeLeft - currPrices[valveName] - 1,
					finished:      true,
					steps:         newSteps,
					finalPressure: p.finalPressure + (p.timeLeft-currPrices[valveName]-1)*valves[valveName].flowRate,
				})
			}
		}
		if !madeNewPath {
			p.finished = true
		}
	}

	var finishedPaths []*path
	for _, p := range paths {
		if p.finished {
			finishedPaths = append(finishedPaths, p)
		}
	}
	paths = finishedPaths

	mostPressureReleased := -1
	for humanPath := 0; humanPath < len(paths); humanPath++ {
		if humanPath%10000 == 0 {
			fmt.Println(humanPath, "/", len(paths))
		}
		for elephantPath := humanPath + 1; elephantPath < len(paths); elephantPath++ {
			if !overlap(paths[humanPath].steps, paths[elephantPath].steps) {
				pressure := paths[humanPath].finalPressure + paths[elephantPath].finalPressure
				if pressure > mostPressureReleased {
					mostPressureReleased = pressure
				}
			}
		}
	}
	fmt.Println(mostPressureReleased)
}

func overlap(a, b []string) bool {
	for _, v := range a {
		if slices.Contains(b, v) {
			return true
		}
	}
	return false
}

func createValves(input string) map[string]*valve {
	lines := conv.SplitNewline(input)
	valves := make(map[string]*valve)
	for _, line := range lines {
		splitted := strings.Fields(line)
		valveID := splitted[1]
		flowRateStr := strings.TrimPrefix(splitted[4], "rate=")
		flowRate := conv.MustAtoi(flowRateStr[:len(flowRateStr)-1])
		leadToValves := splitted[9:]
		var tunnelTo []string
		for _, v := range leadToValves {
			tunnelTo = append(tunnelTo, strings.TrimSuffix(v, ","))
		}

		valves[valveID] = &valve{
			name:     valveID,
			flowRate: flowRate,
			tunnelTo: tunnelTo,
		}
	}
	return valves
}

func dijkstra(valves map[string]*valve, startValve string, endValves []string) map[string]int {
	visited := container.NewSet[string]()
	toVisit := []string{startValve}
	lowestCost := make(map[string]int)
	lowestCost[startValve] = 0

	var curr string
	for len(toVisit) > 0 {
		curr, toVisit = toVisit[0], toVisit[1:]
		if visited.Contains(curr) {
			continue
		}

		worthItAdj := make([]string, 0)
		for _, neighbour := range valves[curr].tunnelTo {
			if !visited.Contains(neighbour) {
				worthItAdj = append(worthItAdj, neighbour)
			}
		}

		toVisit = append(toVisit, worthItAdj...)

		costToCurr := lowestCost[curr]

		for _, neighbour := range worthItAdj {
			newCostToNeighbour := costToCurr + 1
			costToNeighbour := lowestCost[neighbour]
			if costToNeighbour == 0 {
				costToNeighbour = newCostToNeighbour
			}

			if newCostToNeighbour <= costToNeighbour {
				lowestCost[neighbour] = newCostToNeighbour
			}
		}

		visited.Add(curr)
	}

	result := make(map[string]int)
	for _, room := range endValves {
		result[room] = lowestCost[room]
	}
	return result
}
