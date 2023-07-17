package leetcode

import (
	"reflect"
	"testing"
)

/*
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*/
func twoSum(nums []int, target int) []int {
	seen := make(map[int]int, len(nums))
	for i, n := range nums {
		key := target - n
		//fmt.Printf("i = %v, n = %v, key = %v\n", i, n, key)
		if index, exist := seen[key]; exist {
			return []int{index, i}
		}
		seen[n] = i // store index
		//fmt.Printf("seen = %+v\n", seen)
	}
	return nil // not found
}

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := map[string]struct {
		args args
		want []int
	}{
		"normal": {
			args: args{
				nums:   []int{2, 7, 11, 15},
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
