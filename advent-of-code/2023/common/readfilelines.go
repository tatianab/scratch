package common

import (
	"os"
	"strings"
)

func ReadFileLines(filename string) ([]string, error) {
	b, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(b), "\n")
	return lines[:len(lines)-1], nil // Remove the last empty line.
}
