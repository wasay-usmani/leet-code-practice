package graphs

import "testing"

func TestNumIslands(t *testing.T) {
	tests := []struct {
		name   string
		grid   [][]byte
		expect int
	}{
		{
			"example1",
			[][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			1,
		},
		{
			"example2",
			[][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			3,
		},
		{
			"all water",
			[][]byte{
				{'0', '0'},
				{'0', '0'},
			},
			0,
		},
		{
			"all land",
			[][]byte{
				{'1', '1'},
				{'1', '1'},
			},
			1,
		},
		{
			"single island",
			[][]byte{
				{'1'},
			},
			1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := numIslands(tt.grid)
			if got != tt.expect {
				t.Errorf("numIslands() = %v, want %v", got, tt.expect)
			}
		})
	}
}
