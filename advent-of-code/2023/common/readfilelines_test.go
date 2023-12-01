package common

import (
	"os"
	"reflect"
	"testing"
)

func TestReadFileLines(t *testing.T) {
	tmp := t.TempDir()
	f, err := os.CreateTemp(tmp, "input.txt")
	if err != nil {
		t.Fatal(err)
	}
	f.WriteString(`1abc2
pqr3stu8vwx
`)
	filename := f.Name()
	f.Close()

	want := []string{"1abc2", "pqr3stu8vwx"}
	got, err := ReadFileLines(filename)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ReadFileLines(%q) = %v, want %v", filename, got, want)
	}
}
