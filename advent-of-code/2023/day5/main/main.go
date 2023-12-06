package main

import (
	"fmt"
	"log"

	"github.com/tatianab/scratch/advent-of-code/2023/common"
	"github.com/tatianab/scratch/advent-of-code/2023/day5"
)

func main() {
	lines, err := common.ReadFileLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Day 5")

	part1, err := day5.Part1(lines)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Part 1: ", part1)

	part2, err := day5.Part2(lines)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Part 2: ", part2)
}
