package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestCleanInput(t *testing.T) {
	input := "Hello, World! 123"
	expected := "hello world 123"

	got := cleanInput(strings.ToLower(input))
	if got != expected {
		t.Errorf("cleanInput() = %q; want %q", got, expected)
	}
}

func TestWordFrequency(t *testing.T) {
	input := "Go is great. Go is fun! Fun, fun, fun."
	cleaned := cleanInput(strings.ToLower(input))
	words := strings.Fields(cleaned)

	wordFreq := make(map[string]int)
	for _, word := range words {
		wordFreq[word]++
	}

	expected := map[string]int{
		"go":    2,
		"is":    2,
		"great": 1,
		"fun":   4,
	}

	if !reflect.DeepEqual(wordFreq, expected) {
		t.Errorf("wordFreq = %v; want %v", wordFreq, expected)
	}
}
