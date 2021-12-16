package main

import (
	"testing"
)

func Test_getLeastNumbers(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "test1", args: args{arr: []int{3, 2, 1}, k: 2}, want: []int{1, 2}},
		{name: "test1", args: args{arr: []int{0, 1, 2, 1}, k: 1}, want: []int{0}},
		{name: "test1", args: args{arr: []int{0, 1, 2, 1}, k: 4}, want: []int{0, 1, 1, 2}},
		{name: "test1", args: args{arr: []int{0,1,1,1,4,5,3,7,7,8,10,2,7,8,0,5,2,16,12,1,19,15,5,18,2,2,22,15,8,22,17,6,22,6,22,26,32,8,10,11,2,26,9,12,9,7,28,33,20,7,2,17,44,3,52,27,2,23,19,56,56,58,36,31,1,19,19,6,65,49,27,63,29,1,69,47,56,61,40,43,10,71,60,66,42,44,10,12,83,69,73,2,65,93,92,47,35,39,13,75}, k: 75}, 
		want: []int{0,0,1,1,1,1,1,1,2,2,2,2,2,2,2,2,3,3,4,5,5,5,6,6,6,7,7,7,7,7,8,8,8,8,9,9,10,10,10,10,11,12,12,12,13,15,15,16,17,17,18,19,19,19,19,20,22,22,22,22,23,26,26,27,27,28,29,31,32,33,35,36,39,40,42}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getLeastNumbers(tt.args.arr, tt.args.k)
			gotMap:=getMap(got)
			wantMap:=getMap(tt.want)
			for k,v:= range wantMap {
				if gotMap[k]!=wantMap[k] {
					t.Errorf("getLeastNumbers() got= %v, want= %v", gotMap[k],v)
				}
			}
		})
	}
}

func getMap(nums []int) map[int]int {
	m := make(map[int]int, len(nums))
	for _, v := range nums {
		m[v] += 1
	}
	return m
}
