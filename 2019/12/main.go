package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/mathx"
	"fmt"
	"log"
)

type moon struct {
	x, y, z    int
	vx, vy, vz int
}

func main() {
	input, err := download.ReadInput(2019, 12)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	moons := make([]moon, len(lines))
	for i, line := range lines {
		var x, y, z int
		_, err := fmt.Sscanf(line, "<x=%d, y=%d, z=%d>", &x, &y, &z)
		if err != nil {
			log.Fatalf("parsing moon %d failed: %v", i, err)
		}
		moons[i] = moon{x: x, y: y, z: z}
	}

	for i := 0; i < 1000; i++ {
		applyGravity(moons)
		applyVelocity(moons)
	}

	energy := 0
	for _, m := range moons {
		potential := mathx.Abs(m.x) + mathx.Abs(m.y) + mathx.Abs(m.z)
		kinetic := mathx.Abs(m.vx) + mathx.Abs(m.vy) + mathx.Abs(m.vz)
		energy += potential * kinetic
	}

	fmt.Println("Part 1", energy)
}

func applyGravity(moons []moon) {
	for i := range moons {
		for j := range moons {
			if i == j {
				continue
			}
			if moons[i].x < moons[j].x {
				moons[i].vx++
			} else if moons[i].x > moons[j].x {
				moons[i].vx--
			}
			if moons[i].y < moons[j].y {
				moons[i].vy++
			} else if moons[i].y > moons[j].y {
				moons[i].vy--
			}
			if moons[i].z < moons[j].z {
				moons[i].vz++
			} else if moons[i].z > moons[j].z {
				moons[i].vz--
			}
		}
	}
}

func applyVelocity(moons []moon) {
	for i := range moons {
		moons[i].x += moons[i].vx
		moons[i].y += moons[i].vy
		moons[i].z += moons[i].vz
	}
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	moons := make([]moon, len(lines))
	for i, line := range lines {
		var x, y, z int
		_, err := fmt.Sscanf(line, "<x=%d, y=%d, z=%d>", &x, &y, &z)
		if err != nil {
			log.Fatalf("parsing moon %d failed: %v", i, err)
		}
		moons[i] = moon{x: x, y: y, z: z}
	}

	periods := make([]int, 3)
	seen := make([]map[[8]int]int, 3)
	for i := range seen {
		seen[i] = make(map[[8]int]int)
	}

	for i := 0; i < 3; i++ {
		for j := 0; ; j++ {
			applyGravity(moons)
			applyVelocity(moons)
			state := [8]int{}
			if i == 0 {
				state = [8]int{
					moons[0].x, moons[0].vx,
					moons[1].x, moons[1].vx,
					moons[2].x, moons[2].vx,
					moons[3].x, moons[3].vx,
				}
			} else if i == 1 {
				state = [8]int{
					moons[0].y, moons[0].vy,
					moons[1].y, moons[1].vy,
					moons[2].y, moons[2].vy,
					moons[3].y, moons[3].vy,
				}
			} else {
				state = [8]int{
					moons[0].z, moons[0].vz,
					moons[1].z, moons[1].vz,
					moons[2].z, moons[2].vz,
					moons[3].z, moons[3].vz,
				}
			}
			if _, ok := seen[i][state]; ok {
				periods[i] = j
				break
			}
			seen[i][state] = j
		}
	}

	fmt.Println("Part 2", mathx.Lcm(periods))
}
