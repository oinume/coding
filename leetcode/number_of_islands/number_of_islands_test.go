package number_of_islands

import "testing"

func TestNumIslands(t *testing.T) {
	tests := map[string]struct {
		input [][]byte
		want  int
	}{
		"example1": {
			input: [][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			want: 1,
		},
		"example2": {
			input: [][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			want: 3,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if got := numIslands(test.input); got != test.want {
				t.Errorf("numIslands: got=%v, want=%v", got, test.want)
			}
		})
	}
}
