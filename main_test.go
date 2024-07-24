package main

import (
	"testing"
)

func TestHideLinks(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "Here's my spammy page: http://hehefouls.netHAHAHA see you.",
			expected: "Here's my spammy page: http://******************* see you.",
		},
		{
			input:    "",
			expected: "",
		},
		{
			input:    "Here's my spammy page: http:/hehefouls.netHAHAHA see you.",
			expected: "Here's my spammy page: http:/hehefouls.netHAHAHA see you.",
		},
		{
			input:    "http://hehefouls.netHAHAHA http://hehefouls.netHAHAHA",
			expected: "http://******************* http://*******************",
		},
	}

	for _, test := range tests {
		result := hideLinks(test.input)
		if result != test.expected {
			t.Errorf("\n%q is not equal to \n%q", result, test.expected)
		}
	}
}