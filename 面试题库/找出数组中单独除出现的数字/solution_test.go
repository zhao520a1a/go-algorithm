package 找出数组中单独除出现的数字

import "testing"

func Test_FindDifferentNum(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{arr: []int{1, 2, 1}},
			2,
		},
		{
			"2",
			args{arr: []int{-1, 3, 4, -1, 3}},
			4,
		},

		{
			"3",
			args{arr: []int{4, 7, 6, 4, 5, 5, 6}},
			7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findAloneNum(tt.args.arr); got != tt.want {
				t.Errorf("findAloneNum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkAlone(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name     string
		args     args
		wantNum1 int
		wantNum2 int
	}{
		{
			"1",
			args{arr: []int{1, 3}},
			1,
			3,
		},
		{
			"2",
			args{arr: []int{1, 2, 2, 3}},
			1,
			3,
		},
		{
			"3",
			args{arr: []int{2, 1, 3, 2}},
			1,
			3,
		},
		{
			"4",
			args{arr: []int{2, 1, 2, 3}},
			1,
			3,
		},
		{
			"5",
			args{arr: []int{-1, 1, 3, 1, -1, 5}},
			3,
			5,
		},
		{
			"6",
			args{arr: []int{1, 1, 1, 3, 5, 1, 2, 2}},
			3,
			5,
		},
		{
			"7",
			args{arr: []int{1, 1, 1, 1, 1, 2}},
			1,
			2,
		},
		{
			"8",
			args{arr: []int{-1, -1, -1, -4}},
			-4,
			-1,
		},
		{
			"9",
			args{arr: []int{0, 1, 0, 1, 0, 0, 5, 6}},
			5,
			6,
		},
		{
			"10",
			args{arr: []int{100, 100, 101, 103, 102, 102}},
			101,
			103,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNum1, gotNum2 := findAloneTwoNum(tt.args.arr)
			if gotNum1 != tt.wantNum1 {
				t.Errorf("findAloneTwoNum() gotNum1 = %v, want %v", gotNum1, tt.wantNum1)
			}
			if gotNum2 != tt.wantNum2 {
				t.Errorf("findAloneTwoNum() gotNum2 = %v, want %v", gotNum2, tt.wantNum2)
			}
		})
	}
}
