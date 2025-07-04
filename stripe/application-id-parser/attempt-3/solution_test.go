package attempt3

import (
	"reflect"
	"testing"
)

func TestParseApplicationIDs(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []string
	}{
		{
			name:  "test 1",
			input: "10A13414124218B124564356434567430",
			want:  []string{"A134141242", "B12456435643456743"},
		},
		{
			name:  "test 2",
			input: "2A23B124C1235D12340",
			want:  []string{"A2", "B12", "C123", "D1234"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseApplicationIDs(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseApplicationIDs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseWhiteListed(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		whitelisted []string
		want        []string
	}{
		{
			name:        "test 1",
			input:       "10A13414124218B124564356434567430",
			whitelisted: []string{"A134141242"},
			want:        []string{"A134141242"},
		},
		{
			name:        "test 2",
			input:       "2A23B124C1235D12340",
			whitelisted: []string{"A2", "C123"},
			want:        []string{"A2", "C123"},
		},
		{
			name:        "test 2",
			input:       "2A23B124C1235D12340",
			whitelisted: []string{"A2", "F1236"},
			want:        []string{"A2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseWhiteListed(tt.input, tt.whitelisted); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseWhiteListed() = %v, want %v", got, tt.want)
			}
		})
	}
}
