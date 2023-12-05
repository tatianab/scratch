package day4

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func Part1(lines []string) (int, error) {
	cards, err := parseCards(lines)
	if err != nil {
		return 0, err
	}
	total := 0
	for _, c := range cards {
		score := c.score()
		if score > 0 {
			total += int(math.Pow(2, float64(score-1)))
		}
	}
	return total, nil
}

func Part2(lines []string) (int, error) {
	cards, err := parseCards(lines)
	if err != nil {
		return 0, err
	}
	scores := make([]int, len(cards))
	for i, c := range cards {
		scores[i] = c.score()
	}
	total := len(cards)
	for i := range scores {
		total += score(scores, i)
	}
	return total, nil
}

func score(scores []int, i int) int {
	if i >= len(scores) {
		return 0
	}
	sc := scores[i]
	for offset := 1; offset < scores[i]+1; offset++ {
		sc += score(scores, i+offset)
	}
	return sc
}

func parseCards(lines []string) ([]*card, error) {
	cards := []*card{}
	for _, line := range lines {
		c, err := parseCard(line)
		if err != nil {
			return nil, err
		}
		cards = append(cards, c)
	}
	return cards, nil
}

var cardRE = regexp.MustCompile(`Card +(\d+): +(\d.*) +\| +(\d.*)`)

func parseCard(line string) (*card, error) {
	parts := cardRE.FindStringSubmatch(line)
	if len(parts) != 4 {
		return nil, fmt.Errorf("invalid card: %q", line)
	}
	i, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, fmt.Errorf("invalid card index %q: %w", parts[1], err)
	}
	winning, err := parseWinning(parts[2])
	if err != nil {
		return nil, err
	}
	numbers, err := parseNumbers(parts[3])
	if err != nil {
		return nil, err
	}
	return &card{
		i:       i,
		winning: winning,
		numbers: numbers,
	}, nil
}

func parseWinning(s string) (map[int]bool, error) {
	winning := map[int]bool{}
	for _, w := range strings.Fields(s) {
		n, err := strconv.Atoi(w)
		if err != nil {
			return nil, fmt.Errorf("invalid winning number %q: %w", w, err)
		}
		winning[n] = true
	}
	return winning, nil
}

func parseNumbers(s string) ([]int, error) {
	numbers := []int{}
	for _, num := range strings.Fields(s) {
		n, err := strconv.Atoi(num)
		if err != nil {
			return nil, fmt.Errorf("invalid number %q: %w", num, err)
		}
		numbers = append(numbers, n)
	}
	return numbers, nil
}

type card struct {
	i       int // index
	winning map[int]bool
	numbers []int
}

func (c *card) score() int {
	score := 0
	for _, n := range c.numbers {
		if c.winning[n] {
			score++
		}
	}
	return score
}
