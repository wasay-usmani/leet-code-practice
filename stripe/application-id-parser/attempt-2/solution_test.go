package attempt2

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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseApplicationIDs(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseApplicationIDs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterWhiteListed(t *testing.T) {

	tests := []struct {
		name        string
		ids         string
		whitelisted []string
		want        []string
	}{
		{
			name:        "Test 1",
			ids:         "10A13414124218B124564356434567430",
			whitelisted: []string{"A134141242"},
			want:        []string{"A134141242"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterWhiteListed(tt.ids, tt.whitelisted); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterWhiteListed() = %v, want %v", got, tt.want)
			}
		})
	}
}
