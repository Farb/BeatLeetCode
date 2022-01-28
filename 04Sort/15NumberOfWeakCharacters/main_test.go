package main

import "testing"

func Test_numberOfWeakCharacters(t *testing.T) {
	type args struct {
		props [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantRes int
	}{
		{name: "test1", args: args{props: [][]int{{5, 5}, {6, 3}, {3, 6}}}, wantRes: 0},
		{name: "test2", args: args{props: [][]int{{2, 2}, {3, 3}}}, wantRes: 1},
		{name: "test3", args: args{props: [][]int{{1, 5}, {10, 4}, {4, 3}}}, wantRes: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := numberOfWeakCharacters(tt.args.props); gotRes != tt.wantRes {
				t.Errorf("numberOfWeakCharacters() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
