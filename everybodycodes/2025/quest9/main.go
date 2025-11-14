package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"fmt"
	"strings"
)

func main() {
	partI()
	partII()
	partIII()
}

func partI() {
	input := `1:GTGCGGTCCAATATAGGGCTCATGGTGAAGATAGTACTATTCGCCGCAATTGGGTTCGTACGGTCATAGAGAAGATATTGGGAAGATCTCTGGGAGGATATTTATCTTTTTAATTCGTCGTATCAGAG
2:AAATCCGAGCCCAAGTATACGAGGAGGGGCGCCGGGTGCTTGCTTTGGGAATTGACGGCGCGTAGTTGCGCCTGATAGTTCGGGCCTAGGACTATGATAAATTCGTGGATGGGGTGCTCTTCACAGGG
3:AAATCGTACCATAAGGATCCCATGGTGGGGATCGTATTATTGGCTTCAGAAGGGTTGGTGCGTACATGGGGATGATATTGCGAGCCTATCACGGTGGTTATTTCGCTGTTTGGTTCCTCGTAACAGAG`

	lines := conv.SplitNewline(input)
	sequences := make([]string, 3)

	for i, line := range lines {
		parts := strings.Split(line, ":")
		sequences[i] = parts[1]
	}

	for childIdx := range 3 {
		parent1Idx := (childIdx + 1) % 3
		parent2Idx := (childIdx + 2) % 3

		if isChild(sequences[childIdx], sequences[parent1Idx], sequences[parent2Idx]) {
			similarity := calculateSimilarity(sequences[childIdx], sequences[parent1Idx], sequences[parent2Idx])
			fmt.Println(similarity)
			break
		}
	}
}

func isChild(child, parent1, parent2 string) bool {
	for i := range len(child) {
		if child[i] != parent1[i] && child[i] != parent2[i] {
			return false
		}
	}
	return true
}

func calculateSimilarity(child, parent1, parent2 string) int {
	matches1 := 0
	matches2 := 0

	for i := range len(child) {
		if child[i] == parent1[i] {
			matches1++
		}
		if child[i] == parent2[i] {
			matches2++
		}
	}

	return matches1 * matches2
}

func partII() {
	input := `1:TGGCTATGATTCCATGAATCCTCATCGGAAATGCCTCCCATAGAACTCGAAGTGTTCTGATGCCGTAACCCCGGGTGAAGAGAAGTTGGAAAAGGGCCAAGGCTTTCAAATACGCAAGCGGCAAAATG
2:GCCACGATGGGGAACCGTCAGGGTCTCCAGTGTCCACTAGCAGTACGTATGTAGAGCTCGTCGCAGCCTTTGACCCCTCTTTAATAACGCCTCGTTTACCTCCTCTGACGATTGTCACGAGGATGCAG
3:AGTACCAGTACGTGGGAATTGATTCAAAAGGGGCGCTTGGGGCCGACAATAAGTCCCTTTACCTATGATCAAGCATGCTTGAACGGCCAACTGAACGATAGATAACAACCGCCGCGTTACGAGCGGTG
4:CTTCGGCACCAGTTATTATCCAGGTCACCCGAAGGTACCCGGGAGGGGGTTCGGTATCTAGAACGTTAGCGGCGCATAAAGTGAGAGGAGCTGAGCAGTGTTCCAGGGACGGAGATGTACAATTGCGT
5:GCCAAGCTGGCGAAGGATGCGCGTCTACGGTGTTCATTATCAGAACGGGGAAAAGCCTCGGCCCAATGTTAGACATTTCAAACTTAGCCCCTAGGTTACGACCTCCGCCGGATGTGAAGATGTTTCAG
6:CGCTCGCCTATTGCTCCTAATATGCACGGTCAGCAGGGTGTATTCTGAGCGTCACTGCCAGCTAAAGTACTGTGTCAAAGCAATTATAGACCATCATAAAGCGCGCCTAACTCAAACGTAAACACCTC`

	lines := conv.SplitNewline(input)
	sequences := make([]string, len(lines))

	for i, line := range lines {
		parts := strings.Split(line, ":")
		sequences[i] = parts[1]
	}

	totalSimilarity := 0
	childrenFound := container.NewSet[int]()

	for childIdx := range sequences {
		if childrenFound.Contains(childIdx) {
			continue
		}

		for parent1Idx := range sequences {
			if parent1Idx == childIdx {
				continue
			}

			for parent2Idx := parent1Idx + 1; parent2Idx < len(sequences); parent2Idx++ {
				if parent2Idx == childIdx {
					continue
				}

				if isChild(sequences[childIdx], sequences[parent1Idx], sequences[parent2Idx]) {
					similarity := calculateSimilarity(sequences[childIdx], sequences[parent1Idx], sequences[parent2Idx])
					totalSimilarity += similarity
					childrenFound.Add(childIdx)
					break
				}
			}
			if childrenFound.Contains(childIdx) {
				break
			}
		}
	}

	fmt.Println(totalSimilarity)
}

func partIII() {
	input := `1:GCAGAAATAGCAAACTGAGACGGTTAGACGTTCTGACCACCCCTCGGTGAATGAACAATCTATACGTCAGTGATCGATAGAATCTGGTGTGAATATGACTTGGAGACATAGCCGAAGACTTTACCGAG
2:GACACCATAGGACCTCTCTCGATTTTTTGGCACTCCCGATCTTATGGGCGTCGCCTAGGCCAGCGACAAGAACACAACTCCCCTTCTTTAAAGTGGGGAACCTCCACCTTCAACGCAGGGCTAAGGAA`

	lines := conv.SplitNewline(input)
	sequences := make(map[int]string)

	for _, line := range lines {
		parts := strings.Split(line, ":")
		id := conv.MustAtoi(parts[0])
		sequences[id] = parts[1]
	}

	graph := make(map[int]*container.Set[int])
	for id := range sequences {
		graph[id] = container.NewSet[int]()
	}

	for childID, childSeq := range sequences {
		for parent1ID, parent1Seq := range sequences {
			if parent1ID == childID {
				continue
			}
			for parent2ID, parent2Seq := range sequences {
				if parent2ID == childID || parent2ID <= parent1ID {
					continue
				}

				if isChild(childSeq, parent1Seq, parent2Seq) {
					graph[childID].Add(parent1ID)
					graph[childID].Add(parent2ID)
					graph[parent1ID].Add(childID)
					graph[parent1ID].Add(parent2ID)
					graph[parent2ID].Add(childID)
					graph[parent2ID].Add(parent1ID)
				}
			}
		}
	}

	visited := container.NewSet[int]()
	families := [][]int{}

	for id := range sequences {
		if !visited.Contains(id) {
			family := []int{}
			dfs(id, graph, visited, &family)
			families = append(families, family)
		}
	}

	maxSum := 0
	for _, family := range families {
		sum := 0
		for _, id := range family {
			sum += id
		}
		if sum > maxSum {
			maxSum = sum
		}
	}

	fmt.Println(maxSum)
}

func dfs(node int, graph map[int]*container.Set[int], visited *container.Set[int], family *[]int) {
	visited.Add(node)
	*family = append(*family, node)

	for _, neighbor := range graph[node].Values() {
		if !visited.Contains(neighbor) {
			dfs(neighbor, graph, visited, family)
		}
	}
}
