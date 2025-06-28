package hackerrank

import "testing"

func TestCaesarCipher(t *testing.T) {
	tests := []struct {
		name   string
		s      string
		k      int32
		expect string
	}{
		{"example1", "middle-Outz", 2, "okffng-Qwvb"},
		{"wrap around", "xyz", 2, "zab"},
		{"no shift", "abc", 0, "abc"},
		{"large k", "abc", 52, "abc"},
		{"punctuation", "a.b,c!", 1, "b.c,d!"},
		{"upper and lower", "AbC", 1, "BcD"},
		{"all non-alpha", "123!@#", 5, "123!@#"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := caesarCipher(tt.s, tt.k)
			if got != tt.expect {
				t.Errorf("caesarCipher(%q, %d) = %q, want %q", tt.s, tt.k, got, tt.expect)
			}
		})
	}
}
