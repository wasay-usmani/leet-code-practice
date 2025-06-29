package backtracking

import "testing"

func TestExist(t *testing.T) {
	tests := []struct {
		name   string
		board  [][]byte
		word   string
		expect bool
	}{
		{
			"example1",
			[][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
			"ABCCED",
			true,
		},
		{
			"example2",
			[][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
			"SEE",
			true,
		},
		{
			"example3",
			[][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
			"ABCB",
			false,
		},
		{
			"single cell true",
			[][]byte{{'A'}},
			"A",
			true,
		},
		{
			"single cell false",
			[][]byte{{'A'}},
			"B",
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := exist(tt.board, tt.word)
			if got != tt.expect {
				t.Errorf("exist() = %v, want %v", got, tt.expect)
			}
		})
	}
}
