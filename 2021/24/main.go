package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"github.com/aclements/go-z3/z3"
	"log"
	"math"
	"strconv"
	"strings"
)

func main() {
	input, err := download.ReadInput(2021, 24)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	solve(input, false)
}

func part2(input string) {
	solve(input, true)
}

func solve(input string, part2 bool) {
	config := z3.NewContextConfig()
	ctx := z3.NewContext(config)

	solver := z3.NewSolver(ctx)

	lines := conv.SplitNewline(input)

	next := 0
	var inputs []z3.Int

	for i := 0; i < 14; i++ {
		d := ctx.IntConst("i" + strconv.Itoa(i))
		solver.Assert(d.LE(ctx.FromInt(9, ctx.IntSort()).(z3.Int)))
		solver.Assert(d.GE(ctx.FromInt(1, ctx.IntSort()).(z3.Int)))
		inputs = append(inputs, d)
	}

	zero := ctx.FromInt(0, ctx.IntSort()).(z3.Int)
	one := ctx.FromInt(1, ctx.IntSort()).(z3.Int)

	registers := make(map[rune]z3.Int)
	registers['w'] = zero
	registers['x'] = zero
	registers['y'] = zero
	registers['z'] = zero

	for i, line := range lines {
		splitted := strings.Fields(line)
		instruction := splitted[0]
		if instruction == "inp" {
			register := rune(splitted[1][0])
			registers[register] = inputs[next]
			next++
			continue
		}
		c := ctx.IntConst("c" + strconv.Itoa(i))
		aStr := splitted[1]
		bStr := splitted[2]

		a := registers[rune(aStr[0])]
		var b z3.Int
		if bStr[0] >= 'w' && bStr[0] <= 'z' {
			b = registers[rune(bStr[0])]
		} else {
			b = ctx.FromInt(int64(conv.MustAtoi(bStr)), ctx.IntSort()).(z3.Int)
		}

		if instruction == "add" {
			solver.Assert(c.Eq(a.Add(b)))
		} else if instruction == "mul" {
			solver.Assert(c.Eq(a.Mul(b)))
		} else if instruction == "div" {
			solver.Assert(b.NE(zero))
			solver.Assert(c.Eq(a.Div(b)))
		} else if instruction == "mod" {
			solver.Assert(a.GE(zero))
			solver.Assert(b.GT(zero))
			solver.Assert(c.Eq(a.Mod(b)))
		} else if instruction == "eql" {
			solver.Assert(c.Eq(a.Eq(b).IfThenElse(one, zero).(z3.Int)))
		} else {
			panic("unknown instruction")
		}
		registers[rune(aStr[0])] = c
	}

	solver.Assert(registers['z'].Eq(zero))

	var best int64
	if part2 {
		best = int64(99999999999999)
	} else {
		best = 0
	}

	for {
		solver.Push()
		sum := zero
		for i, d := range inputs {
			sum = sum.Add(d.Mul(ctx.FromInt(int64(math.Pow(10, float64(13-i))), ctx.IntSort()).(z3.Int)))
		}

		if part2 {
			solver.Assert(sum.LT(ctx.FromInt(best, ctx.IntSort()).(z3.Int)))
		} else {
			solver.Assert(sum.GT(ctx.FromInt(best, ctx.IntSort()).(z3.Int)))
		}
		ok, err := solver.Check()
		if err != nil {
			panic(err)
		}
		if ok {
			b := solver.Model().Eval(sum, true).(z3.Int)
			best, _, _ = b.AsInt64()
		} else {
			fmt.Println(best)
			break
		}
		solver.Pop()
	}
}
