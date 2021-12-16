package main

import (
	"reflect"
	"testing"
)

func Test_mergeSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "Test1",args: args{nums: []int{1}},want: []int{1}},
		{name: "Test2",args: args{nums: []int{2,1}},want: []int{1,2}},
		{name: "Test2",args: args{nums: []int{1,2}},want: []int{1,2}},
		{name: "Test3",args: args{nums: []int{2,3,1}},want: []int{1,2,3}},
		{name: "Test3",args: args{nums: []int{1,2,3}},want: []int{1,2,3}},
		{name: "Test4",args: args{nums: []int{1,3,2,4}},want: []int{1,2,3,4}},
		{name: "Test5",args: args{nums: []int{5,3,1,4,2}},want: []int{1,2,3,4,5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeSort(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
