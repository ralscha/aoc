package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"container/list"
	"fmt"
	"log"
	"math"
	"strings"
)

func main() {
	inputFile := "./2015/22/input.txt"
	input, err := download.ReadInput(inputFile, 2015, 22)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1and2(input, false)
	part1and2(input, true)
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
	playerTurn  bool
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

func part1and2(input string, part2 bool) {
	splitted := strings.Fields(input)
	bossHp := conv.MustAtoi(splitted[2])
	bossDamage := conv.MustAtoi(splitted[4])

	spells := makeSpells()

	initialState := state{
		playerTurn: true,
		spentMana:  0,
		playerHp:   50,
		playerMana: 500,
		bossHp:     bossHp,
		bossDamage: bossDamage,
	}

	queue := list.New()
	queue.PushBack(initialState)
	minMana := math.MaxInt
	for queue.Len() > 0 {
		currentState := queue.Remove(queue.Front()).(state)
		currentState.playerArmor = 0
		if part2 && currentState.playerTurn {
			currentState.playerHp--
			if currentState.playerHp <= 0 {
				continue
			}
		}

		applyEffects(&currentState)
		if currentState.bossHp <= 0 {
			if currentState.spentMana < minMana {
				minMana = currentState.spentMana
			}
			continue
		}

		if !currentState.playerTurn {
			// boss turn
			applyBossAttack(&currentState)
			if currentState.playerHp <= 0 {
				continue
			}
			currentState.playerTurn = true
			queue.PushBack(currentState)
		} else {
			// player turn
			for _, sp := range spells {
				// already active
				alreadActive := false
				if sp.duration != 0 {
					for _, effect := range currentState.effects {
						if effect.name == sp.name {
							alreadActive = true
							continue
						}
					}
				}
				if alreadActive {
					continue
				}

				// not enough mana
				if currentState.playerMana < sp.cost {
					continue
				}

				if currentState.spentMana+sp.cost <= minMana {
					stateCopy := currentState
					stateCopy.playerMana -= sp.cost
					stateCopy.spentMana += sp.cost
					stateCopy.effects = append(currentState.effects, sp)
					stateCopy.playerTurn = false

					queue.PushBack(stateCopy)
				}
			}
		}
	}

	fmt.Println(minMana)
}

func applyEffects(state *state) {
	for i, effect := range state.effects {
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
	state.effects = newEffects
}

func applyBossAttack(state *state) {
	damage := state.bossDamage - state.playerArmor
	if damage < 1 {
		damage = 1
	}

	state.playerHp -= damage
}
