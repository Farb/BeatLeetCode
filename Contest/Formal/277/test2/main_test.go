package main

import (
	"reflect"
	"testing"
)

func Test_rearrangeArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantRes []int
	}{
		{name: "test1", args: args{nums: []int{3, 1, -2, -5, 2, -4}}, wantRes: []int{3, -2, 1, -5, 2, -4}},
		{name: "test1", args: args{nums: []int{-1, 1}}, wantRes: []int{1, -1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := rearrangeArray(tt.args.nums); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("rearrangeArray() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
