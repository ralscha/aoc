package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/mathx"
	"fmt"
	"golang.org/x/exp/maps"
	"log"
	"strings"
)

func main() {
	input, err := download.ReadInput(2023, 20)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
}

var modules map[string]module

type pulseSend struct {
	caller   string
	receiver string
	pulse    int
}

var pulseQueue []pulseSend

type module interface {
	receive(caller string, pulse int)
	destinations() []string
	name() string
}

type broadcaster struct {
	moduleName         string
	moduleDestinations []string
}

func (b *broadcaster) receive(_ string, pulse int) {
	for _, dest := range b.moduleDestinations {
		pulseQueue = append(pulseQueue, pulseSend{caller: b.moduleName, receiver: dest, pulse: pulse})
	}
}

func (b *broadcaster) destinations() []string {
	return b.moduleDestinations
}

func (b *broadcaster) name() string {
	return b.moduleName
}

type flipFlop struct {
	moduleName         string
	moduleDestinations []string
	state              bool
}

func (ff *flipFlop) receive(_ string, pulse int) {
	if pulse == 0 {
		var nextPulse int
		if !ff.state {
			ff.state = true
			nextPulse = 1
		} else {
			ff.state = false
			nextPulse = 0
		}
		for _, dest := range ff.moduleDestinations {
			pulseQueue = append(pulseQueue, pulseSend{caller: ff.moduleName, receiver: dest, pulse: nextPulse})
		}
	}
}

func (ff *flipFlop) destinations() []string {
	return ff.moduleDestinations
}

func (ff *flipFlop) name() string {
	return ff.moduleName
}

type conjunction struct {
	moduleName         string
	inputStates        map[string]int
	moduleDestinations []string
}

func (c *conjunction) receive(caller string, pulse int) {
	c.inputStates[caller] = pulse
	allHigh := true
	for _, state := range c.inputStates {
		if state == 0 {
			allHigh = false
			break
		}
	}
	nextPulse := 1
	if allHigh {
		nextPulse = 0
	}
	for _, dest := range c.moduleDestinations {
		pulseQueue = append(pulseQueue, pulseSend{caller: c.moduleName, receiver: dest, pulse: nextPulse})
	}
}

func (c *conjunction) destinations() []string {
	return c.moduleDestinations
}

func (c *conjunction) name() string {
	return c.moduleName
}

func part1(input string) {
	lines := conv.SplitNewline(input)

	modules = make(map[string]module)
	var conjunctions []*conjunction
	var flipFlops []*flipFlop

	for _, line := range lines {
		splitted := strings.Split(line, "->")
		if splitted[0] == "broadcaster " {
			broadcasterTargets := strings.Split(splitted[1], ",")
			for i, target := range broadcasterTargets {
				broadcasterTargets[i] = strings.TrimSpace(target)
			}
			modules["broadcaster"] = &broadcaster{moduleName: "broadcaster", moduleDestinations: broadcasterTargets}
		} else {
			module := strings.TrimSpace(splitted[0])
			moduleName := module[1:]
			targets := strings.Split(splitted[1], ",")
			for i, target := range targets {
				targets[i] = strings.TrimSpace(target)
			}
			if module[0] == '&' {
				c := &conjunction{moduleName: moduleName, moduleDestinations: targets, inputStates: make(map[string]int)}
				modules[moduleName] = c
				conjunctions = append(conjunctions, c)
			} else if module[0] == '%' {
				f := &flipFlop{moduleName: moduleName, moduleDestinations: targets}
				modules[moduleName] = f
				flipFlops = append(flipFlops, f)
			}
		}
	}

	for _, c := range conjunctions {
		for _, m := range modules {
			for _, dest := range m.destinations() {
				if dest == c.moduleName {
					c.inputStates[m.name()] = 0
				}
			}
		}
	}

	lowPulseSent := 0
	highPulseSent := 0

	for i := 0; i < 1000; i++ {
		pulseQueue = append(pulseQueue, pulseSend{caller: "", receiver: "broadcaster", pulse: 0})
		for len(pulseQueue) > 0 {
			pulse := pulseQueue[0]
			pulseQueue = pulseQueue[1:]

			/*
				if pulse.receiver == "kc" && pulse.pulse == 1 {
					fmt.Println(pulse.caller, cycleCount)
				}
			*/

			receiver := modules[pulse.receiver]
			if pulse.pulse == 0 {
				lowPulseSent++
			} else {
				highPulseSent++
			}
			if receiver != nil {
				receiver.receive(pulse.caller, pulse.pulse)
			}
		}
	}

	fmt.Println(lowPulseSent * highPulseSent)

	for _, c := range conjunctions {
		for k := range c.inputStates {
			c.inputStates[k] = 0
		}
	}
	for _, f := range flipFlops {
		f.state = false
	}

	// search conjunction that sends to rx
	senderConjunction := ""
	for _, line := range lines {
		if strings.HasSuffix(line, "rx") {
			splitted := strings.Split(line, "->")
			sender := strings.TrimSpace(splitted[0])
			senderConjunction = sender[1:]
			break
		}
	}

	var sender *conjunction
	for _, c := range conjunctions {
		if c.moduleName == senderConjunction {
			sender = c
			break
		}
	}

	// search all inputs of sender conjunction
	inputs := make(map[string]int)
	for k := range sender.inputStates {
		inputs[k] = 0
	}

	cycleCount := 1
outer:
	for {
		pulseQueue = append(pulseQueue, pulseSend{caller: "", receiver: "broadcaster", pulse: 0})
		for len(pulseQueue) > 0 {
			pulse := pulseQueue[0]
			pulseQueue = pulseQueue[1:]

			if pulse.receiver == senderConjunction && pulse.pulse == 1 {
				if _, ok := inputs[pulse.caller]; ok {
					inputs[pulse.caller] = cycleCount
				}
				found := true
				for _, v := range inputs {
					if v == 0 {
						found = false
					}
				}
				if found {
					break outer
				}
			}

			receiver := modules[pulse.receiver]
			if pulse.pulse == 0 {
				lowPulseSent++
			} else {
				highPulseSent++
			}
			if receiver != nil {
				receiver.receive(pulse.caller, pulse.pulse)
			}
		}
		cycleCount++
	}

	fmt.Println(mathx.Lcm(maps.Values(inputs)))

}
