package main

import "testing"

func Test_jumpingOnClouds(t *testing.T) {
	tests := map[string]struct {
		clouds []int32
		want   int32
	}{
		"example1": {
			clouds: []int32{0, 1, 0, 0, 0, 1, 0},
			want:   3,
		},
		"example2": {
			clouds: []int32{0, 0, 1, 0, 0, 1, 0},
			want:   4,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := jumpingOnClouds(test.clouds)
			if test.want != got {
				t.Fatalf("want=%v, got=%v", test.want, got)
			}
		})
	}
}
