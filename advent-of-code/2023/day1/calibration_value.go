package day1

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"

	"github.com/tatianab/scratch/advent-of-code/2023/common"
)

func SumCalibrationValue(filename string) (int, error) {
	lines, err := common.ReadFileLines(filename)
	if err != nil {
		return 0, err
	}
	return sumCalibrationValue(lines)
}

func sumCalibrationValue(lines []string) (int, error) {
	sum := 0
	for i, line := range lines {
		v, err := calibrationValue(line)
		if err != nil {
			return 0, fmt.Errorf("line %d: %w", i, err)
		}
		sum += v
	}
	return sum, nil
}

func calibrationValue(line string) (int, error) {
	fd, err := firstDigit(line)
	if err != nil {
		return 0, err
	}

	ld, err := lastDigit(line)
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(fd + ld)
}

func firstDigit(line string) (string, error) {
	for _, r := range line {
		if unicode.IsDigit(r) {
			return string(r), nil
		}
	}
	return "", errors.New("no digit found")
}

func lastDigit(line string) (string, error) {
	// Assuming ASCII encoding.
	for i := len(line) - 1; i >= 0; i-- {
		r := rune(line[i])
		if unicode.IsDigit(r) {
			return string(r), nil
		}
	}
	return "", errors.New("no digit found")
}
