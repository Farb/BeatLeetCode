package main

import "testing"

func Test_minNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test1", args: args{nums: []int{10, 2}}, want: "102"},
		{name: "test1", args: args{nums: []int{3, 30, 34, 5, 9}}, want: "3033459"},
		{name: "test1", args: args{nums: []int{5, 30, 34, 3, 9}}, want: "3033459"},
		{name: "test1", args: args{nums: []int{9, 5, 30, 34, 3}}, want: "3033459"},
		{name: "test2", args: args{nums: []int{1, 1, 1}}, want: "111"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minNumber(tt.args.nums); got != tt.want {
				t.Errorf("minNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
