package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFactorial(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{
			name:     "factorial of 0",
			input:    0,
			expected: 1,
		},
		{
			name:     "factorial of 1",
			input:    1,
			expected: 1,
		},
		{
			name:     "factorial of 5",
			input:    5,
			expected: 120,
		},
		{
			name:     "factorial of 10",
			input:    10,
			expected: 3628800,
		},
		{
			name:     "factorial of 7",
			input:    7,
			expected: 5040,
		},
		{
			name:     "factorial of 3",
			input:    3,
			expected: 6,
		},
		{
			name:     "factorial of 12",
			input:    12,
			expected: 479001600,
		},
		{
			name:     "factorial of negative number",
			input:    -5,
			expected: 1, // assuming our function returns 0 for negative inputs
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := factorial(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
