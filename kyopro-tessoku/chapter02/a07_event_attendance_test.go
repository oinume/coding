package main

import (
	"reflect"
	"testing"
)

func TestEventAttendance(t *testing.T) {
	type args struct {
		days                 int // D
		attendees            int // N
		attendeeStartEndDays [][2]int
	}

	tests := map[string]struct {
		args args
		want []int
	}{
		"example0": {
			args: args{
				days:      8,
				attendees: 5,
				attendeeStartEndDays: [][2]int{
					{2, 3},
					{3, 6},
					{5, 7},
					{3, 7},
					{1, 5},
				},
			},
			want: []int{1, 2, 4, 3, 4, 3, 2, 0},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ea := NewEventAttendance(test.args.days, test.args.attendees, test.args.attendeeStartEndDays)
			got := ea.Attendees()
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("want=%+v, got=%+v", test.want, got)
			}
		})
	}
}
