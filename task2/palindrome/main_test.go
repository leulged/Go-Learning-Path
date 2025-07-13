package main

import (
    "testing"
    "strings"
)

func TestCleanInput(t *testing.T) {
	input := "A man, a plan, a canal: Panama!"
	expected := "amanaplanacanalpanama"

	result := cleanInput(strings.ToLower(input))
	if result != expected {
		t.Errorf("cleanInput() = %q; want %q", result, expected)
	}
}

func TestPalindrome(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a", true},
		{"abba", true},
		{"abcba", true},
		{"abccba", true},
		{"abc", false},
		{"palindrome", false},
	}

	for _, tt := range tests {
		got := palindrome(tt.input)
		if got != tt.want {
			t.Errorf("palindrome(%q) = %v; want %v", tt.input, got, tt.want)
		}
	}
}
