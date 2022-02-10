package main

import "testing"

func Test_addBinary(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test1", args: args{a: "11", b: "1"}, want: "100"},
		{name: "test2", args: args{a: "1010", b: "1011"}, want: "10101"},
		{name: "test3", args: args{a: "1111", b: "1111"}, want: "11110"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addBinary(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("addBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}
