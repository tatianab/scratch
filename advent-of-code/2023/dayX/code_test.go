package dayX

import "testing"

var input = []string{
	"",
}

func TestPart1(t *testing.T) {
	want := 0
	got, err := Part1(input)
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Errorf("Part1(%q) = %d, want %d", input, got, want)
	}
}

func TestPart2(t *testing.T) {
	want := 0
	got, err := Part2(input)
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Errorf("Part2(%q) = %d, want %d", input, got, want)
	}
}
