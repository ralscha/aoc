package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"slices"
	"sort"
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
	jobDependencies := make(map[string][]string)

	lines := conv.SplitNewline(input)
	var jobs []string

	for _, line := range lines {
		words := strings.Fields(line)
		step := words[7]
		dependsOn := words[1]
		jobDependencies[step] = append(jobDependencies[step], dependsOn)
		if !slices.Contains(jobs, step) {
			jobs = append(jobs, step)
		}
		if !slices.Contains(jobs, dependsOn) {
			jobs = append(jobs, dependsOn)
		}
	}

	var readyJobs []string
	for _, job := range jobs {
		if _, ok := jobDependencies[job]; !ok {
			readyJobs = append(readyJobs, job)
		}
	}

	sort.Strings(readyJobs)

	jobOrder := ""
	doneJobs := container.NewSet[string]()

	for len(readyJobs) > 0 {
		job := readyJobs[0]
		jobOrder += job
		doneJobs.Add(job)
		readyJobs = readyJobs[1:]
		for _, job := range jobs {
			if doneJobs.Contains(job) {
				continue
			}
			dependsOn := jobDependencies[job]
			dependsOnJobsDone := true
			for _, dependsOnJob := range dependsOn {
				if !doneJobs.Contains(dependsOnJob) {
					dependsOnJobsDone = false
					break
				}
			}
			if !dependsOnJobsDone {
				continue
			}

			if !slices.Contains(readyJobs, job) {
				readyJobs = append(readyJobs, job)
			}
		}
		sort.Strings(readyJobs)

	}

	fmt.Println(jobOrder)
}

type worker struct {
	job  string
	time int
}

func part2(input string) {
	jobDependencies := make(map[string][]string)

	lines := conv.SplitNewline(input)
	var jobs []string

	for _, line := range lines {
		words := strings.Fields(line)
		step := words[7]
		dependsOn := words[1]
		jobDependencies[step] = append(jobDependencies[step], dependsOn)
		if !slices.Contains(jobs, step) {
			jobs = append(jobs, step)
		}
		if !slices.Contains(jobs, dependsOn) {
			jobs = append(jobs, dependsOn)
		}
	}

	var readyJobs []string
	for _, job := range jobs {
		if _, ok := jobDependencies[job]; !ok {
			readyJobs = append(readyJobs, job)
		}
	}

	sort.Strings(readyJobs)
	doneJobs := container.NewSet[string]()

	noOfWorkers := 5
	workers := make([]worker, noOfWorkers)

	seconds := 0
	for {
		if doneJobs.Len() == len(jobs) {
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
			for _, job := range jobs {
				if doneJobs.Contains(job) {
					continue
				}
				dependsOn := jobDependencies[job]
				dependsOnJobsDone := true
				for _, dependsOnJob := range dependsOn {
					if !doneJobs.Contains(dependsOnJob) {
						dependsOnJobsDone = false
						break
					}
				}
				if !dependsOnJobsDone {
					continue
				}

				workInProgress := false
				for _, worker := range workers {
					if worker.job == job {
						workInProgress = true
						break
					}
				}

				if !slices.Contains(readyJobs, job) && !workInProgress {
					readyJobs = append(readyJobs, job)
				}
			}
			sort.Strings(readyJobs)
		}

		for i, worker := range workers {
			if worker.job == "" {
				if len(readyJobs) > 0 {
					workers[i].job = readyJobs[0]
					workers[i].time = 60 + int(workers[i].job[0]) - 64
					readyJobs = readyJobs[1:]
				}
			}
		}

		for i, worker := range workers {
			if worker.job != "" {
				workers[i].time--
			}
		}

		seconds++
	}

	fmt.Println(seconds - 1)
}
