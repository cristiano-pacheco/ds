package main

import (
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestFindSmallest(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected *int
	}{
		{
			name:     "Basic array with positive numbers",
			input:    []int{64, 34, 25, 12, 22, 11, 90},
			expected: lo.ToPtr(5), // index of 11
		},
		{
			name:     "Array with negative numbers",
			input:    []int{-5, 3, -10, 1, 0},
			expected: lo.ToPtr(2), // index of -10
		},
		{
			name:     "Array with duplicate numbers",
			input:    []int{5, 5, 5, 5, 5},
			expected: lo.ToPtr(0), // first occurrence
		},
		{
			name:     "Single element array",
			input:    []int{42},
			expected: lo.ToPtr(0),
		},
		{
			name:     "Empty array",
			input:    []int{},
			expected: nil,
		},
		{
			name:     "Array with smallest at the end",
			input:    []int{100, 200, 300, 400, 1},
			expected: lo.ToPtr(4),
		},
		{
			name:     "Array with smallest at the beginning",
			input:    []int{1, 200, 300, 400, 500},
			expected: lo.ToPtr(0),
		},
		{
			name:     "Array with duplicate smallest values",
			input:    []int{5, 2, 8, 2, 1, 1, 9},
			expected: lo.ToPtr(4), // should return first occurrence of 1
		},
		{
			name:     "Array with all negative numbers",
			input:    []int{-1, -5, -2, -8, -3},
			expected: lo.ToPtr(3), // index of -8
		},
		{
			name:     "Array with very large numbers",
			input:    []int{1000000, 999999, 2000000, 1500000},
			expected: lo.ToPtr(1), // index of 999999
		},
		{
			name:     "Array with very small numbers",
			input:    []int{-999999, -1000000, -2000000, -1500000},
			expected: lo.ToPtr(2), // index of -2000000
		},
		{
			name:     "Two element array",
			input:    []int{2, 1},
			expected: lo.ToPtr(1),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findSmallest(tt.input)
			if tt.expected == nil {
				assert.Nil(t, result, "Expected nil for empty array")
			} else {
				assert.NotNil(t, result, "Expected non-nil result")
				assert.Equal(t, *tt.expected, *result, "findSmallest() returned incorrect index")
			}
		})
	}
}
