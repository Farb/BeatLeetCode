package main

import (
	"reflect"
	"testing"
)

func Test_heapify(t *testing.T) {
	type args struct {
		tree []int
		n    int
		i    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "Test1", args: args{tree: []int{1, 2}, n: 2, i: 0}, want: []int{2, 1}},
		{name: "Test1", args: args{tree: []int{1, 2, 3}, n: 3, i: 0}, want: []int{3, 2, 1}},
		{name: "Test1", args: args{tree: []int{3, 2, 4, 1}, n: 4, i: 0}, want: []int{4, 2, 3, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			heapify(tt.args.tree, tt.args.n, tt.args.i)
			if !reflect.DeepEqual(tt.args.tree, tt.want) {
				t.Errorf("heapify()=%v,want=%v", tt.args.tree, tt.want)
			}
		})
	}
}

func Test_buildTree(t *testing.T) {
	type args struct {
		tree []int
		n    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "test1", args: args{tree: []int{1, 2}, n: 2}, want: []int{2, 1}},
		{name: "Test1", args: args{tree: []int{1, 2, 3}, n: 3}, want: []int{3, 2, 1}},
		{name: "Test1", args: args{tree: []int{3, 2, 4, 1}, n: 4}, want: []int{4, 2, 3, 1}},
		{name: "Test2", args: args{tree: []int{1, 2, 3, 4}, n: 4}, want: []int{4, 2, 3, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buildTree(tt.args.tree, tt.args.n)
			if !reflect.DeepEqual(tt.args.tree, tt.want) {
				t.Errorf("buildTree()=%v,want=%v", tt.args.tree, tt.want)
			}
		})
	}
}

func Test_heapSort(t *testing.T) {
	type args struct {
		tree []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "test1", args: args{tree: []int{1}}, want: []int{1}},
		{name: "test2", args: args{tree: []int{2, 1}}, want: []int{1, 2}},
		{name: "test3", args: args{tree: []int{2, 1, 3}}, want: []int{1, 2, 3}},
		{name: "test4", args: args{tree: []int{4, 2, 1, 3}}, want: []int{1, 2, 3, 4}},
		{name: "test5", args: args{tree: []int{4, 2, 5, 7, 9, 8, 1, 3, 6}}, want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			heapSort(tt.args.tree)
			if !reflect.DeepEqual(tt.args.tree, tt.want) {
				t.Errorf("heapSort()=%v,want=%v", tt.args.tree, tt.want)
			}
		})
	}
}
