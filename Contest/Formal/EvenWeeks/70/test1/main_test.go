package main

import "testing"

func Test_minimumCost(t *testing.T) {
	type args struct {
		cost []int
	}
	tests := []struct {
		name    string
		args    args
		wantRes int
	}{
		{name: "test1", args: args{cost: []int{1, 2}}, wantRes: 3},
		{name: "test2", args: args{cost: []int{1, 2, 3}}, wantRes: 5},
		{name: "test3", args: args{cost: []int{6, 5, 7, 9, 2, 2}}, wantRes: 23},
		{name: "test4", args: args{cost: []int{5}}, wantRes: 5},
		{name: "test5", args: args{cost: []int{5, 5}}, wantRes: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := minimumCost(tt.args.cost); gotRes != tt.wantRes {
				t.Errorf("minimumCost() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
