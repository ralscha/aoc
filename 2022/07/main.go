package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"strings"
)

type file struct {
	name string
	size int
}

type directory struct {
	name        string
	files       []*file
	directories []*directory
	parent      *directory
}

func main() {
	inputFile := "./2022/07/input.txt"
	input, err := download.ReadInput(inputFile, 2022, 7)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1and2(input)
}

func part1and2(input string) {
	currentDir := &directory{name: "/", parent: nil}
	root := currentDir

	inputLines := strings.Split(input, "\n")
	ix := 0
	line := inputLines[ix]
	for line != "" {
		if strings.HasPrefix(line, "$ cd") {
			parts := strings.Fields(line)
			dirName := parts[len(parts)-1]
			if dirName == "/" {
				currentDir = root
			} else if dirName == ".." {
				if currentDir.parent != nil {
					currentDir = currentDir.parent
				}
			} else {
				found := false
				for _, dir := range currentDir.directories {
					if dir.name == dirName {
						currentDir = dir
						found = true
						break
					}
				}
				if !found {
					newDir := &directory{name: dirName, parent: currentDir}
					currentDir.directories = append(currentDir.directories, newDir)
					currentDir = newDir
				}
			}
			ix++
			line = inputLines[ix]
		} else if strings.HasPrefix(line, "$ ls") {
			ix++
			line = inputLines[ix]
			for len(line) > 0 && line[0] != '$' {
				parts := strings.Fields(line)
				if parts[0] != "dir" {
					fileSize := conv.MustAtoi(parts[0])
					fileName := parts[1]
					currentDir.files = append(currentDir.files, &file{name: fileName, size: fileSize})
				}
				ix++
				line = inputLines[ix]
			}
		}
	}

	// part 1
	directories := root.findDirectoriesWithSizeLessThan(100000)
	totalSize := 0
	for _, dir := range directories {
		totalSize += dir.totalSize()
	}
	fmt.Println(totalSize)

	// part 2
	totalSize = root.totalSize()
	toBeDeleted := 30000000 - (70000000 - totalSize)

	candidates := root.findDirectoriesWithSizeGreaterThan(toBeDeleted)
	lowest := candidates[0].totalSize()
	for i := 1; i < len(candidates); i++ {
		if candidates[i].totalSize() < lowest {
			lowest = candidates[i].totalSize()
		}
	}
	fmt.Println(lowest)
}

func (d *directory) totalSize() int {
	size := 0
	for _, f := range d.files {
		size += f.size
	}
	for _, subdir := range d.directories {
		size += subdir.totalSize()
	}
	return size
}

func (d *directory) findDirectoriesWithSizeLessThan(n int) []*directory {
	var result []*directory
	for _, subdir := range d.directories {
		if subdir.totalSize() <= n {
			result = append(result, subdir)
		}
		result = append(result, subdir.findDirectoriesWithSizeLessThan(n)...)
	}
	return result
}

func (d *directory) findDirectoriesWithSizeGreaterThan(n int) []*directory {
	var result []*directory
	for _, subdir := range d.directories {
		if subdir.totalSize() >= n {
			result = append(result, subdir)
		}
		result = append(result, subdir.findDirectoriesWithSizeGreaterThan(n)...)
	}
	return result
}
