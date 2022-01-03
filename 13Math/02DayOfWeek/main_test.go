package main

import "testing"

func Test_dayOfTheWeek(t *testing.T) {
	type args struct {
		day   int
		month int
		year  int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test1", args: args{day: 31, month: 8, year: 2019}, want: "Saturday"},
		{name: "test1", args: args{day: 18, month: 7, year: 1999}, want: "Sunday"},
		{name: "test1", args: args{day: 18, month: 7, year: 1999}, want: "Sunday"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dayOfTheWeek(tt.args.day, tt.args.month, tt.args.year); got != tt.want {
				t.Errorf("dayOfTheWeek() = %v, want %v", got, tt.want)
			}
		})
	}
}
