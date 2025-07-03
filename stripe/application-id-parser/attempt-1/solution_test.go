package applicationidparser

import (
	"reflect"
	"testing"
)

func TestParseApplicationIDs(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect []string
	}{
		{"example", "10A13414124218B124564356434567430", []string{"A134141242", "B12456435643456743"}},
		{"single", "5ABCDE0", []string{"ABCDE"}},
		{"empty", "0", []string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ParseApplicationIDs(tt.input)
			if !reflect.DeepEqual(got, tt.expect) {
				t.Errorf("ParseApplicationIDs(%q) = %v, want %v", tt.input, got, tt.expect)
			}
		})
	}
}

func TestFilterWhitelisted(t *testing.T) {
	ids := " 10A13414124218B124564356434567430"
	whitelist := []string{"A134141242"}
	expect := []string{"A134141242"}
	got := FilterWhitelisted(ids, whitelist)
	if !reflect.DeepEqual(got, expect) {
		t.Errorf("FilterWhitelisted(%v, %v) = %v, want %v", ids, whitelist, got, expect)
	}
}
