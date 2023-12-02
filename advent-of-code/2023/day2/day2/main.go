package main

import (
	"fmt"
	"log"

	"github.com/tatianab/scratch/advent-of-code/2023/common"
	"github.com/tatianab/scratch/advent-of-code/2023/day2"
)

func main() {
	lines, err := common.ReadFileLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	part1, err := day2.SumPossibleGames(lines)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Part 1: ", part1)
	part2, err := day2.SumPower(lines)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Part 2: ", part2)
}
