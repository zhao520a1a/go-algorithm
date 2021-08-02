package 识别山脉数组

import "testing"

func Test_validMountainArray(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"1",
			args{arr: []int{1}},
			false,
		},
		{
			"2",
			args{arr: []int{2, 1}},
			false,
		},
		{
			"3",
			args{arr: []int{3, 2, 1}},
			false,
		},
		{
			"4",
			args{arr: []int{5, 4, 5}},
			false,
		},
		{
			"5",
			args{arr: []int{1, 2, 3, 4}},
			false,
		},
		{
			"6",
			args{arr: []int{5, 4, 5, 3}},
			false,
		},
		{
			"7",
			args{arr: []int{2, 2, 2, 2}},
			false,
		},
		{
			"8",
			args{arr: []int{2, 3, 3, 2}},
			false,
		},
		{
			"9",
			args{arr: []int{-3, -1, 0, 1, -1}},
			true,
		},
		{
			"10",
			args{arr: []int{0, 1, 2, 3, 4, 2}},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validMountainArray(tt.args.arr); got != tt.want {
				t.Errorf("validMountainArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
