package codenamegenerator

import (
	"testing"
)

func TestGetSingleWord(t *testing.T) {
	var tests = []struct {
		a, e string
	}{
		{"toronto", "toronto"},
		{"one-toronto", "toronto"},
		{"two toronto", "toronto"},
		{"one two-toronto", "toronto"},
	}

	for _, tt := range tests {
		testname := tt.a
		t.Run(testname, func(t *testing.T) {
			ans := getSingleWord(tt.a)
			if ans != tt.e {
				t.Errorf("got %s, wanted %s", ans, tt.e)
			}
		})
	}
}
