package main

import "testing"

func TestCheck(t *testing.T) {
	tests := []struct {
		score int
		want  bool
	}{
		{-1, false},
		{0, true},
		{50, true},
		{100, true},
		{101, false},
	}

	for _, tt := range tests {
		got := check(tt.score)
		if got != tt.want {
			t.Errorf("check(%d) = %v; want %v", tt.score, got, tt.want)
		}
	}
}

func TestCalculateAverage(t *testing.T) {
	tests := []struct {
		grades map[string]int
		want   float64
	}{
		{map[string]int{}, 0},
		{map[string]int{"Math": 100}, 100},
		{map[string]int{"Math": 100, "English": 80}, 90},
		{map[string]int{"Math": 50, "English": 70, "Science": 90}, 70},
	}

	for _, tt := range tests {
		got := calculateAverage(tt.grades)
		if got != tt.want {
			t.Errorf("calculateAverage(%v) = %v; want %v", tt.grades, got, tt.want)
		}
	}
}
