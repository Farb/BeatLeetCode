package main

import (
	"reflect"
	"testing"
)

func Test_insertSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "test", args: args{nums: []int{3, 2, 1}}, want: []int{1, 2, 3}},
		{name: "test", args: args{nums: []int{2, 1}}, want: []int{1, 2}},
		{name: "test", args: args{nums: []int{1, 2}}, want: []int{1, 2}},
		{name: "test", args: args{nums: []int{1}}, want: []int{1}},
		{name: "test", args: args{nums: []int{4, 2, 9, 1}}, want: []int{1, 2, 4, 9}},
		{name: "test", args: args{nums: []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}}, want: []int{1, 3, 4, 5, 6, 6, 6, 8, 9, 14, 25, 49}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insertSort(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("input=%v,insertSort() = %v, want %v", tt.args.nums,got, tt.want)
			}
		})
	}
}
