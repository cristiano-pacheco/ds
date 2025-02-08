package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelectionSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Basic array with positive numbers",
			input:    []int{64, 34, 25, 12, 22, 11, 90},
			expected: []int{11, 12, 22, 25, 34, 64, 90},
		},
		{
			name:     "Array with negative numbers",
			input:    []int{-5, 3, -10, 1, 0},
			expected: []int{-10, -5, 0, 1, 3},
		},
		{
			name:     "Array with duplicate numbers",
			input:    []int{5, 5, 5, 5, 5},
			expected: []int{5, 5, 5, 5, 5},
		},
		{
			name:     "Single element array",
			input:    []int{42},
			expected: []int{42},
		},
		{
			name:     "Empty array",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Array already sorted",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Array in reverse order",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Array with all negative numbers",
			input:    []int{-1, -5, -2, -8, -3},
			expected: []int{-8, -5, -3, -2, -1},
		},
		{
			name:     "Array with very large numbers",
			input:    []int{1000000, 999999, 2000000, 1500000},
			expected: []int{999999, 1000000, 1500000, 2000000},
		},
		{
			name:     "Two element array",
			input:    []int{2, 1},
			expected: []int{1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a copy of input to ensure we're not testing against a modified slice
			inputCopy := make([]int, len(tt.input))
			copy(inputCopy, tt.input)

			selectionSort(inputCopy)
			assert.Equal(t, tt.expected, inputCopy, "selectionSort() returned incorrect result")
		})
	}
}
