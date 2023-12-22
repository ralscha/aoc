package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/mathx"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input, err := download.ReadInput(2021, 19)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1and2(input)
}

func part1and2(input string) {
	var ll [][][3]int
	var currentScanner [][3]int
	lines := conv.SplitNewline(input)
	for _, line := range lines {
		if line == "" {
			ll = append(ll, currentScanner)
			currentScanner = nil
			continue
		}
		if strings.HasPrefix(line, "---") {
			continue
		}
		coords := strings.Split(line, ",")
		var point [3]int
		for i, coord := range coords {
			val, err := strconv.Atoi(coord)
			if err != nil {
				fmt.Println("Error parsing integer:", err)
				return
			}
			point[i] = val
		}
		currentScanner = append(currentScanner, point)
	}
	ll = append(ll, currentScanner)

	coordRemaps := [][3]int{
		{0, 1, 2}, {0, 2, 1}, {1, 0, 2}, {1, 2, 0}, {2, 0, 1}, {2, 1, 0},
	}
	coordNegations := [][3]int{
		{1, 1, 1}, {1, 1, -1}, {1, -1, 1}, {1, -1, -1}, {-1, 1, 1}, {-1, 1, -1}, {-1, -1, 1}, {-1, -1, -1},
	}

	apply := func(remap, negat [3]int, scan [][3]int) [][3]int {
		ret := make([][3]int, len(scan))
		for i, item := range scan {
			ret[i] = [3]int{
				negat[0] * item[remap[0]],
				negat[1] * item[remap[1]],
				negat[2] * item[remap[2]],
			}
		}
		return ret
	}

	distancesFromScan0 := [][3]int{{0, 0, 0}}
	findAlignment := func(scanA, scanB [][3]int) (bool, [][3]int) {
		inA := make(map[[3]int]struct{})
		for _, x := range scanA {
			inA[x] = struct{}{}
		}
		for _, remap := range coordRemaps {
			for _, negat := range coordNegations {
				a := scanA
				b := apply(remap, negat, scanB)
				for _, aPos := range a {
					for _, bPos := range b {
						remapBy := [3]int{bPos[0] - aPos[0], bPos[1] - aPos[1], bPos[2] - aPos[2]}
						matches := 0
						var allRemapped [][3]int
						for _, otherB := range b {
							remappedToA := [3]int{otherB[0] - remapBy[0], otherB[1] - remapBy[1], otherB[2] - remapBy[2]}
							if _, exists := inA[remappedToA]; exists {
								matches++
							}
							allRemapped = append(allRemapped, remappedToA)
						}
						if matches >= 12 {
							distancesFromScan0 = append(distancesFromScan0, remapBy)
							return true, allRemapped
						}
					}
				}
			}
		}
		return false, nil
	}

	alignedIndices := make(map[int]struct{})
	alignedIndices[0] = struct{}{}
	aligned := make(map[int][][3]int)
	aligned[0] = ll[0]
	allAligned := make(map[[3]int]struct{})
	for _, x := range ll[0] {
		allAligned[x] = struct{}{}
	}
	noalign := make(map[[2]int]struct{})

	for len(alignedIndices) < len(ll) {
		for i := range ll {
			if _, exists := alignedIndices[i]; exists {
				continue
			}
			for j := range alignedIndices {
				if _, exists := noalign[[2]int{i, j}]; exists {
					continue
				}
				ok, remap := findAlignment(aligned[j], ll[i])
				if ok {
					alignedIndices[i] = struct{}{}
					aligned[i] = remap
					for _, x := range remap {
						allAligned[x] = struct{}{}
					}
					break
				}
				noalign[[2]int{i, j}] = struct{}{}
			}
		}
	}
	fmt.Println(len(allAligned))

	var dists []int
	for _, a := range distancesFromScan0 {
		for _, b := range distancesFromScan0 {
			dists = append(dists, mathx.Abs(a[0]-b[0])+mathx.Abs(a[1]-b[1])+mathx.Abs(a[2]-b[2]))
		}
	}
	fmt.Println(slices.Max(dists))
}
