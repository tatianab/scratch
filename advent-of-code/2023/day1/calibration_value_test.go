package day1

import (
	"os"
	"testing"
)

func TestSumCalibrationValue(t *testing.T) {
	tmp := t.TempDir()
	f, err := os.CreateTemp(tmp, "input.txt")
	if err != nil {
		t.Fatal(err)
	}
	f.WriteString(`1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
`)
	filename := f.Name()
	f.Close()

	want := 142
	got, err := SumCalibrationValue(filename)
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Errorf("SumCalibrationValue(%q) = %d, want %d", filename, got, want)
	}
}

func Test_sumCalibrationValue(t *testing.T) {
	lines := []string{"1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet"}
	want := 142
	got, err := sumCalibrationValue(lines)
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Errorf("sumCalibrationValue(%q) = %d, want %d", lines, got, want)
	}
}

func TestCalibrationValue(t *testing.T) {
	tcs := []struct {
		input string
		want  int
	}{
		{"abc123def", 13},
		{"1abc2", 12},
		{"pqr3stu8vwx", 38},
		{"a1b2c3d4e5f", 15},
		{"treb7uchet", 77},
	}
	for _, tc := range tcs {
		t.Run(tc.input, func(t *testing.T) {
			got, err := calibrationValue(tc.input)
			if err != nil {
				t.Fatal(err)
			}
			if got != tc.want {
				t.Errorf("calibrationValue(%q) = %d, want %d", tc.input, got, tc.want)
			}
		})
	}
}
