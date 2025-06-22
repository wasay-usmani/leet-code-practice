package hackerrank

import "testing"

func TestTimeConversion(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Sample PM",
			input:    "07:05:45PM",
			expected: "19:05:45",
		},
		{
			name:     "Sample AM",
			input:    "07:05:45AM",
			expected: "07:05:45",
		},
		{
			name:     "Midnight",
			input:    "12:00:00AM",
			expected: "00:00:00",
		},
		{
			name:     "Noon",
			input:    "12:00:00PM",
			expected: "12:00:00",
		},
		{
			name:     "Single digit hour AM",
			input:    "01:00:00AM",
			expected: "01:00:00",
		},
		{
			name:     "Single digit hour PM",
			input:    "01:00:00PM",
			expected: "13:00:00",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := TimeConversion(tc.input)
			if result != tc.expected {
				t.Errorf("expected %s, got %s", tc.expected, result)
			}
		})
	}
}
