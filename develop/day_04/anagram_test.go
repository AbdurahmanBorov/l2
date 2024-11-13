package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestGetAnagram(t *testing.T) {
	tests := []struct {
		name     string
		words    []string
		expected map[string][]string
	}{
		{
			name:  "simple test",
			words: []string{"топор", "порот", "ротоп", "топор", "шар", "раш"},
			expected: map[string][]string{
				"топор": {"порот", "ротоп", "топор"},
				"шар":   {"раш", "шар"},
			},
		},
		{
			name:     "no anagrams",
			words:    []string{"книга", "яблоко", "машина", "дерево"},
			expected: map[string][]string{},
		},
		{
			name:  "multiple anagrams with same word order",
			words: []string{"листок", "слиток", "столик", "слиток"},
			expected: map[string][]string{
				"листок": {"листок", "слиток", "столик"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getAnagram(tt.words)
			for _, v := range result {
				sort.Strings(v)
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("For input %v, expected %v, but got %v", tt.words, tt.expected, result)
			}
		})
	}
}
