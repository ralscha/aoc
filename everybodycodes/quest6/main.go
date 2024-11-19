package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type node struct {
	name     string
	children []string
}

func findPowerfulFruit(graph map[string]node) []string {
	var paths [][]string
	visited := make(map[string]bool)
	var currentPath []string

	var dfs func(node string)
	dfs = func(node string) {
		currentPath = append(currentPath, node)
		visited[node] = true

		if node == "@" {
			pathCopy := make([]string, len(currentPath))
			copy(pathCopy, currentPath)
			paths = append(paths, pathCopy)
		} else {
			for _, child := range graph[node].children {
				if !visited[child] {
					dfs(child)
				}
			}
		}

		currentPath = currentPath[:len(currentPath)-1]
		visited[node] = false
	}

	dfs("RR")

	pathsByLength := make(map[int][][]string)
	for _, path := range paths {
		length := len(path)
		pathsByLength[length] = append(pathsByLength[length], path)
	}

	var uniquePath []string
	for _, paths := range pathsByLength {
		if len(paths) == 1 {
			uniquePath = paths[0]
			break
		}
	}

	return uniquePath
}

func partI() {
	graph := readGraph("everybodycodes/quest6/partI.txt")
	result := findPowerfulFruit(graph)
	resultStr := strings.Join(result, "")
	fmt.Println("Part I", resultStr)
}

func partII() {
	graph := readGraph("everybodycodes/quest6/partII.txt")
	result := findPowerfulFruit(graph)
	resultStr := ""
	for _, node := range result {
		resultStr += string(node[0])
	}
	fmt.Println("Part II", resultStr)
}

func partIII() {
	graph := readGraph("everybodycodes/quest6/partIII.txt")
	result := findPowerfulFruit(graph)
	resultStr := ""
	for _, node := range result {
		resultStr += string(node[0])
	}
	fmt.Println("Part III", resultStr)
}

func readGraph(fileName string) map[string]node {
	input, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(input), "\n")
	for i, l := range lines {
		lines[i] = strings.TrimRight(l, "\r\n")
	}
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	graph := make(map[string]node)
	for _, line := range lines {
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			continue
		}

		nodeName := strings.TrimSpace(parts[0])
		children := strings.Split(strings.TrimSpace(parts[1]), ",")

		for i := range children {
			children[i] = strings.TrimSpace(children[i])
		}

		graph[nodeName] = node{
			name:     nodeName,
			children: children,
		}
	}
	return graph
}

func main() {
	partI()
	partII()
	partIII()
}
