package main

import (
	"reflect"
	"testing"
)

func Test_quickSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "Test1",args: args{nums: []int{1}},want: []int{1}},
		{name: "Test2",args: args{nums: []int{1,2}},want: []int{1,2}},
		{name: "Test2",args: args{nums: []int{2,1}},want: []int{1,2}},
		{name: "Test3",args: args{nums: []int{3,2,1}},want: []int{1,2,3}},
		{name: "Test3",args: args{nums: []int{2,3,1}},want: []int{1,2,3}},
		{name: "Test3",args: args{nums: []int{1,2,3}},want: []int{1,2,3}},
		{name: "Test3",args: args{nums: []int{2,3,2}},want: []int{2,2,3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			quickSort(tt.args.nums, 0, len(tt.args.nums)-1)
			if !reflect.DeepEqual(tt.args.nums,tt.want) {
				t.Errorf("quickSort()=%v,want=%v",tt.args.nums,tt.want)
			}
		})
	}
}
