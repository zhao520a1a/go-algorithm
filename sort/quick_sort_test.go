package sort

import (
	"fmt"
	"testing"

	"github.com/duke-git/lancet/v2/algorithm"
)

type intComparator struct{}

func (c *intComparator) Compare(v1 any, v2 any) int {
	val1, _ := v1.(int)
	val2, _ := v2.(int)

	// ascending order
	if val1 < val2 {
		return -1
	} else if val1 > val2 {
		return 1
	}
	return 0
}

func TestQuickSort(t *testing.T) {
	intSlice := []int{2, 1, 5, 3, 6, 4}
	comparator := &intComparator{}
	algorithm.QuickSort(intSlice, 0, len(intSlice)-1, comparator)
	fmt.Println(intSlice)

	QuickSort(intSlice, 0, len(intSlice)-1, comparator)
	fmt.Println(intSlice)
}
