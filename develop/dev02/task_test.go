package main

import (
	"testing"
)

type addTest struct {
	arg1, expected string
}

var addTests = []addTest{
	addTest{"", ""},
	addTest{"a4bc2d5e", "aaaabccddddde"},
	addTest{"45", ""},
	addTest{"abcd", "abcd"},
	addTest{"qwe\\4\\5", "qwe45"},
	addTest{"qwe\\45", "qwe44444"},
	addTest{"qwe\\\\5", "qwe\\\\\\\\\\"},
}

func TestUnpacking(t *testing.T) {
	for _, test := range addTests {
		if output, err := Unpacking(test.arg1); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
			t.Errorf("Error: %s", err)
		}
	}
}
