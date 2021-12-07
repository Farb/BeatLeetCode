package main

import (
	"reflect"
	"testing"
)

func Test_dailyTemperatures(t *testing.T) {
	type args struct {
		temperatures []int
	}
	tests := []struct {
		name    string
		args    args
		wantRes []int
	}{
		// TODO: Add test cases.
		{name: "",args: args{temperatures:  []int{73,74,75,71,69,72,76,73}},wantRes: []int{1,1,4,2,1,1,0,0}},
		{name: "",args: args{temperatures:  []int{30,40,50,60}},wantRes: []int{1,1,1,0}},
		{name: "",args: args{temperatures:  []int{30,60,90}},wantRes: []int{1,1,0}},
		{name: "",args: args{temperatures:  []int{90,60,30}},wantRes: []int{0,0,0}},
		{name: "",args: args{temperatures:  []int{90,60}},wantRes: []int{0,0}},
		{name: "",args: args{temperatures:  []int{60,90}},wantRes: []int{1,0}},
		{name: "",args: args{temperatures:  []int{60}},wantRes: []int{0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := dailyTemperatures(tt.args.temperatures); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("dailyTemperatures() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
