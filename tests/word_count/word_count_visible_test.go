package main

import (
	"reflect"
	"testing"
)

func TestLevel6_WordCount_Visible(t *testing.T) {
	tests := []struct {
		name     string
		words    []string
		expected map[string]int
	}{
		{"example_1_repeated_words", []string{"a", "b", "a"}, map[string]int{"a": 2, "b": 1}},
		{"example_2_case_sensitive", []string{"go", "Go", "go"}, map[string]int{"go": 2, "Go": 1}},
		{"example_3_empty_slice", []string{}, map[string]int{}},
		{"nil_slice", nil, map[string]int{}},
		{"single_word", []string{"x"}, map[string]int{"x": 1}},
		{"all_same_word", []string{"hi", "hi", "hi", "hi"}, map[string]int{"hi": 4}},
		{"all_unique", []string{"one", "two", "three"}, map[string]int{"one": 1, "two": 1, "three": 1}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := WordCount(tc.words)
			if got == nil {
				t.Fatalf("WordCount(%v) returned a nil map; want a non-nil map", tc.words)
			}
			if !reflect.DeepEqual(got, tc.expected) {
				t.Fatalf("WordCount(%v) = %v; want %v", tc.words, got, tc.expected)
			}
		})
	}
}
