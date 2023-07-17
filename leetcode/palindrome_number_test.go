package leetcode

import (
	"fmt"
	"testing"
)

/*
Given an integer x, return true if x is a palindrome, and false otherwise.

Example 1:
Input: x = 121
Output: true
Explanation: 121 reads as 121 from left to right and from right to left.

Example 2:
Input: x = -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.

Example 3:
Input: x = 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
*/

func isPalindrome(x int) bool {
	original := fmt.Sprintf("%d", x)
	reversed := []rune(original)
	for i, j := 0, len(reversed)-1; i < j; i, j = i+1, j-1 {
		reversed[i], reversed[j] = reversed[j], reversed[i]
	}
	return original == string(reversed)
}

func Test_isPalindrome(t *testing.T) {
	tests := map[string]struct {
		input int
		want  bool
	}{
		"121": {
			input: 121,
			want:  true,
		},
		"-121": {
			input: -121,
			want:  false,
		},
		"10": {
			input: 10,
			want:  false,
		},
		"15951": {
			input: 10,
			want:  false,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := isPalindrome(test.input)
			if test.want != got {
				t.Errorf("[%s] want=%v, got=%v", name, test.want, got)
			}
		})
	}
}
