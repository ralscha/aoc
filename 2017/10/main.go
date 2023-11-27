package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"encoding/hex"
	"fmt"
	"log"
	"strings"
)

func main() {
	input, err := download.ReadInput(2017, 10)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}
	part1(input)
	part2(input)
}

func part1(input string) {
	s := conv.ToIntSlice(strings.Split(input, ","))
	list := make([]int, 256)
	for i := range list {
		list[i] = i
	}
	pos := 0
	skip := 0
	for _, l := range s {
		reverse(list, pos, l)
		pos += l + skip
		skip++
	}
	fmt.Println(list[0] * list[1])
}

func part2(input string) {
	s := make([]byte, len(input))
	for i, v := range input {
		s[i] = byte(v)
	}
	s = append(s, []byte{17, 31, 73, 47, 23}...)
	list := make([]int, 256)
	for i := range list {
		list[i] = i
	}
	pos := 0
	skip := 0
	for i := 0; i < 64; i++ {
		for _, l := range s {
			reverse(list, pos, int(l))
			pos += int(l) + skip
			skip++
		}
	}
	dense := make([]byte, 16)
	for i := 0; i < 16; i++ {
		for j := 0; j < 16; j++ {
			dense[i] ^= byte(list[i*16+j])
		}
	}

	fmt.Println(hex.EncodeToString(dense))
}

func reverse(list []int, pos, length int) {
	for i := 0; i < length/2; i++ {
		a := (pos + i) % len(list)
		b := (pos + length - i - 1) % len(list)
		list[a], list[b] = list[b], list[a]
	}
}
