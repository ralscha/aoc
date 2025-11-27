package main

import (
	"fmt"
	"maps"
	"os"
	"regexp"
	"strings"

	"aoc/internal/container"
	"aoc/internal/conv"
)

func main() {
	partI()
	partII()
	partIII()
}

type Plant struct {
	id        int
	thickness int
	branches  []Branch
	isFree    bool
}

type Branch struct {
	toPlantID int
	thickness int
}

func partI() {
	content, _ := os.ReadFile("input1")
	text := strings.ReplaceAll(string(content), "\r\n", "\n")
	lines := conv.SplitNewline(strings.TrimSpace(text))

	plants, _ := parseInput(lines)

	referenced := container.NewSet[int]()
	for _, p := range plants {
		for _, b := range p.branches {
			referenced.Add(b.toPlantID)
		}
	}

	var lastPlantID int
	for id, plant := range plants {
		if !referenced.Contains(id) && !plant.isFree {
			lastPlantID = id
			break
		}
	}

	energy := calculateEnergy(plants, lastPlantID, nil)

	fmt.Println(energy)
}

func parseInput(lines []string) (map[int]*Plant, [][]int) {
	plants := make(map[int]*Plant)
	plantRegex := regexp.MustCompile(`Plant (\d+) with thickness (\d+):`)
	freeBranchRegex := regexp.MustCompile(`- free branch with thickness (\d+)`)
	branchRegex := regexp.MustCompile(`- branch to Plant (\d+) with thickness (-?\d+)`)

	var currentPlant *Plant
	var testCases [][]int

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		if matches := plantRegex.FindStringSubmatch(line); matches != nil {
			id := conv.MustAtoi(matches[1])
			thickness := conv.MustAtoi(matches[2])
			currentPlant = &Plant{
				id:        id,
				thickness: thickness,
				branches:  []Branch{},
			}
			plants[id] = currentPlant
		} else if matches := freeBranchRegex.FindStringSubmatch(line); matches != nil {
			currentPlant.isFree = true
		} else if matches := branchRegex.FindStringSubmatch(line); matches != nil {
			toID := conv.MustAtoi(matches[1])
			thickness := conv.MustAtoi(matches[2])
			currentPlant.branches = append(currentPlant.branches, Branch{
				toPlantID: toID,
				thickness: thickness,
			})
		} else if strings.Contains(line, "0") || strings.Contains(line, "1") {
			parts := strings.Fields(line)
			testCase := make([]int, len(parts))
			for i, p := range parts {
				testCase[i] = conv.MustAtoi(p)
			}
			testCases = append(testCases, testCase)
		}
	}

	return plants, testCases
}

func calculateEnergy(plants map[int]*Plant, lastPlantID int, activation map[int]int) int {
	energy := make(map[int]int)

	var calcEnergy func(id int) int
	calcEnergy = func(id int) int {
		if val, ok := energy[id]; ok {
			return val
		}

		plant := plants[id]

		if plant.isFree {
			if activation != nil {
				energy[id] = activation[id]
			} else if plant.thickness <= 1 {
				energy[id] = 1
			} else {
				energy[id] = 0
			}
			return energy[id]
		}

		incomingEnergy := 0
		for _, branch := range plant.branches {
			sourceEnergy := calcEnergy(branch.toPlantID)
			incomingEnergy += branch.thickness * sourceEnergy
		}

		if incomingEnergy >= plant.thickness {
			energy[id] = incomingEnergy
		} else {
			energy[id] = 0
		}

		return energy[id]
	}

	return calcEnergy(lastPlantID)
}

func partII() {
	content, _ := os.ReadFile("input2")
	text := strings.ReplaceAll(string(content), "\r\n", "\n")
	lines := conv.SplitNewline(strings.TrimSpace(text))

	plants, testCases := parseInput(lines)

	var freePlants []int
	for id, plant := range plants {
		if plant.isFree {
			freePlants = append(freePlants, id)
		}
	}

	for i := range freePlants {
		for j := i + 1; j < len(freePlants); j++ {
			if freePlants[i] > freePlants[j] {
				freePlants[i], freePlants[j] = freePlants[j], freePlants[i]
			}
		}
	}

	referenced := container.NewSet[int]()
	for _, p := range plants {
		for _, b := range p.branches {
			referenced.Add(b.toPlantID)
		}
	}

	var lastPlantID int
	for id, plant := range plants {
		if !referenced.Contains(id) && !plant.isFree {
			lastPlantID = id
			break
		}
	}

	totalEnergy := 0
	for _, testCase := range testCases {
		activation := make(map[int]int)
		for i, activated := range testCase {
			if i < len(freePlants) {
				activation[freePlants[i]] = activated
			}
		}

		energy := calculateEnergy(plants, lastPlantID, activation)
		totalEnergy += energy
	}

	fmt.Println(totalEnergy)
}

func partIII() {
	content, _ := os.ReadFile("input3")
	text := strings.ReplaceAll(string(content), "\r\n", "\n")
	lines := conv.SplitNewline(strings.TrimSpace(text))

	plants, testCases := parseInput(lines)

	var freePlants []int
	for id, plant := range plants {
		if plant.isFree {
			freePlants = append(freePlants, id)
		}
	}
	for i := range freePlants {
		for j := i + 1; j < len(freePlants); j++ {
			if freePlants[i] > freePlants[j] {
				freePlants[i], freePlants[j] = freePlants[j], freePlants[i]
			}
		}
	}

	referenced := container.NewSet[int]()
	for _, p := range plants {
		for _, b := range p.branches {
			referenced.Add(b.toPlantID)
		}
	}

	var lastPlantID int
	for id, plant := range plants {
		if !referenced.Contains(id) && !plant.isFree {
			lastPlantID = id
			break
		}
	}

	maxEnergy := findMaxEnergy(plants, lastPlantID, freePlants)

	maxFromTests := 0
	for _, testCase := range testCases {
		activation := make(map[int]int)
		for i, activated := range testCase {
			if i < len(freePlants) {
				activation[freePlants[i]] = activated
			}
		}
		energy := calculateEnergy(plants, lastPlantID, activation)
		if energy > maxFromTests {
			maxFromTests = energy
		}
	}

	if maxFromTests > maxEnergy {
		maxEnergy = maxFromTests
	}

	totalDiff := 0
	activatedCount := 0
	for _, testCase := range testCases {
		activation := make(map[int]int)
		for i, activated := range testCase {
			if i < len(freePlants) {
				activation[freePlants[i]] = activated
			}
		}

		energy := calculateEnergy(plants, lastPlantID, activation)
		if energy > 0 {
			totalDiff += maxEnergy - energy
			activatedCount++
		}
	}

	fmt.Println(totalDiff)
}

func findMaxEnergy(plants map[int]*Plant, lastPlantID int, freePlants []int) int {
	lastPlant := plants[lastPlantID]
	dependencies := make(map[int]*container.Set[int])

	var findDeps func(id int) *container.Set[int]
	findDeps = func(id int) *container.Set[int] {
		if deps, ok := dependencies[id]; ok {
			return deps
		}
		plant := plants[id]
		deps := container.NewSet[int]()
		if plant.isFree {
			deps.Add(id)
		} else {
			for _, b := range plant.branches {
				for _, freeID := range findDeps(b.toPlantID).Values() {
					deps.Add(freeID)
				}
			}
		}
		dependencies[id] = deps
		return deps
	}

	branchDeps := make([]*container.Set[int], len(lastPlant.branches))
	for i, b := range lastPlant.branches {
		branchDeps[i] = findDeps(b.toPlantID)
	}

	hasOverlap := false
	for i := range branchDeps {
		for j := i + 1; j < len(branchDeps); j++ {
			if branchDeps[i].Intersection(branchDeps[j]).Len() > 0 {
				hasOverlap = true
				break
			}
		}
		if hasOverlap {
			break
		}
	}

	if !hasOverlap {
		totalEnergy := 0
		for _, b := range lastPlant.branches {
			maxBranchEnergy := findMaxForSubtree(plants, b.toPlantID)
			totalEnergy += b.thickness * maxBranchEnergy
		}

		if totalEnergy >= lastPlant.thickness {
			return totalEnergy
		}
		return 0
	}

	return findMaxBySearch(plants, lastPlantID, freePlants)
}

func findMaxForSubtree(plants map[int]*Plant, rootID int) int {
	var freePlants []int
	var findFree func(id int)
	findFree = func(id int) {
		plant := plants[id]
		if plant.isFree {
			freePlants = append(freePlants, id)
		} else {
			for _, b := range plant.branches {
				findFree(b.toPlantID)
			}
		}
	}
	findFree(rootID)

	maxEnergy := 0
	numFree := len(freePlants)
	for mask := 0; mask < (1 << numFree); mask++ {
		activation := make(map[int]int)
		for i, freeID := range freePlants {
			if mask&(1<<i) != 0 {
				activation[freeID] = 1
			} else {
				activation[freeID] = 0
			}
		}
		energy := calculateEnergyForSubtree(plants, rootID, activation)
		if energy > maxEnergy {
			maxEnergy = energy
		}
	}
	return maxEnergy
}

func calculateEnergyForSubtree(plants map[int]*Plant, rootID int, activation map[int]int) int {
	var calcEnergy func(id int) int
	energy := make(map[int]int)

	calcEnergy = func(id int) int {
		if val, ok := energy[id]; ok {
			return val
		}

		plant := plants[id]

		if plant.isFree {
			if val, ok := activation[id]; ok {
				energy[id] = val
			} else {
				energy[id] = 0
			}
			return energy[id]
		}

		incomingEnergy := 0
		for _, branch := range plant.branches {
			sourceEnergy := calcEnergy(branch.toPlantID)
			incomingEnergy += branch.thickness * sourceEnergy
		}

		if incomingEnergy >= plant.thickness {
			energy[id] = incomingEnergy
		} else {
			energy[id] = 0
		}

		return energy[id]
	}

	return calcEnergy(rootID)
}

func findMaxBySearch(plants map[int]*Plant, lastPlantID int, freePlants []int) int {

	rowOptimals := make(map[int][]int)

	for plantID := 82; plantID <= 90; plantID++ {
		if p, ok := plants[plantID]; ok {
			var rowFree []int
			for _, b := range p.branches {
				if plants[b.toPlantID].isFree {
					rowFree = append(rowFree, b.toPlantID)
				}
			}

			maxRowEnergy := 0
			bestActivation := make([]int, len(rowFree))
			for mask := 0; mask < (1 << len(rowFree)); mask++ {
				activation := make(map[int]int)
				for i, fid := range rowFree {
					if mask&(1<<i) != 0 {
						activation[fid] = 1
					} else {
						activation[fid] = 0
					}
				}
				energy := calculateEnergyForSubtree(plants, plantID, activation)
				if energy > maxRowEnergy {
					maxRowEnergy = energy
					bestActivation = make([]int, len(rowFree))
					for i := range rowFree {
						if mask&(1<<i) != 0 {
							bestActivation[i] = 1
						} else {
							bestActivation[i] = 0
						}
					}
				}
			}
			rowOptimals[plantID] = bestActivation
		}
	}

	globalActivation := make(map[int]int)
	for plantID, activation := range rowOptimals {
		p := plants[plantID]
		freeIdx := 0
		for _, b := range p.branches {
			if plants[b.toPlantID].isFree {
				globalActivation[b.toPlantID] = activation[freeIdx]
				freeIdx++
			}
		}
	}

	bestEnergy := localSearch(plants, lastPlantID, freePlants, globalActivation)

	for plantID := 91; plantID <= 99; plantID++ {
		if p, ok := plants[plantID]; ok {
			for _, b := range p.branches {
				if plants[b.toPlantID].isFree {
					testActivation := make(map[int]int)
					maps.Copy(testActivation, globalActivation)
					testActivation[b.toPlantID] = 1 - testActivation[b.toPlantID]
					energy := localSearch(plants, lastPlantID, freePlants, testActivation)
					if energy > bestEnergy {
						bestEnergy = energy
						globalActivation[b.toPlantID] = testActivation[b.toPlantID]
					}
				}
			}
		}
	}

	return bestEnergy
}

func localSearch(plants map[int]*Plant, lastPlantID int, freePlants []int, startActivation map[int]int) int {
	activation := make(map[int]int)
	maps.Copy(activation, startActivation)

	bestEnergy := calculateEnergy(plants, lastPlantID, activation)

	improved := true
	for improved {
		improved = false
		for _, fid := range freePlants {
			activation[fid] = 1 - activation[fid]
			energy := calculateEnergy(plants, lastPlantID, activation)
			if energy > bestEnergy {
				bestEnergy = energy
				improved = true
			} else {
				activation[fid] = 1 - activation[fid]
			}
		}
	}

	return bestEnergy
}
