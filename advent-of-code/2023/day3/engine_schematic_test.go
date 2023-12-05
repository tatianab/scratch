package day3

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSumPartNumbers(t *testing.T) {
	lines := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}
	want := 4361
	got, err := SumPartNumbers(lines)
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Errorf("SumPartNumbers(%q) = %d, want %d", lines, got, want)
	}
}

func TestSumGearRatios(t *testing.T) {
	lines := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}
	want := 467835
	got, err := SumGearRatios(lines)
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Errorf("SumGearRatios(%q) = %d, want %d", lines, got, want)
	}
}

func TestParseSchematic(t *testing.T) {
	lines := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}
	want := &schematic{
		numbers: map[pos]*number{
			{0, 0}: {value: 467, pos: pos{line: 0, col: 0}, len: 3},
			{0, 5}: {value: 114, pos: pos{line: 0, col: 5}, len: 3},
			{2, 2}: {value: 35, pos: pos{line: 2, col: 2}, len: 2},
			{2, 6}: {value: 633, pos: pos{line: 2, col: 6}, len: 3},
			{4, 0}: {value: 617, pos: pos{line: 4, col: 0}, len: 3},
			{5, 7}: {value: 58, pos: pos{line: 5, col: 7}, len: 2},
			{6, 2}: {value: 592, pos: pos{line: 6, col: 2}, len: 3},
			{7, 6}: {value: 755, pos: pos{line: 7, col: 6}, len: 3},
			{9, 1}: {value: 664, pos: pos{line: 9, col: 1}, len: 3},
			{9, 5}: {value: 598, pos: pos{line: 9, col: 5}, len: 3},
		},
		symbols: map[pos]*symbol{
			{1, 3}: {pos{1, 3}, '*'},
			{3, 6}: {pos{3, 6}, '#'},
			{4, 3}: {pos{4, 3}, '*'},
			{5, 5}: {pos{5, 5}, '+'},
			{8, 3}: {pos{8, 3}, '$'},
			{8, 5}: {pos{8, 5}, '*'},
		},
	}
	got, err := parseSchematic(lines)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got, cmp.AllowUnexported(schematic{}, number{}, pos{}, symbol{})); diff != "" {
		t.Errorf("parseSchematic(%q) mismatch (-want +got):\n%s", lines, diff)
	}
}

func TestIsSymbol(t *testing.T) {
	tc := []struct {
		symbol rune
		want   bool
	}{
		{'$', true},
		{'*', true},
		{'.', false},
		{'#', true},
		{'1', false},
	}
	for _, tt := range tc {
		got := isSymbol(tt.symbol)
		if got != tt.want {
			t.Errorf("isSymbol(%q) = %t, want %t", tt.symbol, got, tt.want)
		}
	}
}
