package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"math"
	"strings"
)

func main() {
	input, err := download.ReadInput(2015, 21)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1and2(input)
}

type item struct {
	cost   int
	damage int
	armor  int
}

type player struct {
	hitpoints int
	damage    int
	armor     int
}

type boss struct {
	hitpoints int
	damage    int
	armor     int
}

func part1and2(input string) {
	lines := conv.SplitNewline(input)

	bossHitPoints := conv.MustAtoi(strings.Fields(lines[0])[2])
	bossDamage := conv.MustAtoi(strings.Fields(lines[1])[1])
	bossArmor := conv.MustAtoi(strings.Fields(lines[2])[1])
	boss := boss{bossHitPoints, bossDamage, bossArmor}

	player := player{100, 0, 0}
	minCost, maxCost := getMinMaxCost(player, boss)
	fmt.Println("Part 1", minCost)
	fmt.Println("Part 2", maxCost)

}

func getMinMaxCost(player player, boss boss) (int, int) {
	minCost := math.MaxInt32
	maxCost := 0
	for _, weapon := range weaponsInShop() {
		for _, armor := range armorInShop() {
			for _, ring1 := range ringsInShop() {
				for _, ring2 := range ringsInShop() {
					if ring1 == ring2 {
						continue
					}
					player.damage = weapon.damage + ring1.damage + ring2.damage
					player.armor = armor.armor + ring1.armor + ring2.armor
					cost := weapon.cost + armor.cost + ring1.cost + ring2.cost
					if doesPlayerWin(player, boss) && cost < minCost {
						minCost = cost
					}
					if !doesPlayerWin(player, boss) && cost > maxCost {
						maxCost = cost
					}
				}
			}
		}
	}
	return minCost, maxCost
}

func doesPlayerWin(player player, boss boss) bool {
	bossDamagePerTurn := player.damage - boss.armor
	if bossDamagePerTurn < 1 {
		bossDamagePerTurn = 1
	}
	playerDamagePerTurn := max(boss.damage-player.armor, 1)

	bossTurns := boss.hitpoints / bossDamagePerTurn
	if boss.hitpoints%bossDamagePerTurn != 0 {
		bossTurns++
	}
	playerTurns := player.hitpoints / playerDamagePerTurn
	if player.hitpoints%playerDamagePerTurn != 0 {
		playerTurns++
	}

	return playerTurns >= bossTurns
}

func weaponsInShop() []item {
	return []item{
		{8, 4, 0},
		{10, 5, 0},
		{25, 6, 0},
		{40, 7, 0},
		{74, 8, 0},
	}
}

func armorInShop() []item {
	return []item{
		{13, 0, 1},
		{31, 0, 2},
		{53, 0, 3},
		{75, 0, 4},
		{102, 0, 5},
		{0, 0, 0},
	}
}

func ringsInShop() []item {
	return []item{
		{25, 1, 0},
		{50, 2, 0},
		{100, 3, 0},
		{20, 0, 1},
		{40, 0, 2},
		{80, 0, 3},
		{0, 0, 0},
	}
}
