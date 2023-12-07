package day6

import (
	"errors"
	"log"
	"strconv"
	"strings"
)

func Part1(lines []string) (int, error) {
	races, err := parseRaces(lines)
	if err != nil {
		return 0, err
	}
	total := 1
	for _, race := range races {
		w := race.waysToWin()
		log.Printf("ways to win for race %v: %d\n", race, w)
		total *= w
	}
	return total, nil
}

func parseRaces(lines []string) (races, error) {
	var rs races
	if len(lines) != 2 {
		return nil, errors.New("expected 2 lines")
	}
	times := strings.Fields(lines[0])[1:]
	distances := strings.Fields(lines[1])[1:]
	for i, t := range times {
		time, err := strconv.Atoi(t)
		if err != nil {
			return nil, err
		}
		dist, err := strconv.Atoi(distances[i])
		if err != nil {
			return nil, err
		}
		rs = append(rs, &race{time: time, distance: dist})
	}
	return rs, nil
}

func parseRace(lines []string) (*race, error) {
	if len(lines) != 2 {
		return nil, errors.New("expected 2 lines")
	}
	times := strings.Fields(lines[0])[1:]
	distances := strings.Fields(lines[1])[1:]
	var time, distance string
	for i, t := range times {
		time += t
		distance += distances[i]
	}
	timeInt, err := strconv.Atoi(time)
	if err != nil {
		return nil, err
	}
	distInt, err := strconv.Atoi(distance)
	if err != nil {
		return nil, err
	}
	return &race{time: timeInt, distance: distInt}, nil
}

type races []*race

type race struct {
	time     int // time to complete race, in ms
	distance int // record distance, in mm
}

func (r *race) waysToWin() int {
	ways := 0
	for t := 1; t < r.time; t++ {
		if t*(r.time-t) > r.distance {
			ways++
		}
	}
	return ways
}

func Part2(lines []string) (int, error) {
	r, err := parseRace(lines)
	if err != nil {
		return 0, err
	}
	return r.waysToWin(), nil
}
