package main

import (
	"bytes"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	t.Run("Correct", func(t *testing.T) {
		test := []string{"листок", "пятак", "пятка", "столик", "слиток", "тяпка"}
		res := map[string][]string{"листок": {"слиток", "столик"}, "пятак": {"пятка", "тяпка"}}
		out := *(FindAnagrams(&test))

		if len(res) != len(out) {
			t.Error("Test fail")
		}
		for key, _ := range res {
			if len(res[key]) != len(out[key]) {
				t.Error("Test fail")
				break
			}
			for i, _ := range res[key] {
				check := bytes.Equal([]byte(res[key][i]), []byte(out[key][i]))
				if !check {
					t.Error("Test fail")
					break
				}
			}
		}
	})

	t.Run("Incorrect", func(t *testing.T) {
		test := []string{"корова", "чайник", "компьютер", "собака"}
		res := make(map[string][]string)
		out := *(FindAnagrams(&test))

		if len(res) != len(out) {
			t.Error("Test fail")
		}
	})

}
