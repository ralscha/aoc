package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"strings"
)

func main() {
	input, err := download.ReadInput(2019, 6)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	orbits := make(map[string]string)
	for _, line := range lines {
		split := strings.Split(line, ")")
		orbits[split[1]] = split[0]
	}

	count := 0
	for k := range orbits {
		count++
		for orbits[k] != "COM" {
			count++
			k = orbits[k]
		}
	}

	fmt.Println("Result 1:", count)
}

/*
Now, you just need to figure out how many orbital transfers you (YOU) need to take to get to Santa (SAN).

You start at the object YOU are orbiting; your destination is the object SAN is orbiting. An orbital transfer lets you move from any object to an object orbiting or orbited by that object.

For example, suppose you have the following map:

COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L
K)YOU
I)SAN
Visually, the above map of orbits looks like this:

                          YOU
                         /
        G - H       J - K - L
       /           /
COM - B - C - D - E - F
               \
                I - SAN
In this example, YOU are in orbit around K, and SAN is in orbit around I. To move from K to I, a minimum of 4 orbital transfers are required:

K to J
J to E
E to D
D to I
Afterward, the map of orbits looks like this:

        G - H       J - K - L
       /           /
COM - B - C - D - E - F
               \
                I - SAN
                 \
                  YOU
What is the minimum number of orbital transfers required to move from the object YOU are orbiting to the object SAN is orbiting? (Between the objects they are orbiting - not between YOU and SAN.)
*/

func part2(input string) {
	lines := conv.SplitNewline(input)
	orbits := make(map[string]string)
	for _, line := range lines {
		split := strings.Split(line, ")")
		orbits[split[1]] = split[0]
	}

	you := make([]string, 0)
	san := make([]string, 0)

	for k := "YOU"; k != "COM"; k = orbits[k] {
		you = append(you, k)
	}

	for k := "SAN"; k != "COM"; k = orbits[k] {
		san = append(san, k)
	}

	for i, y := range you {
		for j, s := range san {
			if y == s {
				fmt.Println("Result 2:", i+j-2)
				return
			}
		}
	}
}
