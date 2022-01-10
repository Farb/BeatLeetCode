package main

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Test1", args: args{words: []string{"lc", "cl", "gg"}}, want: 6},
		{name: "Test1", args: args{words: []string{"ab", "ty", "yt", "lc", "cl", "ab"}}, want: 8},
		{name: "Test1", args: args{words: []string{"cc", "ll", "xx"}}, want: 2},
		{name: "Test2", args: args{words: []string{"io", "io"}}, want:0},
		{name: "Test3", args: args{words: []string{"lc", "cl", "gg","gg"}}, want: 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.words); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
