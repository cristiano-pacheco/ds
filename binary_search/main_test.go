package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name           string
		array          []int
		target         int
		expectedResult int
	}{
		// Basic cases
		// {
		// 	name:           "Should find element in the middle of array",
		// 	array:          []int{1, 2, 3, 4, 5},
		// 	target:         3,
		// 	expectedResult: 2,
		// },
		{
			name:           "Should find first element",
			array:          []int{1, 2, 3, 4, 5},
			target:         1,
			expectedResult: 0,
		},
		{
			name:           "Should find last element",
			array:          []int{1, 2, 3, 4, 5},
			target:         5,
			expectedResult: 4,
		},
		// Edge cases
		{
			name:           "Should return -1 for element smaller than all",
			array:          []int{1, 2, 3, 4, 5},
			target:         0,
			expectedResult: -1,
		},
		{
			name:           "Should return -1 for element greater than all",
			array:          []int{1, 2, 3, 4, 5},
			target:         6,
			expectedResult: -1,
		},
		{
			name:           "Should work with empty array",
			array:          []int{},
			target:         1,
			expectedResult: -1,
		},
		{
			name:           "Should work with single element array when found",
			array:          []int{1},
			target:         1,
			expectedResult: 0,
		},
		{
			name:           "Should work with single element array when not found",
			array:          []int{1},
			target:         2,
			expectedResult: -1,
		},
		// Large arrays
		{
			name:           "Should work with large sorted array",
			array:          []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			target:         13,
			expectedResult: 12,
		},
		// Negative numbers
		{
			name:           "Should work with negative numbers",
			array:          []int{-5, -4, -3, -2, -1, 0, 1},
			target:         -3,
			expectedResult: 2,
		},
		// Special cases
		{
			name:           "Should work with two elements when target is present",
			array:          []int{1, 2},
			target:         2,
			expectedResult: 1,
		},
		{
			name:           "Should work with two elements when target is not present",
			array:          []int{1, 2},
			target:         3,
			expectedResult: -1,
		},
		{
			name:           "Should work with array of same elements when target is present",
			array:          []int{5, 5, 5, 5, 5},
			target:         5,
			expectedResult: 2,
		},
		{
			name:           "Should work with array of same elements when target is not present",
			array:          []int{5, 5, 5, 5, 5},
			target:         4,
			expectedResult: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := binarySearch(tt.array, tt.target)
			assert.Equal(t, tt.expectedResult, result)
		})
	}
}
