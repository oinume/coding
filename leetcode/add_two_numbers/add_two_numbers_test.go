package add_two_numbers

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		v1 []int
		v2 []int
	}

	tests := map[string]struct{
		args args
		want []int
	}{
		"normal": {
			args: args{
				v1: []int{2, 4, 3},
				v2: []int{5, 6, 4},
			},
			want: []int{7, 0, 8},
		},
		"no carry": {
			args: args{
				v1: []int{1, 2, 3},
				v2: []int{4, 5, 6},
			},
			want: []int{5, 7, 9},
		},
		"normal2": {
			args: args{
				v1: []int{7, 9, 9},
				v2: []int{3, 4, 2},
			},
			want: []int{0, 4, 2, 1},
		},
		"different input ListNode length": {
			args: args{
				v1: []int{1, 8},
				v2: []int{0},
			},
			want: []int{1, 8},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			l1 := createListNodes(test.args.v1...)
			l2 := createListNodes(test.args.v2...)
			want := createListNodes(test.want...)
			if got := addTwoNumbers(l1, l2); !reflect.DeepEqual(got, want) {
				t.Errorf("\n[%s] addTwoNumbers(): \ngot =%v\nwant=%v", name, got, want)
			} else {
				t.Logf("\n[%s]\ngot =%v\nwant=%v\n", name, got, want)
			}
		})
	}
}

func Test_createListNodes(t *testing.T) {
	node := createListNodes(1, 2, 3)
	for n := node; n != nil ; n = n.Next {
		fmt.Printf("n.Val = %d\n", n.Val)
	}
}
