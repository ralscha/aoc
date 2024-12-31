package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"math/big"
	"strings"
)

func main() {
	input, err := download.ReadInput(2019, 22)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	deckSize := 10007
	cardToFind := 2019
	deck := make([]int, deckSize)
	for i := range deckSize {
		deck[i] = i
	}

	lines := conv.SplitNewline(input)
	for _, line := range lines {
		if line == "deal into new stack" {
			for i, j := 0, deckSize-1; i < j; i, j = i+1, j-1 {
				deck[i], deck[j] = deck[j], deck[i]
			}
		} else if strings.HasPrefix(line, "cut ") {
			nStr := line[4:]
			n := conv.MustAtoi(nStr)
			if n > 0 {
				top := deck[:n]
				rest := deck[n:]
				deck = append(rest, top...)
			} else {
				n = -n
				bottom := deck[deckSize-n:]
				rest := deck[:deckSize-n]
				deck = append(bottom, rest...)
			}
		} else if strings.HasPrefix(line, "deal with increment ") {
			nStr := line[20:]
			increment := conv.MustAtoi(nStr)
			newDeck := make([]int, deckSize)
			for i := range deckSize {
				newDeck[i] = -1
			}
			currentPos := 0
			for _, card := range deck {
				newDeck[currentPos] = card
				currentPos = (currentPos + increment) % deckSize
			}
			deck = newDeck
		}
	}

	position := -1
	for i, card := range deck {
		if card == cardToFind {
			position = i
			break
		}
	}

	fmt.Println("Part 1", position)
}

type linFunc struct {
	k, m, n *big.Int
}

func newLinFunc(k, m, n int64) linFunc {
	return linFunc{
		k: big.NewInt(k),
		m: big.NewInt(m),
		n: big.NewInt(n),
	}
}

func (f linFunc) apply(x *big.Int) *big.Int {
	result := new(big.Int)
	result.Mul(f.k, x)
	result.Add(result, f.m)
	result.Mod(result, f.n)
	return result
}

func (f linFunc) compose(g linFunc) linFunc {
	k := new(big.Int).Mul(g.k, f.k)
	k.Mod(k, f.n)

	m := new(big.Int).Mul(g.k, f.m)
	m.Add(m, g.m)
	m.Mod(m, f.n)

	return linFunc{k: k, m: m, n: f.n}
}

func executeTimes(f linFunc, times *big.Int) linFunc {
	if times.Sign() == 0 {
		return newLinFunc(1, 0, f.n.Int64()) // Identity function
	}
	if times.Bit(0) == 0 { // Even number
		half := new(big.Int).Rsh(times, 1)
		g := executeTimes(f, half)
		return g.compose(g)
	}
	g := executeTimes(f, new(big.Int).Sub(times, big.NewInt(1)))
	return f.compose(g)
}

func part2(input string) {
	deckSize := big.NewInt(119315717514047)
	nShuffles := big.NewInt(101741582076661)
	position := big.NewInt(2020)

	shuffle := newLinFunc(1, 0, deckSize.Int64())

	lines := conv.SplitNewline(input)
	for i := len(lines) - 1; i >= 0; i-- {
		line := lines[i]
		var f linFunc

		if line == "deal into new stack" {
			f = newLinFunc(-1, -1, deckSize.Int64())
		} else if strings.HasPrefix(line, "cut ") {
			n := conv.MustAtoi(line[4:])
			f = newLinFunc(1, int64(n), deckSize.Int64())
		} else if strings.HasPrefix(line, "deal with increment ") {
			n := conv.MustAtoi(line[20:])
			nBig := big.NewInt(int64(n))
			inv := new(big.Int).ModInverse(nBig, deckSize)
			f = linFunc{k: inv, m: big.NewInt(0), n: deckSize}
		}

		shuffle = shuffle.compose(f)
	}

	finalFunc := executeTimes(shuffle, nShuffles)
	result := finalFunc.apply(position)

	fmt.Println("Part 2", result)
}
