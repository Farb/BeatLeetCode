package main

import "testing"

func Test_mostWordsFound(t *testing.T) {
	type args struct {
		sentences []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1",args: args{sentences: []string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"}},want:6 },
		{name: "test1",args: args{sentences: []string{"please wait", "continue to fight", "continue to win"}},want:3 },
		{name: "test2",args: args{sentences: []string{"please", "continue", "continue"}},want:1 },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mostWordsFound(tt.args.sentences); got != tt.want {
				t.Errorf("mostWordsFound() = %v, want %v", got, tt.want)
			}
		})
	}
}
