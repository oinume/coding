package main

import "testing"

func TestHowManyGuests(t *testing.T) {
	type args struct {
		days         int   // N
		questions    int   // Q
		guests       []int // A
		startEndDays [][2]int
	}

	tests := map[string]struct {
		args args
		want []int
	}{
		"example0": {
			args: args{
				days:      15,
				questions: 3,
				guests:    []int{62, 65, 41, 13, 20, 11, 18, 44, 53, 12, 18, 17, 14, 10, 39},
				startEndDays: [][2]int{
					{4, 13},
					{3, 10},
					{2, 15},
				},
			},
			want: []int{220, 212, 375},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			g := NewHowManyGuests(test.args.guests)
			for i, startEndDay := range test.args.startEndDays {
				got := g.Sum(startEndDay[0], startEndDay[1])
				if test.want[i] != got {
					t.Errorf(
						"[%s][%d] startDay=%v, endDay=%v, want=%+v, got=%+v",
						name, i, startEndDay[0], startEndDay[1], test.want[i], got,
					)
				}
			}
		})
	}
}
