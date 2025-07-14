package main

import (
	"testing"
)

func TestPalindCheck(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "simple palindrome",
			input:    "racecar",
			expected: true,
		},
		{
			name:     "with punctuation",
			input:    "racecar.;:",
			expected: true,
		},
		{
			name:     "not palindrome",
			input:    "world",
			expected: false,
		},
		{
			name:     "case insensitive",
			input:    "RaceCar",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := palindCheck(tt.input)
			if result != tt.expected {
				t.Errorf("For input '%s', expected %t, got %t", tt.input, tt.expected, result)
			}
		})
	}
}

