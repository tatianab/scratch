package day6

import "testing"

var input = []string{
	"Time:      7  15   30",
	"Distance:  9  40  200",
}

func TestPart1(t *testing.T) {
	want := 288
	got, err := Part1(input)
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Errorf("Part1(%q) = %d, want %d", input, got, want)
	}
}

func TestPart2(t *testing.T) {
	want := 71503
	got, err := Part2(input)
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Errorf("Part2(%q) = %d, want %d", input, got, want)
	}
}
