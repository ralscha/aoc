package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"strings"
)

func main() {
	input, err := download.ReadInput(2017, 7)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}
	part1and2(input)
}

type node struct {
	name        string
	weight      int
	parent      *node
	children    []*node
	totalWeight int
}

func part1and2(input string) {
	lines := conv.SplitNewline(input)
	nodes := make(map[string]*node)
	for _, line := range lines {
		parts := strings.Split(line, " ")
		name := parts[0]
		weight := conv.MustAtoi(parts[1][1 : len(parts[1])-1])
		var children []string
		if len(parts) > 2 {
			for _, child := range parts[3:] {
				child = strings.TrimSuffix(child, ",")
				children = append(children, child)
			}
		}
		if _, ok := nodes[name]; !ok {
			nodes[name] = &node{name: name, weight: weight}
		} else {
			nodes[name].weight = weight
		}
		for _, child := range children {
			if childNode, ok := nodes[child]; ok {
				childNode.parent = nodes[name]
				nodes[name].children = append(nodes[name].children, childNode)
			} else {
				nodes[child] = &node{name: child, parent: nodes[name]}
				nodes[name].children = append(nodes[name].children, nodes[child])
			}
		}
	}

	var rootNode *node
	for _, node := range nodes {
		if node.parent == nil {
			fmt.Println(node.name)
			rootNode = node
		}
	}

	rootNode.totalWeight = calculateTotalWeight(rootNode)
	checkBalance(rootNode)
}

func checkBalance(node *node) bool {
	for _, child := range node.children {
		if child.totalWeight != node.children[0].totalWeight {
			if checkBalance(child) {
				fmt.Println(child.weight - (child.totalWeight - node.children[0].totalWeight))
			}
			return false
		}
	}
	return true
}
func calculateTotalWeight(node *node) int {
	if len(node.children) == 0 {
		return node.weight
	}
	totalWeight := node.weight
	for _, child := range node.children {
		totalWeight += calculateTotalWeight(child)
	}
	node.totalWeight = totalWeight
	return totalWeight
}
