package day1

import (
	"os"
	"testing"
)

func TestSumCalibrationValue2(t *testing.T) {
	tmp := t.TempDir()
	f, err := os.CreateTemp(tmp, "input.txt")
	if err != nil {
		t.Fatal(err)
	}
	f.WriteString(`two1nine
	eightwothree
	abcone2threexyz
	xtwone3four
	4nineeightseven2
	zoneight234
	7pqrstsixteen
`)
	filename := f.Name()
	f.Close()

	want := 281
	got, err := SumCalibrationValue2(filename)
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Errorf("SumCalibrationValue2(%q) = %d, want %d", filename, got, want)
	}
}

func Test_sumCalibrationValue2(t *testing.T) {
	lines := []string{"1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet"}
	want := 142
	got, err := sumCalibrationValue2(lines)
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Errorf("sumCalibrationValue2(%q) = %d, want %d", lines, got, want)
	}
}

func TestCalibrationValue2(t *testing.T) {
	tcs := []struct {
		input string
		want  int
	}{
		{"two1nine", 29},
		{"eightwothree", 83},
		{"abcone2threexyz", 13},
		{"xtwone3four", 24},
		{"4nineeightseven2", 42},
		{"zoneight234", 14},
		{"7pqrstsixteen", 76},
	}
	for _, tc := range tcs {
		t.Run(tc.input, func(t *testing.T) {
			got, err := calibrationValue2(tc.input)
			if err != nil {
				t.Fatal(err)
			}
			if got != tc.want {
				t.Errorf("calibrationValue2(%q) = %d, want %d", tc.input, got, tc.want)
			}
		})
	}
}
