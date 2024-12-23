package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"github.com/aclements/go-z3/z3"
	"log"
	"strconv"
	"strings"
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

	config := z3.NewContextConfig()
	ctx := z3.NewContext(config)
	solver := z3.NewSolver(ctx)

	x := ctx.IntConst("x")
	y := ctx.IntConst("y")
	z := ctx.IntConst("z")
	vx := ctx.IntConst("vx")
	vy := ctx.IntConst("vy")
	vz := ctx.IntConst("vz")

	for i, hs := range hailstones[:3] {
		a := ctx.FromInt(int64(hs.position.x), ctx.IntSort()).(z3.Int)
		va := ctx.FromInt(int64(hs.velocity.x), ctx.IntSort()).(z3.Int)
		b := ctx.FromInt(int64(hs.position.y), ctx.IntSort()).(z3.Int)
		vb := ctx.FromInt(int64(hs.velocity.y), ctx.IntSort()).(z3.Int)
		c := ctx.FromInt(int64(hs.position.z), ctx.IntSort()).(z3.Int)
		vc := ctx.FromInt(int64(hs.velocity.z), ctx.IntSort()).(z3.Int)

		t := ctx.IntConst("t" + strconv.Itoa(i))
		solver.Assert(t.GT(ctx.FromInt(0, ctx.IntSort()).(z3.Int)))
		solver.Assert(x.Add(vx.Mul(t)).Eq(a.Add(va.Mul(t))))
		solver.Assert(y.Add(vy.Mul(t)).Eq(b.Add(vb.Mul(t))))
		solver.Assert(z.Add(vz.Mul(t)).Eq(c.Add(vc.Mul(t))))
	}

	ok, err := solver.Check()
	if err != nil {
		panic(err)
	}
	if ok {
		fmt.Println(solver.Model().Eval(x.Add(y).Add(z), true))
	} else {
		fmt.Println("Failed to solve!")
	}

}
