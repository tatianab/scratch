package day5

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func Part1(lines []string) (int, error) {
	a, err := parseAlmanac(lines)
	if err != nil {
		return 0, err
	}
	min := math.MaxInt64
	for _, seed := range a.seeds {
		loc, err := findLocation(a, seed)
		if err != nil {
			return 0, err
		}
		if loc < min {
			min = loc
		}
	}
	return min, nil
}

func findLocation(a *almanac, seed int) (int, error) {
	val := seed
	src := "seed"
	for src != "location" {
		m, ok := a.maps[src]
		if !ok {
			return 0, fmt.Errorf("no map for %s-to-X", src)
		}
		for _, sdl := range m.mappings {
			offset := val - sdl.srcStart
			if offset >= 0 && offset < sdl.length {
				val = sdl.dstStart + offset
				break
			}
		}
		src = m.dst
	}
	return val, nil
}

type almanac struct {
	seeds []int
	maps  map[string]*srcToDst
}

type srcToDst struct {
	src, dst string
	mappings []*sdl
}

type sdl struct {
	srcStart, dstStart, length int
}

var (
	mapRe = regexp.MustCompile(`^(\w+)-to-(\w+) map:$`)
)

func parseAlmanac(lines []string) (*almanac, error) {
	if len(lines) == 0 {
		return nil, nil
	}

	seeds, err := parseSeeds(lines[0])
	if err != nil {
		return nil, err
	}

	maps := make(map[string]*srcToDst)
	var current *srcToDst
	for _, line := range lines[1:] {
		if line == "" {
			if current != nil {
				maps[current.src] = current

			}
			current = &srcToDst{}
			continue
		}
		if src, dst, ok := parseMapTitle(line); ok {
			current.src = src
			current.dst = dst
			continue
		}
		dstStart, srcStart, length, err := parseMapLine(line)
		if err != nil {
			return nil, err
		}
		current.mappings = append(current.mappings, &sdl{
			srcStart, dstStart, length,
		})
	}
	maps[current.src] = current

	return &almanac{
		seeds: seeds,
		maps:  maps,
	}, nil
}

func parseSeeds(line string) ([]int, error) {
	parts := strings.Fields(line)[1:]

	seeds := make([]int, len(parts))
	for i, part := range parts {
		seed, err := strconv.Atoi(part)
		if err != nil {
			return nil, fmt.Errorf("invalid seeds line: %q", line)
		}
		seeds[i] = seed
	}
	return seeds, nil
}

func parseMapTitle(line string) (string, string, bool) {
	matches := mapRe.FindStringSubmatch(line)
	if len(matches) != 3 {
		return "", "", false
	}
	return matches[1], matches[2], true
}

func parseMapLine(line string) (int, int, int, error) {
	parts := strings.Fields(line)
	if len(parts) != 3 {
		return 0, 0, 0, fmt.Errorf("invalid map line: %q", line)
	}
	dstStart, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, 0, fmt.Errorf("invalid map line: %q", line)
	}
	srcStart, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, 0, 0, fmt.Errorf("invalid map line: %q", line)
	}
	length, err := strconv.Atoi(parts[2])
	if err != nil {
		return 0, 0, 0, fmt.Errorf("invalid map line: %q", line)
	}
	return dstStart, srcStart, length, nil
}

func Part2(lines []string) (int, error) {
	a, err := parseAlmanac(lines)
	if err != nil {
		return 0, err
	}
	min := math.MaxInt64
	for i, seed := range a.seeds {
		if i%2 == 1 {
			continue
		}
		for s := seed; s <= seed+a.seeds[i+1]; s++ {
			loc, err := findLocation(a, s)
			if err != nil {
				return 0, err
			}
			if loc < min {
				min = loc
			}
		}
	}
	return min, nil
}
