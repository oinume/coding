package arrays

import (
	"reflect"
	"testing"
)

func Test_rotateLeft(t *testing.T) {
	tests := map[string]struct {
		d     int32
		array []int32
		want  []int32
	}{
		"example0": {
			d:     2,
			array: []int32{1, 2, 3, 4, 5},
			want:  []int32{3, 4, 5, 1, 2},
		},
		"example1": {
			d:     4,
			array: []int32{1, 2, 3, 4, 5},
			want:  []int32{5, 1, 2, 3, 4},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := rotateLeft(test.d, test.array)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("[%s] want=%+v, got=%+v", name, test.want, got)
			}
		})
	}
}
