package day5

import "testing"

var input = []string{
	"seeds: 79 14 55 13",
	"",
	"seed-to-soil map:",
	"50 98 2",
	"52 50 48",
	"",
	"soil-to-fertilizer map:",
	"0 15 37",
	"37 52 2",
	"39 0 15",
	"",
	"fertilizer-to-water map:",
	"49 53 8",
	"0 11 42",
	"42 0 7",
	"57 7 4",
	"",
	"water-to-light map:",
	"88 18 7",
	"18 25 70",
	"",
	"light-to-temperature map:",
	"45 77 23",
	"81 45 19",
	"68 64 13",
	"",
	"temperature-to-humidity map:",
	"0 69 1",
	"1 0 69",
	"",
	"humidity-to-location map:",
	"60 56 37",
	"56 93 4",
}

func TestPart1(t *testing.T) {
	want := 35
	got, err := Part1(input)
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Errorf("Part1(%q) = %d, want %d", input, got, want)
	}
}

func TestFindLocation(t *testing.T) {
	a, err := parseAlmanac(input)
	if err != nil {
		t.Fatal(err)
	}
	want := 43
	got, err := findLocation(a, 14)
	if err != nil {
		t.Fatal(err)
	}
	// Seed 14, soil 14, fertilizer 53, water 49, light 42, temperature 42, humidity 43, location 43.
	if got != want {
		t.Errorf("findLocation(%q) = %d, want %d", input, got, want)
	}
}
func TestPart2(t *testing.T) {
	want := 46
	got, err := Part2(input)
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Errorf("Part2(%q) = %d, want %d", input, got, want)
	}
}
