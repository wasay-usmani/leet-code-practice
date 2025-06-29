package graphs

import "testing"

func TestMaxAreaOfIsland(t *testing.T) {
	tests := []struct {
		name   string
		grid   [][]int
		expect int
	}{
		{
			"example1",
			[][]int{
				{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
				{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
				{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
			},
			6,
		},
		{
			"all water",
			[][]int{
				{0, 0, 0},
				{0, 0, 0},
			},
			0,
		},
		{
			"all land",
			[][]int{
				{1, 1},
				{1, 1},
			},
			4,
		},
		{
			"single island",
			[][]int{
				{1},
			},
			1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxAreaOfIsland(tt.grid)
			if got != tt.expect {
				t.Errorf("maxAreaOfIsland() = %v, want %v", got, tt.expect)
			}
		})
	}
}
