package main

import "testing"

func Test_sockMerchant(t *testing.T) {
	tests := map[string]struct {
		n     int32
		socks []int32
		want  int32
	}{
		"example": {
			n:     9,
			socks: []int32{10, 20, 20, 10, 10, 30, 50, 10, 20},
			want:  3,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := sockMerchant(test.n, test.socks)
			if test.want != got {
				t.Fatalf("want=%v, got=%v", test.want, got)
			}
		})
	}
}
