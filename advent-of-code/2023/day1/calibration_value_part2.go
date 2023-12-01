package day1

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/tatianab/scratch/advent-of-code/2023/common"
)

func SumCalibrationValue2(filename string) (int, error) {
	lines, err := common.ReadFileLines(filename)
	if err != nil {
		return 0, err
	}
	return sumCalibrationValue2(lines)
}

func sumCalibrationValue2(lines []string) (int, error) {
	sum := 0
	for i, line := range lines {
		v, err := calibrationValue2(line)
		if err != nil {
			return 0, fmt.Errorf("line %d (%q): %w", i, line, err)
		}
		sum += v
	}
	return sum, nil
}

func calibrationValue2(line string) (int, error) {
	fd, err := firstDigitOrNumber(line)
	if err != nil {
		return 0, err
	}

	ld, err := lastDigitOrNumber(line)
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(fd + ld)
}

func firstDigitOrNumber(line string) (string, error) {
	for i, r := range line {
		if unicode.IsDigit(r) {
			return string(r), nil
		}
		if d, ok := englishDigit(line[i:]); ok {
			return d, nil
		}
	}
	return "", errors.New("no digit found")
}

func lastDigitOrNumber(line string) (string, error) {
	// Assuming ASCII encoding.
	for i := len(line) - 1; i >= 0; i-- {
		r := rune(line[i])
		if unicode.IsDigit(r) {
			return string(r), nil
		}
		if d, ok := englishDigit(line[i:]); ok {
			return d, nil
		}
	}
	return "", errors.New("no digit found")
}

var digits = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func englishDigit(s string) (string, bool) {
	for ed, d := range digits {
		if strings.HasPrefix(s, ed) {
			return d, true
		}
	}
	return "", false
}
