package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"strings"

	z3 "github.com/Z3Prover/z3/src/api/go"
)

func main() {
	input, err := download.ReadInput(2022, 21)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1and2(input, false)
	part1and2(input, true)
}

func part1and2(input string, part2 bool) {
	lines := conv.SplitNewline(input)

	ctx := z3.NewContext()
	solver := ctx.NewSolver()

	for _, line := range lines {
		colonIx := strings.Index(line, ":")
		name := line[:colonIx]
		job := line[colonIx+2:]
		yells := 0
		if part2 && name == "humn" {
			ctx.MkIntConst("humn")
			continue
		}
		d := ctx.MkIntConst(name)
		if job[0] >= '0' && job[0] <= '9' {
			yells = conv.MustAtoi(job)
			solver.Assert(ctx.MkEq(d, ctx.MkInt(yells, ctx.MkIntSort())))
		} else {
			splitted := strings.Fields(job)
			operation := splitted[1]
			left := ctx.MkIntConst(splitted[0])
			right := ctx.MkIntConst(splitted[2])
			if part2 && name == "root" {
				solver.Assert(ctx.MkEq(left, right))
			} else {
				switch operation {
				case "+":
					solver.Assert(ctx.MkEq(d, ctx.MkAdd(left, right)))
				case "-":
					solver.Assert(ctx.MkEq(d, ctx.MkSub(left, right)))
				case "*":
					solver.Assert(ctx.MkEq(d, ctx.MkMul(left, right)))
				case "/":
					solver.Assert(ctx.MkEq(d, ctx.MkDiv(left, right)))
				}
			}
		}

	}

	if solver.Check() == z3.Satisfiable {
		if part2 {
			value, _ := solver.Model().Eval(ctx.MkIntConst("humn"), true)
			fmt.Println(value)
		} else {
			value, _ := solver.Model().Eval(ctx.MkIntConst("root"), true)
			fmt.Println(value)
		}
	} else {
		fmt.Println("unsat")
	}
}
