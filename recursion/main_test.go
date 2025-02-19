package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []int
		expected int
	}{
		{
			name:     "empty slice",
			numbers:  []int{},
			expected: 0,
		},
		{
			name:     "single number",
			numbers:  []int{5},
			expected: 5,
		},
		{
			name:     "multiple positive numbers",
			numbers:  []int{1, 2, 3, 4, 5},
			expected: 15,
		},
		{
			name:     "negative numbers",
			numbers:  []int{-1, -2, -3},
			expected: -6,
		},
		{
			name:     "mixed positive and negative numbers",
			numbers:  []int{-1, 2, -3, 4, -5},
			expected: -3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := sum(tt.numbers)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []int
		expected int
	}{
		{
			name:     "single number",
			numbers:  []int{5},
			expected: 5,
		},
		{
			name:     "multiple positive numbers",
			numbers:  []int{1, 2, 3, 4, 5},
			expected: 5,
		},
		{
			name:     "negative numbers",
			numbers:  []int{-1, -2, -3},
			expected: -1,
		},
		{
			name:     "mixed positive and negative numbers",
			numbers:  []int{-1, 2, -3, 4, -5},
			expected: 4,
		},
		{
			name:     "duplicate maximum",
			numbers:  []int{1, 5, 3, 5, 2},
			expected: 5,
		},
		{
			name:     "all same numbers",
			numbers:  []int{4, 4, 4, 4},
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maximum(tt.numbers)
			assert.Equal(t, tt.expected, result)
		})
	}
}
