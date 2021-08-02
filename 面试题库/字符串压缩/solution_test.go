package 字符串压缩

import "testing"

func Test_compressString(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"1",
			args{"abcd"},
			"abcd",
		},
		{
			"2",
			args{"abbcd"},
			"abbcd",
		},
		{
			"3",
			args{"aabbbcc"},
			"a2b3c2",
		},
		{
			"4",
			args{"abc"},
			"abc",
		},
		{
			"5",
			args{"ddffddcccff"},
			"d2f2d2c3f2",
		},
		{
			"6",
			args{"xxxyyy"},
			"x3y3",
		},
		{
			"7",
			args{"11ff222555"},
			"12f22353",
		},
		{
			"8",
			args{"zzzz"},
			"z4",
		},
		{
			"9",
			args{"ff"},
			"ff",
		},
		{
			"10",
			args{"testing"},
			"testing",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compressString(tt.args.str); got != tt.want {
				t.Errorf("compressString() = %v, want %v", got, tt.want)
			}
		})
	}
}
