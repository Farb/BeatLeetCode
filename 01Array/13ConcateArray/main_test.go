package main

import (
	"reflect"
	"testing"
)

func Test_getConcatenation(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "Test1", args: args{nums: []int{1, 2, 1}}, want: []int{1, 2, 1, 1, 2, 1}},
		{name: "Test1", args: args{nums: []int{1, 3, 2, 1}}, want: []int{1, 3, 2, 1, 1, 3, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getConcatenation(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getConcatenation() = %v, want %v", got, tt.want)
			}
		})
	}
}
