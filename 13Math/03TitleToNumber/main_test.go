package main

import "testing"

func Test_titleToNumber(t *testing.T) {
	type args struct {
		columnTitle string
	}
	tests := []struct {
		name    string
		args    args
		wantRes int
	}{
		{name: "test1", args: args{columnTitle: "A"}, wantRes: 1},
		{name: "test2", args: args{columnTitle: "AB"}, wantRes: 28},
		{name: "test3", args: args{columnTitle: "ZY"}, wantRes: 701},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := titleToNumber(tt.args.columnTitle); gotRes != tt.wantRes {
				t.Errorf("titleToNumber() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
