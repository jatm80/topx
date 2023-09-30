package main

import (
	"bufio"
	"os"
	"testing"
)

func TestInSlice(t *testing.T) {
	var someSlice = []int{1, 2, 3, 4}
	var line = map[int]bool{
		1: true,
		2: true,
		3: true,
		4: true,
		5: false,
		6: false,
		7: false,
	}

	for k, v := range line {
		got := inSlice(someSlice, k)
		AssertBool(t, got, v)

	}

}

func TestOpenFile(t *testing.T) {

	var got string

	data := []byte("some_string")
	os.WriteFile("unit_test", data, 0666)

	file := openFile("unit_test")
	fileScan := bufio.NewScanner(file)

	for fileScan.Scan() {
		got = fileScan.Text()
	}
	file.Close()
	os.Remove("unit_test")

	AssertStrings(t, got, "some_string")

}

func AssertStrings(t testing.TB, got string, want string) {
	t.Helper()

	if got != want {
		t.Errorf("Got %s but you were expecting %s", got, want)
	}
}

func AssertBool(t testing.TB, got bool, want bool) {
	t.Helper()

	if got != want {
		t.Errorf("Got %v but you were expecting %v", got, want)
	}
}
