package day3

import (
	"cmp"
	"fmt"
	"slices"
	"strconv"
	"unicode"
)

func SumPartNumbers(lines []string) (int, error) {
	es, err := parseSchematic(lines)
	if err != nil {
		return 0, fmt.Errorf("error parsing schematic: %v", err)
	}
	sum := 0
	for _, n := range es.numbers {
		if es.isPartNumber(n) {
			sum += n.value
		}
	}
	return sum, nil
}

func SumGearRatios(lines []string) (int, error) {
	es, err := parseSchematic(lines)
	if err != nil {
		return 0, fmt.Errorf("error parsing schematic: %v", err)
	}
	for _, n := range es.numbers {
		_ = es.isPartNumber(n) // for side effects
	}
	sum := 0
	for _, numbers := range es.gearsToNumbers {
		// dedupe numbers
		slices.SortFunc(numbers, func(n1, n2 *number) int {
			lc := cmp.Compare(n1.line, n2.line)
			if lc == 0 {
				return cmp.Compare(n1.col, n2.col)
			}
			return lc
		})
		numbers = slices.Compact(numbers)
		if len(numbers) == 2 {
			sum += numbers[0].value * numbers[1].value
		}
	}
	return sum, nil
}

type schematic struct {
	numbers     map[pos]*number
	symbols     map[pos]*symbol
	lines, cols int

	// for part 2 - map of '*'s (possible gears)
	// to adjacent numbers
	gearsToNumbers map[pos][]*number
}

type symbol struct {
	pos
	value rune
}

func (s *schematic) isPartNumber(n *number) bool {
	for i := n.col; i < n.col+n.len; i++ {
		if s.adjacentSymbols(n, i) {
			return true
		}
	}
	return false
}

func (s *schematic) adjacentSymbols(n *number, col int) bool {
	line := n.line
	above, below := line-1, line+1
	left, right := col-1, col+1
	lastline, lastcol := s.lines-1, s.cols-1
	found := false
	checkPos := func(line, col int) {
		p := pos{line: line, col: col}
		if sym, ok := s.symbols[p]; ok {
			found = true
			// add gear
			if sym.value == '*' {
				s.gearsToNumbers[sym.pos] = append(s.gearsToNumbers[sym.pos], n)
			}
		}
	}
	if line != 0 {
		// check above
		checkPos(above, col)
		if col != 0 {
			// check above left
			checkPos(above, left)
		}
		if col != lastcol {
			// check above right
			checkPos(above, right)
		}
	}
	if line != lastline {
		// check below
		checkPos(below, col)
		if col != 0 {
			// check below left
			checkPos(below, left)
		}
		if col != lastcol {
			// check below right
			checkPos(below, right)
		}
	}
	if col != 0 {
		// check left
		checkPos(line, left)
	}
	if col != lastcol {
		// check right
		checkPos(line, right)
	}
	return found
}

type number struct {
	pos
	value int
	len   int // length of the number
}

type pos struct {
	line, col int
}

func parseSchematic(lines []string) (*schematic, error) {
	s := &schematic{
		numbers:        make(map[pos]*number),
		symbols:        make(map[pos]*symbol),
		gearsToNumbers: make(map[pos][]*number),
	}
	for i, line := range lines {
		vStr := ""
		for j, r := range line {
			if isSymbol(r) {
				p := pos{line: i, col: j}
				s.symbols[p] = &symbol{
					pos:   p,
					value: r,
				}
			}
			if unicode.IsDigit(r) {
				vStr += string(r)
			}
			if (!unicode.IsDigit(r) || j == len(line)-1) && vStr != "" {
				v, err := strconv.Atoi(vStr)
				if err != nil {
					return nil, fmt.Errorf("error parsing number: %v", err)
				}
				pos := pos{line: i, col: j - len(vStr)}
				s.numbers[pos] = &number{
					value: v,
					pos:   pos,
					len:   len(vStr),
				}
				vStr = ""
			}
		}
	}
	return s, nil
}

func isSymbol(r rune) bool {
	return !unicode.IsDigit(r) && r != '.'
}
