package main

import "testing"

func Test_minMoves(t *testing.T) {
	type args struct {
		target     int
		maxDoubles int
	}
	tests := []struct {
		name    string
		args    args
		wantRes int
	}{
		{name: "test1", args: args{target: 1, maxDoubles: 1}, wantRes: 0},
		{name: "test1", args: args{target: 5, maxDoubles: 0}, wantRes: 4},
		{name: "test1", args: args{target: 19, maxDoubles: 2}, wantRes: 7},
		{name: "test1", args: args{target: 10, maxDoubles: 4}, wantRes: 4},
		{name: "test2", args: args{target: 1e9, maxDoubles: 0}, wantRes: 1e9 - 1},
		{name: "test2", args: args{target: 1, maxDoubles: 0}, wantRes: 0},
		{name: "test2", args: args{target: 1, maxDoubles: 1}, wantRes: 0},
		{name: "test2", args: args{target: 2, maxDoubles: 0}, wantRes: 1},
		{name: "test2", args: args{target: 2, maxDoubles: 1}, wantRes: 1},
		{name: "test3", args: args{target: 3, maxDoubles: 0}, wantRes: 2},
		{name: "test3", args: args{target: 3, maxDoubles: 1}, wantRes: 2},
		{name: "test4", args: args{target: 1e9 - 1, maxDoubles: 1}, wantRes: 500000000},
		{name: "test4", args: args{target: 1e9 - 1, maxDoubles: 100}, wantRes: 49},
		{name: "test4", args: args{target: 1e9, maxDoubles: 100}, wantRes: 41},
		{name: "test5", args: args{target: 10, maxDoubles: 100}, wantRes: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := minMoves(tt.args.target, tt.args.maxDoubles); gotRes != tt.wantRes {
				t.Errorf("minMoves() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
