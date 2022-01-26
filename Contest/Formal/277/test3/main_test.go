package main

import (
	"reflect"
	"testing"
)

func Test_findLonely(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantRes []int
	}{
		{name: "test1",args: args{nums: []int{10,6,5,8}},wantRes: []int{10,8}},
		{name: "test1",args: args{nums: []int{1,3,5,3}},wantRes: []int{1,5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := findLonely(tt.args.nums); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("findLonely() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
