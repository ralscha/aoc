package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"strconv"
	"strings"

	z3 "github.com/Z3Prover/z3/src/api/go"
)

func main() {
	input, err := download.ReadInput(2023, 24)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1and2(input)
}

type vector struct {
	x, y, z float64
}

type hailstone struct {
	position, velocity vector
}

func part1and2(input string) {
	lines := conv.SplitNewline(input)
	var hailstones []hailstone

	for _, line := range lines {

		parts := strings.Split(line, " @ ")
		posStr := strings.Split(strings.TrimSpace(parts[0]), ",")
		velStr := strings.Split(strings.TrimSpace(parts[1]), ",")

		posX := conv.MustAtoi(strings.TrimSpace(posStr[0]))
		posY := conv.MustAtoi(strings.TrimSpace(posStr[1]))
		posZ := conv.MustAtoi(strings.TrimSpace(posStr[2]))
		velX := conv.MustAtoi(strings.TrimSpace(velStr[0]))
		velY := conv.MustAtoi(strings.TrimSpace(velStr[1]))
		velZ := conv.MustAtoi(strings.TrimSpace(velStr[2]))

		h := hailstone{
			position: vector{x: float64(posX), y: float64(posY), z: float64(posZ)},
			velocity: vector{x: float64(velX), y: float64(velY), z: float64(velZ)},
		}

		hailstones = append(hailstones, h)
	}

	var lower float64 = 200000000000000
	var upper float64 = 400000000000000

	intersections := 0

	for i := range len(hailstones) - 1 {
		for j := i + 1; j < len(hailstones); j++ {

			h1 := hailstones[i]
			h2 := hailstones[j]

			apx := h1.position.x
			apy := h1.position.y
			avx := h1.velocity.x
			avy := h1.velocity.y
			y := (h2.position.x - (h2.velocity.x / h2.velocity.y * h2.position.y) + (h1.velocity.x / h1.velocity.y * h1.position.y) - h1.position.x) / (h1.velocity.x/h1.velocity.y - h2.velocity.x/h2.velocity.y)
			x := ((y-h1.position.y)/h1.velocity.y)*h1.velocity.x + h1.position.x
			if lower <= x && x <= upper && lower <= y && y <= upper {
				if (x-apx)/avx > 0 && (y-apy)/avy > 0 && (x-h2.position.x)/h2.velocity.x > 0 && (y-h2.position.y)/h2.velocity.y > 0 {
					intersections++
				}
			}
		}
	}

	fmt.Println(intersections)

	ctx := z3.NewContext()
	solver := ctx.NewSolver()

	x := ctx.MkIntConst("x")
	y := ctx.MkIntConst("y")
	z := ctx.MkIntConst("z")
	vx := ctx.MkIntConst("vx")
	vy := ctx.MkIntConst("vy")
	vz := ctx.MkIntConst("vz")

	for i, hs := range hailstones[:3] {
		a := ctx.MkInt64(int64(hs.position.x), ctx.MkIntSort())
		va := ctx.MkInt64(int64(hs.velocity.x), ctx.MkIntSort())
		b := ctx.MkInt64(int64(hs.position.y), ctx.MkIntSort())
		vb := ctx.MkInt64(int64(hs.velocity.y), ctx.MkIntSort())
		c := ctx.MkInt64(int64(hs.position.z), ctx.MkIntSort())
		vc := ctx.MkInt64(int64(hs.velocity.z), ctx.MkIntSort())

		t := ctx.MkIntConst("t" + strconv.Itoa(i))
		solver.Assert(ctx.MkGt(t, ctx.MkInt(0, ctx.MkIntSort())))
		solver.Assert(ctx.MkEq(ctx.MkAdd(x, ctx.MkMul(vx, t)), ctx.MkAdd(a, ctx.MkMul(va, t))))
		solver.Assert(ctx.MkEq(ctx.MkAdd(y, ctx.MkMul(vy, t)), ctx.MkAdd(b, ctx.MkMul(vb, t))))
		solver.Assert(ctx.MkEq(ctx.MkAdd(z, ctx.MkMul(vz, t)), ctx.MkAdd(c, ctx.MkMul(vc, t))))
	}

	if solver.Check() == z3.Satisfiable {
		value, _ := solver.Model().Eval(ctx.MkAdd(x, y, z), true)
		fmt.Println(value)
	} else {
		fmt.Println("Failed to solve!")
	}

}
