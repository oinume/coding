package two_sum

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := map[string]struct{
		args args
		want []int
	}{
		"normal": {
			args: args{
				nums: []int{2, 7, 11, 15},
				target: 9,
			},
			want: []int{0, 1},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if got := twoSum(test.args.nums, test.args.target); !reflect.DeepEqual(got, test.want) {
				t.Errorf("twoSum(): got=%v, want=%v", got, test.want)
			}
		})
	}
}
