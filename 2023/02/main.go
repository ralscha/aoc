package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"strings"
)

func main() {
	input, err := download.ReadInput(2023, 2)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1and2(input)
}

func part1and2(input string) {
	lines := conv.SplitNewline(input)
	sum := 0
	sumMult := 0
	for _, line := range lines {
		maxRed := 0
		maxBlue := 0
		maxGreen := 0

		splitted := strings.Split(line, ":")
		game := splitted[0]
		plays := strings.Split(splitted[1], ";")
		possible := true
		for _, play := range plays {

			red := 0
			blue := 0
			green := 0
			for _, color := range strings.Split(play, ",") {
				color = strings.TrimSpace(color)
				colorSplitted := strings.Split(color, " ")
				number := conv.MustAtoi(strings.TrimSpace(colorSplitted[0]))
				color := strings.TrimSpace(colorSplitted[1])
				switch color {
				case "red":
					red += number
					if number > maxRed {
						maxRed = number
					}
				case "blue":
					blue += number
					if number > maxBlue {
						maxBlue = number
					}
				case "green":
					green += number
					if number > maxGreen {
						maxGreen = number
					}
				}
			}
			if red > 12 || green > 13 || blue > 14 {
				possible = false
			}
		}
		if possible {
			gameNumber := conv.MustAtoi(strings.Split(game, " ")[1])
			sum += gameNumber
		}

		maxMult := maxRed * maxBlue * maxGreen
		sumMult += maxMult
	}
	fmt.Println(sum)
	fmt.Println(sumMult)
}
