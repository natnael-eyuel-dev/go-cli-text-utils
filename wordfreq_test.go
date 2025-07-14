package main

import (
	"testing"
)

func resetWordStore() {
	storeWords = WordsFreq{wordsFreq: make(map[string]int)}
}

func TestWordFreqCount(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected map[string]int
	}{
		{
			name:  "simple words",
			input: "hello world hello",
			expected: map[string]int{
				"hello": 2,
				"world": 1,
			},
		},
		{
			name:  "with punctuation",
			input: "Hello, world! Hello?",
			expected: map[string]int{
				"hello": 2,
				"world": 1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resetWordStore()
			result := wordFreqCount(tt.input)
			
			if len(result) != len(tt.expected) {
				t.Errorf("Expected %d words, got %d", len(tt.expected), len(result))
			}

			for word, count := range tt.expected {
				if result[word] != count {
					t.Errorf("For word '%s', expected count %d, got %d", word, count, result[word])
				}
			}
		})
	}
}