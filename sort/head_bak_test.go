package sort

import (
	"reflect"
	"testing"
)

func Test_heapSort(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want []int
	}{
		{
			"1",
			[]int{3, 5, 1, 6},
			[]int{1, 3, 5, 6},
		},
		{
			"2",
			[]int{1, 23, 3, 4, 532, 13},
			[]int{1, 3, 4, 13, 23, 532},
		},
		{
			"3",
			[]int{-1, -3, -8, -1, 0, 2, 4},
			[]int{-8, -3, -1, -1, 0, 2, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := heapSort(tt.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("heapSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
