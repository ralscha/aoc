package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
)

func main() {
	input, err := download.ReadInput(2024, 9)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	line := lines[0]

	blocks := parseDiskMap(line)
	compactBlocks(blocks)
	checksum := calculateChecksum(blocks)
	fmt.Println("Part 1", checksum)
}

func parseDiskMap(diskMap string) []int {
	var blocks []int
	fileID := 0

	for i := 0; i < len(diskMap); i++ {
		id := fileID
		if i%2 == 1 {
			id = -1
		} else {
			fileID++
		}
		size := conv.MustAtoi(string(diskMap[i]))
		for j := 0; j < size; j++ {
			blocks = append(blocks, id)
		}
	}

	return blocks
}

func compactBlocks(blocks []int) {
	back := len(blocks) - 1
	front := 0

	for front < back {
		if blocks[front] == -1 {
			if blocks[back] != -1 {
				blocks[front] = blocks[back]
				blocks[back] = -1
				front++
			}
			back--
		} else {
			front++
		}
	}
}

func calculateChecksum(blocks []int) int {
	checksum := 0
	for i, block := range blocks {
		if block != -1 {
			checksum += i * block
		}
	}
	return checksum
}

type Block struct {
	index int
	size  int
}

func compactBlocks2(blocks []int) {
	largestId := 0
	blockSizes := make(map[int]Block)
	for i, block := range blocks {
		if block != -1 {
			b, exists := blockSizes[block]
			if exists {
				b.size++
				blockSizes[block] = b
			} else {
				blockSizes[block] = Block{index: i, size: 1}
			}
		}
		if block > largestId {
			largestId = block
		}
	}

	for id := largestId; id >= 0; id-- {
		currentIndex := 0
		for currentIndex < len(blocks) {
			if blockSizes[id].index < currentIndex {
				break
			}
			if blocks[currentIndex] == -1 {
				freeSize := 1
				for i := currentIndex + 1; i < len(blocks); i++ {
					if blocks[i] == -1 {
						freeSize++
					} else {
						break
					}
				}
				size := blockSizes[id].size
				if freeSize >= size {
					for j := 0; j < size; j++ {
						blocks[currentIndex+j] = id
						blocks[blockSizes[id].index+j] = -1
					}
					currentIndex += size
					break
				} else {
					currentIndex += freeSize
				}
			} else {
				currentIndex++
			}
		}
	}
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	line := lines[0]

	blocks := parseDiskMap(line)
	compactBlocks2(blocks)
	checksum := calculateChecksum(blocks)
	fmt.Println("Part 2", checksum)
}
