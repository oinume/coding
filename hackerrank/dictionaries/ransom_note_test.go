package main

import "testing"

func TestCheckMagazine(t *testing.T) {
	tests := map[string]struct {
		magazine []string
		note     []string
		want     string
	}{
		"sample0": {
			magazine: []string{
				"give", "me", "one", "grand", "today", "night",
			},
			note: []string{
				"give", "one", "grand", "today",
			},
			want: "Yes",
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := checkMagazine(test.magazine, test.note)
			if test.want != got {
				t.Fatalf("want=%v, got=%v", test.want, got)
			}
		})
	}
}
