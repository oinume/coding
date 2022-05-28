package main

import "testing"

func Test_twoStrings(t *testing.T) {
	tests := map[string]struct {
		s1, s2 string
		want   string
	}{
		"example": {
			s1:   "and",
			s2:   "art",
			want: "YES",
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := twoStrings(test.s1, test.s2)
			if test.want != got {
				t.Errorf("want=%q, got=%q", test.want, got)
			}
		})
	}
}
