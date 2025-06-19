package validparentheses

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		s        string
		expected bool
		name     string
	}{
		{"()", true, "simple"},
		{"()[]{}", true, "all types"},
		{"(]", false, "mismatch"},
		{"([)]", false, "wrong order"},
		{"{[]}", true, "nested"},
		{"", true, "empty"},
		{"(((((((", false, "all open"},
		{")))))))", false, "all close"},
		{"([{}])", true, "deep nested"},
		{"[", false, "single open"},
		{"]", false, "single close"},
	}

	for _, tc := range tests {
		got := IsValid(tc.s)
		if got != tc.expected {
			t.Errorf("%s: got %v, want %v", tc.name, got, tc.expected)
		}
	}
}
