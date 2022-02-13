package main

import (
	"reflect"
	"testing"
)

func Test_kWeakestRows(t *testing.T) {
	type args struct {
		mat [][]int
		k   int
	}
	tests := []struct {
		name    string
		args    args
		wantRes []int
	}{
		{name: "test1", args: args{mat: [][]int{
			{1, 1, 0, 0, 0},
			{1, 1, 1, 1, 0},
			{1, 0, 0, 0, 0},
			{1, 1, 0, 0, 0},
			{1, 1, 1, 1, 1}}, k: 3}, wantRes: []int{2, 0, 3}},
		{name: "Test2", args: args{mat: [][]int{
			{1, 0, 0, 0},
			{1, 1, 1, 1},
			{1, 0, 0, 0},
			{1, 0, 0, 0},
		},k: 2,}, wantRes: []int{0, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := kWeakestRows(tt.args.mat, tt.args.k); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("kWeakestRows() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
