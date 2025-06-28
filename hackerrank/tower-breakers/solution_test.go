package hackerrank

import "testing"

func TestTowerBreakers(t *testing.T) {
	tests := []struct {
		name   string
		n, m   int32
		expect int32
	}{
		{"example1", 3, 1, 2},
		{"example2", 1, 4, 1},
		{"example3", 4, 7, 2},
		{"all towers height 1", 5, 1, 2},
		{"single tower height > 1", 1, 2, 1},
		{"even towers, height > 1", 2, 2, 2},
		{"odd towers, height > 1", 3, 2, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := towerBreakers(tt.n, tt.m)
			if got != tt.expect {
				t.Errorf("towerBreakers(%d, %d) = %d, want %d", tt.n, tt.m, got, tt.expect)
			}
		})
	}
}
