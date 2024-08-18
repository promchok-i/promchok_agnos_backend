package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrongPasswordChecker(t *testing.T) {
	tests := []struct {
		name     string
		password string
		expected int
	}{
		{
			name:     "Short password",
			password: "aA1",
			expected: 3,
		},
		{
			name:     "Missing lowercase",
			password: "ABC123",
			expected: 1,
		},
		{
			name:     "Missing uppercase",
			password: "abc123",
			expected: 1,
		},
		{
			name:     "Missing digit",
			password: "Abcdef",
			expected: 1,
		},
		{
			name:     "Repeating characters",
			password: "aaaBBB111",
			expected: 3,
		},
		{
			name:     "Long password",
			password: "Abc123Abc123Abc123Abc123",
			expected: 4,
		},
		{
			name:     "Strong password",
			password: "1445D1cd",
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := strongPasswordChecker(tt.password)
			assert.Equal(t, tt.expected, result)
		})
	}
}
