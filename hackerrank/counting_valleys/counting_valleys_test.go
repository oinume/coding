package main

import "testing"

func Test_countingValleys(t *testing.T) {
	tests := map[string]struct {
		steps int32
		path  string
		want  int32
	}{
		"example": {
			steps: 8,
			path:  "UDDDUDUU",
			want:  1,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := countingValleys(test.steps, test.path)
			if test.want != got {
				t.Fatalf("%s: want=%v, got=%v", name, test.want, got)
			}
		})
	}
}
