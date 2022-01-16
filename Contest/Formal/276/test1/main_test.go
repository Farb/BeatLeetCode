package main

import (
	"reflect"
	"testing"
)

func Test_divideString(t *testing.T) {
	type args struct {
		s    string
		k    int
		fill byte
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "test1", args: args{s: "abcdefghi", k: 3, fill: 'x'}, want: []string{"abc", "def", "ghi"}},
		{name: "test1", args: args{s: "abcdefghij", k: 3, fill: 'x'}, want: []string{"abc","def","ghi","jxx"}},
		{name: "test1", args: args{s: "ctoyjrwtngqwt", k: 8, fill: 'n'}, want: []string{"ctoyjrwt","ngqwtnnn"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divideString(tt.args.s, tt.args.k, tt.args.fill); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("divideString() = %v, want %v", got, tt.want)
			}
		})
	}
}
