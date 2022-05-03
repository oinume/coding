package main

import "testing"

func Test_repeatedString(t *testing.T) {
	tests := map[string]struct {
		s    string
		n    int64
		want int64
	}{
		"example": {
			s:    "abcac",
			n:    10,
			want: 4,
		},
		"example0": {
			s:    "a",
			n:    1000000000000,
			want: 1000000000000,
		},
		"example1": {
			s:    "aba",
			n:    10,
			want: 7,
		},
		// 10 / len(aba):3 = 3 mod 1
		// (2 * 3 = 6) + substr(1, "aba")
		//
		// A = s の中にある`a`の数
		// X = n / len(s)
		// Y = n mod len(s)
		// として、A * X + Yにaが含まれている数
		/*
			example0: A=1, X=1000000000000, Y=0 -> 1 * 1000000000000 + 0
			example: A=2, X=2, Y=0 -> 2 * 2 + 0
			example1: A=2, X=3, Y=1 -> 2 * 3 + 1
		*/
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
