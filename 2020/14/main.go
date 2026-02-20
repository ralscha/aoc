package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input, err := download.ReadInput(2020, 14)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	memory := make(map[int]int64)
	mask := ""

	maskRegex := regexp.MustCompile(`^mask = ([01X]{36})$`)
	memRegex := regexp.MustCompile(`^mem\[(\d+)] = (\d+)$`)

	for _, line := range lines {
		if maskMatch := maskRegex.FindStringSubmatch(line); len(maskMatch) > 0 {
			mask = maskMatch[1]
		} else if memMatch := memRegex.FindStringSubmatch(line); len(memMatch) > 0 {
			addressStr := memMatch[1]
			valueStr := memMatch[2]
			address := conv.MustAtoi(addressStr)
			value := conv.MustAtoi64(valueStr)

			maskedValue := applyMaskPart1(value, mask)
			memory[address] = maskedValue
		}
	}

	var sum int64
	for _, val := range memory {
		sum += val
	}

	fmt.Println("Part 1", sum)
}

func applyMaskPart1(value int64, mask string) int64 {
	valueBits := fmt.Sprintf("%036b", value)
	var maskedBits strings.Builder
	for i := range 36 {
		switch mask[i] {
		case '0':
			maskedBits.WriteString("0")
		case '1':
			maskedBits.WriteString("1")
		case 'X':
			maskedBits.WriteString(string(valueBits[i]))
		}
	}
	maskedValue, _ := strconv.ParseInt(maskedBits.String(), 2, 64)
	return maskedValue
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	memory := make(map[int64]int64)
	mask := ""

	maskRegex := regexp.MustCompile(`^mask = ([01X]{36})$`)
	memRegex := regexp.MustCompile(`^mem\[(\d+)] = (\d+)$`)

	for _, line := range lines {
		if maskMatch := maskRegex.FindStringSubmatch(line); len(maskMatch) > 0 {
			mask = maskMatch[1]
		} else if memMatch := memRegex.FindStringSubmatch(line); len(memMatch) > 0 {
			addressStr := memMatch[1]
			valueStr := memMatch[2]
			address := conv.MustAtoi64(addressStr)
			value := conv.MustAtoi64(valueStr)

			maskedAddresses := applyMaskPart2(address, mask)
			for _, maskedAddress := range maskedAddresses {
				memory[maskedAddress] = value
			}
		}
	}

	var sum int64
	for _, val := range memory {
		sum += val
	}

	fmt.Println("Part 2", sum)
}

func applyMaskPart2(address int64, mask string) []int64 {
	addressBits := fmt.Sprintf("%036b", address)
	var maskedBits strings.Builder
	var floatingIndices []int
	for i := range 36 {
		switch mask[i] {
		case '0':
			maskedBits.WriteString(string(addressBits[i]))
		case '1':
			maskedBits.WriteString("1")
		case 'X':
			maskedBits.WriteString("X")
			floatingIndices = append(floatingIndices, i)
		}
	}

	var addresses []int64
	numFloating := len(floatingIndices)
	for i := range 1 << numFloating {
		currentBits := maskedBits.String()
		temp := i
		for j := range numFloating {
			bit := temp & 1
			temp >>= 1
			index := floatingIndices[j]
			if bit == 0 {
				currentBits = currentBits[:index] + "0" + currentBits[index+1:]
			} else {
				currentBits = currentBits[:index] + "1" + currentBits[index+1:]
			}
		}
		addr, _ := strconv.ParseInt(currentBits, 2, 64)
		addresses = append(addresses, addr)
	}
	return addresses
}
