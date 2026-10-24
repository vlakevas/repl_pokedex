package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {

	type Case struct {
		input    string
		expected []string
	}

	cases := []Case{
		{
			input:    " ",
			expected: []string{},
		},
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "   Escape    from     LA",
			expected: []string{"escape", "from", "la"},
		},
		{
			input:    " WaSTeD TimES",
			expected: []string{"wasted", "times"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("lengths don't match: '%v' vs '%v'", actual, c.expected)
			continue
		}

		for i := range actual {
			actualWord := actual[i]
			expectedWord := c.expected[i]

			if actualWord != expectedWord {
				t.Errorf("cleanInput(%v) == %v, expected %v", c.input, actual, c.expected)

			}
		}
	}
}
