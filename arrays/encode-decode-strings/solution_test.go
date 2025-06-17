package encodedecodestrings

import (
	"reflect"
	"testing"
)

func TestEncodeDecode(t *testing.T) {
	cases := []struct {
		input []string
		name  string
	}{
		{[]string{"neet", "code", "love", "you"}, "basic"},
		{[]string{"", ""}, "empty strings"},
		{[]string{"a|b", "c:d", "e\\f"}, "special chars"},
		{[]string{}, "empty list"},
		{[]string{"one"}, "single string"},
		{[]string{"", "abc"}, "empty and non-empty"},
		{[]string{"abc", ""}, "non-empty and empty"},
		{[]string{"a\nb", "c\td"}, "escape chars"},
	}

	for _, tc := range cases {
		encoded := Encode(tc.input)
		decoded := Decode(encoded)
		if !reflect.DeepEqual(decoded, tc.input) {
			t.Errorf("Test %s failed: got %v, want %v", tc.name, decoded, tc.input)
		}
	}
}
