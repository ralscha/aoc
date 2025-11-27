package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/graphutil"
	"fmt"
	"log"
	"slices"
	"strings"
)

func main() {
	input, err := download.ReadInput(2018, 7)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	graph := graphutil.NewGraph()
	lines := conv.SplitNewline(input)

	for _, line := range lines {
		words := strings.Fields(line)
		dependsOn, step := words[1], words[7]
		graph.AddNode(step)
		graph.AddNode(dependsOn)
		graph.AddEdge(dependsOn, step, 1)
	}

	readyNodes := make([]string, 0)
	for _, node := range graph.GetNeighbors("") {
		hasIncoming := false
		for _, otherNode := range graph.GetNeighbors("") {
			for _, neighbor := range graph.GetNeighbors(otherNode.ID) {
				if neighbor.ID == node.ID {
					hasIncoming = true
					break
				}
			}
			if hasIncoming {
				break
			}
		}
		if !hasIncoming {
			readyNodes = append(readyNodes, node.ID)
		}
	}

	slices.Sort(readyNodes)
	jobOrder := ""
	doneJobs := container.NewSet[string]()

	for len(readyNodes) > 0 {
		job := readyNodes[0]
		jobOrder += job
		doneJobs.Add(job)
		readyNodes = readyNodes[1:]

		for _, node := range graph.GetNeighbors("") {
			if doneJobs.Contains(node.ID) {
				continue
			}

			allDependenciesDone := true
			for _, otherNode := range graph.GetNeighbors("") {
				for _, neighbor := range graph.GetNeighbors(otherNode.ID) {
					if neighbor.ID == node.ID && !doneJobs.Contains(otherNode.ID) {
						allDependenciesDone = false
						break
					}
				}
				if !allDependenciesDone {
					break
				}
			}

			if allDependenciesDone && !slices.Contains(readyNodes, node.ID) {
				readyNodes = append(readyNodes, node.ID)
			}
		}
		slices.Sort(readyNodes)
	}

	fmt.Println("Part 1", jobOrder)
}

type worker struct {
	job  string
	time int
}

func part2(input string) {
	graph := graphutil.NewGraph()
	lines := conv.SplitNewline(input)

	for _, line := range lines {
		words := strings.Fields(line)
		dependsOn, step := words[1], words[7]
		graph.AddNode(step)
		graph.AddNode(dependsOn)
		graph.AddEdge(dependsOn, step, 1)
	}

	readyNodes := make([]string, 0)
	for _, node := range graph.GetNeighbors("") {
		hasIncoming := false
		for _, otherNode := range graph.GetNeighbors("") {
			for _, neighbor := range graph.GetNeighbors(otherNode.ID) {
				if neighbor.ID == node.ID {
					hasIncoming = true
					break
				}
			}
			if hasIncoming {
				break
			}
		}
		if !hasIncoming {
			readyNodes = append(readyNodes, node.ID)
		}
	}

	slices.Sort(readyNodes)
	doneJobs := container.NewSet[string]()

	noOfWorkers := 5
	workers := make([]worker, noOfWorkers)

	seconds := 0
	for {
		if doneJobs.Len() == len(graph.GetNeighbors("")) {
			break
		}

		completed := false
		for i, worker := range workers {
			if worker.time == 0 && worker.job != "" {
				doneJobs.Add(worker.job)
				workers[i].job = ""
				completed = true
			}
		}

		if completed {
			for _, node := range graph.GetNeighbors("") {
				if doneJobs.Contains(node.ID) {
					continue
				}

				allDependenciesDone := true
				for _, otherNode := range graph.GetNeighbors("") {
					for _, neighbor := range graph.GetNeighbors(otherNode.ID) {
						if neighbor.ID == node.ID && !doneJobs.Contains(otherNode.ID) {
							allDependenciesDone = false
							break
						}
					}
					if !allDependenciesDone {
						break
					}
				}

				workInProgress := false
				for _, worker := range workers {
					if worker.job == node.ID {
						workInProgress = true
						break
					}
				}

				if allDependenciesDone && !slices.Contains(readyNodes, node.ID) && !workInProgress {
					readyNodes = append(readyNodes, node.ID)
				}
			}
			slices.Sort(readyNodes)
		}

		for i, worker := range workers {
			if worker.job == "" && len(readyNodes) > 0 {
				workers[i].job = readyNodes[0]
				workers[i].time = 60 + int(readyNodes[0][0]) - 64
				readyNodes = readyNodes[1:]
			}
		}

		for i, worker := range workers {
			if worker.job != "" {
				workers[i].time--
			}
		}

		seconds++
	}

	fmt.Println("Part 2", seconds-1)
}
