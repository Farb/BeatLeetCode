package main

import (
	"math"
	"testing"
)

func Test_judgeSquareSum(t *testing.T) {
	type args struct {
		c int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Test1",args: args{c: 0},want: true},
		{name: "Test1",args: args{c: 1},want: true},
		{name: "Test1",args: args{c: 2},want: true},
		{name: "Test1",args: args{c: 3},want: false},
		{name: "Test1",args: args{c: 11},want: false},
		{name: "Test1",args: args{c: 100},want: true},
		{name: "Test1",args: args{c: math.MaxInt32},want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := judgeSquareSum(tt.args.c); got != tt.want {
				t.Errorf("judgeSquareSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
