package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"regexp"
	"slices"
	"sort"
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

func part2(input string) {
	fmt.Println("Part 2", "not yet implemented")
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

	immuneLines := strings.Split(parts[0], "\n")[1:]
	infectionLines := strings.Split(parts[1], "\n")[1:]

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
	allGroups := append(immuneSystem, infection...)

	slices.SortFunc(allGroups, func(a, b *group) int {
		if a.initiative > b.initiative {
			return -1
		}
		if a.initiative < b.initiative {
			return 1
		}
		return 0
	})

	targets := make(map[*group]*group)
	attackedBy := make(map[*group]*group)

	var selectableGroups []*group
	for _, g := range immuneSystem {
		if g.units > 0 {
			selectableGroups = append(selectableGroups, g)
		}
	}
	for _, g := range infection {
		if g.units > 0 {
			selectableGroups = append(selectableGroups, g)
		}
	}

	slices.SortFunc(selectableGroups, func(a, b *group) int {
		if a.effectivePower() > b.effectivePower() {
			return -1
		}
		if a.effectivePower() < b.effectivePower() {
			return 1
		}
		if a.initiative > b.initiative {
			return -1
		}
		if a.initiative < b.initiative {
			return 1
		}
		return 0
	})

	availableDefenders := make(map[*group]bool)
	for _, g := range infection {
		if g.units > 0 {
			availableDefenders[g] = true
		}
	}
	availableImmuneDefenders := make(map[*group]bool)
	for _, g := range immuneSystem {
		if g.units > 0 {
			availableImmuneDefenders[g] = true
		}
	}

	for _, attacker := range selectableGroups {
		var target *group
		maxDmg := -1

		var defenders []*group
		if attacker.isImmune {
			for defender := range availableDefenders {
				defenders = append(defenders, defender)
			}
		} else {
			for defender := range availableImmuneDefenders {
				defenders = append(defenders, defender)
			}
		}

		sort.Slice(defenders, func(i, j int) bool {
			dmg1 := attacker.damageTo(defenders[i])
			dmg2 := attacker.damageTo(defenders[j])
			if dmg1 != dmg2 {
				return dmg1 > dmg2
			}
			ep1 := defenders[i].effectivePower()
			ep2 := defenders[j].effectivePower()
			if ep1 != ep2 {
				return ep1 > ep2
			}
			return defenders[i].initiative > defenders[j].initiative
		})

		for _, defender := range defenders {
			dmg := attacker.damageTo(defender)
			if dmg > maxDmg {
				maxDmg = dmg
				target = defender
			}
		}

		if target != nil {
			if attacker.isImmune {
				if _, ok := attackedBy[target]; !ok {
					targets[attacker] = target
					delete(availableDefenders, target)
					attackedBy[target] = attacker
				}
			} else {
				if _, ok := attackedBy[target]; !ok {
					targets[attacker] = target
					delete(availableImmuneDefenders, target)
					attackedBy[target] = attacker
				}
			}
		}
	}

	slices.SortFunc(allGroups, func(a, b *group) int {
		if a.initiative > b.initiative {
			return -1
		}
		if a.initiative < b.initiative {
			return 1
		}
		return 0
	})

	unitsKilled := false
	for _, attacker := range allGroups {
		if attacker.units <= 0 {
			continue
		}
		defender, ok := targets[attacker]
		if ok {
			damage := attacker.damageTo(defender)
			unitsLost := damage / defender.hp
			if unitsLost > 0 {
				unitsKilled = true
			}
			defender.units -= unitsLost
			if defender.units < 0 {
				defender.units = 0
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
