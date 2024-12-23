package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"slices"
)

func main() {
	input, err := download.ReadInput(2023, 7)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

var cardOrder = map[rune]int{
	'A': 14, 'K': 13, 'Q': 12, 'J': 11, 'T': 10,
	'9': 9, '8': 8, '7': 7, '6': 6, '5': 5,
	'4': 4, '3': 3, '2': 2,
}

var cardOrderJoker = map[rune]int{
	'A': 14, 'K': 13, 'Q': 12, 'T': 10,
	'9': 9, '8': 8, '7': 7, '6': 6, '5': 5,
	'4': 4, '3': 3, '2': 2, 'J': 1,
}

type handType int

const (
	HighCard handType = iota + 1
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

type hand struct {
	cards string
	bid   int
	hType handType
}

func handOrder(a, b hand) int {
	if a.hType == b.hType {
		for i := range 5 {
			if a.cards[i] != b.cards[i] {
				return cardOrder[rune(a.cards[i])] - cardOrder[rune(b.cards[i])]
			}
		}
		return 0
	}
	return int(a.hType - b.hType)
}

func handOrderJoker(a, b hand) int {
	if a.hType == b.hType {
		for i := range 5 {
			if a.cards[i] != b.cards[i] {
				return cardOrderJoker[rune(a.cards[i])] - cardOrderJoker[rune(b.cards[i])]
			}
		}
		return 0
	}
	return int(a.hType - b.hType)
}

func determineHandType(hand hand) handType {
	cardBag := container.NewBag[rune]()
	for _, card := range hand.cards {
		cardBag.Add(card)
	}

	cardCounts := cardBag.Values()
	switch len(cardCounts) {
	case 1:
		return FiveOfAKind
	case 2:
		for _, count := range cardCounts {
			if count == 4 {
				return FourOfAKind
			}
		}
		return FullHouse
	case 3:
		for _, count := range cardCounts {
			if count == 3 {
				return ThreeOfAKind
			}
		}
		return TwoPair
	case 4:
		return OnePair
	}
	return HighCard
}

func determineHandTypeJoker(hand hand) handType {
	cardBag := container.NewBag[rune]()
	numberOfJokers := 0
	for _, card := range hand.cards {
		if card == 'J' {
			numberOfJokers++
			continue
		}
		cardBag.Add(card)
	}

	cardCounts := cardBag.Values()
	ht := HighCard

	for card, count := range cardCounts {
		switch {
		case count == 5:
			ht = FiveOfAKind
		case count == 4:
			ht = FourOfAKind
			if numberOfJokers == 1 {
				ht = FiveOfAKind
			}
		case count == 3:
			ht = ThreeOfAKind
			if numberOfJokers == 2 {
				ht = FiveOfAKind
			} else if numberOfJokers == 1 {
				ht = FourOfAKind
			} else {
				for otherCard, otherCount := range cardCounts {
					if card == otherCard {
						continue
					}
					if otherCount == 2 {
						ht = FullHouse
						break
					}
				}
			}
		case count == 2:
			ht = OnePair
			if numberOfJokers == 3 {
				ht = FiveOfAKind
			} else if numberOfJokers == 2 {
				ht = FourOfAKind
			} else {
				for otherCard, otherCount := range cardCounts {
					if card == otherCard {
						continue
					}
					if otherCount == 3 {
						ht = FullHouse
						break
					} else if otherCount == 2 {
						ht = TwoPair
						if numberOfJokers == 1 {
							ht = FullHouse
						}
						break
					}
				}
			}

			if ht == OnePair && numberOfJokers == 1 {
				ht = ThreeOfAKind
			}
		}

		if ht != HighCard {
			break
		}
	}

	if ht == HighCard {
		switch numberOfJokers {
		case 4, 5:
			ht = FiveOfAKind
		case 3:
			ht = FourOfAKind
		case 2:
			ht = ThreeOfAKind
		case 1:
			ht = OnePair
		}
	}
	return ht
}

func part1(input string) {
	lines := conv.SplitNewline(input)

	var hands []hand
	for _, line := range lines {
		hand := hand{
			cards: line[:5],
			bid:   conv.MustAtoi(line[6:]),
		}
		hand.hType = determineHandType(hand)
		hands = append(hands, hand)
	}

	slices.SortFunc(hands, handOrder)

	totalWinnings := 0
	for i, hand := range hands {
		totalWinnings += (i + 1) * hand.bid
	}
	fmt.Println(totalWinnings)
}

func part2(input string) {
	lines := conv.SplitNewline(input)

	var hands []hand
	for _, line := range lines {
		hand := hand{
			cards: line[:5],
			bid:   conv.MustAtoi(line[6:]),
		}
		hand.hType = determineHandTypeJoker(hand)
		hands = append(hands, hand)
	}

	slices.SortFunc(hands, handOrderJoker)

	totalWinnings := 0
	for i, hand := range hands {
		totalWinnings += (i + 1) * hand.bid
	}
	fmt.Println(totalWinnings)
}
