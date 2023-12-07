package main

import (
	"fmt"
	"log"

	"github.com/tatianab/scratch/advent-of-code/2023/common"
	"github.com/tatianab/scratch/advent-of-code/2023/day7"
)

func main() {
	lines, err := common.ReadFileLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Day 7")

	part1, err := day7.Part1(lines)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Part 1: ", part1)

	part2, err := day7.Part2(lines)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Part 2: ", part2)
}
