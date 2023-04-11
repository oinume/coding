package main

/*
P.52 A06 How many guest

ある遊園地では N 日間にわたるイベントが開催され、i 日目には Ai 人が来場しました。以下のQ個の質問に答えるプログラムを作成してください。
● 質問 1:L1 日目から R1 日目までの総来場者数は?
● ..
● 質問 Q:LQ 日目から RQ 日目までの総来場者数は?
入力形式
N  Q
A1 A2 ... AN
L1 R1
..
LQ RQ

Q 行にわたって出力してください。j 行目には、質問 j の答えを出力してください。
制約
● 1 ≤ N ≤ 100000
● 1 ≤ Q ≤ 100000
● 1 ≤ Ai ≤ 10000
● 1≤Lj ≤Rj ≤N

入力例 1
15 3
62 65 41 13 20 11 18 44 53 12 18 17 14 10 39
4 13
3 10
2 15

出力例 1
220 212 375
*/

type HowManyGuests struct {
	guests    []int
	guestsSum []int
}

func NewHowManyGuests(guests []int) *HowManyGuests {
	guestsSum := make([]int, len(guests)+1)
	guestsSum[0] = 0
	for i := 1; i <= len(guests); i++ {
		guestsSum[i] = guestsSum[i-1] + guests[i-1]
	}
	//fmt.Printf("guestsSum = %+v\n", guestsSum)
	return &HowManyGuests{
		guests:    guests,
		guestsSum: guestsSum,
	}
}

func (hmg *HowManyGuests) Sum(startDay, endDay int) int {
	return hmg.guestsSum[endDay] - hmg.guestsSum[startDay-1]
}
