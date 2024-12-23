package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"strings"
)

func main() {
	input, err := download.ReadInput(2018, 8)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

type node struct {
	children []*node
	metadata []int
}

func (n *node) len() int {
	sum := 2
	for _, child := range n.children {
		sum += child.len()
	}
	sum += len(n.metadata)
	return sum
}

func (n *node) value() int {
	if len(n.children) == 0 {
		return n.sumMetadata()
	}

	sum := 0
	for _, metadata := range n.metadata {
		if metadata > 0 && metadata <= len(n.children) {
			sum += n.children[metadata-1].value()
		}
	}
	return sum
}

func parseNode(values []int) *node {
	childCount := values[0]
	metadataCount := values[1]

	n := &node{}
	values = values[2:]

	for range childCount {
		child := parseNode(values)
		n.children = append(n.children, child)
		values = values[child.len():]
	}

	for i := range metadataCount {
		metadata := values[i]
		n.metadata = append(n.metadata, metadata)
	}

	return n
}

func (n *node) sumMetadata() int {
	sum := 0
	for _, metadata := range n.metadata {
		sum += metadata
	}

	for _, child := range n.children {
		sum += child.sumMetadata()
	}

	return sum
}

func part1(input string) {
	license := conv.ToIntSlice(strings.Fields(input))
	root := parseNode(license)
	fmt.Println(root.sumMetadata())
}

func part2(input string) {
	license := conv.ToIntSlice(strings.Fields(input))
	root := parseNode(license)
	fmt.Println(root.value())
}
