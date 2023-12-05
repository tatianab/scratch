package main

import (
	"fmt"
	"log"

	"github.com/tatianab/scratch/advent-of-code/2023/common"
	"github.com/tatianab/scratch/advent-of-code/2023/dayX"
)

func main() {
	lines, err := common.ReadFileLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Day X")

	part1, err := dayX.Part1(lines)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Part 1: ", part1)

	part2, err := dayX.Part2(lines)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Part 2: ", part2)
}
