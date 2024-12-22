package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"strings"
)

func main() {
	input, err := download.ReadInput(2020, 4)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	passports := strings.Split(input, "\n\n")
	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	validCount := 0

	for _, passportStr := range passports {
		fields := make(map[string]string)
		parts := strings.Fields(passportStr)
		for _, part := range parts {
			kv := strings.Split(part, ":")
			if len(kv) == 2 {
				fields[kv[0]] = kv[1]
			}
		}

		isValid := true
		for _, field := range requiredFields {
			if _, ok := fields[field]; !ok {
				isValid = false
				break
			}
		}

		if isValid {
			validCount++
		}
	}

	fmt.Println("Part 1", validCount)
}

func part2(input string) {
	passports := strings.Split(input, "\n\n")
	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	validCount := 0

	for _, passportStr := range passports {
		fields := make(map[string]string)
		parts := strings.Fields(passportStr)
		for _, part := range parts {
			kv := strings.Split(part, ":")
			if len(kv) == 2 {
				fields[kv[0]] = kv[1]
			}
		}

		isValid := true
		for _, field := range requiredFields {
			if _, ok := fields[field]; !ok {
				isValid = false
				break
			}
		}

		if !isValid {
			continue
		}

		byr := conv.MustAtoi(fields["byr"])
		if byr < 1920 || byr > 2002 {
			continue
		}

		iyr := conv.MustAtoi(fields["iyr"])
		if iyr < 2010 || iyr > 2020 {
			continue
		}

		eyr := conv.MustAtoi(fields["eyr"])
		if eyr < 2020 || eyr > 2030 {
			continue
		}

		hgt := fields["hgt"]
		if strings.HasSuffix(hgt, "cm") {
			hgt = hgt[:len(hgt)-2]
			if h := conv.MustAtoi(hgt); h < 150 || h > 193 {
				continue
			}
		} else if strings.HasSuffix(hgt, "in") {
			hgt = hgt[:len(hgt)-2]
			if h := conv.MustAtoi(hgt); h < 59 || h > 76 {
				continue
			}
		} else {
			continue
		}

		hcl := fields["hcl"]
		if len(hcl) != 7 || hcl[0] != '#' {
			continue
		}
		for i := 1; i < 7; i++ {
			if !strings.Contains("0123456789abcdef", string(hcl[i])) {
				continue
			}
		}

		ecl := fields["ecl"]
		if ecl != "amb" && ecl != "blu" && ecl != "brn" && ecl != "gry" && ecl != "grn" && ecl != "hzl" && ecl != "oth" {
			continue
		}

		pid := fields["pid"]
		if len(pid) != 9 {
			continue
		}

		validCount++

	}

	fmt.Println("Part 2", validCount)
}
