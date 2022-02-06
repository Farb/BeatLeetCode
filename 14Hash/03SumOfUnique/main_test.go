package main

import "testing"

func Test_sumOfUnique(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantRes int
	}{
		{name: "test1", args: args{nums: []int{1, 2, 3, 2}}, wantRes: 4},
		{name: "test2", args: args{nums: []int{1, 1, 1, 1, 1}}, wantRes: 0},
		{name: "test3", args: args{nums: []int{1, 2, 3, 4, 5}}, wantRes: 15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := sumOfUnique(tt.args.nums); gotRes != tt.wantRes {
				t.Errorf("sumOfUnique() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
