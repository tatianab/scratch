package main

import (
	"fmt"
	"log"

	"github.com/tatianab/scratch/advent-of-code/2023/day1"
)

func main() {
	sum1, err := day1.SumCalibrationValue("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Part 1: ", sum1)

	sum2, err := day1.SumCalibrationValue2("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Part 2: ", sum2)
}
