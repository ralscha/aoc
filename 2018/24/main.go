package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"regexp"
	"slices"
	"strings"
)

func main() {
	input, err := download.ReadInput(2018, 24)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	immuneSystem, infection := parseInput(input)

	for len(immuneSystem) > 0 && len(infection) > 0 {
		newImmune, newInfection := simulateFight(immuneSystem, infection)
		if newImmune == nil && newInfection == nil {
			break
		}
		immuneSystem = newImmune
		infection = newInfection
	}

	winningUnits := 0
	for _, g := range immuneSystem {
		winningUnits += g.units
	}
	for _, g := range infection {
		winningUnits += g.units
	}

	fmt.Println("Part 1", winningUnits)
}

func simulateWithBoost(input string, boost int) int {
	immuneSystem, infection := parseInput(input)

	for _, g := range immuneSystem {
		g.attackDmg += boost
	}

	for len(immuneSystem) > 0 && len(infection) > 0 {
		newImmune, newInfection := simulateFight(immuneSystem, infection)
		if newImmune == nil && newInfection == nil {
			return 0
		}
		immuneSystem = newImmune
		infection = newInfection
	}

	immuneUnits := 0
	for _, g := range immuneSystem {
		immuneUnits += g.units
	}
	if immuneUnits > 0 && len(infection) == 0 {
		return immuneUnits
	}
	return 0
}

func part2(input string) {
	lo := 0
	loBool := simulateWithBoost(input, lo) > 0

	offset := 1
	for simulateWithBoost(input, lo+offset) == 0 {
		offset *= 2
	}
	hi := lo + offset

	bestSoFar := hi
	for lo <= hi {
		mid := (hi + lo) / 2
		result := simulateWithBoost(input, mid)
		if result > 0 {
			bestSoFar = mid
		}
		if (result > 0) == loBool {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}

	result := simulateWithBoost(input, bestSoFar)
	fmt.Println("Part 2", result)
}

type group struct {
	id         int
	units      int
	hp         int
	attackDmg  int
	attackType string
	initiative int
	weaknesses *container.Set[string]
	immunities *container.Set[string]
	isImmune   bool
}

func (g *group) effectivePower() int {
	return g.units * g.attackDmg
}

func (g *group) damageTo(opponent *group) int {
	if opponent.immunities.Contains(g.attackType) {
		return 0
	}
	dmg := g.effectivePower()
	if opponent.weaknesses.Contains(g.attackType) {
		dmg *= 2
	}
	return dmg
}

func parseGroup(line string, id int, isImmune bool) *group {
	re := regexp.MustCompile(`(\d+) units each with (\d+) hit points(?: \((.+)\))? with an attack that does (\d+) (\w+) damage at initiative (\d+)`)
	matches := re.FindStringSubmatch(line)
	if len(matches) == 0 {
		return nil
	}

	units := conv.MustAtoi(matches[1])
	hp := conv.MustAtoi(matches[2])
	attackDmg := conv.MustAtoi(matches[4])
	attackType := matches[5]
	initiative := conv.MustAtoi(matches[6])

	weaknesses := container.NewSet[string]()
	immunities := container.NewSet[string]()

	if matches[3] != "" {
		parts := strings.Split(matches[3], "; ")
		for _, part := range parts {
			if strings.HasPrefix(part, "weak to ") {
				types := strings.Split(part[8:], ", ")
				for _, t := range types {
					weaknesses.Add(t)
				}
			} else if strings.HasPrefix(part, "immune to ") {
				types := strings.Split(part[10:], ", ")
				for _, t := range types {
					immunities.Add(t)
				}
			}
		}
	}

	return &group{id: id, units: units, hp: hp, attackDmg: attackDmg, attackType: attackType, initiative: initiative, weaknesses: weaknesses, immunities: immunities, isImmune: isImmune}
}

func parseInput(input string) ([]*group, []*group) {
	parts := strings.Split(input, "\n\n")
	if len(parts) != 2 {
		log.Fatalf("invalid input format")
	}

	immuneLines := conv.SplitNewline(parts[0])[1:]
	infectionLines := conv.SplitNewline(parts[1])[1:]

	var immuneSystem []*group
	for i, line := range immuneLines {
		if line == "" {
			continue
		}
		immuneSystem = append(immuneSystem, parseGroup(line, i+1, true))
	}

	var infection []*group
	for i, line := range infectionLines {
		if line == "" {
			continue
		}
		infection = append(infection, parseGroup(line, i+1, false))
	}

	return immuneSystem, infection
}

func simulateFight(immuneSystem []*group, infection []*group) ([]*group, []*group) {
	for _, team := range [][]*group{immuneSystem, infection} {
		slices.SortFunc(team, func(a, b *group) int {
			if a.effectivePower() != b.effectivePower() {
				if a.effectivePower() > b.effectivePower() {
					return -1
				}
				return 1
			}
			if a.initiative > b.initiative {
				return -1
			}
			return 1
		})
	}

	immuneTargets := make([]*group, len(immuneSystem))
	infectionTargets := make([]*group, len(infection))

	selectTargets := func(attackers []*group, defenders []*group, targets []*group) {
		availableTargets := make([]*group, 0, len(defenders))
		for _, g := range defenders {
			if g.units > 0 {
				availableTargets = append(availableTargets, g)
			}
		}

		for i, attacker := range attackers {
			if attacker.units <= 0 {
				continue
			}

			slices.SortFunc(availableTargets, func(a, b *group) int {
				dmgA := attacker.damageTo(a)
				dmgB := attacker.damageTo(b)
				if dmgA != dmgB {
					if dmgA > dmgB {
						return -1
					}
					return 1
				}
				if a.effectivePower() != b.effectivePower() {
					if a.effectivePower() > b.effectivePower() {
						return -1
					}
					return 1
				}
				if a.initiative > b.initiative {
					return -1
				}
				return 1
			})

			for k, target := range availableTargets {
				if attacker.damageTo(target) > 0 {
					targets[i] = target
					availableTargets = append(availableTargets[:k], availableTargets[k+1:]...)
					break
				}
			}
		}
	}

	selectTargets(immuneSystem, infection, immuneTargets)
	selectTargets(infection, immuneSystem, infectionTargets)

	type attack struct {
		attacker   *group
		defender   *group
		isImmune   bool
		initiative int
	}

	var attacks []attack
	for i, attacker := range immuneSystem {
		if attacker.units > 0 && immuneTargets[i] != nil {
			attacks = append(attacks, attack{
				attacker:   attacker,
				defender:   immuneTargets[i],
				isImmune:   true,
				initiative: attacker.initiative,
			})
		}
	}
	for i, attacker := range infection {
		if attacker.units > 0 && infectionTargets[i] != nil {
			attacks = append(attacks, attack{
				attacker:   attacker,
				defender:   infectionTargets[i],
				isImmune:   false,
				initiative: attacker.initiative,
			})
		}
	}

	slices.SortFunc(attacks, func(a, b attack) int {
		if a.initiative > b.initiative {
			return -1
		}
		return 1
	})

	unitsKilled := false
	for _, atk := range attacks {
		if atk.attacker.units <= 0 {
			continue
		}
		damage := atk.attacker.damageTo(atk.defender)
		unitsLost := damage / atk.defender.hp
		if unitsLost > 0 {
			unitsKilled = true
			atk.defender.units -= unitsLost
			if atk.defender.units < 0 {
				atk.defender.units = 0
			}
		}
	}

	var newImmuneSystem []*group
	for _, g := range immuneSystem {
		if g.units > 0 {
			newImmuneSystem = append(newImmuneSystem, g)
		}
	}

	var newInfection []*group
	for _, g := range infection {
		if g.units > 0 {
			newInfection = append(newInfection, g)
		}
	}

	if !unitsKilled {
		return nil, nil
	}

	return newImmuneSystem, newInfection
}
