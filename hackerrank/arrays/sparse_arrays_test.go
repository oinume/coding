package main

import (
	"reflect"
	"testing"
)

func Test_matchingStrings(t *testing.T) {
	tests := map[string]struct {
		strs    []string
		queries []string
		want    []int32
	}{
		"example": {
			strs:    []string{"ab", "ab", "abc"},
			queries: []string{"ab", "abc", "bc"},
			want:    []int32{2, 1, 0},
		},
		"sample1": {
			strs:    []string{"aba", "baba", "aba", "xzxb"},
			queries: []string{"aba", "xzxb", "ab"},
			want:    []int32{2, 1, 0},
		},
		"sample2": {
			strs:    []string{"def", "de", "fgh"},
			queries: []string{"de", "lmn", "fgh"},
			want:    []int32{1, 0, 1},
		},
		"sample3": {
			strs:    []string{"abcde", "sdaklfj", "asdjf", "na", "basdn", "sdaklfj", "asdjf", "na", "asdjf", "na", "basdn", "sdaklfj", "asdjf"},
			queries: []string{"abcde", "sdaklfj", "asdjf", "na", "basdn"},
			want:    []int32{1, 3, 4, 3, 2},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := matchingStrings(test.strs, test.queries)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("want=%+v, got=%+v", test.want, got)
			}
		})
	}
}
