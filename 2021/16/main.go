package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"strconv"
)

func main() {
	input, err := download.ReadInput(2021, 16)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1and2(input)
}

func part1and2(input string) {
	lines := conv.SplitNewline(input)
	binaryInput := hexToBinary(lines[0])
	_, value, versionSum := parsePackets(binaryInput, 0)
	fmt.Println("Part 1", versionSum)
	fmt.Println("Part 2", value)
}

func hexToBinary(hexStr string) string {
	binaryStr := ""
	for _, hexChar := range hexStr {
		bits, err := strconv.ParseUint(string(hexChar), 16, 4)
		if err != nil {
			log.Fatal(err)
		}
		binaryStr += fmt.Sprintf("%04b", bits)
	}
	return binaryStr
}

func parsePackets(binaryStr string, vTotal int64) (string, int64, int64) {
	if len(binaryStr) < 11 {
		return binaryStr, 0, vTotal
	}

	version, _ := strconv.ParseInt(binaryStr[:3], 2, 64)
	vTotal += version
	typeID, _ := strconv.ParseInt(binaryStr[3:6], 2, 64)
	binaryStr = binaryStr[6:]

	switch typeID {
	case 4:
		value := int64(0)
		for {
			group := binaryStr[:5]
			binaryStr = binaryStr[5:]
			value = (value << 4) | int64(group[1]-'0')<<3 | int64(group[2]-'0')<<2 | int64(group[3]-'0')<<1 | int64(group[4]-'0')
			if group[0] == '0' {
				break
			}
		}
		return binaryStr, value, vTotal
	default:
		var values []int64
		lengthTypeID := binaryStr[0]
		binaryStr = binaryStr[1:]
		if lengthTypeID == '0' {
			totalLength, _ := strconv.ParseInt(binaryStr[:15], 2, 64)
			binaryStr = binaryStr[15:]
			subPackets := binaryStr[:totalLength]
			binaryStr = binaryStr[totalLength:]
			for len(subPackets) > 0 {
				var value int64
				subPackets, value, vTotal = parsePackets(subPackets, vTotal)
				values = append(values, value)
			}
		} else {
			numSubPackets, _ := strconv.ParseInt(binaryStr[:11], 2, 64)
			binaryStr = binaryStr[11:]
			for range int(numSubPackets) {
				var value int64
				binaryStr, value, vTotal = parsePackets(binaryStr, vTotal)
				values = append(values, value)
			}
		}
		return binaryStr, evaluate(typeID, values), vTotal
	}
}

func evaluate(typeID int64, values []int64) int64 {
	switch typeID {
	case 0:
		sum := int64(0)
		for _, v := range values {
			sum += v
		}
		return sum
	case 1:
		product := int64(1)
		for _, v := range values {
			product *= v
		}
		return product
	case 2:
		minValue := values[0]
		for _, v := range values[1:] {
			if v < minValue {
				minValue = v
			}
		}
		return minValue
	case 3:
		maxValue := values[0]
		for _, v := range values[1:] {
			if v > maxValue {
				maxValue = v
			}
		}
		return maxValue
	case 5:
		if values[0] > values[1] {
			return 1
		}
		return 0
	case 6:
		if values[0] < values[1] {
			return 1
		}
		return 0
	case 7:
		if values[0] == values[1] {
			return 1
		}
		return 0
	default:
		log.Fatal("Unknown type ID: ", typeID)
	}
	return 0
}
