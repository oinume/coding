package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

/*
https://leetcode.com/problems/add-two-numbers/

You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func (ln *ListNode) String() string {
	return fmt.Sprintf("{Val=%d, Next=%+v}", ln.Val, ln.Next)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	dummyHead := &ListNode{}
	current := dummyHead
	for n1, n2 := l1, l2; n1 != nil || n2 != nil; {
		v1, v2 := 0, 0
		if n1 != nil {
			v1 = n1.Val
		}
		if n2 != nil {
			v2 = n2.Val
		}
		sum := v1 + v2 + carry
		carry = sum / 10

		current.Next = &ListNode{
			Val: sum % 10,
		}
		current = current.Next

		if n1 != nil {
			n1 = n1.Next
		}
		if n2 != nil {
			n2 = n2.Next
		}
	}

	if carry > 0 {
		current.Next = &ListNode{
			Val: carry,
		}
	}

	return dummyHead.Next
}

func createListNodes(numbers ...int) *ListNode {
	if len(numbers) == 0 {
		return nil
	}
	nodes := make([]*ListNode, len(numbers))
	for i, n := range numbers {
		nodes[i] = &ListNode{
			Val: n,
		}
	}
	for i, n := range nodes {
		if i < len(nodes)-1 {
			n.Next = nodes[i+1]
		}
	}
	return nodes[0]
}

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		v1 []int
		v2 []int
	}

	tests := map[string]struct {
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
	for n := node; n != nil; n = n.Next {
		fmt.Printf("n.Val = %d\n", n.Val)
	}
}
