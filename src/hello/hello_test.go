package main

import "testing"

func Test_fullname(t *testing.T) {
	type args struct {
		fname string
		lname string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fullname(tt.args.fname, tt.args.lname); got != tt.want {
				t.Errorf("fullname() = %v, want %v", got, tt.want)
			}
		})
	}
}
