package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"github.com/aclements/go-z3/z3"
	"log"
	"strings"
)

func main() {
	inputFile := "./2022/21/input.txt"
	input, err := download.ReadInput(inputFile, 2022, 21)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1and2(input, false)
	part1and2(input, true)
}

func part1and2(input string, part2 bool) {
	lines := conv.SplitNewline(input)

	config := z3.NewContextConfig()
	ctx := z3.NewContext(config)
	solver := z3.NewSolver(ctx)

	for _, line := range lines {
		colonIx := strings.Index(line, ":")
		name := line[:colonIx]
		job := line[colonIx+2:]
		yells := 0
		if part2 && name == "humn" {
			ctx.IntConst("humn")
			continue
		}
		d := ctx.IntConst(name)
		if job[0] >= '0' && job[0] <= '9' {
			yells = conv.MustAtoi(job)
			solver.Assert(d.Eq(ctx.FromInt(int64(yells), ctx.IntSort()).(z3.Int)))
		} else {
			splitted := strings.Fields(job)
			operation := splitted[1]
			left := ctx.IntConst(splitted[0])
			right := ctx.IntConst(splitted[2])
			if part2 && name == "root" {
				solver.Assert(left.Eq(right))
			} else {
				switch operation {
				case "+":
					solver.Assert(d.Eq(left.Add(right)))
				case "-":
					solver.Assert(d.Eq(left.Sub(right)))
				case "*":
					solver.Assert(d.Eq(left.Mul(right)))
				case "/":
					solver.Assert(d.Eq(left.Div(right)))
				}
			}
		}

	}

	ok, err := solver.Check()
	if err != nil {
		panic(err)
	}
	if ok {
		if part2 {
			fmt.Println(solver.Model().Eval(ctx.IntConst("humn"), true))
		} else {
			fmt.Println(solver.Model().Eval(ctx.IntConst("root"), true))
		}
	} else {
		fmt.Println("unsat")
	}
}
