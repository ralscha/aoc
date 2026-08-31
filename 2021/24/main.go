package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	z3 "github.com/Z3Prover/z3/src/api/go"
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
	ctx := z3.NewContext()
	solver := ctx.NewSolver()

	lines := conv.SplitNewline(input)

	next := 0
	var inputs []*z3.Expr

	for i := range 14 {
		d := ctx.MkIntConst("i" + strconv.Itoa(i))
		solver.Assert(ctx.MkLe(d, ctx.MkInt(9, ctx.MkIntSort())))
		solver.Assert(ctx.MkGe(d, ctx.MkInt(1, ctx.MkIntSort())))
		inputs = append(inputs, d)
	}

	zero := ctx.MkInt(0, ctx.MkIntSort())
	one := ctx.MkInt(1, ctx.MkIntSort())

	registers := make(map[rune]*z3.Expr)
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
		c := ctx.MkIntConst("c" + strconv.Itoa(i))
		aStr := splitted[1]
		bStr := splitted[2]

		a := registers[rune(aStr[0])]
		var b *z3.Expr
		if bStr[0] >= 'w' && bStr[0] <= 'z' {
			b = registers[rune(bStr[0])]
		} else {
			b = ctx.MkInt(conv.MustAtoi(bStr), ctx.MkIntSort())
		}

		if instruction == "add" {
			solver.Assert(ctx.MkEq(c, ctx.MkAdd(a, b)))
		} else if instruction == "mul" {
			solver.Assert(ctx.MkEq(c, ctx.MkMul(a, b)))
		} else if instruction == "div" {
			solver.Assert(ctx.MkNot(ctx.MkEq(b, zero)))
			solver.Assert(ctx.MkEq(c, ctx.MkDiv(a, b)))
		} else if instruction == "mod" {
			solver.Assert(ctx.MkGe(a, zero))
			solver.Assert(ctx.MkGt(b, zero))
			solver.Assert(ctx.MkEq(c, ctx.MkMod(a, b)))
		} else if instruction == "eql" {
			equal := ctx.MkEq(a, b)
			solver.Assert(ctx.MkOr(
				ctx.MkAnd(equal, ctx.MkEq(c, one)),
				ctx.MkAnd(ctx.MkNot(equal), ctx.MkEq(c, zero)),
			))
		} else {
			panic("unknown instruction")
		}
		registers[rune(aStr[0])] = c
	}

	solver.Assert(ctx.MkEq(registers['z'], zero))

	var best int64
	if part2 {
		best = int64(99999999999999)
	} else {
		best = 0
	}

search:
	for {
		solver.Push()
		sum := zero
		for i, d := range inputs {
			place := ctx.MkInt64(int64(math.Pow(10, float64(13-i))), ctx.MkIntSort())
			sum = ctx.MkAdd(sum, ctx.MkMul(d, place))
		}

		if part2 {
			solver.Assert(ctx.MkLt(sum, ctx.MkInt64(best, ctx.MkIntSort())))
		} else {
			solver.Assert(ctx.MkGt(sum, ctx.MkInt64(best, ctx.MkIntSort())))
		}
		switch status := solver.Check(); status {
		case z3.Satisfiable:
			value, ok := solver.Model().Eval(sum, true)
			if !ok {
				panic("failed to evaluate model number")
			}
			parsedBest, err := strconv.ParseInt(value.String(), 10, 64)
			if err != nil {
				panic(err)
			}
			best = parsedBest
		case z3.Unsatisfiable:
			fmt.Println(best)
			break search
		default:
			panic("Z3 returned unknown")
		}
		solver.Pop(1)
	}
}
