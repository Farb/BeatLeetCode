package main

import "testing"

func Test_numWaterBottles(t *testing.T) {
	type args struct {
		numBottles  int
		numExchange int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		 {name: "test1", args: args{numBottles: 9, numExchange: 3}, want: 13},
		{name: "test1", args: args{numBottles: 15, numExchange: 4}, want: 19},
		{name: "test1", args: args{numBottles: 5, numExchange: 5}, want: 6},
		{name: "test1", args: args{numBottles: 2, numExchange: 3}, want: 2},
		{name: "test2", args: args{numBottles: 1, numExchange: 2}, want: 1},
		{name: "test2", args: args{numBottles: 100, numExchange: 100}, want: 101},
		{name: "test3", args: args{numBottles: 15, numExchange: 8}, want: 17},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numWaterBottles(tt.args.numBottles, tt.args.numExchange); got != tt.want {
				t.Errorf("numWaterBottles() = %v, want %v", got, tt.want)
			}
		})
	}
}
