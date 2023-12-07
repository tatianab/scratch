package day7

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func Part1(lines []string) (int, error) {
	hands, err := parseHands(lines, false)
	if err != nil {
		return 0, err
	}
	slices.SortFunc(hands, cmpHands)
	return sumHands(hands), nil
}

func Part2(lines []string) (int, error) {
	hands, err := parseHands(lines, true)
	if err != nil {
		return 0, err
	}
	slices.SortFunc(hands, cmpHandsJoker)
	return sumHands(hands), nil
}

func sumHands(hands []*hand) int {
	total := 0
	for i, h := range hands {
		rank := i + 1
		winnings := h.bid * rank
		total += winnings
		fmt.Printf("%d. %s (%s): %d * %d = %d\n", rank, h.raw, h.strength, h.bid, rank, winnings)
	}
	return total
}

type strength int

const (
	highCard strength = iota
	onePair
	twoPairs
	threeOfAKind
	fullHouse
	fourOfAKind
	fiveOfAKind
)

func (s strength) String() string {
	switch s {
	case highCard:
		return "high card"
	case onePair:
		return "one pair"
	case twoPairs:
		return "two pairs"
	case threeOfAKind:
		return "three of a kind"
	case fullHouse:
		return "full house"
	case fourOfAKind:
		return "four of a kind"
	case fiveOfAKind:
		return "five of a kind"
	}
	return "unknown"
}

type card int

const (
	two card = iota + 2
	three
	four
	five
	six
	seven
	eight
	nine
	ten
	jack
	queen
	king
	ace
)

var cardMap = map[rune]card{
	'2': two,
	'3': three,
	'4': four,
	'5': five,
	'6': six,
	'7': seven,
	'8': eight,
	'9': nine,
	'T': ten,
	'J': jack,
	'Q': queen,
	'K': king,
	'A': ace,
}

type hand struct {
	raw      string
	cards    []card
	strength strength
	bid      int
}

func parseHands(lines []string, withJoker bool) ([]*hand, error) {
	var hands []*hand
	for _, line := range lines {
		h, err := parseHand(line, withJoker)
		if err != nil {
			return nil, err
		}
		hands = append(hands, h)
	}
	return hands, nil
}

func cmpHands(a, b *hand) int {
	if a.strength != b.strength {
		return int(a.strength) - int(b.strength)
	}
	for i, c := range a.cards {
		if c != b.cards[i] {
			return int(c) - int(b.cards[i])
		}
	}
	return 0
}

func cmpHandsJoker(a, b *hand) int {
	if a.strength != b.strength {
		return int(a.strength) - int(b.strength)
	}
	for i, c := range a.cards {
		if c != b.cards[i] {
			// jack is always lowest
			if c == jack {
				return -1
			} else if b.cards[i] == jack {
				return 1
			}
			return int(c) - int(b.cards[i])
		}
	}
	return 0
}

func parseHand(line string, withJoker bool) (*hand, error) {
	parts := strings.Fields(line)
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid input: %q", line)
	}
	cards, err := parseCards(parts[0])
	if err != nil {
		return nil, err
	}
	bid, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, fmt.Errorf("invalid input: %q", line)
	}
	strength, err := calcStrength(cards, withJoker)
	if err != nil {
		return nil, err
	}
	return &hand{
		raw:      line,
		cards:    cards,
		strength: strength,
		bid:      bid,
	}, nil
}

func parseCards(s string) ([]card, error) {
	var cards []card
	for _, r := range s {
		c, ok := cardMap[r]
		if !ok {
			return nil, fmt.Errorf("invalid card: %q", r)
		}
		cards = append(cards, c)
	}
	return cards, nil
}

func calcStrength(cards []card, withJoker bool) (strength, error) {
	if withJoker {
		return calcStrengthWithJoker(cards)
	}

	groups := make(map[card]int)
	for _, c := range cards {
		groups[c]++
	}
	switch len(groups) {
	case 5:
		return highCard, nil
	case 4:
		return onePair, nil
	case 3:
		for _, n := range groups {
			if n == 3 {
				return threeOfAKind, nil
			}
		}
		return twoPairs, nil
	case 2:
		for _, n := range groups {
			if n == 4 {
				return fourOfAKind, nil
			}
		}
		return fullHouse, nil
	case 1:
		return fiveOfAKind, nil
	}
	return highCard, fmt.Errorf("invalid hand: %v", cards)
}

func calcStrengthWithJoker(cards []card) (strength, error) {
	groups := make(map[card]int)
	for _, c := range cards {
		groups[c]++
	}
	switch len(groups) {
	case 5:
		if _, ok := groups[jack]; ok {
			return onePair, nil
		}
		return highCard, nil
	case 4:
		if _, ok := groups[jack]; ok {
			return threeOfAKind, nil
		}
		return onePair, nil
	case 3:
		for _, n := range groups {
			if n == 3 {
				if _, ok := groups[jack]; ok {
					return fourOfAKind, nil
				}
				return threeOfAKind, nil
			}
		}
		if n, ok := groups[jack]; ok {
			if n == 1 {
				return fullHouse, nil
			}
			if n == 2 {
				return fourOfAKind, nil
			}
		}
		return twoPairs, nil
	case 2:
		for _, n := range groups {
			if n == 4 {
				if _, ok := groups[jack]; ok {
					return fiveOfAKind, nil
				}
				return fourOfAKind, nil
			}
		}
		if _, ok := groups[jack]; ok {
			return fiveOfAKind, nil
		}
		return fullHouse, nil
	case 1:
		return fiveOfAKind, nil
	}
	return highCard, fmt.Errorf("invalid hand: %v", cards)
}
