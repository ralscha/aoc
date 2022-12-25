package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"container/list"
	"fmt"
	"log"
	"strings"
)

func main() {
	inputFile := "./2015/22/input.txt"
	input, err := download.ReadInput(inputFile, 2015, 22)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	//part2(input)
}

type spell struct {
	name     string
	cost     int
	duration int
	damage   int
	heal     int
	armor    int
	mana     int
}

type state struct {
	spentMana   int
	playerHp    int
	playerMana  int
	playerArmor int
	bossHp      int
	bossDamage  int
	effects     []spell
}

func makeSpells() []spell {
	return []spell{
		{name: "Magic Missile", cost: 53, damage: 4},
		{name: "Drain", cost: 73, damage: 2, heal: 2},
		{name: "Shield", cost: 113, duration: 6, armor: 7},
		{name: "Poison", cost: 173, duration: 6, damage: 3},
		{name: "Recharge", cost: 229, duration: 5, mana: 101},
	}
}

func part1(input string) {
	splitted := strings.Fields(input)
	bossHp := conv.MustAtoi(splitted[2])
	bossDamage := conv.MustAtoi(splitted[4])

	spells := makeSpells()

	initialState := state{
		spentMana:  0,
		playerHp:   50,
		playerMana: 500,
		bossHp:     bossHp,
		bossDamage: bossDamage,
	}

	queue := list.New()
	queue.PushBack(initialState)

	minMana := 999999999
	for queue.Len() > 0 {
		currentState := queue.Remove(queue.Front()).(state)
		if currentState.playerHp <= 0 {
			continue
		}

		if currentState.spentMana >= minMana {
			continue
		}

		for _, spell := range spells {
			if spell.cost > currentState.playerMana {
				continue
			}

			newState := currentState
			newState.spentMana += spell.cost
			newState.playerMana -= spell.cost

			newState = applyEffects(newState)
			if newState.bossHp <= 0 {
				if newState.spentMana < minMana {
					minMana = newState.spentMana
				}
				continue
			}

			newState = applySpell(newState, spell)
			if newState.bossHp <= 0 {
				if newState.spentMana < minMana {
					minMana = newState.spentMana
				}
				continue
			}

			newState = applyEffects(newState)
			if newState.bossHp <= 0 {
				if newState.spentMana < minMana {
					minMana = newState.spentMana
				}
				continue
			}

			newState = applyBossAttack(newState)
			if newState.bossHp <= 0 {
				if newState.spentMana < minMana {
					minMana = newState.spentMana
				}
				continue
			}

			queue.PushBack(newState)
		}
	}

	fmt.Println(minMana)
}

func applyEffects(state *state) *state {
	for i, effect := range state.effects {
		if effect.duration == 0 {
			continue
		}

		state.playerArmor += effect.armor
		state.playerMana += effect.mana
		state.playerHp += effect.heal
		state.bossHp -= effect.damage
		state.effects[i].duration--
	}

	// remove expired effects
	newEffects := make([]spell, 0)
	for _, effect := range state.effects {
		if effect.duration > 0 {
			newEffects = append(newEffects, effect)
		}
	}

	return state
}

func applySpell(state *state, spell spell) *state {
	state.playerArmor += spell.armor
	state.playerMana += spell.mana
	state.playerHp += spell.heal
	state.bossHp -= spell.damage
	state.effects = append(state.effects, spell)

	return state
}

func applyBossAttack(state *state) *state {
	damage := state.bossDamage - state.playerArmor
	if damage < 1 {
		damage = 1
	}

	state.playerHp -= damage

	return state
}
