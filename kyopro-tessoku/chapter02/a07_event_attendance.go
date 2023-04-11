package main

/*
ある会社では D 日間にわたってイベントが開催され、N 人が出席します。参加者 i (i = 1,2,...,N ) は Li 日目から Ri 日目まで出席する予定です。各日の出席者数を出力するプログラムを作成してください。

入力形式
D
N
L1 R1
 :
LN RN

D 行にわたって出力してください。 d 行目には、d 日目の出席者数を出力してください。

制約
● 1 ≤ D ≤ 100000
● 1 ≤ N ≤ 100000
● 1≤Li ≤Ri ≤D

入力例 1
8
5
2 3
3 6
5 7
3 7
1 5

出力例 1
1
2
4
3
4
3
2
0
*/

type EventAttendance struct {
	days                 int
	attendees            int
	attendeeStartEndDays [][2]int
	diffEachDay          []int
}

func NewEventAttendance(days, attendees int, attendeeStartEndDays [][2]int) *EventAttendance {
	attendeeStartDays := make([]int, attendees)
	attendeeEndDays := make([]int, attendees)
	for i, startEnd := range attendeeStartEndDays {
		attendeeStartDays[i] = startEnd[0]
		attendeeEndDays[i] = startEnd[1]
	}

	diffEachDay := make([]int, days)
	for i := 0; i < attendees; i++ {
		startDayIndex := attendeeStartDays[i] - 1
		endDayIndex := attendeeEndDays[i]
		diffEachDay[startDayIndex]++
		diffEachDay[endDayIndex]--
	}
	return &EventAttendance{
		days:                 days,
		attendees:            attendees,
		attendeeStartEndDays: attendeeStartEndDays,
		diffEachDay:          diffEachDay,
	}
}

func (ea *EventAttendance) Attendees() []int {
	//	fmt.Printf("diffEachDay = %+v\n", ea.diffEachDay)
	attendees := make([]int, ea.days)
	for i := range attendees {
		attendees[i] += ea.diffEachDay[i]
		if i > 0 {
			attendees[i] += attendees[i-1]
		}
	}
	return attendees
}
