package main

import (
	"testing"
)

func TestReverse(t *testing.T) {
	testcases := []struct {
		input, expected string
	}{
		{"Hello world", "dlrow olleH"},
		{" ", " "},
		{"12345", "54321"},
	}
	for _, tc := range testcases {
		rev, err := Reverse(tc.input)
		if err != nil {
			t.Errorf("Error: Reverse(%q) == %q, expected %q", tc.input, rev, tc.expected)
		}
		if rev != tc.expected {
			t.Errorf("Reverse(%q) == %q, expected %q", tc.input, rev, tc.expected)
		}
	}
}
