package main

import (
	"testing"
)

/*
小さい順に並べられている、要素数 N の配列 A = [A1,A2, ... ,AN ] があります。要素 X は配列 A の何番目に存在するかを出力してください。なお、この問題は単純な全探索(→ 1.2 節)でも解けますが、ここでは二分探索法を使って実装してください。

入力形式
N X
A 1  A 2 ... A N

出力形式
要素 X は配列 A の何番目に存在するかを出力してください。
制約
● 1 ≤ N ≤ 100000
● 1 ≤ A1 < A2 < … < AN ≤ 109
● 整数 X は A1,A2, ... ,AN のいずれかである

入力例 1
15 47
11 13 17 19 23 29 31 37 41 43 47 53 59 61 67

出力例 1
11
A11 の値が 47 になっています。

入力例 2
10 80
10 20 30 40 50 60 70 80 90 100

出力例 2
8
*/

func TestBinarySearch(t *testing.T) {
	tests := map[string]struct {
		size    int
		target  int
		numbers []int
		want    int
	}{
		"example0": {
			size:    5,
			target:  10000,
			numbers: []int{1, 10, 100, 1000, 10000},
			want:    4,
		},
		"example1": {
			size:    15,
			target:  47,
			numbers: []int{11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67},
			want:    10,
		},
		"example2": {
			size:    10,
			target:  80,
			numbers: []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100},
			want:    7,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := BinarySearch(test.size, test.target, test.numbers)
			if test.want != got {
				t.Errorf("[%s] want=%v, got=%v", name, test.want, got)
			}
		})
	}
}

func BinarySearch(size, target int, numbers []int) int {
	// zero index based
	firstIndex := 0
	lastIndex := size - 1
	for firstIndex <= lastIndex {
		centerIndex := (firstIndex + lastIndex) / 2
		if target < numbers[centerIndex] {
			lastIndex = centerIndex - 1
		}
		if target == numbers[centerIndex] {
			return centerIndex
		}
		if target > numbers[centerIndex] {
			firstIndex = centerIndex + 1
		}
	}
	return -1
}
