package main

import (
	"reflect"
	"testing"
)

func Test_highestRankedKItems(t *testing.T) {
	type args struct {
		grid    [][]int
		pricing []int
		start   []int
		k       int
	}
	tests := []struct {
		name    string
		args    args
		wantRes [][]int
	}{
		// {name: "test1", args: args{grid: [][]int{{1, 2, 0, 1}, {1, 3, 3, 1}, {0, 2, 5, 1}}, pricing: []int{2, 3}, start: []int{2, 3}, k: 2}, wantRes: [][]int{{2, 1}, {1, 2}}},
		{name: "test2", args: args{grid: [][]int{{1, 1, 1}, {0, 0, 1}, {2, 3, 4}}, pricing: []int{2, 3}, start: []int{0, 0}, k: 3}, wantRes: [][]int{{2, 1}, {2, 0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := highestRankedKItems(tt.args.grid, tt.args.pricing, tt.args.start, tt.args.k); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("highestRankedKItems() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
