package day2

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func SumPossibleGames(lines []string) (int, error) {
	sum := 0
	for _, line := range lines {
		g, err := parseGame(line)
		if err != nil {
			return 0, err
		}
		if g.possibleGame() {
			sum += g.index
		}
	}
	return sum, nil
}

type game struct {
	index  int
	trials []*trial
}

type trial struct {
	red   int
	blue  int
	green int
}

var gameRegex = regexp.MustCompile(`Game (\d+): (.*)`)

func parseGame(line string) (*game, error) {
	match := gameRegex.FindStringSubmatch(line)
	if len(match) != 3 {
		return nil, fmt.Errorf("invalid game line: %q", line)
	}
	index, err := strconv.Atoi(match[1])
	if err != nil {
		return nil, fmt.Errorf("game index: %w", err)
	}
	g := &game{
		index: index,
	}
	trials, err := parseTrials(match[2])
	if err != nil {
		return nil, fmt.Errorf("game %d: %w", g.index, err)
	}
	g.trials = trials
	return g, nil
}

func parseTrials(s string) (trials []*trial, err error) {
	ts := strings.Split(s, "; ")
	for _, t := range ts {
		trial, err := parseTrial(strings.Split(t, ", "))
		if err != nil {
			return nil, err
		}
		trials = append(trials, trial)
	}
	return trials, nil
}

func parseTrial(s []string) (*trial, error) {
	t := &trial{}
	for _, colorAmount := range s {
		n, color, err := parseColorAmount(colorAmount)
		if err != nil {
			return nil, err
		}
		switch color {
		case "red":
			t.red = n
		case "blue":
			t.blue = n
		case "green":
			t.green = n
		default:
			return nil, fmt.Errorf("invalid color: %q", color)
		}
	}
	return t, nil
}

func parseColorAmount(s string) (int, string, error) {
	parts := strings.Split(s, " ")
	if len(parts) != 2 {
		return 0, "", fmt.Errorf("invalid color amount: %q", s)
	}
	n, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, "", fmt.Errorf("amount: %w", err)
	}
	return n, parts[1], nil
}

// A game is possible if the bag could have contained
// only 12 red cubes, 13 green cubes, and 14 blue cubes.
func (g *game) possibleGame() bool {
	for _, t := range g.trials {
		if t.red > 12 || t.green > 13 || t.blue > 14 {
			return false
		}
	}
	return true
}
