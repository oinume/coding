package main

import "testing"

func Test_repeatedString(t *testing.T) {
	tests := map[string]struct {
		s    string
		n    int64
		want int64
	}{
		"example0": {
			s:    "abcac",
			n:    10,
			want: 4,
		},
		"example1": {
			s:    "aba",
			n:    10,
			want: 7,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := repeatedString(test.s, test.n)
			if test.want != got {
				t.Fatalf("want=%v, got=%v", test.want, got)
			}
		})
	}
}
